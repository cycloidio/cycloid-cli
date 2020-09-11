package members

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/youdeploy-cli/printer"
	"github.com/cycloidio/youdeploy-cli/printer/factory"
)

func NewListCommand() *cobra.Command {
	var (
		example = `
	# List all the members within my-org organization
	cy --org my-org members list
	`
		short = "Get the list of organization members"
		long  = short
	)

	var cmd = &cobra.Command{
		Use:     "list",
		Example: example,
		Short:   short,
		Long:    long,
		RunE:    listMembers,
	}

	return cmd
}

func listMembers(cmd *cobra.Command, args []string) error {
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

	mbs, err := m.ListMembers(org)
	if err != nil {
		return err
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// print the result on the standard output
	if err := p.Print(mbs, printer.Options{}, os.Stdout); err != nil {
		return errors.Wrap(err, "unable to print result")
	}

	return nil
}
