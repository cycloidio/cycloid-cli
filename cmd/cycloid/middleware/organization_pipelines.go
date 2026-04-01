package middleware

import (
	"net/http"
	"net/url"

	"github.com/cycloidio/cycloid-cli/client/models"
)

// GetOrgPipelines list all pipeline of a designated org.
func (m *middleware) GetOrgPipelines(org string, concoursePipeline, project, env *string, statuses []string) ([]*models.Pipeline, *http.Response, error) {
	query := url.Values{}
	if concoursePipeline != nil {
		query.Set("concourse_pipeline_name", *concoursePipeline)
	}
	if project != nil {
		query.Set("project_canonical", *project)
	}
	if env != nil {
		query.Set("environment_canonical", *env)
	}
	for _, s := range statuses {
		query.Add("statuses", s)
	}

	var result []*models.Pipeline
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "pipelines"},
		Query:        query,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
