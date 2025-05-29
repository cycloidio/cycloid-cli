package middleware

// 	import (
// 		strfmt "github.com/go-openapi/strfmt"
//
// 		"github.com/cycloidio/cycloid-cli/client/client/organization_pipelines"
// 		"github.com/cycloidio/cycloid-cli/client/client/organization_pipelines_jobs_build"
// 		"github.com/cycloidio/cycloid-cli/client/models"
// 		"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
// 	)
//
// 	func (m *middleware) PausePipeline(org, project, env string) error {
// 		pipelineName := common.GetPipelineName(project, env)
//
// 		params := organization_pipelines.NewPausePipelineParams()
// 		params.SetOrganizationCanonical(org)
// 		params.SetProjectCanonical(project)
// 		params.SetInpathPipelineName(pipelineName)
//
// 		_, err := m.api.OrganizationPipelines.PausePipeline(params, m.api.Credentials(&org))
// 		if err != nil {
// 			return NewApiError(err)
// 		}
//
// 		return nil
// 	}
//
// 	func (m *middleware) UnpausePipeline(org, project, env string) error {
// 		pipelineName := common.GetPipelineName(project, env)
//
// 		params := organization_pipelines.NewUnpausePipelineParams()
// 		params.SetOrganizationCanonical(org)
// 		params.SetProjectCanonical(project)
// 		params.SetInpathPipelineName(pipelineName)
//
// 		_, err := m.api.OrganizationPipelines.UnpausePipeline(params, m.api.Credentials(&org))
// 		if err != nil {
// 			return NewApiError(err)
// 		}
//
// 		return nil
// 	}
//
// 	func (m *middleware) UpdatePipeline(org, project, env, pipeline, variables string) (*models.Pipeline, error) {
// 		var body *models.UpdatePipeline
//
// 		cyCtx := common.CycloidContext{Env: env,
// 		Org:     org,
// 		Project: project}
//
// 		pipelineName := common.GetPipelineName(project, env)
//
//
// 	params := organization_pipelines.NewUpdatePipelineParams()
// 	params.SetOrganizationCanonical(org)
// 	params.SetProjectCanonical(project)
// 	params.SetInpathPipelineName(pipelineName)
//
// 	body = &models.UpdatePipeline{
// 		PassedConfig: &pipeline,
// 		YamlVars:     vars,
// 	}
// 	err := body.Validate(strfmt.Default)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	params.SetBody(body)
//
// 	resp, err := m.api.OrganizationPipelines.UpdatePipeline(params, m.api.Credentials(&org))
// 	if err != nil {
// 		return nil, NewApiError(err)
// 	}
//
// 	p := resp.GetPayload()
//
// 	d := p.Data
// 	return d, nil
// }
//
// func (m *middleware) DiffPipeline(org, project, env, pipeline, variables string) (*models.PipelineDiffs, error) {
// 	cyCtx := common.CycloidContext{Env: env,
// 		Org:     org,
// 		Project: project}
// 	vars := common.ReplaceCycloidVarsString(cyCtx, variables)
// 	pipelineName := common.GetPipelineName(project, env)
//
// 	params := organization_pipelines.NewDiffPipelineParams()
// 	params.SetOrganizationCanonical(org)
// 	params.SetInpathPipelineName(pipelineName)
//
// 	body := &models.UpdatePipeline{
// 		PassedConfig: &pipeline,
// 		YamlVars:     vars,
// 	}
//
// 	params.SetBody(body)
// 	err := body.Validate(strfmt.Default)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	resp, err := m.api.OrganizationPipelines.DiffPipeline(params, m.api.Credentials(&org))
// 	if err != nil {
// 		return nil, NewApiError(err)
// 	}
//
// 	p := resp.GetPayload()
// 	err = p.Validate(strfmt.Default)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	d := p.Data
// 	return d, err
// }
//
// func (m *middleware) CreatePipeline(org, project, env, pipeline, variables, usecase string) (*models.Pipeline, error) {
// 	cyCtx := common.CycloidContext{Env: env,
// 		Org:     org,
// 		Project: project}
//
// 	params := organization_pipelines.NewCreatePipelineParams()
// 	params.SetOrganizationCanonical(org)
// 	params.SetProjectCanonical(project)
//
// 	vars := common.ReplaceCycloidVarsString(cyCtx, variables)
//
// 	pipelineName := common.GetPipelineName(project, env)
//
// 	body := &models.NewPipeline{
// 		Environment: &models.NewEnvironment{
// 			// TODO: https://github.com/cycloidio/cycloid-cli/issues/67
// 			Canonical: &env,
// 		},
// 		PipelineName: &pipelineName,
// 		UseCase:      usecase,
// 		PassedConfig: pipeline,
// 		YamlVars:     vars,
// 	}
// 	err := body.Validate(strfmt.Default)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	params.SetBody(body)
// 	resp, err := m.api.OrganizationPipelines.CreatePipeline(params, m.api.Credentials(&org))
// 	if err != nil {
// 		return nil, NewApiError(err)
// 	}
//
// 	p := resp.GetPayload()
// 	err = p.Validate(strfmt.Default)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	d := p.Data
// 	return d, nil
// }
//
// func (m *middleware) ListPipelineJobsBuilds(org, project, env, job string) ([]*models.Build, error) {
// 	pipelineName := common.GetPipelineName(project, env)
//
// 	params := organization_pipelines_jobs_build.NewGetBuildsParams()
// 	params.SetOrganizationCanonical(org)
// 	params.SetProjectCanonical(project)
// 	params.SetInpathPipelineName(pipelineName)
// 	params.SetJobName(job)
//
// 	resp, err := m.api.OrganizationPipelinesJobsBuild.GetBuilds(params, m.api.Credentials(&org))
// 	if err != nil {
// 		return nil, NewApiError(err)
// 	}
//
// 	p := resp.GetPayload()
// 	//err = p.Validate(strfmt.Default)
// 	//if err != nil {
// 	//	return nil, err
// 	//}
//
// 	d := p.Data
// 	return d, nil
// }
//
// func (m *middleware) TriggerPipelineBuild(org, project, env, job string) error {
//
// 	pipelineName := common.GetPipelineName(project, env)
//
// 	params := organization_pipelines_jobs_build.NewCreateBuildParams()
// 	params.SetOrganizationCanonical(org)
// 	params.SetProjectCanonical(project)
// 	params.SetInpathPipelineName(pipelineName)
// 	params.SetJobName(job)
//
// 	_, err := m.api.OrganizationPipelinesJobsBuild.CreateBuild(params, m.api.Credentials(&org))
// 	if err != nil {
// 		return NewApiError(err)
// 	}
//
// 	return nil
// }
//
// func (m *middleware) ListPipelines(org string) ([]*models.Pipeline, error) {
//
// 	params := organization_pipelines.NewGetPipelinesParams()
// 	params.SetOrganizationCanonical(org)
//
// 	resp, err := m.api.OrganizationPipelines.GetPipelines(params, m.api.Credentials(&org))
// 	if err != nil {
// 		return nil, NewApiError(err)
// 	}
//
// 	p := resp.GetPayload()
// 	// err = p.Validate(strfmt.Default)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
//
// 	d := p.Data
// 	return d, nil
// }
//
// func (m *middleware) SyncedPipeline(org, project, env string) (*models.PipelineStatus, error) {
//
// 	pipelineName := common.GetPipelineName(project, env)
//
// 	params := organization_pipelines.NewSyncedPipelineParams()
// 	params.SetOrganizationCanonical(org)
// 	params.SetInpathPipelineName(pipelineName)
//
// 	resp, err := m.api.OrganizationPipelines.SyncedPipeline(params, m.api.Credentials(&org))
// 	if err != nil {
// 		return nil, NewApiError(err)
// 	}
//
// 	p := resp.GetPayload()
// 	err = p.Validate(strfmt.Default)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	d := p.Data
// 	// In case of nil data, add an empty PipelineStatus model produce an homogen output
// 	if d == nil {
// 		d = &models.PipelineStatus{}
// 	}
// 	return d, nil
// }
