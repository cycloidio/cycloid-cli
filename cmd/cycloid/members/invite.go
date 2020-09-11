package members

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
)

func NewInviteCommand() *cobra.Command {
	var (
		example = `
	# Invite a member within my-org organization
	cy --org my-org members invite --email user@email.com
	`
		short = "Invite the organization member"
		long  = short
	)

	var cmd = &cobra.Command{
		Use:     "invite",
		Example: example,
		Short:   short,
		Long:    long,
		RunE:    inviteMember,
	}

	common.RequiredFlag(WithFlagEmail, cmd)
	common.RequiredFlag(WithFlagRoleID, cmd)

	return cmd
}

// /organizations/{organization_canonical}/members-invitations
// put: inviteUserToOrgMember

func inviteMember(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	email, err := cmd.Flags().GetString("email")
	if err != nil {
		return err
	}

	roleID, err := cmd.Flags().GetUint32("role-id")
	if err != nil {
		return err
	}

	err = m.InviteMember(org, email, roleID)
	if err != nil {
		return err
	}

	return nil
}
