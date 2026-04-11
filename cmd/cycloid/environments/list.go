package environments

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

var environmentTableOptions = printer.Options{
	Columns:    []string{"Canonical", "Name"},
	Identifier: "Canonical",
}

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "list",
		Args:    cobra.NoArgs,
		Short:   "list the environments of a project",
		Example: `cy --org my-org environments list -p project -o json`,
		RunE:    list,
	}

	cyargs.AddProjectFlag(cmd)
	cyout.RegisterModel(cmd, models.Environment{})
	return cmd
}

func list(cmd *cobra.Command, args []string) error {
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

	environments, _, err := m.ListProjectsEnv(org, project)
	return cyout.PrintWithOptions(cmd, environments, err, "unable to list environments", environmentTableOptions)
}
