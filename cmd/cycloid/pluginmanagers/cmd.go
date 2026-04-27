package pluginmanagers

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "plugin-manager",
		Aliases: []string{
			"plugin-managers",
			"pm",
		},
		Short: "manage plugin managers",
	}

	cmd.AddCommand(
		NewListCommand(),
		NewCreateCommand(),
	)

	return cmd
}
