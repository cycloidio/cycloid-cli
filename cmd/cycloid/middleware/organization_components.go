package middleware

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
)

func (m *middleware) GetComponentConfig(org, project, env, component string) (models.FormVariables, *http.Response, error) {
	var result models.FormVariables
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "config"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) GetComponent(org, project, env, component string) (*models.Component, *http.Response, error) {
	var result *models.Component
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) ListComponents(org, project, env string) ([]*models.Component, *http.Response, error) {
	var result []*models.Component
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// newComponentBody is a local body struct used for component upsert.
// It includes service_catalog_source_version_id which the backend still requires
// but which was removed from the generated models.NewComponent.
type newComponentBody struct {
	Canonical                             string               `json:"canonical,omitempty"`
	CloudProviderCanonical                string               `json:"cloud_provider_canonical,omitempty"`
	Description                           string               `json:"description,omitempty"`
	Name                                  *string              `json:"name"`
	ServiceCatalogRef                     *string              `json:"service_catalog_ref"`
	ServiceCatalogSourceVersionID         *uint32              `json:"service_catalog_source_version_id,omitempty"`
	ServiceCatalogSourceVersionCommitHash *string              `json:"service_catalog_source_version_commit_hash,omitempty"`
	UseCase                               *string              `json:"use_case,omitempty"`
	Vars                                  models.FormVariables `json:"vars,omitempty"`
}

// CreateOrUpdateComponent creates or updates a component with the provided configuration,
// including syncing the Concourse pipeline. Uses PUT on the collection endpoint.
func (m *middleware) CreateOrUpdateComponent(org, project, env, component, description, name, stackRef, versionTag, versionBranch, versionCommitHash, useCase, cloudProvider string, vars models.FormVariables) (*models.Component, *http.Response, error) {
	versionID, commitHash, err := m.resolveStackVersion(org, stackRef, versionTag, versionBranch, versionCommitHash)
	if err != nil {
		return nil, nil, err
	}
	var useCasePtr *string
	if useCase != "" {
		useCasePtr = ptr.Ptr(useCase)
	}
	body := &newComponentBody{
		Canonical:                             component,
		CloudProviderCanonical:                cloudProvider,
		Description:                           description,
		Name:                                  ptr.Ptr(name),
		ServiceCatalogRef:                     ptr.Ptr(stackRef),
		ServiceCatalogSourceVersionID:         ptr.Ptr(versionID),
		ServiceCatalogSourceVersionCommitHash: ptr.Ptr(commitHash),
		UseCase:                               useCasePtr,
		Vars:                                  vars,
	}

	var result *models.Component
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create or update component: %w", err)
	}
	return result, resp, nil
}

func (m *middleware) MigrateComponent(org, project, env, component, targetProject, targetEnv, newCanonical, newName string) (*models.Component, *http.Response, error) {
	body := models.MigrateComponent{
		DestinationProjectCanonical:     targetProject,
		DestinationEnvironmentCanonical: targetEnv,
		DestinationComponentCanonical:   newCanonical,
		DestinationComponentName:        newName,
	}

	var result *models.Component
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "migrate"},
		Body:         &body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to migrate component: %w", err)
	}
	return result, resp, nil
}

func (m *middleware) DeleteComponent(org, project, env, component string) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component},
	}, nil)
	return resp, err
}

func (m *middleware) GetComponentStackConfig(org, project, env, component, useCase, versionTag, versionBranch, versionCommitHash string) (models.ServiceCatalogConfigs, *http.Response, error) {
	// Need to get component to determine stack ref
	comp, _, err := m.GetComponent(org, project, env, component)
	if err != nil {
		return nil, nil, err
	}

	stackRef := *comp.ServiceCatalog.Ref

	// Resolve version parameters to ID and commit hash
	versionID, commitHash, err := m.resolveStackVersion(org, stackRef, versionTag, versionBranch, versionCommitHash)
	if err != nil {
		return nil, nil, err
	}

	query := url.Values{
		"service_catalog_source_version_id":          []string{strconv.FormatUint(uint64(versionID), 10)},
		"service_catalog_source_version_commit_hash": []string{commitHash},
	}
	if useCase != "" {
		query.Set("use_case", useCase)
	}

	var result models.ServiceCatalogConfigs
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "stack_config"},
		Query:        query,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
