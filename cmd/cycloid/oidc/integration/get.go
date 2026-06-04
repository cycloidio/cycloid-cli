package integration

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

var integrationTableOptions = printer.Options{
	Columns: []string{
		"Enabled",
		"OidcDisplayName",
		"OidcClientID",
		"OidcIssuer",
		"OidcDiscoveryURL",
		"OidcGroupsClaimName",
		"OidcSessionTTLSeconds",
		"HasSecret",
		"HasCaCertificate",
	},
}

// NewGetCommand returns the `integration get` command.
func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Args:  cobra.NoArgs,
		Short: "Get the org's OIDC SSO integration config",
		Example: `
  # Get the OIDC SSO integration for my-org
  cy --org my-org oidc integration get
`,
		RunE: getIntegration,
	}

	cyout.RegisterModel(cmd, middleware.OIDCIntegration{})
	return cmd
}

func getIntegration(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	result, _, err := m.GetOIDCIntegration(org)
	return cyout.PrintWithOptions(cmd, result, err, "unable to get OIDC integration", integrationTableOptions)
}
