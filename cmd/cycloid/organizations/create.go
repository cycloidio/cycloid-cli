package organizations

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

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "create an organization",
		Example: `
	# create an organization foo
	cy organization create --name foo
`,
		RunE:    create,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredFlag(WithFlagName, cmd)

	return cmd
}

func create(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	o, err := m.CreateOrganization(name)
	if err != nil {
		return errors.Wrap(err, "unable to create organization")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// print the result on the standard output
	if err := p.Print(o, printer.Options{}, os.Stdout); err != nil {
		return errors.Wrap(err, "unable to print result")
	}

	return nil
}
