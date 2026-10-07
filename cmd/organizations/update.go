package organizations

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

// This command have been Hidden because it is not compatible with API key login.
// Advanced user still can use it passing a user token in CY_API_KEY env var during a login.
func NewUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Args:  cobra.NoArgs,
		Short: "update an organization",
		Example: `
	# update an organization foo
	cy organization update --org org --name foo

	# allow child organizations to manage OIDC mappings
	cy organization update --org org --name foo --can-children-manage-oidc-mapping=true

	# hide the "out of sync" stack version indicator from the component and project lists
	cy organization update --org org --name foo --hide-stack-version-out-of-sync=true

	# allow two users to impersonate other users (root organization only)
	cy organization update --org org --name foo --impersonation-emails alice@example.com,bob@example.com

	# clear the impersonation allowlist
	cy organization update --org org --name foo --impersonation-emails ""

	# reject signup everywhere on the platform unless the email is already invited (root organization only)
	cy organization update --org org --name foo --disable-signup=true
`,
		RunE: update,
	}

	cmd.MarkFlagRequired(cyargs.AddOrgNameFlag(cmd))
	cmd.Flags().Bool("can-children-manage-oidc-mapping", true, "Whether child organizations are allowed to manage their own OIDC group mappings")
	cmd.Flags().Bool("hide-stack-version-out-of-sync", false, "Whether the component and project lists hide the indicator shown when a version's commit no longer matches its reference commit")
	cmd.Flags().StringSlice("impersonation-emails", nil, "Comma-separated emails of the users allowed to impersonate other users (root organization only). Pass an empty string to clear the list")
	cmd.Flags().Bool("disable-signup", false, "Reject signup everywhere on the platform (plain email/password, AWS Marketplace, and SSO/social auto-enrollment) unless the email is already invited (root organization only)")

	return cmd
}

func update(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	name, err := cyargs.GetOrgName(cmd)
	if err != nil {
		return err
	}

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return errors.Wrap(err, "unable get org flag")
	}

	var opts apiclient.UpdateOrganizationOpts
	if cmd.Flags().Changed("can-children-manage-oidc-mapping") {
		v, _ := cmd.Flags().GetBool("can-children-manage-oidc-mapping")
		opts.CanChildrenManageOidcMapping = &v
	}
	if cmd.Flags().Changed("hide-stack-version-out-of-sync") {
		v, _ := cmd.Flags().GetBool("hide-stack-version-out-of-sync")
		opts.HideStackVersionOutOfSync = &v
	}
	if cmd.Flags().Changed("impersonation-emails") {
		v, _ := cmd.Flags().GetStringSlice("impersonation-emails")
		// non-nil even when empty: an empty slice clears the list. pflag keeps the
		// space after a comma ("a@b.io, c@d.io"), which the API pattern rejects
		emails := make([]string, 0, len(v))
		for _, e := range v {
			if e = strings.TrimSpace(e); e != "" {
				emails = append(emails, e)
			}
		}
		opts.ImpersonationEmails = emails
	}
	if cmd.Flags().Changed("disable-signup") {
		v, _ := cmd.Flags().GetBool("disable-signup")
		opts.DisableSignup = &v
	}

	o, _, err := m.UpdateOrganization(org, name, opts)
	return cyout.PrintWithOptions(cmd, o, err, "unable to update organization", printer.Options{})
}
