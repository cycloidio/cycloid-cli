package projects

import (
	"fmt"
	"io/ioutil"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_config_repositories"
	"github.com/cycloidio/youdeploy-cli/client/client/organization_pipelines"
	"github.com/cycloidio/youdeploy-cli/client/client/organization_projects"
	"github.com/cycloidio/youdeploy-cli/client/models"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/pipelines"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/spf13/cobra"
)

func NewCreateEnvCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create-env",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  createEnv,
	}
	// common.RequiredFlag(WithFlagStackRef, cmd)
	// common.RequiredFlag(WithFlagConfigRepository, cmd)
	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredFlag(WithFlagPipeline, cmd)
	common.RequiredFlag(WithFlagVars, cmd)
	WithFlagConfig(cmd)
	WithFlagUsecase(cmd)

	// common.RequiredFlag(WithFlagPipeline, cmd)
	// common.RequiredFlag(WithFlagVars, cmd)
	//
	// WithFlagUsecase(cmd)
	// WithFlagDescription(cmd)
	// WithFlagCanonical(cmd)
	// WithFlagCloudProvider(cmd)
	// WithFlagConfig(cmd)

	// TODO
	// Request URL: https://http-api-staging.cycloid.io/organizations/seraf/projects/gael/pipelines
	// {pipeline_name: "gael-test2",…}
	// environment: "test2"
	// passed_config: "{"jobs":[{"name":"job-hello-world","build_logs_to_retain":3,"plan":[{"task":"hello-world","config":{"platform":"linux","image_resource":{"type":"docker-image","source":{"repository":"busybox"}},"run":{"path":"/bin/sh","args":["-ec","echo ${MESSAGE}\n"]},"params":{"MESSAGE":"((message))"}}}]}]}"
	// pipeline_name: "gael-test2"
	// use_case: "default"
	// yaml_vars: "#↵# This file has been generated by Cycloid, please DO NOT modify.↵# Any manual modifications done to this file, WILL be lost on the↵# next project edition via the forms.↵#↵# Please note that comments in sample files will have been dropped↵# due to some limitations upon files' generation.↵#↵# Any extra variables not found in the original file have been ignored.↵#↵message: hello world and especially to seraf↵"
	//
	// update project
	// {cloud_provider: "google", config_repository_id: 307, created_at: 1597852668,…}
	// cloud_provider: "google"
	// config_repository_id: 307
	// created_at: 1597852668
	// environments: ["test", "test2"]
	// 0: "test"
	// 1: "test2"
	// id: 1280
	// name: "Gael"
	// owner: "cycloid_seraf"
	// service_catalog_ref: "seraf:stack-dummy"
	// updated_at: -62135596800
	//
	// + call to config

	return cmd
}

func createEnv(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()
	//
	var err error
	var body *models.UpdateProject
	// var pipelines []*models.NewPipeline

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	project, err := cmd.Flags().GetString("project")
	if err != nil {
		return err
	}
	env, err := cmd.Flags().GetString("env")
	if err != nil {
		return err
	}
	usecase, err := cmd.Flags().GetString("usecase")
	if err != nil {
		return err
	}
	varsPath, err := cmd.Flags().GetString("vars")
	if err != nil {
		return err
	}
	pipelinePath, err := cmd.Flags().GetString("pipeline")
	if err != nil {
		return err
	}
	configs, err := cmd.Flags().GetStringToString("config")
	if err != nil {
		return err
	}

	cyCtx := common.CycloidContext{Env: env,
		Org:     org,
		Project: project}

	projectData, err := Get(api, org, project)
	if err != nil {
		return err
	}

	if common.IsInList(env, projectData.Environments) {
		return fmt.Errorf("Environment %s already exist for this project %s", env, project)
	}
	envs := append(projectData.Environments, env)
	//
	// UPDATE PROJECT
	//
	params := organization_projects.NewUpdateProjectParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)

	body = &models.UpdateProject{
		Name:               projectData.Name,
		Description:        projectData.Description,
		CloudProvider:      *projectData.CloudProvider.Canonical,
		ServiceCatalogRef:  projectData.ServiceCatalogRef,
		ConfigRepositoryID: projectData.ConfigRepositoryID,
		Environments:       envs,
		Owner:              projectData.Owner.Username,
	}

	err = body.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	params.SetBody(body)
	resp, err := api.OrganizationProjects.UpdateProject(params, root.ClientCredentials())
	// TODO create a error handeling function to format our error with a better display
	if err != nil {
		return err
	}

	//
	// CREATE PIPELINE
	//

	paramsP := organization_pipelines.NewCreatePipelineParams()
	paramsP.SetOrganizationCanonical(org)
	paramsP.SetProjectCanonical(project)

	rawPipeline, err := ioutil.ReadFile(pipelinePath)
	if err != nil {
		return fmt.Errorf("Pipeline file reading error : %s", err.Error())
	}
	pipelineTemplate := string(rawPipeline)

	rawVars, err := ioutil.ReadFile(varsPath)
	if err != nil {
		return fmt.Errorf("Pipeline variables file reading error : %s", err.Error())
	}
	vars := common.ReplaceCycloidVarsString(cyCtx, string(rawVars))

	pipelineName := fmt.Sprintf("%s-%s", project, env)

	bodyP := &models.NewPipeline{
		Environment:  &env,
		PipelineName: &pipelineName,
		UseCase:      usecase,
		PassedConfig: &pipelineTemplate,
		YamlVars:     vars,
	}
	err = bodyP.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	paramsP.SetBody(bodyP)
	_, err = api.OrganizationPipelines.CreatePipeline(paramsP, root.ClientCredentials())
	if err != nil {
		return err
	}

	//
	// PUSH CONFIG If project creation succeeded we push the config files
	//

	if len(configs) > 0 {
		paramsC := organization_config_repositories.NewCreateConfigRepositoryConfigParams()
		paramsC.SetOrganizationCanonical(org)
		paramsC.SetConfigRepositoryID(projectData.ConfigRepositoryID)

		var cfs []*models.ConfigFile

		for fp, dest := range configs {
			var c strfmt.Base64
			p := common.ReplaceCycloidVarsString(cyCtx, dest)
			c, err = ioutil.ReadFile(fp)
			if err != nil {
				return fmt.Errorf("Config file reading error : %s", err.Error())
			}
			c = common.ReplaceCycloidVars(cyCtx, c)

			cf := &models.ConfigFile{
				Content: &c,
				Path:    &p,
			}
			err = cf.Validate(strfmt.Default)
			if err != nil {
				return err
			}

			cfs = append(cfs, cf)
		}

		bodyC := &models.SCConfig{Configs: cfs}

		err = bodyC.Validate(strfmt.Default)
		if err != nil {
			return err
		}

		paramsC.SetBody(bodyC)
		_, err = api.OrganizationConfigRepositories.CreateConfigRepositoryConfig(paramsC, root.ClientCredentials())
		if err != nil {
			return err
		}

	}

	//
	// PIPELINE UNPAUSE
	//
	err = pipelines.Unpause(api, org, project, env)
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}

// /organizations/{organization_canonical}/projects
// post: createProject
// Create a new project with envs and pipelines in the organization.
