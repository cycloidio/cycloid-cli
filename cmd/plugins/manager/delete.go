package manager

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewDeleteCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "delete <id-or-name>...",
		Args:              cobra.MinimumNArgs(1),
		ValidArgsFunction: cyargs.CompletePluginManagerID,
		Short:             "Delete a plugin manager",
		Example: `
  cy plugin manager delete 1
  cy plugin manager delete my-manager
  cy plugin manager delete my-manager-a my-manager-b
`,
		RunE: deletePluginManager,
	}
}

func deletePluginManager(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	for _, arg := range args {
		id, err := cyargs.ResolvePluginManagerID(org, arg, m)
		if err != nil {
			return err
		}

		_, err = m.DeletePluginManager(org, id)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "unable to delete plugin manager "+arg, printer.Options{})
		}
	}
	return cyout.PrintWithOptions(cmd, nil, nil, "", printer.Options{})
}
