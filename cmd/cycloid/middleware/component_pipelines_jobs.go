package middleware

import (
	"net/http"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) GetJobs(org, project, env, component, pipeline string) ([]*models.Job, *http.Response, error) {
	var result []*models.Job
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines", pipeline, "jobs"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) GetJob(org, project, env, component, pipeline, job string) (*models.Job, *http.Response, error) {
	var result *models.Job
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines", pipeline, "jobs", job},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) PauseJob(org, project, env, component, pipeline, job string) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines", pipeline, "jobs", job, "pause"},
	}, nil)
	return resp, err
}

func (m *middleware) UnPauseJob(org, project, env, component, pipeline, job string) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines", pipeline, "jobs", job, "unpause"},
	}, nil)
	return resp, err
}

func (m *middleware) ClearTaskCache(org, project, env, component, pipeline, job, step string) (*models.ClearTaskCache, *http.Response, error) {
	var result *models.ClearTaskCache
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines", pipeline, "jobs", job, "tasks", step, "cache"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
