package middleware

import (
	"fmt"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/client/client/component_pipelines"
	"github.com/cycloidio/cycloid-cli/client/models"
)

func componentsPipelineName(project, env string) string {
	return project + "-" + env
}

func (m *middleware) PausePipeline(org, project, env, component, pipelineName string) error {
	params := component_pipelines.NewPausePipelineParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipelineName)

	_, err := m.api.ComponentPipelines.PausePipeline(params, m.api.Credentials(&org))
	if err != nil {
		return NewApiError(err)
	}

	return nil
}

func (m *middleware) UnpausePipeline(org, project, env, component, pipelineName string) error {
	params := component_pipelines.NewUnpausePipelineParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipelineName)

	_, err := m.api.ComponentPipelines.UnpausePipeline(params, m.api.Credentials(&org))
	if err != nil {
		return NewApiError(err)
	}

	return nil
}

func (m *middleware) DiffPipeline(org, project, env, component, pipelineName, yamlPipeline, yamlVariables string, checkCredentials bool) (*models.PipelineDiffs, error) {
	params := component_pipelines.NewDiffPipelineParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipelineName)

	body := &models.UpdatePipeline{
		PassedConfig:     &yamlPipeline,
		YamlVars:         yamlVariables,
		CheckCredentials: checkCredentials,
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	resp, err := m.api.ComponentPipelines.DiffPipeline(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	if err := payload.Validate(strfmt.Default); err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}

func (m *middleware) CreatePipeline(org, project, env, pipeline, component, yamlPipeline, yamlVariables string, checkCredentials bool) (*models.Pipeline, error) {
	params := component_pipelines.NewCreatePipelineParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)

	pipelineName := componentsPipelineName(project, env)

	body := &models.NewPipeline{
		PipelineName:     &pipelineName,
		PassedConfig:     yamlPipeline,
		YamlVars:         yamlVariables,
		CheckCredentials: checkCredentials,
	}
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	params.SetBody(body)
	resp, err := m.api.ComponentPipelines.CreatePipeline(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	if err := payload.Validate(strfmt.Default); err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}

func (m *middleware) RenamePipeline(org, project, env, component, pipeline, newName string) error {
	params := component_pipelines.NewRenamePipelineParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipeline)
	params.SetPipelineName(newName)

	_, err := m.api.ComponentPipelines.RenamePipeline(params, m.api.Credentials(&org))
	if err != nil {
		return NewApiError(err)
	}

	return nil
}

func (m *middleware) SyncedPipeline(org, project, env, component, pipeline string) (*models.PipelineStatus, error) {
	params := component_pipelines.NewSyncedPipelineParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipeline)

	resp, err := m.api.ComponentPipelines.SyncedPipeline(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	if err := payload.Validate(strfmt.Default); err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}

func (m *middleware) GetPipeline(org, project, env, component, pipeline string) (*models.Pipeline, error) {
	params := component_pipelines.NewGetPipelineParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipeline)

	resp, err := m.api.ComponentPipelines.GetPipeline(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	if err := payload.Validate(strfmt.Default); err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}

func (m *middleware) UpdatePipeline(org, project, env, component, pipelineName, yamlPipeline, yamlVariables string, checkCredentials bool) (*models.Pipeline, error) {
	params := component_pipelines.NewUpdatePipelineParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipelineName)

	body := &models.UpdatePipeline{
		PassedConfig:     &yamlPipeline,
		YamlVars:         yamlVariables,
		CheckCredentials: checkCredentials,
	}
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	params.SetBody(body)

	resp, err := m.api.ComponentPipelines.UpdatePipeline(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	if err := payload.Validate(strfmt.Default); err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}

func (m *middleware) DeletePipeline(org, project, env, component, pipeline string) error {
	params := component_pipelines.NewDeletePipelineParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipeline)

	_, err := m.api.ComponentPipelines.DeletePipeline(params, m.api.Credentials(&org))
	if err != nil {
		return NewApiError(err)
	}

	return nil
}
