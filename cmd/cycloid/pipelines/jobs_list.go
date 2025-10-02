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

func NewJobsListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "list",
		Short:   "list a pipeline's jobs",
		Example: `cy --org my-org pp list-jobs --project my-project --env env --component component -o json`,
		RunE:    listJobs,
		Args:    cobra.NoArgs,
	}

	cyargs.AddCyContext(cmd)
	cyargs.AddPipeline(cmd)
	return cmd
}

func listJobs(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cyargs.GetCyContext(cmd)
	if err != nil {
		return err
	}

	pipeline, err := cyargs.GetPipeline(cmd)
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

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	jobs, err := m.GetJobs(org, project, env, component, pipeline)
	if err != nil {
		return fmt.Errorf("failed to fetch jobs for pipeline %q: %w", pipeline, err)
	}

	return printer.SmartPrint(p, jobs, nil, "", printer.Options{}, cmd.OutOrStdout())
}
