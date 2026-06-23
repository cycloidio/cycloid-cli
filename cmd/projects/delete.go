package projects

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
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
	cyargs.AddDeleteFlags(cmd)
	return cmd
}

func deleteProject(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	force, skipHooks, ignoreConfigFilesErr, err := cyargs.GetDeleteFlags(cmd)
	if err != nil {
		return err
	}

	if projectFlag, err := cyargs.GetProjectOrEmpty(cmd); err != nil {
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

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	opts := apiclient.DeleteOptions{Force: force, SkipHooks: skipHooks, IgnoreConfigFilesErr: ignoreConfigFilesErr}
	deleted := make([]string, 0, len(args))
	for _, project := range args {
		_, err = m.DeleteProject(org, project, opts)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "unable to delete project "+project, printer.Options{})
		}
		deleted = append(deleted, project)
	}
	return cyout.PrintWithOptions(cmd, deleted, nil, "", printer.Options{Columns: []string{"Canonical"}})
}
