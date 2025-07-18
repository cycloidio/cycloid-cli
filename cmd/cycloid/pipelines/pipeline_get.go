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

func NewPipelineGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "get",
		Short:   "get a pipeline",
		Example: `cy pipeline get --project my-project --env env --component component --pipeline pipeline_name`,
		RunE:    getPipeline,
		Args:    cobra.NoArgs,
	}

	cyargs.AddCyContext(cmd)
	cyargs.AddPipeline(cmd)
	return cmd
}

func getPipeline(cmd *cobra.Command, args []string) error {
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

	outPipeline, err := m.GetPipeline(org, project, env, component, pipeline)
	if err != nil {
		return fmt.Errorf("failed to get pipeline '%s' in context project '%s', env '%s', component '%s': %s", pipeline, project, env, component, err)
	}

	return printer.SmartPrint(p, outPipeline, nil, "", printer.Options{}, cmd.OutOrStdout())
}
