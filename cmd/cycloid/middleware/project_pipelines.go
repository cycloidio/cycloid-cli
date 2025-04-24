package middleware

import (
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
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()
	err = p.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return p.Data, nil
}
