package registry

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "update <id-or-name-or-url>",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginRegistryID,
		Short:             "Update a plugin registry",
		Example: `
  cy plugin registry update 1 --name new-name
  cy plugin registry update my-registry --name renamed
`,
		RunE: updatePluginRegistry,
	}

	cmd.MarkFlagRequired(cyargs.AddNameFlag(cmd))
	return cmd
}

func updatePluginRegistry(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	id, err := cyargs.ResolvePluginRegistryID(org, args[0], m)
	if err != nil {
		return err
	}

	name, err := cyargs.GetName(cmd)
	if err != nil {
		return err
	}

	result, _, err := m.UpdatePluginRegistry(org, id, name)
	return cyout.PrintWithOptions(cmd, result, err, "unable to update plugin registry", printer.Options{})
}
