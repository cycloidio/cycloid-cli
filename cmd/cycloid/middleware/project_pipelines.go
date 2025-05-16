package middleware

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/client/client/project_pipelines"
	"github.com/cycloidio/cycloid-cli/client/models"
	strfmt "github.com/go-openapi/strfmt"
)

// GetEnvPipelines list all pipeline of a designated env.
func (m *middleware) GetProjectPipelines(org, project string) ([]*models.Pipeline, error) {
	params := project_pipelines.NewGetProjectPipelinesParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)

	resp, err := m.api.ProjectPipelines.GetProjectPipelines(params, m.api.Credentials(&org))
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
