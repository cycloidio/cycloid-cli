package middleware

import (
	"net/http"
	"net/url"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func componentsPipelineName(project, env string) string {
	return project + "-" + env
}

func (m *middleware) PausePipeline(org, project, env, component, pipelineName string) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines", pipelineName, "pause"},
	}, nil)
	return resp, err
}

func (m *middleware) UnpausePipeline(org, project, env, component, pipelineName string) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines", pipelineName, "unpause"},
	}, nil)
	return resp, err
}

func (m *middleware) DiffPipeline(org, project, env, component, pipelineName, yamlPipeline, yamlVariables string, checkCredentials bool) (*models.PipelineDiffs, *http.Response, error) {
	body := &models.UpdatePipeline{
		PassedConfig:     &yamlPipeline,
		YamlVars:         yamlVariables,
		CheckCredentials: checkCredentials,
	}

	var result *models.PipelineDiffs
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines", pipelineName, "diff"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) CreatePipeline(org, project, env, pipeline, component, yamlPipeline, yamlVariables string, checkCredentials bool) (*models.Pipeline, *http.Response, error) {
	pipelineName := componentsPipelineName(project, env)

	body := &models.NewPipeline{
		PipelineName:     &pipelineName,
		PassedConfig:     yamlPipeline,
		YamlVars:         yamlVariables,
		CheckCredentials: checkCredentials,
	}

	var result *models.Pipeline
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) RenamePipeline(org, project, env, component, pipeline, newName string) (*http.Response, error) {
	// renamePipeline uses a query param for the new name
	query := url.Values{"pipeline_name": []string{newName}}

	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines", pipeline, "rename"},
		Query:        query,
	}, nil)
	return resp, err
}

func (m *middleware) SyncedPipeline(org, project, env, component, pipeline string) (*models.PipelineStatus, *http.Response, error) {
	var result *models.PipelineStatus
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines", pipeline, "synced"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) GetPipeline(org, project, env, component, pipeline string) (*models.Pipeline, *http.Response, error) {
	var result *models.Pipeline
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines", pipeline},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) UpdatePipeline(org, project, env, component, pipelineName, yamlPipeline, yamlVariables string, checkCredentials bool) (*models.Pipeline, *http.Response, error) {
	body := &models.UpdatePipeline{
		PassedConfig:     &yamlPipeline,
		YamlVars:         yamlVariables,
		CheckCredentials: checkCredentials,
	}

	var result *models.Pipeline
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines", pipelineName},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) DeletePipeline(org, project, env, component, pipeline string) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "pipelines", pipeline},
	}, nil)
	return resp, err
}
