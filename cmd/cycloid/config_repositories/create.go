package config_repositories

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Args:  cobra.NoArgs,
		Short: "create a config repository",
		Example: `
	# create a config repository and set up as default
	cy --org my-org config-repo create --branch config --cred my-cred --url "git@github.com:my/repo.git" --name default-config --default
`,
		RunE:    createConfigRepository,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredFlag(common.WithFlagCred, cmd)
	common.RequiredFlag(WithFlagName, cmd)
	common.RequiredFlag(WithFlagBranch, cmd)
	common.RequiredFlag(WithFlagURL, cmd)
	cyargs.AddConfigRepositoryFlag(cmd)
	WithFlagDefault(cmd)

	return cmd
}

func createConfigRepository(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	canonical, err := cyargs.GetConfigRepository(cmd, org)
	if err != nil {
		return err
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	if name == "" {
		name = canonical
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

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	cr, err := m.CreateConfigRepository(org, name, canonical, url, branch, cred, setDefault)
	return printer.SmartPrint(p, cr, err, "unable to create config repository", printer.Options{}, cmd.OutOrStdout())
}
