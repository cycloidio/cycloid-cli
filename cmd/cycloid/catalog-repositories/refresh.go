package catalogRepositories

import (
	"fmt"
	"time"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/spf13/cobra"
)

func NewRefreshCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "refresh",
		Short: "refresh a catalog repository",
		Long:  "refresh action can be used if the .cycloid.yml definition has been updated",
		Example: `
	# refresh a catalog repository with the ID 123
	cy --org my-org catalog-repo refresh --id 123
`,
		RunE: refreshCatalogRepository,
	}

	common.RequiredFlag(common.WithFlagID, cmd)

	return cmd
}

// /organizations/{organization_canonical}/service_catalog_sources/{service_catalog_source_id}/refresh
// post: refreshServiceCatalogSource
// refresh a Service catalog source
func refreshCatalogRepository(cmd *cobra.Command, args []string) error {
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

	cr, err := m.RefreshCatalogRepository(org, id)
	if err != nil {
		return err
	}
	fmt.Printf("id: %d    name: %s    url: %s    branch: %s    credential_id: %d\n", *cr.ID, *cr.Name, *cr.URL, cr.Branch, cr.CredentialID)
	fmt.Printf("created_at: %v    updated_at: %v\n", time.Unix(*cr.CreatedAt, 0), time.Unix(*cr.UpdatedAt, 0))

	//TODO: Wait PR merged https://github.com/cycloidio/youdeploy-http-api/pull/2066
	// output is not available yet
	for stack := range cr.ServiceCatalogs {
		_ = stack
	}

	fmt.Println(cr)
	fmt.Printf("%+v\n", err)

	return nil
}
