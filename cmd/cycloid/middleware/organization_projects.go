package middleware

import (
	"fmt"
	"net/http"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) ListProjects(org string) ([]*models.Project, *http.Response, error) {
	var result []*models.Project
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) ListProjectsEnv(org, project string) ([]*models.Environment, *http.Response, error) {
	var result []*models.Environment
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) GetProject(org, project string) (*models.Project, *http.Response, error) {
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

func (m *middleware) CreateProject(org, projectName, project, description, configRepository, owner, team, color, icon string) (*models.Project, *http.Response, error) {
	projectName, project, err := NameOrCanonical(&projectName, &project)
	if err != nil {
		return nil, nil, err
	}

	body := &models.NewProject{
		Name:                      &projectName,
		Description:               description,
		Canonical:                 project,
		ConfigRepositoryCanonical: &configRepository,
		Owner:                     owner,
		Icon:                      icon,
		Color:                     color,
		TeamCanonical:             team,
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

func (m *middleware) UpdateProject(org, projectName, project, description, configRepository, owner, team, color, icon, cloudProvider string) (*models.Project, *http.Response, error) {
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

func (m *middleware) DeleteProject(org, project string) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project},
	}, nil)
	return resp, err
}
