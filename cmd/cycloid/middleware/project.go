package middleware

import (
	"regexp"
	"strings"

	"github.com/cycloidio/cycloid-cli/client/client/organization_projects"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	strfmt "github.com/go-openapi/strfmt"
)

func (m *middleware) ListProjects(org string) ([]*models.Project, error) {

	params := organization_projects.NewGetProjectsParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationProjects.GetProjects(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()
	err = p.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	d := p.Data

	return d, nil
}

func (m *middleware) GetProject(org, project string) (*models.Project, error) {

	params := organization_projects.NewGetProjectParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)

	resp, err := m.api.OrganizationProjects.GetProject(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()
	err = p.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	d := p.Data

	return d, nil
}

func (m *middleware) CreateProject(org, projectName, projectCanonical, env, pipelineTemplate, variables, description, stackRef, usecase, configRepo string) (*models.Project, error) {

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
		Environment: &models.NewEnvironment{
			// TODO: https://github.com/cycloidio/cycloid-cli/issues/67
			Canonical: &env,
		},
		PipelineName: &pipelineName,
		UseCase:      usecase,
		PassedConfig: pipelineTemplate,
		YamlVars:     vars,
	}
	err := pipeline.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	pipelines = append(pipelines, pipeline)

	body = &models.NewProject{
		Name:                      &projectName,
		Description:               description,
		Canonical:                 projectCanonical,
		ServiceCatalogRef:         &stackRef,
		ConfigRepositoryCanonical: &configRepo,
		Pipelines:                 pipelines,
	}

	err = body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	params.SetBody(body)
	resp, err := m.api.OrganizationProjects.CreateProject(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()
	err = p.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	d := p.Data
	return d, err
}

func (m *middleware) UpdateProject(org, projectName, projectCanonical string, envs []*models.NewEnvironment, description, stackRef, owner, configRepo string, updatedAt uint64) (*models.Project, error) {

	params := organization_projects.NewUpdateProjectParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(projectCanonical)

	body := &models.UpdateProject{
		Name:                      &projectName,
		Description:               description,
		ServiceCatalogRef:         &stackRef,
		ConfigRepositoryCanonical: configRepo,
		Environments:              envs,
		Owner:                     owner,
		UpdatedAt:                 &updatedAt,
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	params.SetBody(body)
	resp, err := m.api.OrganizationProjects.UpdateProject(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()
	err = p.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	d := p.Data
	return d, err
}

func (m *middleware) DeleteProjectEnv(org, project, env string) error {

	params := organization_projects.NewDeleteProjectEnvironmentParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)

	_, err := m.api.OrganizationProjects.DeleteProjectEnvironment(params, common.ClientCredentials(&org))
	if err != nil {
		return NewApiError(err)
	}
	return nil
}

func (m *middleware) DeleteProject(org, project string) error {

	params := organization_projects.NewDeleteProjectParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)

	_, err := m.api.OrganizationProjects.DeleteProject(params, common.ClientCredentials(&org))
	if err != nil {
		return NewApiError(err)
	}
	return nil
}
