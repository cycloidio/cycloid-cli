package registry

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewGetCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "get <id-or-name-or-url>",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginRegistryID,
		Short:             "Get a plugin registry",
		Example: `
  cy plugin registry get 1
  cy plugin registry get my-registry
  cy plugin registry get https://registry.example.com
`,
		RunE: getPluginRegistry,
	}
}

func getPluginRegistry(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	id, err := cyargs.ResolvePluginRegistryID(org, args[0], m)
	if err != nil {
		return err
	}

	result, _, err := m.GetPluginRegistry(org, id)
	return cyout.PrintWithOptions(cmd, result, err, "unable to get plugin registry", printer.Options{})
}
