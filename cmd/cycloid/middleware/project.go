package middleware

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/pkg/errors"

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
	// err = p.Validate(strfmt.Default)
	// if err != nil {
	// 	return nil, err
	// }

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
	err = p.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	d := p.Data

	return d, nil
}

func (m *middleware) GetProjectConfig(org string, project string, env string) (*models.ProjectEnvironmentConfig, error) {
	params := organization_projects.NewGetProjectConfigParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)
	params.WithEnvironmentCanonical(env)
	params.WithDefaults()

	resp, err := m.api.OrganizationProjects.GetProjectConfig(
		params,
		m.api.Credentials(&org),
		organization_projects.WithContentType("application/vnd.cycloid.io.v1+json"),
		organization_projects.WithAccept("application/json"),
	)
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return payload.Data, nil
}

func (m *middleware) CreateProject(org, projectName, projectCanonical, description, stackRef, configRepo, owner string) (*models.Project, error) {
	params := organization_projects.NewCreateProjectParams()
	params.WithOrganizationCanonical(org)

	body := &models.NewProject{
		Name:                      &projectName,
		Description:               description,
		Canonical:                 projectCanonical,
		ServiceCatalogRef:         &stackRef,
		ConfigRepositoryCanonical: &configRepo,
		Owner:                     owner,
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate body input for createProject")
	}

	params.WithBody(body)

	response, err := m.api.OrganizationProjects.CreateProject(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := response.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate response from API after project creation.")
	}

	return payload.Data, nil
}

/*
Warning: This method is deprecated. You can't create envs from this route anymore, use the env CRUD
*/
func (m *middleware) UpdateProject(org, projectName, projectCanonical string, envs []*models.NewEnvironment, description, stackRef, owner, configRepo string, inputs []*models.FormInput, updatedAt uint64) (*models.Project, error) {
	params := organization_projects.NewUpdateProjectParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(projectCanonical)

	body := &models.UpdateProject{
		Name:                      &projectName,
		Description:               description,
		ServiceCatalogRef:         &stackRef,
		ConfigRepositoryCanonical: configRepo,
		Environments:              envs,
		Owner:                     owner,
		Inputs:                    inputs,
		UpdatedAt:                 &updatedAt,
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	params.SetBody(body)
	resp, err := m.api.OrganizationProjects.UpdateProject(
		params,
		m.api.Credentials(&org),
		// organization_projects.WithContentType("application/vnd.cycloid.io.v1+json"),
		// organization_projects.WithAccept("application/json"),
	)
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()
	err = p.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	d := p.Data
	return d, err
}

func (m *middleware) CreateEnv(org, project, env string) error {
	return errors.New("not implemented")
}
func (m *middleware) UpdateEnv(org, project, env string) error {
	return errors.New("not implemented")
}

func (m *middleware) DeleteEnv(org, project, env string) error {
	params := organization_projects.NewDeleteEnvironmentParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)

	_, err := m.api.OrganizationProjects.DeleteEnvironment(params, m.api.Credentials(&org))
	if err != nil {
		return NewApiError(err)
	}
	return nil
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
