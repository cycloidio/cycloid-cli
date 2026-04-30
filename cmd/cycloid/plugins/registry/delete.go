package registry

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewDeleteCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "delete <id-or-name-or-url>...",
		Args:              cobra.MinimumNArgs(1),
		ValidArgsFunction: cyargs.CompletePluginRegistryID,
		Short:             "Delete a plugin registry",
		Example: `
  cy plugin registry delete 1
  cy plugin registry delete my-registry
  cy plugin registry delete my-registry-a my-registry-b
`,
		RunE: deletePluginRegistry,
	}
}

func deletePluginRegistry(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	for _, arg := range args {
		id, err := cyargs.ResolvePluginRegistryID(org, arg, m)
		if err != nil {
			return err
		}

		_, err = m.DeletePluginRegistry(org, id)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "unable to delete plugin registry "+arg, printer.Options{})
		}
	}
	return cyout.PrintWithOptions(cmd, nil, nil, "", printer.Options{})
}
