package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client/organization_pipelines"
	"github.com/cycloidio/cycloid-cli/client/client/organization_pipelines_jobs"
	"github.com/cycloidio/cycloid-cli/client/client/organization_pipelines_jobs_build"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	strfmt "github.com/go-openapi/strfmt"
)

func (m *middleware) PausePipeline(org, project, env string) error {

	pipelineName := common.GetPipelineName(project, env)

	params := organization_pipelines.NewPausePipelineParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetInpathPipelineName(pipelineName)

	_, err := m.api.OrganizationPipelines.PausePipeline(params, common.ClientCredentials(&org))

	return err
}

func (m *middleware) ClearTaskCachePipeline(org, project, env, job, task string) error {

	pipelineName := common.GetPipelineName(project, env)

	params := organization_pipelines_jobs.NewClearTaskCacheParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetInpathPipelineName(pipelineName)
	params.SetJobName(job)
	params.SetStepName(task)

	_, err := m.api.OrganizationPipelinesJobs.ClearTaskCache(params, common.ClientCredentials(&org))
	if err != nil {
		return err
	}

	return err
}

func (m *middleware) UnpausePipeline(org, project, env string) error {

	pipelineName := common.GetPipelineName(project, env)

	params := organization_pipelines.NewUnpausePipelineParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetInpathPipelineName(pipelineName)

	_, err := m.api.OrganizationPipelines.UnpausePipeline(params, common.ClientCredentials(&org))
	// if err != nil {
	// 	return nil, err
	// }

	return err
}

func (m *middleware) UpdatePipeline(org, project, env, pipeline, variables string) (*models.Pipeline, error) {

	var body *models.UpdatePipeline

	cyCtx := common.CycloidContext{Env: env,
		Org:     org,
		Project: project}

	pipelineName := common.GetPipelineName(project, env)

	vars := common.ReplaceCycloidVarsString(cyCtx, variables)

	params := organization_pipelines.NewUpdatePipelineParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetInpathPipelineName(pipelineName)

	body = &models.UpdatePipeline{
		PassedConfig: &pipeline,
		YamlVars:     vars,
	}
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	params.SetBody(body)

	resp, err := m.api.OrganizationPipelines.UpdatePipeline(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()
	// TODO this validate have been removed https://github.com/cycloidio/youdeploy-http-api/issues/2262
	// err = p.Validate(strfmt.Default)
	// if err != nil {
	// 	return err
	// }

	d := p.Data
	return d, err
}

func (m *middleware) DiffPipeline(org, project, env, pipeline, variables string) (*models.PipelineDiffs, error) {

	cyCtx := common.CycloidContext{Env: env,
		Org:     org,
		Project: project}
	vars := common.ReplaceCycloidVarsString(cyCtx, variables)
	pipelineName := common.GetPipelineName(project, env)

	params := organization_pipelines.NewDiffPipelineParams()
	params.SetOrganizationCanonical(org)
	params.SetInpathPipelineName(pipelineName)

	body := &models.UpdatePipeline{
		PassedConfig: &pipeline,
		YamlVars:     vars,
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	resp, err := m.api.OrganizationPipelines.DiffPipeline(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()
	// TODO this validate have been removed https://github.com/cycloidio/youdeploy-http-api/issues/2262
	// err = p.Validate(strfmt.Default)
	// if err != nil {
	// 	return err
	// }

	d := p.Data
	return d, err
}

func (m *middleware) CreatePipeline(org, project, env, pipeline, variables, usecase string) (*models.Pipeline, error) {

	cyCtx := common.CycloidContext{Env: env,
		Org:     org,
		Project: project}

	params := organization_pipelines.NewCreatePipelineParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)

	vars := common.ReplaceCycloidVarsString(cyCtx, string(variables))

	pipelineName := common.GetPipelineName(project, env)

	body := &models.NewPipeline{
		Environment:  &env,
		PipelineName: &pipelineName,
		UseCase:      usecase,
		PassedConfig: &pipeline,
		YamlVars:     vars,
	}
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	params.SetBody(body)
	resp, err := m.api.OrganizationPipelines.CreatePipeline(params, common.ClientCredentials(&org))

	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()
	// TODO this validate have been removed https://github.com/cycloidio/youdeploy-http-api/issues/2262
	// err = p.Validate(strfmt.Default)
	// if err != nil {
	// 	return err
	// }

	d := p.Data
	return d, err
}

func (m *middleware) GetPipelineJob(org, project, env, job string) (*models.Job, error) {

	pipelineName := common.GetPipelineName(project, env)

	params := organization_pipelines_jobs.NewGetJobParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetInpathPipelineName(pipelineName)
	params.SetJobName(job)

	resp, err := m.api.OrganizationPipelinesJobs.GetJob(params, common.ClientCredentials(&org))

	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()
	// TODO this validate have been removed https://github.com/cycloidio/youdeploy-http-api/issues/2262
	// err = p.Validate(strfmt.Default)
	// if err != nil {
	// 	return err
	// }

	d := p.Data
	return d, err
}

func (m *middleware) ListPipelineJobsBuilds(org, project, env, job string) ([]*models.Build, error) {

	pipelineName := common.GetPipelineName(project, env)

	params := organization_pipelines_jobs_build.NewGetBuildsParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetInpathPipelineName(pipelineName)
	params.SetJobName(job)

	resp, err := m.api.OrganizationPipelinesJobsBuild.GetBuilds(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()
	// TODO this validate have been removed https://github.com/cycloidio/youdeploy-http-api/issues/2262
	// err = p.Validate(strfmt.Default)
	// if err != nil {
	// 	return err
	// }

	d := p.Data
	return d, err
}

func (m *middleware) ListPipelineJobs(org, project, env string) ([]*models.Job, error) {

	pipelineName := common.GetPipelineName(project, env)

	params := organization_pipelines_jobs.NewGetJobsParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetInpathPipelineName(pipelineName)

	resp, err := m.api.OrganizationPipelinesJobs.GetJobs(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()
	// TODO this validate have been removed https://github.com/cycloidio/youdeploy-http-api/issues/2262
	// err = p.Validate(strfmt.Default)
	// if err != nil {
	// 	return err
	// }

	d := p.Data
	return d, err
}

func (m *middleware) PausePipelineJob(org, project, env, job string) error {

	pipelineName := common.GetPipelineName(project, env)

	params := organization_pipelines_jobs.NewPauseJobParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetInpathPipelineName(pipelineName)
	params.SetJobName(job)

	_, err := m.api.OrganizationPipelinesJobs.PauseJob(params, common.ClientCredentials(&org))

	return err
}

func (m *middleware) UnpausePipelineJob(org, project, env, job string) error {

	pipelineName := common.GetPipelineName(project, env)

	params := organization_pipelines_jobs.NewUnpauseJobParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetInpathPipelineName(pipelineName)
	params.SetJobName(job)

	_, err := m.api.OrganizationPipelinesJobs.UnpauseJob(params, common.ClientCredentials(&org))

	return err
}

func (m *middleware) TriggerPipelineBuild(org, project, env, job string) error {

	pipelineName := common.GetPipelineName(project, env)

	params := organization_pipelines_jobs_build.NewCreateBuildParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetInpathPipelineName(pipelineName)
	params.SetJobName(job)

	_, err := m.api.OrganizationPipelinesJobsBuild.CreateBuild(params, common.ClientCredentials(&org))

	return err
}
