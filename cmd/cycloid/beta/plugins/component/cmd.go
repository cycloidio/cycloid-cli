package component

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "component",
		Aliases: []string{"components"},
		Short:   "[beta] Manage component-plugin relations",
	}

	cmd.AddCommand(
		NewListCommand(),
		NewRelationSetCommand(),
	)
	return cmd
}
