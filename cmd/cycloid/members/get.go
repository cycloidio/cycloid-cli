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

func NewGetCommand() *cobra.Command {
	var (
		example = `
	# Get a member within my-org organization
	cy --org my-org members get --name my-username
	`
		short = "Get the organization member"
		long  = short
	)

	var cmd = &cobra.Command{
		Use:     "get",
		Example: example,
		Short:   short,
		Long:    long,
		RunE:    getMember,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredFlag(WithFlagName, cmd)

	return cmd
}

func getMember(cmd *cobra.Command, args []string) error {
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

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	mb, err := m.GetMember(org, name)
	return printer.SmartPrint(p, mb, err, "unable to get members", printer.Options{}, cmd.OutOrStdout())
}
