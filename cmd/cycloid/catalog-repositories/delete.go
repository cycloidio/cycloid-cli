package catalogRepositories

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "delete a catalog repository",
		Example: `
	# delete a catalog repository with the ID 123
	cy  --org my-org catalog-repository delete --canonical my-catalog-repository
`,
		RunE: deleteCatalogRepository,
	}

	common.RequiredFlag(common.WithFlagCan, cmd)

	return cmd
}

// /organizations/{organization_canonical}/service_catalog_sources/{service_catalog_source_id}
// delete: deleteServiceCatalogSource
// delete a Service catalog source

func deleteCatalogRepository(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	canonical, err := cmd.Flags().GetString("canonical")
	if err != nil {
		return err
	}

	if err := m.DeleteCatalogRepository(org, canonical); err != nil {
		return errors.Wrap(err, "unable to delete repository")
	}
	return nil
}
