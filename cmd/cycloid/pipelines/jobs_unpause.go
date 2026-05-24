package pipelines

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewJobsUnpauseCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "unpause",
		Short:   "unpause a pipeline job",
		RunE:    unpauseJob,
		Example: `cy pp job unpause --project my-project --env env --component component --pipeline pipeline --job my-job`,
		Args:    cobra.NoArgs,
	}

	cyargs.AddCyContext(cmd)
	cyargs.AddPipeline(cmd)
	cyargs.AddPipelineJob(cmd)
	return cmd
}

func unpauseJob(cmd *cobra.Command, args []string) error {
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

	_, err = m.UnPauseJob(org, project, env, component, pipeline, job)
	return cyout.PrintWithOptions(cmd, nil, err, "unable to unpause the job", printer.Options{})
}
