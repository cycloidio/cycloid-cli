package oidc

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/oidc/integration"
	"github.com/cycloidio/cycloid-cli/cmd/oidc/mappings"
	"github.com/cycloidio/cycloid-cli/cmd/oidc/settings"
)

// NewCommands returns the top-level `oidc` cobra command.
func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oidc",
		Short: "Manage OIDC group mappings and settings",
		Long: `Manage OIDC group-to-team mappings and per-organization reconciliation
settings.`,
	}

	cmd.AddCommand(
		integration.NewCommands(),
		mappings.NewCommands(),
		settings.NewCommands(),
	)

	return cmd
}
