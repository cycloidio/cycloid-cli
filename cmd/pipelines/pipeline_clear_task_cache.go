package pipelines

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewPipelineClearTaskCacheCommand() *cobra.Command {
	cmd := &cobra.Command{
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

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	out, _, err := m.ClearTaskCache(org, project, env, component, pipeline, job, step)
	return cyout.PrintWithOptions(cmd, out, err, "unable to clear task cache", printer.Options{})
}
