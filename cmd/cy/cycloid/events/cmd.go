package events

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "event",
		Aliases: []string{
			"events",
			"e",
		},
		Short: "Manage the events",
	}

	cmd.AddCommand(
		NewSendCommand(),
		NewListCommand(),
	)

	return cmd
}
