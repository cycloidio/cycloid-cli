package catalogRepositories

import (
	"os"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get",
		Short: "get a catalog repository",
		Example: `
	# get the catalog repository with the canonical 123 and display the result in YAML
	cy  --org my-org cr get --canonical my-catalog-repository -o yaml
`,
		RunE: getCatalogRepository,
	}

	common.RequiredFlag(common.WithFlagCan, cmd)

	return cmd
}

func getCatalogRepository(cmd *cobra.Command, args []string) error {
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

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	cr, err := m.GetCatalogRepository(org, can)
	if err != nil {
		// print the result on the standard output
		if err := p.Print(err, printer.Options{}, os.Stdout); err != nil {
			return errors.Wrap(err, "unable to print result")
		}
		return err
	}

	// print the result on the standard output
	if err := p.Print(cr, printer.Options{}, os.Stdout); err != nil {
		return errors.Wrap(err, "unable to print result")
	}

	return nil
}
