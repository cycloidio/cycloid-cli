package middleware

import (
	"regexp"
	"strings"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_projects"
	"github.com/cycloidio/youdeploy-cli/client/models"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	strfmt "github.com/go-openapi/strfmt"
)

func (m *middleware) ListProjects(org string) ([]*models.ProjectsItem, error) {

	params := organization_projects.NewGetProjectsParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationProjects.GetProjects(params, common.ClientCredentials(&org))
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

func (m *middleware) GetProject(org, project string) (*models.Project, error) {

	params := organization_projects.NewGetProjectParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)

	resp, err := m.api.OrganizationProjects.GetProject(params, common.ClientCredentials(&org))
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

	pipelineName := common.GetPipelineName(projectCanonical, env)

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
	resp, err := m.api.OrganizationProjects.CreateProject(params, common.ClientCredentials(&org))
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

func (m *middleware) UpdateProject(org, projectName, projectCanonical string, envs []string, description, cloudProvider, stackRef, owner string, configRepo uint32) (*models.Project, error) {

	params := organization_projects.NewUpdateProjectParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(projectCanonical)

	body := &models.UpdateProject{
		Name:               &projectName,
		Description:        description,
		CloudProvider:      cloudProvider,
		ServiceCatalogRef:  &stackRef,
		ConfigRepositoryID: configRepo,
		Environments:       envs,
		Owner:              owner,
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	params.SetBody(body)
	resp, err := m.api.OrganizationProjects.UpdateProject(params, common.ClientCredentials(&org))
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

func (m *middleware) DeleteProjectEnv(org, project, env string) error {

	params := organization_projects.NewDeleteProjectEnvironmentParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)

	_, err := m.api.OrganizationProjects.DeleteProjectEnvironment(params, common.ClientCredentials(&org))

	return err
}

func (m *middleware) DeleteProject(org, project string) error {

	params := organization_projects.NewDeleteProjectParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)

	_, err := m.api.OrganizationProjects.DeleteProject(params, common.ClientCredentials(&org))

	return err
}
