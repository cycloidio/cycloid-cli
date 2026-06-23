package configrepositories

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

var configRepoTableOptions = printer.Options{
	Columns:    []string{"Canonical", "Name", "URL", "Branch", "Default"},
	Identifier: "Canonical",
}

func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Args:  cobra.NoArgs,
		Short: "list the config repositories",
		Example: `
	# list the config repositories in the org 'my-org' and display the result in JSON format
	cy  --org my-org config-repo list -o json
`,
		RunE: listConfigRepositories,
	}
	cyout.RegisterModel(cmd, models.ConfigRepository{})
	return cmd
}

func listConfigRepositories(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	crs, _, err := m.ListConfigRepositories(org)
	return cyout.PrintWithOptions(cmd, crs, err, "unable to list config repositories", configRepoTableOptions)
}
