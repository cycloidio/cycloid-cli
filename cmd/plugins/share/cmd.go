package share

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "share",
		Short: "Manage plugin sharing with child organizations",
	}

	cmd.AddCommand(
		NewGetCommand(),
		NewSetCommand(),
	)
	return cmd
}
