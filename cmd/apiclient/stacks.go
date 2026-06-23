package apiclient

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cycloidio/cycloid-cli/gen/models"
	"github.com/cycloidio/cycloid-cli/utils/ptr"
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

// ListStacks lists service catalog stacks for an organization.
//
// Supported LHS filter attributes: service_catalog_ref, service_catalog_visibility
// (values: local, shared, hidden), service_catalog_author, service_catalog_blueprint,
// service_catalog_form_enabled, service_catalog_source_canonical, user_canonical.
func (m *middleware) ListStacks(org string, filters ...LHSFilter) ([]*models.ServiceCatalog, *http.Response, error) {
	var result []*models.ServiceCatalog
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "service_catalogs"},
		LHSFilters:   filters,
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
			return 0, "", fmt.Errorf("no stack catalog version found")
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

// ListStackUseCases lists use cases for a stack version.
//
// NOTE: the backend handler for this route does not call lhs.ParseQuery, so
// LHS filters are accepted by the middleware but silently ignored server-side.
func (m *middleware) ListStackUseCases(org, ref, versionTag, versionBranch, versionCommitHash string, filters ...LHSFilter) ([]*StackUseCase, *http.Response, error) {
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
		LHSFilters:   filters,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// ListStackVersions lists versions for a stack.
//
// NOTE: the backend handler for this route does not call lhs.ParseQuery, so
// LHS filters are accepted by the middleware but silently ignored server-side.
func (m *middleware) ListStackVersions(org, ref string, filters ...LHSFilter) ([]*StackVersion, *http.Response, error) {
	var result []*StackVersion
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "service_catalogs", ref, "versions"},
		LHSFilters:   filters,
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

// ListBlueprints lists stacks flagged as blueprints (service_catalog_blueprint=true).
// Additional caller-supplied filters are merged in.
//
// Supported LHS filter attributes: same as ListStacks (service_catalog_ref,
// service_catalog_visibility, service_catalog_author, service_catalog_source_canonical,
// user_canonical). service_catalog_blueprint is always set to true internally.
func (m *middleware) ListBlueprints(org string, filters ...LHSFilter) ([]*models.ServiceCatalog, *http.Response, error) {
	lhsFilters := append(
		[]LHSFilter{{Attribute: "service_catalog_blueprint", Condition: "eq", Value: "true"}},
		filters...,
	)

	var result []*models.ServiceCatalog
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "service_catalogs"},
		LHSFilters:   lhsFilters,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
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
		return nil, resp, fmt.Errorf("failed to create stack from blueprint: %w", err)
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
		return result, resp, fmt.Errorf(
			"maintainer with canonical %q may not exist, maintainer on stack ref %q has been removed, please check your team canonical argument and ensure that the team exists",
			teamCanonical, ref,
		)
	}

	return result, resp, nil
}
