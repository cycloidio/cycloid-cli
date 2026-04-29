package plugin

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "update <registry-id-or-name> <plugin-id-or-name>",
		Args:              cobra.ExactArgs(2),
		ValidArgsFunction: cyargs.CompleteRegistryPluginID,
		Short:             "[beta] Update a plugin in a registry",
		Example: `
  cy beta plugin registry plugin update my-registry my-plugin --name new-name
  cy beta plugin registry plugin update 1 42 --name renamed
`,
		RunE: updateRegistryPlugin,
	}

	cmd.MarkFlagRequired(cyargs.AddNameFlag(cmd))
	return cmd
}

func updateRegistryPlugin(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	registryID, err := cyargs.ResolvePluginRegistryID(org, args[0], m)
	if err != nil {
		return err
	}

	pluginID, err := cyargs.ResolveRegistryPluginID(org, registryID, args[1], m)
	if err != nil {
		return err
	}

	name, err := cyargs.GetName(cmd)
	if err != nil {
		return err
	}

	result, _, err := m.UpdateRegistryPlugin(org, registryID, pluginID, name)
	return cyout.PrintWithOptions(cmd, result, err, "unable to update registry plugin", printer.Options{})
}
