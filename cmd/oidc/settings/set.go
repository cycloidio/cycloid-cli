package settings

import (
	"errors"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

// NewSetCommand returns the `settings set` command.
func NewSetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Args:  cobra.NoArgs,
		Short: "Set OIDC organization settings",
		Long: `Update OIDC reconciliation settings for the organization.

Only the flags you pass are changed; unset flags keep their current value
(read-merge-write). Note: --no-match-policy=eject requires --oidc-managed=true;
the API returns HTTP 422 otherwise.`,
		Example: `
  # Enable OIDC-managed mode and eject members with no matching group
  cy --org my-org oidc settings set --oidc-managed --no-match-policy eject

  # Set a default role for OIDC-provisioned members (other settings untouched)
  cy --org my-org oidc settings set --default-role organization-member
`,
		RunE: setSettings,
	}

	cyargs.AddOIDCManagedFlag(cmd)
	cyargs.AddOIDCNoMatchPolicyFlag(cmd)
	cyargs.AddOIDCDefaultRoleFlag(cmd)

	cyout.RegisterModel(cmd, apiclient.OIDCOrganizationSettings{})
	return cmd
}

func setSettings(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	oidcManaged, err := cyargs.GetOIDCManaged(cmd)
	if err != nil {
		return err
	}

	noMatchPolicy, err := cyargs.GetOIDCNoMatchPolicy(cmd)
	if err != nil {
		return err
	}

	defaultRole, err := cyargs.GetOIDCDefaultRole(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	// The settings endpoint is a full-replace PUT. Read the current settings and
	// override only the flags the user explicitly set, so an unspecified flag
	// never silently resets live state. A not-yet-configured org (404) starts
	// from defaults.
	// Default for fresh-create (404 on read). Must stay in sync with the
	// server-side default in the OIDC organization settings endpoint.
	merged := apiclient.UpdateOIDCOrganizationSettings{
		OIDCNoMatchPolicy: "keep_membership",
	}
	current, _, err := m.GetOIDCOrganizationSettings(org)
	if err != nil {
		var apiErr *apiclient.APIResponseError
		if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusNotFound {
			return cyout.Print(cmd, nil, err, "unable to read current OIDC organization settings")
		}
	} else if current != nil {
		merged.OIDCManaged = current.OIDCManaged
		merged.OIDCNoMatchPolicy = current.OIDCNoMatchPolicy
		merged.DefaultRoleCanonical = current.DefaultRoleCanonical
	}

	if cmd.Flags().Changed(cyargs.OIDCManagedFlagName) {
		merged.OIDCManaged = oidcManaged
	}
	if cmd.Flags().Changed(cyargs.OIDCNoMatchPolicyFlagName) {
		merged.OIDCNoMatchPolicy = noMatchPolicy
	}
	if cmd.Flags().Changed(cyargs.OIDCDefaultRoleFlagName) {
		merged.DefaultRoleCanonical = defaultRole
	}

	result, _, err := m.UpdateOIDCOrganizationSettings(org, merged)
	return cyout.PrintWithOptions(cmd, result, err, "unable to update OIDC organization settings", settingsTableOptions)
}
