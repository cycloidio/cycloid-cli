package orgplugins

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "plugin",
		Aliases: []string{"plugins"},
		Short:   "manage installed organization plugins",
	}

	cmd.AddCommand(
		NewListCommand(),
		NewGetCommand(),
		NewUpdateCommand(),
		NewDeleteCommand(),
		NewLogsCommand(),
		NewWidgetCommands(),
	)

	return cmd
}
