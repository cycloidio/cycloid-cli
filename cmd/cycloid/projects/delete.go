package projects

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:               "delete [canonical...]",
		Args:              cyargs.RequireArgsOrFlag("project"),
		Aliases:           []string{"del", "rm"},
		ValidArgsFunction: cyargs.CompleteProject,
		Short:             "delete a project",
		Example: `
	# delete a project by canonical
	cy --org my-org project delete my-project

	# delete multiple projects at once
	cy --org my-org project delete proj-one proj-two

	# delete using the --project flag
	cy --org my-org project delete --project my-project
`,
		RunE: deleteProject,
	}

	cyargs.AddProjectFlag(cmd)
	return cmd
}

func deleteProject(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	// If no positional args, fall back to --project flag / CY_PROJECT
	if len(args) == 0 {
		project, err := cyargs.GetProject(cmd)
		if err != nil {
			return err
		}
		args = []string{project}
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return err
	}

	if output == "table" {
		output = "json"
	}
	p, err := factory.GetPrinter(output)
	if err != nil {
		return err
	}

	deleted := make([]string, 0, len(args))
	for _, project := range args {
		_, err = m.DeleteProject(org, project)
		if err != nil {
			return printer.SmartPrint(p, nil, err, "unable to delete project "+project, printer.Options{}, cmd.OutOrStderr())
		}
		deleted = append(deleted, project)
	}
	return printer.SmartPrint(p, deleted, nil, "", printer.Options{}, cmd.OutOrStdout())
}
