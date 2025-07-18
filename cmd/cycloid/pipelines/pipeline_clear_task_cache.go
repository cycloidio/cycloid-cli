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

func NewPipelineClearTaskCacheCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "clear-task-cache",
		Short:   "clear cache for a task",
		Example: `cy pp clear-task-cache --project my-project --job my-job --env my-env --step my-task`,
		RunE:    cleartaskCache,
		Args:    cobra.NoArgs,
	}

	cyargs.AddCyContext(cmd)
	cyargs.AddPipeline(cmd)
	cyargs.AddPipelineJob(cmd)
	cyargs.AddPipelineStep(cmd)
	return cmd
}

func cleartaskCache(cmd *cobra.Command, args []string) error {
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

	step, err := cyargs.GetPipelineStep(cmd)
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

	out, err := m.ClearTaskCache(org, project, env, component, pipeline, job, step)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to clear task cache", printer.Options{}, cmd.OutOrStdout())
	}

	return printer.SmartPrint(p, out, nil, "", printer.Options{}, cmd.OutOrStdout())
}
