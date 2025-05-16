package middleware

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/client/client/environment_pipelines"
	"github.com/cycloidio/cycloid-cli/client/models"
	strfmt "github.com/go-openapi/strfmt"
)

// GetEnvPipelines list all pipeline of a designated env.
func (m *middleware) GetEnvPipelines(org, project, env string) ([]*models.Pipeline, error) {
	params := environment_pipelines.NewGetEnvironmentPipelinesParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)

	resp, err := m.api.EnvironmentPipelines.GetEnvironmentPipelines(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}
