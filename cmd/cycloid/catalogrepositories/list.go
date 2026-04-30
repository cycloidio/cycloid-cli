package catalogrepositories

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

var catalogSourceTableOptions = printer.Options{
	Columns:    []string{"Canonical", "Name", "URL", "Branch", "StackCount"},
	Identifier: "Canonical",
}

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Args:  cobra.NoArgs,
		Short: "list the catalog repositories",
		Example: `
	# list the catalog repositories in the org 'my-org' and display the result in JSON format
	cy --org my-org cr list -o json
`,
		RunE: listCatalogRepositories,
	}
	cyout.RegisterModel(cmd, models.ServiceCatalogSource{})
	return cmd
}

func listCatalogRepositories(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	crs, _, err := m.ListCatalogRepositories(org)
	return cyout.PrintWithOptions(cmd, crs, err, "unable to list catalog repositories", catalogSourceTableOptions)
}
