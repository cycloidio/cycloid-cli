package members

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

var memberTableOptions = printer.Options{
	Columns:    []string{"Username", "Email", "GivenName", "FamilyName"},
	Identifier: "Username",
}

func NewListCommand() *cobra.Command {
	var (
		example = `
	# List all the members within my-org organization
	cy --org my-org members list
	`
		short = "Get the list of organization members"
		long  = short
	)

	cmd := &cobra.Command{
		Use:     "list",
		Args:    cobra.NoArgs,
		Example: example,
		Short:   short,
		Long:    long,
		RunE:    listMembers,
	}
	cyout.RegisterModel(cmd, models.User{})
	return cmd
}

func listMembers(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	mbs, _, err := m.ListMembers(org)
	return cyout.PrintWithOptions(cmd, mbs, err, "unable to list members", memberTableOptions)
}
