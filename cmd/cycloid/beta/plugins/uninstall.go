package plugins

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewUninstallCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "uninstall <id-or-name>",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginInstallID,
		Short:             "[beta] Uninstall a plugin",
		Example: `
  cy beta plugin uninstall 42
  cy beta plugin uninstall my-plugin
`,
		RunE: uninstallPlugin,
	}
}

func uninstallPlugin(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	id, err := cyargs.ResolvePluginInstallID(org, args[0], m)
	if err != nil {
		return err
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return err
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return err
	}

	_, err = m.DeletePlugin(org, id)
	return printer.SmartPrint(p, nil, err, "unable to uninstall plugin", printer.Options{}, cmd.OutOrStdout())
}
