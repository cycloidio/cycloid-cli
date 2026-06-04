package mappings

import (
	"github.com/spf13/cobra"
)

// NewCommands returns the `mappings` subcommand group.
func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "mappings",
		Aliases: []string{"mapping"},
		Short:   "Manage OIDC group-to-team mappings",
	}

	cmd.AddCommand(
		NewListCommand(),
		NewCreateCommand(),
		NewDeleteCommand(),
	)

	return cmd
}
