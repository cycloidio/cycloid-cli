package projects

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewListEnvCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "list-env",
		Args:    cobra.NoArgs,
		Short:   "List environments in the current project",
		Example: `cy --org my-org projects list-env -p project -o json`,
		RunE:    listEnv,
	}

	cyargs.AddProjectFlag(cmd)
	return cmd
}

func listEnv(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
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

	projects, err := m.ListProjectsEnv(org, project)
	return printer.SmartPrint(p, projects, err, "unable to listenv project", printer.Options{}, cmd.OutOrStdout())
}
