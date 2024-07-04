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

func NewDeleteCommand() *cobra.Command {
	var (
		example = `
	# Remove a member from my-org organization
	cy --org my-org members delete --name my-username
	`
		short = "Remove a user from the organization"
		long  = short
	)

	var cmd = &cobra.Command{
		Use:     "delete",
		Example: example,
		Short:   short,
		Long:    long,
		RunE:    deleteMember,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredFlag(WithFlagName, cmd)

	return cmd
}

func deleteMember(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := common.GetOrg(cmd)
	if err != nil {
		return err
	}

	name, err := cmd.Flags().GetString("name")
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

	err = m.DeleteMember(org, name)
	return printer.SmartPrint(p, nil, err, "unable to remove member", printer.Options{}, cmd.OutOrStdout())
}
