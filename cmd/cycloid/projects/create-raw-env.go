package projects

import (
	"fmt"
	"os"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/pipelines"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewCreateRawEnvCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "create-raw-env",
		Short:   "create an environment within a project with the old configuration format.",
		Aliases: []string{"create-env"},
		Example: `
	# create 'prod' environment in 'my-project'
	cy --org my-org project create-raw-env \
		--project my-project \
		--env prod \
		--usecase usecase-1 \
		--pipeline /my/pipeline.yml \
		--vars /my/pipeline/vars.yml \
		--config /path/to/config=/path/in/config_repo
`,
		Hidden: false,
		// TODO: Bring back deprecation when switch will be operated
		// If we depreciate it, it will disappear from the help
		//Deprecated: "This command will be changed to use stackforms in the future.\n" +
		//	"If your still want to use this command as is, use `cy project create-env` instead.\n" +
		//	"Please see https://github.com/cycloidio/cycloid-cli/issues/268 for more information.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			internal.Warning(cmd.ErrOrStderr(),
				"This command will be changed to use stackforms in the future.\n"+
					"If your still want to use this command as is, use `cy project create-raw-env` instead.\n"+
					"Please see https://github.com/cycloidio/cycloid-cli/issues/268 for more information.\n")
			return internal.CheckAPIAndCLIVersion(cmd, args)
		},
		RunE: createRawEnv,
	}
	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredFlag(WithFlagPipeline, cmd)
	common.RequiredFlag(WithFlagVars, cmd)
	WithFlagConfig(cmd)
	WithFlagUsecase(cmd)

	return cmd
}

func createRawEnv(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	var err error

	org, err := common.GetOrg(cmd)
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

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// need to conver the environment to "new environment" as required
	// by the API
	envs := make([]*models.NewEnvironment, len(projectData.Environments))

	for i, e := range projectData.Environments {
		if *e.Canonical == env {
			return fmt.Errorf("environment %s exists already in %s", env, project)
		}
		envs[i] = &models.NewEnvironment{
			Canonical: e.Canonical,
		}
	}

	// finally add the new environment
	envs = append(envs, &models.NewEnvironment{
		// TODO: https://github.com/cycloidio/cycloid-cli/issues/67
		Canonical: &env,
	})

	//
	// UPDATE PROJECT
	//
	resp, err := m.UpdateProject(org,
		*projectData.Name,
		project,
		envs,
		projectData.Description,
		*projectData.ServiceCatalog.Ref,
		*projectData.Owner.Username,
		projectData.ConfigRepositoryCanonical,
		nil,
		*projectData.UpdatedAt)

	err = printer.SmartPrint(p, nil, err, "unable to update project", printer.Options{}, cmd.OutOrStdout())
	if err != nil {
		return err
	}

	//
	// CREATE PIPELINE
	//

	rawPipeline, err := os.ReadFile(pipelinePath)
	if err != nil {
		return errors.Wrap(err, "unable to read pipeline file")
	}
	pipelineTemplate := string(rawPipeline)

	rawVars, err := os.ReadFile(varsPath)
	if err != nil {
		return errors.Wrap(err, "unable to read variables file")
	}

	variables := string(rawVars)

	newPP, err := m.CreatePipeline(org, project, env, pipelineTemplate, variables, usecase)
	err = printer.SmartPrint(p, nil, err, "unable to create pipeline", printer.Options{}, cmd.OutOrStdout())
	if err != nil {
		return err
	}

	//
	// PUSH CONFIG If project creation succeeded we push the config files
	//
	// Pipeline vars file
	crVarsPath, err := pipelines.GetPipelineVarsPath(m, org, project, *newPP.UseCase)
	if err != nil {
		printer.SmartPrint(p, nil, err, "unable to get pipeline variables destination path", printer.Options{}, cmd.OutOrStdout())
	}
	cfs := make(map[string]strfmt.Base64)
	cfs[crVarsPath] = rawVars

	// Additionals config files
	if len(configs) > 0 {

		for fp, dest := range configs {
			var c strfmt.Base64
			c, err = os.ReadFile(fp)
			if err != nil {
				return errors.Wrap(err, "unable to read config file")
			}
			cfs[dest] = c
		}
	}

	err = m.PushConfig(org, project, env, cfs)
	err = printer.SmartPrint(p, nil, err, "unable to push config", printer.Options{}, cmd.OutOrStdout())
	if err != nil {
		return err
	}

	//
	// PIPELINE UNPAUSE
	//
	err = m.UnpausePipeline(org, project, env)
	err = printer.SmartPrint(p, nil, err, "unable to unpause pipeline", printer.Options{}, cmd.OutOrStdout())
	if err != nil {
		return err
	}

	return printer.SmartPrint(p, resp, err, "", printer.Options{}, cmd.OutOrStdout())
}
