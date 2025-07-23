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
	var (
		example = `
	# Invite a member within my-org organization
	cy --org my-org members invite --role organization-member --email user@email.com
	`
		short = "Invite the organization member"
		long  = short
	)

	var cmd = &cobra.Command{
		Use:     "invite",
		Args:    cobra.NoArgs,
		Example: example,
		Short:   short,
		Long:    long,
		RunE:    inviteMember,
	}

	common.RequiredFlag(WithFlagEmail, cmd)
	common.RequiredFlag(WithFlagRoleCanonical, cmd)

	return cmd
}

func inviteMember(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	email, err := cmd.Flags().GetString("email")
	if err != nil {
		return err
	}

	role, err := cmd.Flags().GetString("role")
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	_, err = m.InviteMember(org, email, role)
	return printer.SmartPrint(p, nil, err, "unable to invite member", printer.Options{}, cmd.OutOrStdout())
}
