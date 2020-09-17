package projects

import (
	"fmt"
	"io/ioutil"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
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
	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredFlag(WithFlagPipeline, cmd)
	common.RequiredFlag(WithFlagVars, cmd)
	WithFlagConfig(cmd)
	WithFlagUsecase(cmd)

	return cmd
}

func createEnv(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	var err error
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

	projectData, err := m.GetProject(org, project)
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
	resp, err := m.UpdateProject(org,
		*projectData.Name,
		project,
		envs,
		projectData.Description,
		*projectData.CloudProvider.Canonical,
		*projectData.ServiceCatalogRef,
		*projectData.Owner.Username,
		projectData.ConfigRepositoryID)

	// TODO create a error handeling function to format our error with a better display
	if err != nil {
		return err
	}

	//
	// CREATE PIPELINE
	//

	rawPipeline, err := ioutil.ReadFile(pipelinePath)
	if err != nil {
		return fmt.Errorf("Pipeline file reading error : %s", err.Error())
	}
	pipelineTemplate := string(rawPipeline)

	rawVars, err := ioutil.ReadFile(varsPath)
	if err != nil {
		return fmt.Errorf("Pipeline variables file reading error : %s", err.Error())
	}
	variables := string(rawVars)

	_, err = m.CreatePipeline(org, project, env, pipelineTemplate, variables, usecase)
	if err != nil {
		return err
	}

	//
	// PUSH CONFIG If project creation succeeded we push the config files
	//
	if len(configs) > 0 {

		cfs := make(map[string]strfmt.Base64)

		for fp, dest := range configs {
			var c strfmt.Base64
			c, err = ioutil.ReadFile(fp)
			if err != nil {
				return fmt.Errorf("Config file reading error : %s", err.Error())
			}
			cfs[dest] = c
		}

		err = m.PushConfig(org, project, env, cfs)
		if err != nil {
			return err
		}

	}

	//
	// PIPELINE UNPAUSE
	//
	err = m.UnpausePipeline(org, project, env)
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}
