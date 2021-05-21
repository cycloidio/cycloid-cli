package organizations

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewListWorkersCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "list-workers",
		Short:   "list the organization workers",
		RunE:    listWorkers,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}
	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	return cmd
}

func listWorkers(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
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

	ws, err := m.ListOrganizationWorkers(org)
	return printer.SmartPrint(p, ws, err, "unable to list organization workers", printer.Options{}, cmd.OutOrStdout())
}
