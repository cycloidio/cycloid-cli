package stacks

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "stack",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
	}
	cmd.AddCommand(NewListCommand(),
		NewGetCommand())

	return cmd
}
