package configRepositories

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "create a config repository",
		Example: `
	# create a config repository and set up as default
	cy --org my-org config-repo create --branch stacks --cred my-cred --url "git@github.com:my/repo.git" --name my-catalog-name --default
`,
		RunE:    createConfigRepository,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	// create --branch test --cred 105 --url "git@github.com:foo/bla.git"  --name configname  --default

	common.RequiredFlag(common.WithFlagCred, cmd)
	common.RequiredFlag(WithFlagName, cmd)
	common.RequiredFlag(WithFlagBranch, cmd)
	common.RequiredFlag(WithFlagURL, cmd)
	WithFlagDefault(cmd)

	return cmd
}

// /organizations/{organization_canonical}/config_repositories
// post: createConfigRepository
// Creates a config repository

func createConfigRepository(cmd *cobra.Command, args []string) error {
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

	setDefault, err := cmd.Flags().GetBool("default")
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

	cr, err := m.CreateConfigRepository(org, name, url, branch, cred, setDefault)
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
