package config

import "github.com/spf13/cobra"

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "config",
		Short:   "Configuration related commands",
		Aliases: []string{"cfg"},
	}
	cmd.AddCommand(
		NewInterpolateCmd(),
	)
	return cmd
}
