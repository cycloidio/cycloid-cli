package configRepositories

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/youdeploy-cli/printer"
	"github.com/cycloidio/youdeploy-cli/printer/factory"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "list the config repositories",
		Example: `
	# list the config repositories in the org 'my-org' and display the result in JSON format
	cy  --org my-org config-repo list -o json
`,
		RunE: listConfigRepositories,
	}

	return cmd
}

// /organizations/{organization_canonical}/config_repositories
// get: getConfigRepositories
// Return all the config repositories
func listConfigRepositories(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	crs, err := m.ListConfigRepositories(org)
	if err != nil {
		return errors.Wrap(err, "unable to list config repositories")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// print the result on the standard output
	if err := p.Print(crs, printer.Options{}, os.Stdout); err != nil {
		return errors.Wrap(err, "unable to print result")
	}

	return nil
}
