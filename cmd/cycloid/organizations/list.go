package organizations

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:    "list",
		Args:   cobra.NoArgs,
		Short:  "list the organizations",
		RunE:   list,
		Hidden: true,
		Example: `
	# list the organizations
	cy o list --output json
`,
	}
	cyout.RegisterModel(cmd, models.Organization{})
	return cmd

}

func list(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	orgs, _, err := m.ListOrganizations()
	return cyout.PrintWithOptions(cmd, orgs, err, "unable to list organizations", orgTableOptions)
}
