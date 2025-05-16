package middleware

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/client/client/organization_pipelines"
	"github.com/cycloidio/cycloid-cli/client/models"
	strfmt "github.com/go-openapi/strfmt"
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
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}
