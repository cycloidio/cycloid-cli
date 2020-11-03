package projects

import (
	"fmt"
	"io/ioutil"
	"os"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewCreateEnvCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create-env",
		Short: "create an environment within a project",
		Example: `
	# create 'prod' environment in 'my-project'
	cy --org my-org project create-env \
		--project my-project \
		--env prod \
		--usecase usecase-1 \
		--pipeline /my/pipeline.yml \
		--vars /my/pipeline/vars.yml \
		--config /path/to/config=/path/in/config_repo
`,
		PreRunE: internal.CheckAPIAndCLIVersion,
		RunE:    createEnv,
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
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	if common.IsInList(env, projectData.Environments) {
		return fmt.Errorf("environment: %s already exists for this project: %s", env, project)
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

	if err != nil {
		return errors.Wrap(err, "unable to update project")
	}

	//
	// CREATE PIPELINE
	//

	rawPipeline, err := ioutil.ReadFile(pipelinePath)
	if err != nil {
		return errors.Wrap(err, "unable to read pipeline file")
	}
	pipelineTemplate := string(rawPipeline)

	rawVars, err := ioutil.ReadFile(varsPath)
	if err != nil {
		return errors.Wrap(err, "unable to read variables file")
	}

	variables := string(rawVars)

	if _, err := m.CreatePipeline(org, project, env, pipelineTemplate, variables, usecase); err != nil {
		return errors.Wrap(err, "unable to create pipeline")
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
				return errors.Wrap(err, "unable to read config file")
			}
			cfs[dest] = c
		}

		if err := m.PushConfig(org, project, env, cfs); err != nil {
			return errors.Wrap(err, "unable to push config")
		}
	}

	//
	// PIPELINE UNPAUSE
	//
	if err := m.UnpausePipeline(org, project, env); err != nil {
		return errors.Wrap(err, "unable to unpause pipeline")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// print the result on the standard output
	if err := p.Print(resp, printer.Options{}, os.Stdout); err != nil {
		return errors.Wrap(err, "unable to print result")
	}

	return nil
}
