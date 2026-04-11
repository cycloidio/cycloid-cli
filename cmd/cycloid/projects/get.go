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
		Args:              cobra.MaximumNArgs(1),
		ValidArgsFunction: cyargs.CompleteProject,
		Short:             "get a project",
		Example: `
	# get a project by canonical
	cy --org my-org project get my-project

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

	var project string
	if len(args) == 1 {
		project = args[0]
	} else {
		project, err = cyargs.GetProject(cmd)
		if err != nil {
			return err
		}
	}

	proj, _, err := m.GetProject(org, project)
	return cyout.PrintWithOptions(cmd, proj, err, "unable to get project", projectTableOptions)
}
