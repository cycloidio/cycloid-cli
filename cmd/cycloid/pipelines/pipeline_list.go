package pipelines

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewPipelineListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "list",
		Short:   "list all pipelines of the organization",
		Args:    cobra.NoArgs,
		PreRunE: internal.CheckAPIAndCLIVersion,
		RunE:    list,
	}

	cyargs.AddProjectFlag(cmd)
	cyargs.AddEnvFlag(cmd)
	cyargs.AddPipelineStatuses(cmd)
	cyargs.AddPipeline(cmd)
	return cmd
}

func list(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	// Thoses args are filter and optionnal
	project, _ := cyargs.GetProject(cmd)
	env, _ := cyargs.GetEnv(cmd)
	statuses, _ := cyargs.GetPipelineStatuses(cmd)
	pipelineName, _ := cyargs.GetPipeline(cmd)

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

	pps, err := m.GetOrgPipelines(org, &pipelineName, &project, &env, statuses)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to list pipelines", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, pps, nil, "", printer.Options{}, cmd.OutOrStdout())
}
