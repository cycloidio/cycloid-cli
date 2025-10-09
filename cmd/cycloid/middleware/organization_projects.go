package middleware

import (
	"fmt"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/client/client/organization_projects"
	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) ListProjects(org string) ([]*models.Project, error) {
	params := organization_projects.NewGetProjectsParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationProjects.GetProjects(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	p := resp.GetPayload()
	d := p.Data

	return d, nil
}

func (m *middleware) ListProjectsEnv(org, project string) ([]*models.Environment, error) {
	params := organization_projects.NewGetProjectEnvironmentsParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)

	resp, err := m.api.OrganizationProjects.GetProjectEnvironments(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	p := resp.GetPayload()
	return p.Data, nil
}

func (m *middleware) GetProject(org, project string) (*models.Project, error) {
	params := organization_projects.NewGetProjectParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)

	resp, err := m.api.OrganizationProjects.GetProject(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	p := resp.GetPayload()
	d := p.Data

	return d, nil
}

func (m *middleware) CreateProject(org, projectName, project, description, configRepository, owner, team, color, icon string) (*models.Project, error) {
	params := organization_projects.NewCreateProjectParams()
	params.WithOrganizationCanonical(org)

	if projectName == "" {
		// If name is empty, use the canonical
		projectName = project
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

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, fmt.Errorf("failed to validate body input for createProject: %w", err)
	}
	params.WithBody(body)

	resp, err := m.api.OrganizationProjects.CreateProject(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) UpdateProject(org, projectName, project, description, configRepository, owner, team, color, icon, cloudProvider string) (*models.Project, error) {
	params := organization_projects.NewUpdateProjectParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)

	body := &models.UpdateProject{
		Name:                      &projectName,
		Description:               description,
		ConfigRepositoryCanonical: configRepository,
		Owner:                     owner,
		Icon:                      icon,
		Color:                     color,
		CloudProvider:             cloudProvider,
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	params.SetBody(body)
	resp, err := m.api.OrganizationProjects.UpdateProject(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) DeleteProject(org, project string) error {
	params := organization_projects.NewDeleteProjectParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)

	_, err := m.api.OrganizationProjects.DeleteProject(params, m.api.Credentials(&org))
	if err != nil {
		return NewAPIError(err)
	}
	return nil
}
