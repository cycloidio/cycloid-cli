package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client/organization_pipelines"
	"github.com/cycloidio/cycloid-cli/client/models"
)

// GetOrgPipelines list all pipeline of a designated org.
func (m *middleware) GetOrgPipelines(org string, concoursePipeline, project, env *string, statuses []string) ([]*models.Pipeline, error) {
	params := organization_pipelines.NewGetPipelinesParams()
	params.SetOrganizationCanonical(org)
	params.WithConcoursePipelineName(concoursePipeline)
	params.WithProjectCanonical(project)
	params.WithEnvironmentCanonical(env)
	params.WithStatuses(statuses)

	resp, err := m.api.OrganizationPipelines.GetPipelines(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}
