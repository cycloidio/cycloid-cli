package widget

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "widget",
		Aliases: []string{"widgets"},
		Short:   "Manage component-level plugin widgets",
	}

	cmd.AddCommand(
		NewListCommand(),
		NewQueryCommand(),
	)
	return cmd
}
