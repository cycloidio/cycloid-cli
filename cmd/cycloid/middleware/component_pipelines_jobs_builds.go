package middleware

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) GetBuilds(org, project, env, component, pipeline, job string) ([]*models.Build, *http.Response, error) {
	var result []*models.Build
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines", pipeline, "jobs", job, "builds"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) GetBuild(org, project, env, component, pipeline, job, buildID string) (*models.Build, *http.Response, error) {
	var result *models.Build
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines", pipeline, "jobs", job, "builds", buildID},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) CreateBuild(org, project, env, component, pipeline, job string) (*models.Build, *http.Response, error) {
	var result *models.Build
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines", pipeline, "jobs", job, "builds"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) RerunBuild(org, project, env, component, pipeline, job, buildID string) (*models.Build, *http.Response, error) {
	var result *models.Build
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines", pipeline, "jobs", job, "builds", buildID},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) AbortBuild(org, project, env, component, pipeline, job, buildID string) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines", pipeline, "jobs", job, "builds", buildID, "abort"},
	}, nil)
	return resp, err
}

// GetBuildEvents returns the build events as a raw string (text/event-stream).
func (m *middleware) GetBuildEvents(org, project, env, component, pipeline, buildID string) (*string, *http.Response, error) {
	baseURL, err := url.Parse(m.api.Config.URL)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse base url: %w", err)
	}

	baseURL.Path = path.Join(
		baseURL.Path,
		"organizations", org,
		"projects", project,
		"environments", env,
		"components", component,
		"pipelines", pipeline,
		"builds", buildID,
		"events",
	)

	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Authorization", "Bearer "+m.api.GetToken(&org))

	resp, err := m.GenericClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, resp, newAPIResponseError(resp, body)
	}

	s := string(body)
	return &s, resp, nil
}

func (m *middleware) GetBuildPlan(org, project, env, component, pipeline, job, buildID string) (*models.PublicPlan, *http.Response, error) {
	var result *models.PublicPlan
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines", pipeline, "jobs", job, "builds", buildID, "plan"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) GetBuildPreparation(org, project, env, component, pipeline, job, buildID string) (*models.Preparation, *http.Response, error) {
	var result *models.Preparation
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines", pipeline, "jobs", job, "builds", buildID, "preparation"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) GetBuildResources(org, project, env, component, pipeline, job, buildID string) (*models.BuildInputsOutputs, *http.Response, error) {
	var result *models.BuildInputsOutputs
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines", pipeline, "jobs", job, "builds", buildID, "resources"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
