package pipelines

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewSyncedCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:    "synced",
		Hidden: true,
		Short:  "Check if a pipeline is up-to-date",
		Example: `
	# Check if a the running pipeline is synced with the template in the stack
	cy --org my-org pipeline synced --project my-project --env env

	# Parse the result in a script: | jq -r '[.[]] | all(. == null)'
`, PreRunE: internal.CheckAPIAndCLIVersion,
		RunE: synced,
	}

	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)

	return cmd
}

func synced(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	project, err := cmd.Flags().GetString("project")
	if err != nil {
		return err
	}
	env, err := cmd.Flags().GetString("env")
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

	pp, err := m.SyncedPipeline(org, project, env)
	return printer.SmartPrint(p, pp, err, "unable to get pipeline", printer.Options{}, cmd.OutOrStdout())
}
