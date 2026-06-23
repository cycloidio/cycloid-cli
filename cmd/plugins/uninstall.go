package plugins

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewUninstallCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "uninstall <id-or-name>...",
		Args:              cobra.MinimumNArgs(1),
		ValidArgsFunction: cyargs.CompletePluginInstallID,
		Short:             "Uninstall a plugin",
		Example: `
  cy plugin uninstall 42
  cy plugin uninstall my-plugin
  cy plugin uninstall my-plugin-a my-plugin-b
`,
		RunE: uninstallPlugin,
	}
}

func uninstallPlugin(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	for _, arg := range args {
		id, err := cyargs.ResolvePluginInstallID(org, arg, m)
		if err != nil {
			return err
		}

		_, err = m.DeletePlugin(org, id)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "unable to uninstall plugin "+arg, printer.Options{})
		}
	}
	return cyout.PrintWithOptions(cmd, nil, nil, "", printer.Options{})
}
