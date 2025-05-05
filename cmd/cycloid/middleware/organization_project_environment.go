package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client/organization_projects"
	"github.com/cycloidio/cycloid-cli/client/models"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
)

func (m *middleware) CreateEnv(org, project, env, envName, color string) (*models.Environment, error) {
	params := organization_projects.NewCreateEnvironmentParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)

	var envBody models.NewEnvironment
	envBody = models.NewEnvironment{
		Name:      &envName,
		Canonical: env,
		Color:     color,
	}

	err := envBody.Validate(strfmt.Default)
	if err != nil {
		return nil, errors.Wrap(err, "Validation failed for createEnv argument, check API docs for allow values in createEnv")
	}

	params.WithBody(&envBody)

	resp, err := m.api.OrganizationProjects.CreateEnvironment(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	return resp.Payload.Data, nil
}

func (m *middleware) UpdateEnv(org, project, env, envName, color string) (*models.Environment, error) {
	params := organization_projects.NewUpdateEnvironmentParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)
	params.WithEnvironmentCanonical(env)

	envBody := models.UpdateEnvironment{
		Name:  &envName,
		Color: color,
	}
	params.WithBody(&envBody)

	response, err := m.api.OrganizationProjects.UpdateEnvironment(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	return response.Payload.Data, nil
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
