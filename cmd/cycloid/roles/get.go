package roles

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

func NewGetCommand() *cobra.Command {
	var (
		example = `
	# Get a role within my-org organization
	cy --org my-org roles get --id my-role-id
	`
		short = "Get the organization role"
		long  = short
	)

	var cmd = &cobra.Command{
		Use:     "get",
		Example: example,
		Short:   short,
		Long:    long,
		RunE:    getRole,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredFlag(common.WithFlagID, cmd)

	return cmd
}

func getRole(cmd *cobra.Command, args []string) error {
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

	id, err := cmd.Flags().GetUint32("id")
	if err != nil {
		return err
	}

	mb, err := m.GetRole(org, id)
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
