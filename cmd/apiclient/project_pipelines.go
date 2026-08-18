package apiclient

import (
	"net/http"

	"github.com/cycloidio/cycloid-cli/gen/models"
)

// GetProjectPipelines lists all pipelines of a designated project.
func (m *apiClient) GetProjectPipelines(org, project string) ([]*models.Pipeline, *http.Response, error) {
	var result []*models.Pipeline
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "pipelines"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
