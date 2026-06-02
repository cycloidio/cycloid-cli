package mappings

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

var mappingTableOptions = printer.Options{
	Columns:    []string{"ID", "GroupName", "Team.Canonical", "Team.Name"},
	Identifier: "ID",
}

// NewListCommand returns the `mappings list` command.
func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Args:  cobra.NoArgs,
		Short: "List OIDC group-to-team mappings",
		Example: `
  # List all OIDC group mappings for my-org
  cy --org my-org beta oidc mappings list
`,
		RunE: listMappings,
	}

	cyout.RegisterModel(cmd, middleware.OIDCGroupMapping{})
	return cmd
}

func listMappings(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	result, _, err := m.ListOIDCGroupMappings(org)
	return cyout.PrintWithOptions(cmd, result, err, "unable to list OIDC group mappings", mappingTableOptions)
}
