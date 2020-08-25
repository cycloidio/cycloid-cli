package catalogRepositories

import (
	"fmt"

	"time"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_service_catalog_sources"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

func NewRefreshCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "refresh",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  refreshCatalogRepository,
	}

	common.RequiredFlag(common.WithFlagID, cmd)

	return cmd
}

// /organizations/{organization_canonical}/service_catalog_sources/{service_catalog_source_id}/refresh
// post: refreshServiceCatalogSource
// refresh a Service catalog source
func refreshCatalogRepository(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	id, err := cmd.Flags().GetUint32("id")
	if err != nil {
		return err
	}

	params := organization_service_catalog_sources.NewRefreshServiceCatalogSourceParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogSourceID(id)

	resp, err := api.OrganizationServiceCatalogSources.RefreshServiceCatalogSource(params, root.ClientCredentials())
	if err != nil {
		return err
	}

	cr := resp.GetPayload()
	fmt.Printf("id: %d    name: %s    url: %s    branch: %s    credential_id: %d\n", *cr.Data.ID, *cr.Data.Name, *cr.Data.URL, cr.Data.Branch, cr.Data.CredentialID)
	fmt.Printf("created_at: %v    updated_at: %v\n", time.Unix(*cr.Data.CreatedAt, 0), time.Unix(*cr.Data.UpdatedAt, 0))

	for stack := range cr.Data.ServiceCatalogs {
		spew.Dump(stack)
	}

	fmt.Println(resp)
	fmt.Printf("%+v\n", err)

	return nil
}
