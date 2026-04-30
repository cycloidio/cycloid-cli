package plugin

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewDeleteCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "delete <registry-id-or-name> <plugin-id-or-name>",
		Args:              cobra.ExactArgs(2),
		ValidArgsFunction: cyargs.CompleteRegistryPluginID,
		Short:             "Delete a plugin from a registry",
		Example: `
  cy plugin registry plugin delete my-registry my-plugin
`,
		RunE: deleteRegistryPlugin,
	}
}

func deleteRegistryPlugin(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	registryID, err := cyargs.ResolvePluginRegistryID(org, args[0], m)
	if err != nil {
		return err
	}

	pluginID, err := cyargs.ResolveRegistryPluginID(org, registryID, args[1], m)
	if err != nil {
		return err
	}

	_, err = m.DeleteRegistryPlugin(org, registryID, pluginID)
	return cyout.PrintWithOptions(cmd, nil, err, "unable to delete registry plugin", printer.Options{})
}
