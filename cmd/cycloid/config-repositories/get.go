package configRepositories

import (
	"os"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get",
		Short: "get a config repository",
		Example: `
	# get the config repository with the id 123 and display the result in YAML
	cy  --org my-org config-repo get --id 123 -o yaml
`,
		RunE:    getConfigRepository,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredFlag(common.WithFlagCan, cmd)

	return cmd
}

// /organizations/{organization_canonical}/config_repositories/{config_repository_id}
// get: getConfigRepository
// Return the Config Repository

func getConfigRepository(cmd *cobra.Command, args []string) error {
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

	cr, err := m.GetConfigRepository(org, can)
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
