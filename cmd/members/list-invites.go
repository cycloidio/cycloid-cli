package members

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
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

	cmd := &cobra.Command{
		Use:     "list-invites",
		Args:    cobra.NoArgs,
		Example: example,
		Short:   short,
		Long:    long,
		RunE:    listInvites,
	}

	return cmd
}

func listInvites(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	mbs, _, err := m.ListInvites(org)
	return cyout.PrintWithOptions(cmd, mbs, err, "unable to list invites", printer.Options{})
}
