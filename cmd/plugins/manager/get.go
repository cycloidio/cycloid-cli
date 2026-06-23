package manager

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
		ValidArgsFunction: cyargs.CompletePluginManagerID,
		Short:             "Get a plugin manager",
		Example: `
  cy plugin manager get 1
  cy plugin manager get my-manager
`,
		RunE: getPluginManager,
	}
}

func getPluginManager(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	id, err := cyargs.ResolvePluginManagerID(org, args[0], m)
	if err != nil {
		return err
	}

	result, _, err := m.GetPluginManager(org, id)
	return cyout.PrintWithOptions(cmd, result, err, "unable to get plugin manager", printer.Options{})
}
