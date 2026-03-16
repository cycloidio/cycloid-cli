package middleware

import (
	"net/http"

	"github.com/cycloidio/cycloid-cli/client/models"
)

// GetEnvPipelines list all pipeline of a designated env.
func (m *middleware) GetEnvPipelines(org, project, env string) ([]*models.Pipeline, *http.Response, error) {
	var result []*models.Pipeline
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "pipelines"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
