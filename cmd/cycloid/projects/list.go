package projects

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cy_args"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "list the projects within the organization",
		Example: `# list projects in 'my-org' and display result in JSON
cy --org my-org projects list -o json`,
		RunE:    list,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}
	return cmd

}

func list(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cy_args.GetOrg(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get org")
	}

	output, err := cy_args.GetOutput(cmd)
	if err != nil {
		return err
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return err
	}

	projects, err := m.ListProjects(org)
	return printer.SmartPrint(p, projects, err, "unable to list project", printer.Options{}, cmd.OutOrStdout())
}
