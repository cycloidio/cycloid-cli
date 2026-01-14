package pipelines

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewBuildCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "trigger",
		Aliases: []string{
			"create",
			"run",
		},
		Short:   "trigger a pipeline build",
		Example: `cy pipeline build trigger --project my-project --env my-env --component my-component --job my-job`,
		RunE:    createBuild,
		Args:    cobra.NoArgs,
	}

	cyargs.AddCyContext(cmd)
	cyargs.AddPipeline(cmd)
	cyargs.AddPipelineJob(cmd)
	return cmd
}

func createBuild(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, project, env, component, err := cyargs.GetCyContext(cmd)
	if err != nil {
		return err
	}

	pipeline, err := cyargs.GetPipeline(cmd)
	if err != nil {
		return err
	}

	job, err := cyargs.GetPipelineJob(cmd)
	if err != nil {
		return err
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	build, err := m.CreateBuild(org, project, env, component, pipeline, job)
	if err != nil {
		return fmt.Errorf("failed to trigger build in context project %q, env %q, component %q with pipeline %q in job %q: %w", project, env, component, pipeline, job, err)
	}

	return printer.SmartPrint(p, build, nil, "", printer.Options{}, cmd.OutOrStdout())
}
