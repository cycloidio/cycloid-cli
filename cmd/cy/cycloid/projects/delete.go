package projects

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cy/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cy/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cy/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cy_args"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "delete",
		Args:    cobra.NoArgs,
		Aliases: []string{"del", "rm"},
		Short:   "delete a project",
		Example: `cy --org my-org project delete --project my-project`,
		RunE:    deleteProject,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	cy_args.AddProjectFlag(cmd)
	return cmd
}

func deleteProject(cmd *cobra.Command, args []string) error {
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

	err = m.DeleteProject(org, project)
	return printer.SmartPrint(p, nil, err, "unable to delete project", printer.Options{}, cmd.OutOrStdout())
}
