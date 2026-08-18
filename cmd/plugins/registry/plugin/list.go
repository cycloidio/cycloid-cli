package plugin

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewListCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "list <registry-id-or-name>",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginRegistryID,
		Short:             "List plugins in a registry",
		Example: `
  cy plugin registry plugin list my-registry
  cy plugin registry plugin list 1
`,
		RunE: listRegistryPlugins,
	}
}

func listRegistryPlugins(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	registryID, err := cyargs.ResolvePluginRegistryID(org, args[0], m)
	if err != nil {
		return err
	}

	result, _, err := m.ListRegistryPlugins(org, registryID)
	return cyout.PrintWithOptions(cmd, result, err, "unable to list registry plugins", printer.Options{})
}
