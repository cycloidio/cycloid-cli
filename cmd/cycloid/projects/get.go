package projects

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:               "get [canonical...]",
		Args:              cyargs.RequireArgsOrFlag("project"),
		ValidArgsFunction: cyargs.CompleteProject,
		Short:             "get a project",
		Example: `
	# get a project by canonical
	cy --org my-org project get my-project

	# get multiple projects
	cy --org my-org project get proj-a proj-b

	# get a project using the --project flag or CY_PROJECT env var
	cy --org my-org project get --project my-project -o yaml
`,
		RunE: get,
	}

	cyargs.AddProjectFlag(cmd)
	cyout.RegisterModel(cmd, models.Project{})
	return cmd
}

func get(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
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

	if len(args) == 1 {
		proj, _, err := m.GetProject(org, args[0])
		return cyout.PrintWithOptions(cmd, proj, err, "unable to get project", projectTableOptions)
	}

	results := make([]*models.Project, 0, len(args))
	for _, canonical := range args {
		proj, _, err := m.GetProject(org, canonical)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "unable to get project "+canonical, projectTableOptions)
		}
		results = append(results, proj)
	}
	return cyout.PrintWithOptions(cmd, results, nil, "", projectTableOptions)
}
