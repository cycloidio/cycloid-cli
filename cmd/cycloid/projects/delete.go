package projects

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
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

	if projectFlag, err := cyargs.GetProject(cmd); err != nil {
		return err
	} else if projectFlag != "" {
		found := false
		for _, a := range args {
			if a == projectFlag {
				found = true
				break
			}
		}
		if !found {
			args = append(args, projectFlag)
		}
	}

	deleted := make([]string, 0, len(args))
	for _, project := range args {
		_, err = m.DeleteProject(org, project)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "unable to delete project "+project, printer.Options{})
		}
		deleted = append(deleted, project)
	}
	return cyout.PrintWithOptions(cmd, deleted, nil, "", printer.Options{Columns: []string{"Canonical"}})
}
