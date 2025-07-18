package pipelines

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewPipelineSyncedCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "synced",
		Short:   "Check if a pipeline is synced with its stacks.",
		Example: `cy --org my-org pipeline synced --project my-project --env env`,
		RunE:    synced,
		Args:    cobra.NoArgs,
	}

	cyargs.AddCyContext(cmd)
	cyargs.AddPipeline(cmd)
	return cmd
}

func synced(cmd *cobra.Command, args []string) error {
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
		return err
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	pp, err := m.SyncedPipeline(org, project, env, component, pipeline)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to pipeline sync status", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, pp, nil, "", printer.Options{}, cmd.OutOrStdout())
}
