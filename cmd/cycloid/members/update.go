package members

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewUpdateCommand() *cobra.Command {
	var (
		example = `
	# Update a member within my-org organization
	cy --org my-org members update --name my_user --role my-role
	`
		short = "Update a member"
		long  = short
	)

	var cmd = &cobra.Command{
		Use:     "update",
		Example: example,
		Short:   short,
		Long:    long,
		RunE:    updateConfigRepository,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredFlag(WithFlagName, cmd)
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

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	role, err := cmd.Flags().GetString("role")
	if err != nil {
		return err
	}

	mb, err := m.UpdateMembers(org, name, role)
	if err != nil {
		return err
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// print the result on the standard output
	if err := p.Print(mb, printer.Options{}, os.Stdout); err != nil {
		return errors.Wrap(err, "unable to print result")
	}

	return nil
}
