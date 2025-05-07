package projects

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cy_args"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewListEnvCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "list-env",
		Short:   "List environments in the current project",
		Example: `cy --org my-org projects list-env -p project -o json`,
		RunE:    listEnv,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	cy_args.AddProjectFlag(cmd)
	return cmd
}

func listEnv(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cy_args.GetOrg(cmd)
	if err != nil {
		return err
	}

	project, err := cy_args.GetProject(cmd)
	if err != nil {
		return err
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

	projects, err := m.ListProjectsEnv(org, project)
	return printer.SmartPrint(p, projects, err, "unable to listenv project", printer.Options{}, cmd.OutOrStdout())
}
