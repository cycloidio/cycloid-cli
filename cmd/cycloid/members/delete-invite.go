package members

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewDeleteInviteCommand() *cobra.Command {
	var (
		example = `
	# Delete an invite within my-org organization
	cy --org my-org members delete-invite
	`
		short = "Delete an organization members invite"
		long  = short
	)

	var cmd = &cobra.Command{
		Use:     "delete-invite",
		Example: example,
		Short:   short,
		Long:    long,
		RunE:    deleteInvite,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredFlag(WithFlagInvite, cmd)
	return cmd
}

func deleteInvite(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	org, err := common.GetOrg(cmd)
	if err != nil {
		return err
	}

	invite, err := cmd.Flags().GetString("invite")
	if err != nil {
		return err
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	err = m.DeleteInvite(org, invite)
	return printer.SmartPrint(p, nil, err, "unable to delete invite", printer.Options{}, cmd.OutOrStdout())
}
