package plugins

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
		Use:               "get <id-or-name>",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginInstallID,
		Short:             "Get a plugin install",
		Example: `
  cy plugin get 42
  cy plugin get my-plugin
`,
		RunE: getPlugin,
	}
}

func getPlugin(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	id, err := cyargs.ResolvePluginInstallID(org, args[0], m)
	if err != nil {
		return err
	}

	result, _, err := m.GetPlugin(org, id)
	return cyout.PrintWithOptions(cmd, result, err, "unable to get plugin", printer.Options{})
}
