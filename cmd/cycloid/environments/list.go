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
		Short:   "list environments",
		Example: `cy --org my-org environments list -o json`,
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

	project, err := cyargs.GetProjectOrEmpty(cmd)
	if err != nil {
		return err
	}

	if project != "" {
		projectEnvs, _, err := m.ListProjectEnvs(org, project)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "unable to list project environments", environmentTableOptions)
		}
		environments := make([]*models.Environment, 0, len(projectEnvs))
		for _, env := range projectEnvs {
			environments = append(environments, projectEnvToEnvironment(env))
		}
		return cyout.PrintWithOptions(cmd, environments, nil, "", environmentTableOptions)
	}

	environments, _, err := m.ListOrgEnvs(org)
	return cyout.PrintWithOptions(cmd, environments, err, "unable to list environments", environmentTableOptions)
}
