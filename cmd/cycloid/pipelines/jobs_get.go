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

func NewJobsGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "get",
		Short:   "get a pipeline's job",
		Example: `cy --org my-org pp get-job --project my-project --env env --component component --job my-job -o json`,
		RunE:    getJob,
		Args:    cobra.NoArgs,
	}

	cyargs.AddCyContext(cmd)
	cyargs.AddPipeline(cmd)
	cyargs.AddPipelineJob(cmd)
	return cmd
}

func getJob(cmd *cobra.Command, args []string) error {
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

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	outJob, err := m.GetJob(org, project, env, component, pipeline, job)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to fetch job: "+job, printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, outJob, nil, "", printer.Options{}, cmd.OutOrStdout())
}
