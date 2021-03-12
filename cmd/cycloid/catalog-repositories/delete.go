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

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "delete a catalog repository",
		Example: `
	# delete a catalog repository with the canonical my-catalog-repository
	cy  --org my-org catalog-repository delete --canonical my-catalog-repository
`,
		RunE: deleteCatalogRepository,
	}

	common.RequiredFlag(common.WithFlagCan, cmd)

	return cmd
}

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

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	err = m.DeleteCatalogRepository(org, canonical)
	return printer.SmartPrint(p, nil, err, "unable to delete catalog repository", printer.Options{}, os.Stdout)
}
