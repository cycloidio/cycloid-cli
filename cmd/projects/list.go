package projects

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

var projectTableOptions = printer.Options{
	Columns:    []string{"Canonical", "Name", "Description", "Owner.Username"},
	Identifier: "Canonical",
}

func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Args:  cobra.NoArgs,
		Short: "list the projects within the organization",
		Example: `# list projects in 'my-org' and display result in JSON
cy --org my-org projects list -o json`,
		RunE: list,
	}
	cyout.RegisterModel(cmd, models.Project{})
	return cmd
}

func list(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	projects, _, err := m.ListProjects(org)
	return cyout.PrintWithOptions(cmd, projects, err, "unable to list projects", projectTableOptions)
}
