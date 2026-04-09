package members

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewInviteCommand() *cobra.Command {
	var cmd = &cobra.Command{
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
	m := middleware.NewMiddleware(api)

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

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	mb, _, err := m.InviteMember(org, email, role)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to invite member", printer.Options{}, cmd.OutOrStderr())
	}
	return printer.SmartPrint(p, mb, nil, "", printer.Options{}, cmd.OutOrStdout())
}
