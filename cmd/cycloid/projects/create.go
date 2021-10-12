package projects

import (
	"io/ioutil"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/pipelines"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "create a project",
		Example: `
	# create a project
	cy --org my-org project create \
		--name my-project \
		--description "an awesome project" \
		--stack-ref my-stack-ref \
		--config-repo config-repo-canonical \
		--env environment-name \
		--usecase usecase-1 \
		--vars /path/to/variables.yml \
		--pipeline /path/to/pipeline.yml \
		--config /path/to/config=/path/in/config_repo
`,
		RunE:    create,
		PreRunE: internal.CheckAPIAndCLIVersion,
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
	WithFlagConfig(cmd)
	return cmd
}

func create(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
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
	stackRef, err := cmd.Flags().GetString("stack-ref")
	if err != nil {
		return err
	}
	configRepo, err := cmd.Flags().GetString("config-repo")
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
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	rawPipeline, err := ioutil.ReadFile(pipelinePath)
	if err != nil {
		return errors.Wrap(err, "unable to read pipeline file")
	}
	pipelineTemplate := string(rawPipeline)

	rawVars, err := ioutil.ReadFile(varsPath)
	if err != nil {
		return errors.Wrap(err, "unable to read variables file")
	}
	vars := string(rawVars)

	project, err := m.CreateProject(org, name, canonical, env, pipelineTemplate, vars, description, stackRef, usecase, configRepo)
	err = printer.SmartPrint(p, nil, err, "unable to create project", printer.Options{}, cmd.OutOrStdout())
	if err != nil {
		return err
	}

	//
	// PUSH CONFIG If project creation succeeded we push the config files
	//
	// Pipeline vars file
	crVarsPath, err := pipelines.GetPipelineVarsPath(m, org, *project.Canonical, usecase)
	if err != nil {
		printer.SmartPrint(p, nil, err, "unable to get pipeline variables destination path", printer.Options{}, cmd.OutOrStdout())
	}
	cfs := make(map[string]strfmt.Base64)
	cfs[crVarsPath] = rawVars

	// Additionals config files
	if len(configs) > 0 {
		for fp, dest := range configs {
			var c strfmt.Base64
			c, err = ioutil.ReadFile(fp)
			if err != nil {
				return errors.Wrap(err, "unable to read config file")
			}
			cfs[dest] = c
		}
	}

	err = m.PushConfig(org, *project.Canonical, env, cfs)
	err = printer.SmartPrint(p, nil, err, "unable to push config", printer.Options{}, cmd.OutOrStdout())
	if err != nil {
		return err
	}

	return printer.SmartPrint(p, project, err, "", printer.Options{}, cmd.OutOrStdout())
}
