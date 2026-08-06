package view

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "view",
		Aliases: []string{"views"},
		Short:   "Manage plugin widget views",
	}

	cmd.AddCommand(
		NewListCommand(),
		NewUpdateCommand(),
	)
	return cmd
}
