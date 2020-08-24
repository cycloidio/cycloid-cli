package catalogRepositories

import (
	"fmt"
	"time"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_service_catalog_sources"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  getCatalogRepository,
	}

	common.RequiredFlag(common.WithFlagID, cmd)

	return cmd
}

// /organizations/{organization_canonical}/service_catalog_sources/{service_catalog_source_id}
// get: getServiceCatalogSource
// Return the Service Catalog Source

func getCatalogRepository(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	id, err := cmd.Flags().GetUint32("id")
	if err != nil {
		return err
	}

	params := organization_service_catalog_sources.NewGetServiceCatalogSourceParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogSourceID(id)

	resp, err := api.OrganizationServiceCatalogSources.GetServiceCatalogSource(params, root.ClientCredentials())
	if err != nil {
		return err
	}

	cr := resp.GetPayload()

	fmt.Printf("id: %d    name: %s    url: %s    branch: %s    credential_id: %d\n", *cr.Data.ID, *cr.Data.Name, *cr.Data.URL, cr.Data.Branch, cr.Data.CredentialID)
	fmt.Printf("created_at: %v    updated_at: %v\n", time.Unix(*cr.Data.CreatedAt, 0), time.Unix(*cr.Data.UpdatedAt, 0))

	fmt.Println(resp)
	fmt.Printf("%+v\n", err)

	return nil
}
