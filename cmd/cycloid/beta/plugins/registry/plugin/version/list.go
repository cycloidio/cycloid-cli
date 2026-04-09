package version

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
		Use:               "list <registry> <plugin>",
		Args:              cobra.ExactArgs(2),
		ValidArgsFunction: cyargs.CompleteRegistryPluginID,
		Short:             "[beta] List versions of a plugin",
		Example: `
  cy beta plugin registry plugin version list my-registry my-plugin
`,
		RunE: listVersions,
	}
}

func listVersions(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	registryID, pluginID, err := resolveRegistryAndPlugin(org, args, m)
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

	result, _, err := m.ListPluginVersions(org, registryID, pluginID)
	return printer.SmartPrint(p, result, err, "unable to list plugin versions", printer.Options{}, cmd.OutOrStdout())
}
