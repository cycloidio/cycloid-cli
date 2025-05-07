package middleware

import (
	"fmt"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/sanity-io/litter"

	"github.com/cycloidio/cycloid-cli/client/client/organization_projects"
	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) ListProjects(org string) ([]*models.Project, error) {
	params := organization_projects.NewGetProjectsParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationProjects.GetProjects(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()
	d := p.Data

	return d, nil
}

func (m *middleware) GetProject(org, project string) (*models.Project, error) {
	params := organization_projects.NewGetProjectParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)

	resp, err := m.api.OrganizationProjects.GetProject(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()
	d := p.Data

	return d, nil
}

// TODO:
//func (m *middleware) GetProjectConfig(org string, project string, env string) (*models.ProjectEnvironmentConfig, error) {
//	params := organization_projects.NewGetProjectConfigParams()
//	params.WithOrganizationCanonical(org)
//	params.WithProjectCanonical(project)
//	params.WithEnvironmentCanonical(env)
//	params.WithDefaults()
//
//	resp, err := m.api.OrganizationProjects.GetProjectConfig(
//		params,
//		m.api.Credentials(&org),
//		// organization_projects.WithContentType("application/vnd.cycloid.io.v1+json"),
//		// organization_projects.WithAccept("application/json"),
//	)
//
//	if err != nil {
//		return nil, NewApiError(err)
//	}
//
//	payload := resp.GetPayload()
//	err = payload.Validate(strfmt.Default)
//	if err != nil {
//		return nil, err
//	}
//
//	return payload.Data, nil
//}

func (m *middleware) CreateProject(org, projectName, project, description, configRepository, owner, team, color, icon string) (*models.Project, error) {
	params := organization_projects.NewCreateProjectParams()
	params.WithOrganizationCanonical(org)
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
		return nil, fmt.Errorf("failed to validate body input for createProject, payload:\n%s\n%v", litter.Sdump(body), err)
	}
	params.WithBody(body)

	response, err := m.api.OrganizationProjects.CreateProject(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := response.GetPayload()
	return payload.Data, nil
}

func (m *middleware) UpdateProject(org, projectName, project, description, configRepository, owner, team, color, icon, cloudProvider string, updatedAt *uint64) (*models.Project, error) {
	params := organization_projects.NewUpdateProjectParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)

	body := &models.UpdateProject{
		Name:                      &projectName,
		Description:               description,
		ConfigRepositoryCanonical: configRepository,
		Owner:                     owner,
		UpdatedAt:                 updatedAt,
		Icon:                      icon,
		Color:                     color,
		CloudProvider:             cloudProvider,
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	params.SetBody(body)
	resp, err := m.api.OrganizationProjects.UpdateProject(
		params,
		m.api.Credentials(&org),
	)
	if err != nil {
		return nil, NewApiError(fmt.Errorf("params:\n%s\nresp:\n%s\nerr:\n%s", litter.Sdump(params), litter.Sdump(resp), err))
	}

	p := resp.GetPayload()
	err = p.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	d := p.Data
	return d, err
}

func (m *middleware) DeleteProject(org, project string) error {
	params := organization_projects.NewDeleteProjectParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)

	_, err := m.api.OrganizationProjects.DeleteProject(params, m.api.Credentials(&org))
	if err != nil {
		return NewApiError(err)
	}
	return nil
}
