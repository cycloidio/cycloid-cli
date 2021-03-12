package members

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewListInvitesCommand() *cobra.Command {
	var (
		example = `
	# List all invites within my-org organization
	cy --org my-org members list-invites
	`
		short = "Get the list of organization members invites"
		long  = short
	)

	var cmd = &cobra.Command{
		Use:     "list-invites",
		Example: example,
		Short:   short,
		Long:    long,
		RunE:    listMembers,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	return cmd
}

func listInvites(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	mbs, err := m.ListInvites(org)
	return printer.SmartPrint(p, mbs, err, "unable to list invites", printer.Options{}, os.Stdout)
}
