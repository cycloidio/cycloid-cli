package catalogRepositories

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewRefreshCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "refresh",
		Short: "refresh a catalog repository",
		Long:  "refresh action can be used if the .cycloid.yml definition has been updated",
		Example: `
	# refresh a catalog repository with the ID 123
	cy --org my-org catalog-repo refresh --canonical my-catalog-repository
`,
		RunE: refreshCatalogRepository,
	}

	common.RequiredFlag(common.WithFlagCan, cmd)

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

	can, err := cmd.Flags().GetString("canonical")
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	cr, err := m.RefreshCatalogRepository(org, can)
	if err != nil {
		return err
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// print the result on the standard output
	if err := p.Print(cr, printer.Options{}, os.Stdout); err != nil {
		return errors.Wrap(err, "unable to print result")
	}

	return nil
}
