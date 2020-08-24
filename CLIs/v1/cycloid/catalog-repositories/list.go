package catalogRepositories

import (
	"fmt"
	"time"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_service_catalog_sources"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
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
	var err error
	var org string

	org, err = cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	params := organization_service_catalog_sources.NewGetServiceCatalogSourcesParams()
	params.SetOrganizationCanonical(org)

	resp, err := api.OrganizationServiceCatalogSources.GetServiceCatalogSources(params, root.ClientCredentials())
	if err != nil {
		return err
	}

	crs := resp.GetPayload()

	for _, cr := range crs.Data {
		fmt.Printf("id: %d    name: %s    url: %s    branch: %s    credential_id: %d\n", *cr.ID, *cr.Name, *cr.URL, cr.Branch, cr.CredentialID)
		fmt.Printf("created_at: %v    updated_at: %v\n", time.Unix(*cr.CreatedAt, 0), time.Unix(*cr.UpdatedAt, 0))
	}
	fmt.Println(resp)
	fmt.Printf("%+v\n", err)

	return nil
}
