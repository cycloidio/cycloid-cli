package projects

import (
	"fmt"
	"io/ioutil"

	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/spf13/cobra"
)

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  create,
	}
	common.RequiredFlag(WithFlagName, cmd)
	common.RequiredFlag(WithFlagStackRef, cmd)
	common.RequiredFlag(WithFlagConfigRepository, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredFlag(WithFlagPipeline, cmd)
	common.RequiredFlag(WithFlagVars, cmd)

	WithFlagUsecase(cmd)
	WithFlagDescription(cmd)
	WithFlagCanonical(cmd)
	WithFlagCloudProvider(cmd)
	WithFlagConfig(cmd)
	return cmd
}

func create(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}
	canonical, err := cmd.Flags().GetString("canonical")
	if err != nil {
		return err
	}
	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return err
	}
	cloudProvider, err := cmd.Flags().GetString("cloud-provider")
	if err != nil {
		return err
	}
	stackRef, err := cmd.Flags().GetString("stack-ref")
	if err != nil {
		return err
	}
	configRepo, err := cmd.Flags().GetUint32("config-repo")
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

	rawPipeline, err := ioutil.ReadFile(pipelinePath)
	if err != nil {
		return fmt.Errorf("Pipeline file reading error : %s", err.Error())
	}
	pipelineTemplate := string(rawPipeline)

	rawVars, err := ioutil.ReadFile(varsPath)
	if err != nil {
		return fmt.Errorf("Pipeline variables file reading error : %s", err.Error())
	}
	vars := string(rawVars)

	project, err := m.CreateProject(org, name, canonical, env, pipelineTemplate, vars, description, cloudProvider, stackRef, usecase, configRepo)
	if err != nil {
		return err
	}

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

		err = m.PushConfig(org, *project.Canonical, env, cfs)
		if err != nil {
			return err
		}
	}

	fmt.Println(project)

	return nil
}

// /organizations/{organization_canonical}/projects
// post: createProject
// Create a new project with envs and pipelines in the organization.
