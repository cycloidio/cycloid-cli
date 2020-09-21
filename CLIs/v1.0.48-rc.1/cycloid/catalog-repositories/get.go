package catalogRepositories

import (
	"fmt"
	"time"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
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
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	id, err := cmd.Flags().GetUint32("id")
	if err != nil {
		return err
	}

	cr, err := m.GetCatalogRepository(org, id)
	if err != nil {
		return err
	}

	fmt.Printf("id: %d    name: %s    url: %s    branch: %s    credential_id: %d\n", *cr.ID, *cr.Name, *cr.URL, cr.Branch, cr.CredentialID)
	fmt.Printf("created_at: %v    updated_at: %v\n", time.Unix(*cr.CreatedAt, 0), time.Unix(*cr.UpdatedAt, 0))

	fmt.Println(cr)
	fmt.Printf("%+v\n", err)

	return nil
}
