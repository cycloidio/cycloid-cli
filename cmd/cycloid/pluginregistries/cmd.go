package pluginregistries

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "plugin-registry",
		Aliases: []string{
			"plugin-registries",
			"pr",
		},
		Short: "manage plugin registries",
	}

	cmd.AddCommand(
		NewListCommand(),
		NewCreateCommand(),
		NewUpdateCommand(),
		NewDeleteCommand(),
		NewPluginCommands(),
		NewVersionCommands(),
	)

	return cmd
}
