package catalogRepositories

import (
	"fmt"
	"time"

	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  listCatalogRepositories,
	}

	return cmd
}

// /organizations/{organization_canonical}/service_catalog_sources
// get: getServiceCatalogSources
// Return all the private service catalogs
func listCatalogRepositories(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	crs, err := m.ListCatalogRepositories(org)

	for _, cr := range crs {
		fmt.Printf("id: %d    name: %s    url: %s    branch: %s    credential_id: %d\n", *cr.ID, *cr.Name, *cr.URL, cr.Branch, cr.CredentialID)
		fmt.Printf("created_at: %v    updated_at: %v\n", time.Unix(*cr.CreatedAt, 0), time.Unix(*cr.UpdatedAt, 0))
	}
	fmt.Println(crs)
	fmt.Printf("%+v\n", err)

	return nil
}
