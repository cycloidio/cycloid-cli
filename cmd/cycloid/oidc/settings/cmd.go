package settings

import (
	"github.com/spf13/cobra"
)

// NewCommands returns the `settings` subcommand group.
func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "settings",
		Short: "Manage OIDC organization settings",
	}

	cmd.AddCommand(
		NewGetCommand(),
		NewSetCommand(),
	)

	return cmd
}
