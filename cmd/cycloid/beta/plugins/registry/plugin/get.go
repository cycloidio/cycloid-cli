package plugin

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewGetCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "get <registry-id-or-name> <plugin-id-or-name>",
		Args:              cobra.ExactArgs(2),
		ValidArgsFunction: cyargs.CompleteRegistryPluginID,
		Short:             "[beta] Get a plugin from a registry",
		Example: `
  cy beta plugin registry plugin get my-registry my-plugin
  cy beta plugin registry plugin get 1 42
`,
		RunE: getRegistryPlugin,
	}
}

func getRegistryPlugin(cmd *cobra.Command, args []string) error {
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

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	result, _, err := m.GetRegistryPlugin(org, registryID, pluginID)
	return printer.SmartPrint(p, result, err, "unable to get registry plugin", printer.Options{}, cmd.OutOrStdout())
}
