package admin

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

func NewDeleteUserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Args:    cobra.NoArgs,
		Short:   "Delete a user of the instance, with their memberships and the API keys they own",
		Long: `Delete a user of the instance, with their memberships and the API keys they own.
It works for users without an organization too. --org has to be the root organization, and the
caller one of its admins. The deletion is refused when the user is the last admin of an
organization; the error names those organizations.`,
		Example: `
	# Delete the user with this email
	cy --org my-root-org admin users delete --email user@example.com
`,
		RunE: deleteUser,
	}

	cmd.Flags().String("email", "", "email of the user to delete")
	cmd.MarkFlagRequired("email")

	return cmd
}

func deleteUser(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	email, err := cmd.Flags().GetString("email")
	if err != nil {
		return err
	}

	_, err = m.AdminDeleteUser(org, email)
	if err != nil {
		return cyout.Print(cmd, nil, err, "unable to delete user "+email)
	}
	return nil
}
