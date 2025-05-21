package members

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewUpdateCommand() *cobra.Command {
	var (
		example = `
	# Update a member within my-org organization
	cy --org my-org members update --id 50 --role my-role
	`
		short = "Update a member"
		long  = short
	)

	var cmd = &cobra.Command{
		Use:     "update",
		Args:    cobra.NoArgs,
		Example: example,
		Short:   short,
		Long:    long,
		RunE:    updateConfigRepository,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredFlag(WithFlagID, cmd)
	common.RequiredFlag(WithFlagRoleCanonical, cmd)

	//TODO : dont Required flags and if not set, use value from the getConfigRepository

	return cmd
}

func updateConfigRepository(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	id, err := cmd.Flags().GetUint32("id")
	if err != nil {
		return err
	}

	role, err := cmd.Flags().GetString("role")
	if err != nil {
		return err
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	mb, err := m.UpdateMember(org, id, role)
	return printer.SmartPrint(p, mb, err, "unable to update member", printer.Options{}, cmd.OutOrStdout())
}
