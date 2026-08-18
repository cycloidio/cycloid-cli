package settings

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
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
  cy --org my-org oidc settings get
`,
		RunE: getSettings,
	}

	cyout.RegisterModel(cmd, apiclient.OIDCOrganizationSettings{})
	return cmd
}

func getSettings(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	result, _, err := m.GetOIDCOrganizationSettings(org)
	return cyout.PrintWithOptions(cmd, result, err, "unable to get OIDC organization settings", settingsTableOptions)
}
