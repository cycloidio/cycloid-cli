package events

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "event",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
	}

	cmd.AddCommand(NewSendCommand())

	return cmd
}
