package middleware

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_projects"
	"github.com/cycloidio/youdeploy-cli/client/models"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	strfmt "github.com/go-openapi/strfmt"
)

func (m *middleware) GetProject(org string, project string) (*models.Project, error) {

	params := organization_projects.NewGetProjectParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)

	resp, err := m.api.OrganizationProjects.GetProject(params, root.ClientCredentials())
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

func (m *middleware) CreateProject(org, projectName, projectCanonical, env, pipelineTemplate, variables, description, cloudProvider, stackRef, usecase string, configRepo uint32) (*models.Project, error) {

	var body *models.NewProject
	var pipelines []*models.NewPipeline

	if projectCanonical == "" {
		re := regexp.MustCompile(`[^a-z0-9\-_]`)
		projectCanonical = re.ReplaceAllString(strings.ToLower(projectName), "-")
	}

	cyCtx := common.CycloidContext{Env: env,
		Org:     org,
		Project: projectCanonical}

	params := organization_projects.NewCreateProjectParams()
	params.SetOrganizationCanonical(org)

	pipelineName := fmt.Sprintf("%s-%s", projectCanonical, env)

	vars := common.ReplaceCycloidVarsString(cyCtx, variables)

	pipeline := &models.NewPipeline{
		Environment:  &env,
		PipelineName: &pipelineName,
		UseCase:      usecase,
		PassedConfig: &pipelineTemplate,
		YamlVars:     vars,
	}
	err := pipeline.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	pipelines = append(pipelines, pipeline)

	body = &models.NewProject{
		Name:               &projectName,
		Description:        description,
		Canonical:          &projectCanonical,
		CloudProvider:      cloudProvider,
		ServiceCatalogRef:  &stackRef,
		ConfigRepositoryID: &configRepo,
		Pipelines:          pipelines,
	}

	err = body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	params.SetBody(body)
	resp, err := m.api.OrganizationProjects.CreateProject(params, root.ClientCredentials())
	// TODO create a error handeling function to format our error with a better display
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
