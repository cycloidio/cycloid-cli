package middleware

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"github.com/pkg/errors"
)

func (m *middleware) GetStack(org, ref string) (*models.ServiceCatalog, *http.Response, error) {
	var result *models.ServiceCatalog
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "service_catalogs", ref},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) ListStacks(org string) ([]*models.ServiceCatalog, *http.Response, error) {
	var result []*models.ServiceCatalog
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "service_catalogs"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
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
	versions, _, err := m.ListStackVersions(org, stackRef)
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

func (m *middleware) ListStackUseCases(org, ref, versionTag, versionBranch, versionCommitHash string) ([]*StackUseCase, *http.Response, error) {
	// Resolve version parameters to ID
	versionID, _, err := m.resolveStackVersion(org, ref, versionTag, versionBranch, versionCommitHash)
	if err != nil {
		return nil, nil, err
	}

	query := url.Values{
		"service_catalog_source_version_id": []string{strconv.FormatUint(uint64(versionID), 10)},
	}

	var result []*StackUseCase
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "service_catalogs", ref, "use_cases"},
		Query:        query,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) ListStackVersions(org, ref string) ([]*StackVersion, *http.Response, error) {
	var result []*StackVersion
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "service_catalogs", ref, "versions"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// ResolveStackVersion is the public API for resolving a stack version string to (versionID, commitHash).
// The stackVersion parameter is treated as a tag name; if empty, the default version is used.
func (m *middleware) ResolveStackVersion(org, ref, stackVersion string) (uint32, string, error) {
	return m.resolveStackVersion(org, ref, stackVersion, "", "")
}

// getDefaultCatalogVersion returns the default catalog version for a stack based on priority:
// 1. If a version with is_latest=true and type="tag" exists, use that
// 2. Otherwise, use the latest commit of the branch of the catalog repository of the stack
func (m *middleware) getDefaultCatalogVersion(org, ref string) (*StackVersion, error) {
	stack, _, err := m.GetStack(org, ref)
	if err != nil {
		return nil, err
	}

	var catalogRepoBranch string
	if stack.ServiceCatalogSourceCanonical != "" {
		catalogRepo, _, err := m.GetCatalogRepository(org, stack.ServiceCatalogSourceCanonical)
		if err != nil {
			return nil, err
		}
		catalogRepoBranch = catalogRepo.Branch
	}

	versions, _, err := m.ListStackVersions(org, ref)
	if err != nil {
		return nil, err
	}

	var branchVersion *StackVersion
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
func (m *middleware) ListBlueprints(org string) ([]*models.ServiceCatalog, *http.Response, error) {
	// This method uses custom LHS filter param like the frontend does: `service_catalog_blueprint[eq]=true`
	// We use url.Values directly here - the encodeQuery will encode this as-is
	// Note: url.Values{} Encode() will URL-encode the brackets, so we manually build the query string
	query := url.Values{}
	query.Set("organization_canonical", org)
	query.Set("service_catalog_blueprint[eq]", "true")

	var stacks []*models.ServiceCatalog
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "service_catalogs"},
		Query:        query,
	}, &stacks)
	if err != nil {
		return nil, resp, err
	}

	var validBlueprints []*models.ServiceCatalog
	for _, catalog := range stacks {
		if catalog.Blueprint {
			validBlueprints = append(validBlueprints, catalog)
		}
	}

	return validBlueprints, resp, nil
}

func (m *middleware) CreateStackFromBlueprint(org, blueprintRef, name, stack, catalogRepository, useCase string) (*models.ServiceCatalog, *http.Response, error) {
	body := &models.NewServiceCatalogFromTemplate{
		Name:                          &name,
		Canonical:                     &stack,
		ServiceCatalogSourceCanonical: &catalogRepository,
		UseCase:                       &useCase,
	}

	var result *models.ServiceCatalog
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "service_catalogs", blueprintRef, "template"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, errors.Wrap(err, "failed to create stack from blueprint")
	}
	return result, resp, nil
}

func (m *middleware) UpdateStack(
	org, ref, teamCanonical string,
	visibility *string,
) (*models.ServiceCatalog, *http.Response, error) {
	body := &models.UpdateServiceCatalog{
		TeamCanonical: teamCanonical,
		Visibility:    visibility,
	}

	var result *models.ServiceCatalog
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "service_catalogs", ref},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}

	// TODO: This is a local fix for https://github.com/cycloidio/youdeploy-http-api/issues/5020
	// Remove this condition when backend will be fixed
	// If the team attribute is nil, this means that the backend did not found the maitainer canonical
	if teamCanonical != "" && result.Team == nil {
		return result, resp, errors.Errorf(
			"maintainer with canonical '%s' may not exists, maintainer on stack ref '%s' has been removed, please check you team canonical argument and ensure that the team exists.",
			teamCanonical, ref,
		)
	}

	return result, resp, nil
}
