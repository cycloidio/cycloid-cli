package stacks

import "github.com/spf13/cobra"

func NewConfigCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Args:  cobra.NoArgs,
		Short: "manage stack configuration.",
	}

	cmd.AddCommand(
		NewConfigGetCommand(),
	)

	return cmd
}
