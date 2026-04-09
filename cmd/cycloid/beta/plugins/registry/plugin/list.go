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

func NewListCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "list <registry-id-or-name>",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginRegistryID,
		Short:             "[beta] List plugins in a registry",
		Example: `
  cy beta plugin registry plugin list my-registry
  cy beta plugin registry plugin list 1
`,
		RunE: listRegistryPlugins,
	}
}

func listRegistryPlugins(cmd *cobra.Command, args []string) error {
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

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	result, _, err := m.ListRegistryPlugins(org, registryID)
	return printer.SmartPrint(p, result, err, "unable to list registry plugins", printer.Options{}, cmd.OutOrStdout())
}
