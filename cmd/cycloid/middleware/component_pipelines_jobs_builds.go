package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client/component_pipelines_jobs_builds"
	"github.com/cycloidio/cycloid-cli/client/models"
	strfmt "github.com/go-openapi/strfmt"
)

func (m *middleware) GetBuilds(org, project, env, component, pipeline, job string, ccPageSince, ccPageUntil, ccPageLimit *uint64) ([]*models.Build, error) {
	params := component_pipelines_jobs_builds.NewGetBuildsParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipeline)
	params.SetJobName(job)

	resp, err := m.api.ComponentPipelinesJobsBuilds.GetBuilds(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return payload.Data, nil
}

func (m *middleware) GetBuild(org, project, env, component, pipeline, job string, ccPageSince, ccPageUntil, ccPageLimit *uint64) (*models.Build, error) {
	params := component_pipelines_jobs_builds.NewGetBuildParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipeline)
	params.SetJobName(job)

	resp, err := m.api.ComponentPipelinesJobsBuilds.GetBuild(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return payload.Data, nil
}

func (m *middleware) CreateBuild(org, project, env, component, pipeline, job string, ccPageSince, ccPageUntil, ccPageLimit *uint64) (*models.Build, error) {
	params := component_pipelines_jobs_builds.NewCreateBuildParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipeline)
	params.SetJobName(job)

	resp, err := m.api.ComponentPipelinesJobsBuilds.CreateBuild(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return payload.Data, nil
}

func (m *middleware) RerunBuild(org, project, env, component, pipeline, job, buildID string, ccPageSince, ccPageUntil, ccPageLimit *uint64) (*models.Build, error) {
	params := component_pipelines_jobs_builds.NewRerunBuildParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipeline)
	params.SetJobName(job)
	params.SetBuildID(buildID)

	resp, err := m.api.ComponentPipelinesJobsBuilds.RerunBuild(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return payload.Data, nil
}

func (m *middleware) AbortBuild(org, project, env, component, pipeline, job, buildID string, ccPageSince, ccPageUntil, ccPageLimit *uint64) error {
	params := component_pipelines_jobs_builds.NewAbortBuildParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipeline)
	params.SetJobName(job)
	params.SetBuildID(buildID)

	_, err := m.api.ComponentPipelinesJobsBuilds.AbortBuild(params, m.api.Credentials(&org))
	if err != nil {
		return NewApiError(err)
	}

	return nil
}

// GetBuildEvents will return the Content-Type as string.
func (m *middleware) GetBuildEvents(org, project, env, component, pipeline, buildID string) (*string, error) {
	params := component_pipelines_jobs_builds.NewGetBuildEventsParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipeline)
	params.SetBuildID(buildID)

	resp, err := m.api.ComponentPipelinesJobsBuilds.GetBuildEvents(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	return &resp.ContentType, nil
}

func (m *middleware) GetBuildPlan(org, project, env, component, pipeline, job, buildID string) (*models.PublicPlan, error) {
	params := component_pipelines_jobs_builds.NewGetBuildPlanParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipeline)
	params.SetBuildID(buildID)

	resp, err := m.api.ComponentPipelinesJobsBuilds.GetBuildPlan(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return payload.Data, nil
}

func (m *middleware) GetBuildPreparation(org, project, env, component, pipeline, job, buildID string) (*models.Preparation, error) {
	params := component_pipelines_jobs_builds.NewGetBuildPreparationParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipeline)
	params.SetBuildID(buildID)

	resp, err := m.api.ComponentPipelinesJobsBuilds.GetBuildPreparation(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return payload.Data, nil
}

func (m *middleware) GetBuildResources(org, project, env, component, pipeline, job, buildID string) (*models.BuildInputsOutputs, error) {
	params := component_pipelines_jobs_builds.NewGetBuildResourcesParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetInpathPipelineName(pipeline)
	params.SetBuildID(buildID)

	resp, err := m.api.ComponentPipelinesJobsBuilds.GetBuildResources(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return payload.Data, nil
}
