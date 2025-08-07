package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client/environment_pipelines"
	"github.com/cycloidio/cycloid-cli/client/models"
)

// GetEnvPipelines list all pipeline of a designated env.
func (m *middleware) GetEnvPipelines(org, project, env string) ([]*models.Pipeline, error) {
	params := environment_pipelines.NewGetEnvironmentPipelinesParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)

	resp, err := m.api.EnvironmentPipelines.GetEnvironmentPipelines(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}
