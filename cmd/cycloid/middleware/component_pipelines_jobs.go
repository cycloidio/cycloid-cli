package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client/component_pipelines_jobs"
	"github.com/cycloidio/cycloid-cli/client/models"
	strfmt "github.com/go-openapi/strfmt"
)

func (m *middleware) GetJobs(org, project, env, component, pipeline string) ([]*models.Job, error) {
	params := component_pipelines_jobs.NewGetJobsParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipeline)

	resp, err := m.api.ComponentPipelinesJobs.GetJobs(params, m.api.Credentials(&org))
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

func (m *middleware) GetJob(org, project, env, component, pipeline, job string) (*models.Job, error) {
	params := component_pipelines_jobs.NewGetJobParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipeline)
	params.SetJobName(job)

	resp, err := m.api.ComponentPipelinesJobs.GetJob(params, m.api.Credentials(&org))
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

func (m *middleware) PauseJob(org, project, env, component, pipeline, job string) error {
	params := component_pipelines_jobs.NewPauseJobParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipeline)
	params.SetJobName(job)

	_, err := m.api.ComponentPipelinesJobs.PauseJob(params, m.api.Credentials(&org))
	if err != nil {
		return NewApiError(err)
	}

	return nil
}

func (m *middleware) UnPauseJob(org, project, env, component, pipeline, job string) error {
	params := component_pipelines_jobs.NewUnpauseJobParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipeline)
	params.SetJobName(job)

	_, err := m.api.ComponentPipelinesJobs.UnpauseJob(params, m.api.Credentials(&org))
	if err != nil {
		return NewApiError(err)
	}

	return nil
}

func (m *middleware) ClearTaskCache(org, project, env, component, pipeline, job, step string) (*models.ClearTaskCache, error) {
	params := component_pipelines_jobs.NewClearTaskCacheParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipeline)
	params.SetJobName(job)
	params.SetStepName(step)

	resp, err := m.api.ComponentPipelinesJobs.ClearTaskCache(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()
	return p.Data, nil
}
