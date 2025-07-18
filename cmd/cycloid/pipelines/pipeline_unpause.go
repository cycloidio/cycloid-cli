package pipelines

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewPipelineUnpauseCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "unpause",
		Short:   "unpause a pipeline",
		Example: `cy --org my-org pipeline unpause --project my-project --env env --component component --pipeline pipeline-name`,
		RunE:    unpause,
		Args:    cobra.NoArgs,
	}

	cyargs.AddCyContext(cmd)
	cyargs.AddPipeline(cmd)
	return cmd
}

func unpause(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cyargs.GetCyContext(cmd)
	if err != nil {
		return err
	}

	pipeline, err := cyargs.GetPipeline(cmd)
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

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	err = m.UnpausePipeline(org, project, env, component, pipeline)
	if err != nil {
		printer.SmartPrint(p, nil, err, "failed to unpause pipeline", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, nil, nil, "", printer.Options{}, cmd.OutOrStdout())
}
