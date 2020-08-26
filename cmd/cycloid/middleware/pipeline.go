package middleware

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_pipelines"
	"github.com/cycloidio/youdeploy-cli/client/models"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	strfmt "github.com/go-openapi/strfmt"
)

func (m *middleware) PausePipeline(org string, project string, env string) error {

	pipelineName := fmt.Sprintf("%s-%s", project, env)

	params := organization_pipelines.NewPausePipelineParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetInpathPipelineName(pipelineName)

	_, err := m.api.OrganizationPipelines.PausePipeline(params, root.ClientCredentials())

	return err
}

func (m *middleware) UnpausePipeline(org string, project string, env string) error {

	pipelineName := fmt.Sprintf("%s-%s", project, env)

	params := organization_pipelines.NewUnpausePipelineParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetInpathPipelineName(pipelineName)

	_, err := m.api.OrganizationPipelines.UnpausePipeline(params, root.ClientCredentials())
	// if err != nil {
	// 	return nil, err
	// }

	return err
}

func (m *middleware) UpdatePipeline(org string, project string, env string, pipeline string, variables string) (*models.Pipeline, error) {

	var body *models.UpdatePipeline

	cyCtx := common.CycloidContext{Env: env,
		Org:     org,
		Project: project}

	pipelineName := fmt.Sprintf("%s-%s", project, env)

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

	resp, err := m.api.OrganizationPipelines.UpdatePipeline(params, root.ClientCredentials())
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
