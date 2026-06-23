package plugin

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "get <plugin-id-or-name>...",
		Args:              cobra.MinimumNArgs(1),
		ValidArgsFunction: cyargs.CompletePluginIDFromRegistryFlag,
		Short:             "Get one or more plugins from a registry",
		Example: `
  cy plugin registry plugin get my-plugin --registry my-registry
  cy plugin registry plugin get plugin-a plugin-b --registry my-registry
  cy plugin registry plugin get 42 --registry 1
`,
		RunE: getRegistryPlugin,
	}
	_ = cmd.MarkFlagRequired(cyargs.AddRegistryFlag(cmd))
	return cmd
}

func getRegistryPlugin(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	registryStr, err := cyargs.GetRegistry(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	registryID, err := cyargs.ResolvePluginRegistryID(org, registryStr, m)
	if err != nil {
		return err
	}

	if len(args) == 1 {
		pluginID, err := cyargs.ResolveRegistryPluginID(org, registryID, args[0], m)
		if err != nil {
			return err
		}
		result, _, err := m.GetRegistryPlugin(org, registryID, pluginID)
		return cyout.PrintWithOptions(cmd, result, err, "unable to get registry plugin", printer.Options{})
	}

	results := make([]*models.Plugin, 0, len(args))
	for _, nameOrID := range args {
		pluginID, err := cyargs.ResolveRegistryPluginID(org, registryID, nameOrID, m)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "unable to resolve plugin "+nameOrID, printer.Options{})
		}
		p, _, err := m.GetRegistryPlugin(org, registryID, pluginID)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "unable to get plugin "+nameOrID, printer.Options{})
		}
		results = append(results, p)
	}
	return cyout.PrintWithOptions(cmd, results, nil, "", printer.Options{})
}
