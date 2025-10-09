package middleware

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/pkg/errors"

	"github.com/cycloidio/cycloid-cli/client/client/organization_projects"
	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) GetEnv(org, project, env string) (*models.Environment, error) {
	params := organization_projects.NewGetEnvironmentParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)

	resp, err := m.api.OrganizationProjects.GetEnvironment(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) CreateEnv(org, project, env, envName, color string) (*models.Environment, error) {
	params := organization_projects.NewCreateEnvironmentParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)

	envBody := models.NewEnvironment{
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
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) UpdateEnv(org, project, env, envName, color string) (*models.Environment, error) {
	params := organization_projects.NewUpdateEnvironmentParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)

	envBody := models.UpdateEnvironment{
		Name:  &envName,
		Color: color,
	}
	params.WithBody(&envBody)

	resp, err := m.api.OrganizationProjects.UpdateEnvironment(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) DeleteEnv(org, project, env string) error {
	params := organization_projects.NewDeleteEnvironmentParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)

	_, err := m.api.OrganizationProjects.DeleteEnvironment(params, m.api.Credentials(&org))
	if err != nil {
		return NewAPIError(err)
	}

	return nil
}
