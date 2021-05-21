package catalogRepositories

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "create a catalog repository",
		Example: `
	# create a catalog repository using credential canonical 123, branch 'stacks' and git URL
	cy --org my-org catalog-repo create --branch stacks --cred my-cred --url "git@github.com:my/repo.git" --name my-catalog-name

	# create a catalog repository using public git repository
	cy --org my-org catalog-repo create --branch stacks --url "https://github.com:my/repo.git" --name my-catalog-name
`,
		RunE: createCatalogRepository,
	}

	// create --branch test --cred my-cred --url "git@github.com:foo/bla.git"  --name catalogname
	common.WithFlagCred(cmd)

	common.RequiredFlag(WithFlagName, cmd)
	common.RequiredFlag(WithFlagBranch, cmd)
	common.RequiredFlag(WithFlagURL, cmd)

	return cmd
}

// /organizations/{organization_canonical}/service_catalog_sources
// post: createServiceCatalogSource
// Creates a Service catalog source

func createCatalogRepository(cmd *cobra.Command, args []string) error {

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	url, err := cmd.Flags().GetString("url")
	if err != nil {
		return err
	}

	branch, err := cmd.Flags().GetString("branch")
	if err != nil {
		return err
	}

	cred, err := cmd.Flags().GetString("cred")
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

	cr, err := m.CreateCatalogRepository(org, name, url, branch, cred)
	return printer.SmartPrint(p, cr, err, "unable to create catalog repository", printer.Options{}, cmd.OutOrStdout())
}
