package environments

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "list",
		Args:    cobra.NoArgs,
		Short:   "list the environments of a project",
		Example: `cy --org my-org environments list -p project -o json`,
		RunE:    list,
	}

	cyargs.AddProjectFlag(cmd)
	return cmd
}

func list(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get org")
	}

	project, err := cyargs.GetProject(cmd)
	if err != nil {
		return err
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return err
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return err
	}

	environments, err := m.ListProjectsEnv(org, project)
	return printer.SmartPrint(p, environments, err, "unable to list environments", printer.Options{}, cmd.OutOrStdout())
}
