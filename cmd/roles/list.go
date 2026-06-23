package roles

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

var roleTableOptions = printer.Options{
	Columns:    []string{"Canonical", "Name", "Description", "Default"},
	Identifier: "Canonical",
}

func NewListCommand() *cobra.Command {
	var (
		example = `cy --org my-org roles list`
		short   = "Get the list of the current organization roles"
		long    = short
	)

	cmd := &cobra.Command{
		Use:     "list",
		Args:    cobra.NoArgs,
		Example: example,
		Short:   short,
		Long:    long,
		RunE:    listRoles,
	}
	cyout.RegisterModel(cmd, models.Role{})
	return cmd
}

func listRoles(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	roles, _, err := m.ListRoles(org)
	return cyout.PrintWithOptions(cmd, roles, err, "unable to list roles", roleTableOptions)
}
