package projects

import (
	"os"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/youdeploy-cli/printer"
	"github.com/cycloidio/youdeploy-cli/printer/factory"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "list the projects within the organization",
		Example: `
	# list projects in 'my-org' and display result in JSON
	cy --org my-org projects list -o json
`,
		RunE: list,
	}
	return cmd

}

func list(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	projects, err := m.ListProjects(org)
	if err != nil {
		return errors.Wrap(err, "unable to list project")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// print the result on the standard output
	if err := p.Print(projects, printer.Options{}, os.Stdout); err != nil {
		return errors.Wrap(err, "unable to print result")
	}
	return nil
}
