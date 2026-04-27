package orgplugins

import (
	"github.com/spf13/cobra"
)

func NewWidgetCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "widget",
		Aliases: []string{"widgets"},
		Short:   "manage plugin widgets",
	}

	cmd.AddCommand(
		NewWidgetListCommand(),
	)

	return cmd
}
