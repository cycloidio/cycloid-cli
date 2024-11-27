package projects

import (
	"fmt"
	"os"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/pipelines"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func createRawEnv(cmd *cobra.Command, org, project, env, useCase, varsPath, pipelinePath, output string, configs map[string]string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	var err error

	projectData, err := m.GetProject(org, project)
	if err != nil {
		return err
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
