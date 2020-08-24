package catalogRepositories

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_service_catalog_sources"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/spf13/cobra"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  deleteCatalogRepository,
	}

	common.RequiredFlag(common.WithFlagID, cmd)

	return cmd
}

// /organizations/{organization_canonical}/service_catalog_sources/{service_catalog_source_id}
// delete: deleteServiceCatalogSource
// delete a Service catalog source

func deleteCatalogRepository(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	id, err := cmd.Flags().GetUint32("id")
	if err != nil {
		return err
	}

	params := organization_service_catalog_sources.NewDeleteServiceCatalogSourceParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogSourceID(id)

	resp, err := api.OrganizationServiceCatalogSources.DeleteServiceCatalogSource(params, root.ClientCredentials())
	if err != nil {
		return err
	}

	fmt.Println(resp)
	fmt.Printf("%+v\n", err)
	return nil
}
