package members

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewInviteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invite",
		Args:  cobra.NoArgs,
		Short: "Invite a user to the organization",
		Example: `
	# Invite a member within my-org organization
	cy --org my-org members invite --role organization-member --email user@email.com
`,
		RunE: inviteMember,
	}

	cmd.MarkFlagRequired(cyargs.AddMemberEmailFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddMemberRoleFlag(cmd))

	return cmd
}

func inviteMember(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	email, err := cyargs.GetMemberEmail(cmd)
	if err != nil {
		return err
	}

	role, err := cyargs.GetMemberRole(cmd)
	if err != nil {
		return err
	}

	mb, _, err := m.InviteMember(org, email, role)
	return cyout.PrintWithOptions(cmd, mb, err, "unable to invite member", printer.Options{})
}
