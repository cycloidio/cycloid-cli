package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cycloidio/cycloid-cli/client/client/service_catalogs"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
)

func (m *middleware) GetStack(org, ref string) (*models.ServiceCatalog, error) {
	params := service_catalogs.NewGetServiceCatalogParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogRef(ref)

	resp, err := m.api.ServiceCatalogs.GetServiceCatalog(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) ListStacks(org string) ([]*models.ServiceCatalog, error) {
	params := service_catalogs.NewListServiceCatalogsParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.ServiceCatalogs.ListServiceCatalogs(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

// resolveStackVersion resolves versionTag/versionBranch/versionCommitHash to version ID and commit hash.
// Priority: tag > branch > commitHash (whichever is provided first)
// If all parameters are empty, uses the default catalog version (latest tag or branch HEAD).
func (m *middleware) resolveStackVersion(org, stackRef, versionTag, versionBranch, versionCommitHash string) (versionID uint32, commitHash string, err error) {
	// If all are empty, use default version
	if versionTag == "" && versionBranch == "" && versionCommitHash == "" {
		defaultVersion, err := m.getDefaultCatalogVersion(org, stackRef)
		if err != nil {
			return 0, "", fmt.Errorf("failed to get default stack version: %w", err)
		}

		if defaultVersion == nil || defaultVersion.ID == nil || defaultVersion.CommitHash == nil {
			return 0, "", errors.New("no stack catalog version found")
		}

		return *defaultVersion.ID, *defaultVersion.CommitHash, nil
	}

	// List all versions for the stack
	versions, err := m.ListStackVersions(org, stackRef)
	if err != nil {
		return 0, "", fmt.Errorf("failed to list stack versions: %w", err)
	}

	// Priority 1: tag
	if versionTag != "" {
		for _, version := range versions {
			if ptr.Value(version.Type) == "tag" &&
				ptr.Value(version.Name) == versionTag {
				return *version.ID, *version.CommitHash, nil
			}
		}
		return 0, "", fmt.Errorf("stack version tag %q not found", versionTag)
	}

	// Priority 2: branch
	if versionBranch != "" {
		for _, version := range versions {
			if ptr.Value(version.Type) == "branch" &&
				ptr.Value(version.Name) == versionBranch {
				return *version.ID, *version.CommitHash, nil
			}
		}
		return 0, "", fmt.Errorf("stack version branch %q not found", versionBranch)
	}

	// Priority 3: commit hash
	for _, version := range versions {
		if ptr.Value(version.CommitHash) == versionCommitHash {
			return *version.ID, *version.CommitHash, nil
		}
	}
	return 0, "", fmt.Errorf("stack version commit hash %q not found", versionCommitHash)
}

func (m *middleware) ListStackUseCases(org, ref, versionTag, versionBranch, versionCommitHash string) ([]*models.StackUseCase, error) {
	// Resolve version parameters to ID
	versionID, _, err := m.resolveStackVersion(org, ref, versionTag, versionBranch, versionCommitHash)
	if err != nil {
		return nil, err
	}

	params := service_catalogs.NewGetServiceCatalogUseCasesParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogRef(ref)
	params.SetServiceCatalogSourceVersionID(versionID)

	resp, err := m.api.ServiceCatalogs.GetServiceCatalogUseCases(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) ListStackVersions(org, ref string) ([]*models.ServiceCatalogSourceVersion, error) {
	params := service_catalogs.NewGetServiceCatalogVersionsParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogRef(ref)

	resp, err := m.api.ServiceCatalogs.GetServiceCatalogVersions(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

// getDefaultCatalogVersion returns the default catalog version for a stack based on priority:
// 1. If a version with is_latest=true and type="tag" exists, use that
// 2. Otherwise, use the latest commit of the branch of the catalog repository of the stack
func (m *middleware) getDefaultCatalogVersion(org, ref string) (*models.ServiceCatalogSourceVersion, error) {
	stack, err := m.GetStack(org, ref)
	if err != nil {
		return nil, err
	}

	var catalogRepoBranch string
	if stack.ServiceCatalogSourceCanonical != "" {
		catalogRepo, err := m.GetCatalogRepository(org, stack.ServiceCatalogSourceCanonical)
		if err != nil {
			return nil, err
		}
		catalogRepoBranch = catalogRepo.Branch
	}

	versions, err := m.ListStackVersions(org, ref)
	if err != nil {
		return nil, err
	}

	var branchVersion *models.ServiceCatalogSourceVersion
	// Default to default catalog branch
	for _, version := range versions {
		if ptr.Value(version.Type) == "branch" &&
			ptr.Value(version.Name) == catalogRepoBranch {
			branchVersion = version
		}
	}

	if branchVersion != nil {
		return branchVersion, nil
	}

	return nil, fmt.Errorf("failed to find the default version")
}

// ListBlueprints will list stacks that are flagged as blueprint. Uses the same route as ListStack.
// TODO: Merge this route with ListStack once we find a way to add LHS filter params to the client.
func (m *middleware) ListBlueprints(org string) ([]*models.ServiceCatalog, error) {
	// This method use a custom request because we use the (undocumented)
	// LHS filter param like the frontend does: `service_catalog_blueprint[eq]=true`
	url := fmt.Sprintf("%s/organizations/%s/service_catalogs?organization_canonical=%s&service_catalog_blueprint%%5Beq%%5D=true",
		m.api.Config.URL, org, org)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	token := m.api.GetToken(&org)
	if token == "" {
		return nil, errors.New("missing API key, please provide valid authentication using CY_API_KEY env var")
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var stacks struct {
		Data []*models.ServiceCatalog `json:"data"`
	}

	err = json.Unmarshal(body, &stacks)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	var validBlueprints []*models.ServiceCatalog
	for _, catalog := range stacks.Data {
		if catalog.Blueprint {
			validBlueprints = append(validBlueprints, catalog)
		}
	}

	// Don't validate payload on this route, not supported atm.
	return validBlueprints, nil
}

func (m *middleware) CreateStackFromBlueprint(org, blueprintRef, name, stack, catalogRepository, useCase string) (*models.ServiceCatalog, error) {
	params := service_catalogs.NewCreateServiceCatalogFromTemplateParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogRef(blueprintRef)
	body := &models.NewServiceCatalogFromTemplate{
		Name:                          &name,
		Canonical:                     &stack,
		ServiceCatalogSourceCanonical: &catalogRepository,
		UseCase:                       &useCase,
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, errors.Wrap(err, "validation failed for createStackFromBlueprint input")
	}
	params.WithBody(body)

	resp, err := m.api.ServiceCatalogs.CreateServiceCatalogFromTemplate(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}
	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) UpdateStack(
	org, ref, teamCanonical string,
	visibility *string,
) (*models.ServiceCatalog, error) {
	params := service_catalogs.NewUpdateServiceCatalogParams()
	params.WithOrganizationCanonical(org)
	params.WithServiceCatalogRef(ref)

	body := &models.UpdateServiceCatalog{
		TeamCanonical: teamCanonical,
		Visibility:    visibility,
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, errors.Wrap(err, "validation failed for updateServiceCatalog input")
	}

	params.WithBody(body)

	resp, err := m.api.ServiceCatalogs.UpdateServiceCatalog(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	// TODO: This is a local fix for https://github.com/cycloidio/youdeploy-http-api/issues/5020
	// Remove this condition when backend will be fixed
	// If the team attribute is nil, this means that the backend did not found the maitainer canonical
	if teamCanonical != "" && payload.Data.Team == nil {
		return payload.Data, errors.Errorf(
			"maintainer with canonical '%s' may not exists, maintainer on stack ref '%s' has been removed, please check you team canonical argument and ensure that the team exists.",
			teamCanonical, ref,
		)
	}

	return payload.Data, nil
}
