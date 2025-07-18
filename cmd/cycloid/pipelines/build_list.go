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

func NewBuildListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "list",
		Short:   "list builds of a pipeline",
		Example: `cy --org my-org pp list-builds --project my-project --env env --component component --job my-job`,
		RunE:    listBuilds,
		Args:    cobra.NoArgs,
	}

	cyargs.AddCyContext(cmd)
	cyargs.AddPipeline(cmd)
	cyargs.AddPipelineJob(cmd)
	return cmd
}

func listBuilds(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	pipeline, err := cyargs.GetPipeline(cmd)
	if err != nil {
		return err
	}

	// Those args are optional filters, we don't care about err
	_, project, env, component, _ := cyargs.GetCyContext(cmd)
	job, _ := cyargs.GetPipelineJob(cmd)

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	builds, err := m.GetBuilds(org, project, env, component, pipeline, job)
	if err != nil {
		return printer.SmartPrint(p, nil, err,
			fmt.Sprintf("failed to get builds of pipeline '%s', in project '%s', in env '%s', in component '%s'.",
				pipeline, project, env, component),
			printer.Options{}, cmd.OutOrStderr(),
		)
	}

	return printer.SmartPrint(p, builds, nil, "", printer.Options{}, cmd.OutOrStdout())
}
