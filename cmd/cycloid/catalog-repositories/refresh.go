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

func NewRefreshCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "refresh",
		Short: "refresh a catalog repository",
		Long:  "refresh action can be used if the .cycloid.yml definition has been updated",
		Example: `
	# refresh a catalog repository with the canonical my-catalog-repository
	cy --org my-org catalog-repo refresh --canonical my-catalog-repository
`,
		RunE: refreshCatalogRepository,
	}

	common.RequiredFlag(common.WithFlagCan, cmd)

	return cmd
}

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

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	cr, err := m.RefreshCatalogRepository(org, can)
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

	// fmt.Printf("id: %d    name: %s    url: %s    branch: %s    credential_id: %d\n", *cr.ID, *cr.Name, *cr.URL, cr.Branch, cr.CredentialCanonical)
	// fmt.Printf("created_at: %v    updated_at: %v\n", time.Unix(int64(*cr.CreatedAt), 0), time.Unix(int64(*cr.UpdatedAt), 0))
	//
	// //TODO: Wait PR merged https://github.com/cycloidio/youdeploy-http-api/pull/2066
	// // output is not available yet
	// for stack := range cr.ServiceCatalogs {
	// 	_ = stack
	// }
	//
	// fmt.Println(cr)
	// fmt.Printf("%+v\n", err)

	return nil
}
