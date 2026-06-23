package integration

import (
	"github.com/spf13/cobra"
)

// NewCommands returns the `integration` subcommand group.
func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "integration",
		Aliases: []string{"config"},
		Short:   "Manage the org's OIDC SSO integration",
		Long:    `Create or update the AuthenticationOIDC SSO integration for the organization.`,
	}

	cmd.AddCommand(
		NewGetCommand(),
		NewSetCommand(),
	)

	return cmd
}
