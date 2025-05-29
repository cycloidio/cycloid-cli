package config_repositories

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cy/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cy/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cy/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cy_args"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Args:  cobra.NoArgs,
		Short: "list the config repositories",
		Example: `
	# list the config repositories in the org 'my-org' and display the result in JSON format
	cy  --org my-org config-repo list -o json
`,
		RunE:    listConfigRepositories,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	return cmd
}

func listConfigRepositories(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cy_args.GetOrg(cmd)
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

	crs, err := m.ListConfigRepositories(org)
	return printer.SmartPrint(p, crs, err, "unable to list config repository", printer.Options{}, cmd.OutOrStdout())
}
