package apiclient

import (
	"fmt"
	"net/http"

	"github.com/cycloidio/cycloid-cli/gen/models"
)

// ListProjects lists projects for an organization.
//
// Supported LHS filter attributes: project_canonical, project_name,
// project_description, project_created_at, project_config_repository_canonical,
// environment_canonical, user_canonical.
func (m *apiClient) ListProjects(org string, filters ...LHSFilter) ([]*models.Project, *http.Response, error) {
	var result []*models.Project
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects"},
		LHSFilters:   filters,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// ListProjectEnvs lists environments linked to a project.
//
// Supported LHS filter attributes: environment_canonical, environment_created_at.
func (m *apiClient) ListProjectEnvs(org, project string, filters ...LHSFilter) ([]*models.ProjectEnvironment, *http.Response, error) {
	result, resp, err := paginatedList[*models.ProjectEnvironment](m, Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments"},
		LHSFilters:   filters,
	}, defaultPageSize)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) GetProject(org, project string) (*models.Project, *http.Response, error) {
	var result *models.Project
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) CreateProject(org, projectName, project, description, configRepository, owner, team, color, icon string) (*models.Project, *http.Response, error) {
	projectName, project, err := NameOrCanonical(&projectName, &project)
	if err != nil {
		return nil, nil, err
	}

	body := &models.NewProject{
		Name:                      &projectName,
		Description:               description,
		Canonical:                 project,
		ConfigRepositoryCanonical: configRepository,
		Owner:                     owner,
		Icon:                      icon,
		Color:                     color,
	}

	var result *models.Project
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "projects"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create project: %w", err)
	}
	return result, resp, nil
}

func (m *apiClient) UpdateProject(org, projectName, project, description, configRepository, owner, team, color, icon, cloudProvider string) (*models.Project, *http.Response, error) {
	body := &models.UpdateProject{
		Name:                      &projectName,
		Description:               description,
		ConfigRepositoryCanonical: configRepository,
		Owner:                     owner,
		Icon:                      icon,
		Color:                     color,
		CloudProvider:             cloudProvider,
	}

	var result *models.Project
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) DeleteProject(org, project string, opts DeleteOptions) (*http.Response, error) {
	req := Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project},
	}
	if q := opts.Resolve(); q.SkipHooks || q.IgnoreConfigFilesErr {
		req.Query = q
	}
	return m.GenericRequest(req, nil)
}
