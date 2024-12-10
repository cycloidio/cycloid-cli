package projects

import (
	"os"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/pipelines"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func createLegacyEnv(cmd *cobra.Command, org, project, env, useCase, varsPath, pipelinePath, output string, configs map[string]string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	err = m.CreateEnv(
		org,
		project,
		env,
		useCase,
		"",  // color
		"",  // icon // TODO add color and icon handling
		"",  // cloudProviderCanonical
		nil, // inputs
	)

	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to create env "+env, printer.Options{}, cmd.OutOrStdout())
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

	newPP, err := m.CreatePipeline(org, project, env, pipelineTemplate, variables, useCase)
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
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to push config", printer.Options{}, cmd.OutOrStdout())
	}

	//
	// PIPELINE UNPAUSE
	//
	err = m.UnpausePipeline(org, project, env)
	return printer.SmartPrint(p, nil, err, "unable to unpause pipeline", printer.Options{}, cmd.OutOrStdout())
}
