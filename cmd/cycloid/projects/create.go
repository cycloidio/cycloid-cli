package projects

import (
	"io/ioutil"
	"os"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/youdeploy-cli/printer"
	"github.com/cycloidio/youdeploy-cli/printer/factory"
)

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "create a project",
		Example: `
	# create a project
	cy --org my-org project create \
		--name my-project \
		--canonical my-project \
		--description "an awesome project" \
		--cloud-provider gcp|aws|... \
		--stack-ref my-stack-ref \
		--config-repo config-repo-id \
		--env environment-name \
		--usecase usecase-1 \
		--vars /path/to/variables.yml \
		--pipeline /path/to/pipeline.yml \
		--config /path/to/config
`,
		RunE: create,
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
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
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
				return errors.Wrap(err, "unable to read config file")
			}
			cfs[dest] = c
		}

		if err := m.PushConfig(org, *project.Canonical, env, cfs); err != nil {
			return errors.Wrap(err, "unable to push config")
		}
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// print the result on the standard output
	if err := p.Print(project, printer.Options{}, os.Stdout); err != nil {
		return errors.Wrap(err, "unable to print result")
	}

	return nil
}
