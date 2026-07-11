package organizations

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
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
	m := apiclient.NewAPIClient(api)

	orgs, _, err := m.ListOrganizations()
	return cyout.PrintWithOptions(cmd, orgs, err, "unable to list organizations", orgTableOptions)
}
