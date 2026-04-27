package pluginregistries

import (
	"github.com/spf13/cobra"
)

func NewPluginCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "plugin",
		Aliases: []string{"plugins"},
		Short:   "manage plugins in a registry",
	}

	cmd.AddCommand(
		NewPluginGetCommand(),
		NewPluginCreateCommand(),
		NewPluginUpdateCommand(),
		NewPluginDeleteCommand(),
	)

	return cmd
}
