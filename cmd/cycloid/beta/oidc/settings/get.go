package settings

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

var settingsTableOptions = printer.Options{
	Columns: []string{"DefaultRoleCanonical", "OIDCManaged", "OIDCNoMatchPolicy"},
}

// NewGetCommand returns the `settings get` command.
func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Args:  cobra.NoArgs,
		Short: "Get OIDC organization settings",
		Example: `
  # Get OIDC settings for my-org
  cy --org my-org beta oidc settings get
`,
		RunE: getSettings,
	}

	cyout.RegisterModel(cmd, middleware.OIDCOrganizationSettings{})
	return cmd
}

func getSettings(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	result, _, err := m.GetOIDCOrganizationSettings(org)
	return cyout.PrintWithOptions(cmd, result, err, "unable to get OIDC organization settings", settingsTableOptions)
}
