package organizations

import (
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
`,
		RunE: update,
	}

	cmd.MarkFlagRequired(cyargs.AddOrgNameFlag(cmd))
	cmd.Flags().Bool("can-children-manage-oidc-mapping", true, "Whether child organizations are allowed to manage their own OIDC group mappings")

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

	o, _, err := m.UpdateOrganization(org, name, opts)
	return cyout.PrintWithOptions(cmd, o, err, "unable to update organization", printer.Options{})
}
