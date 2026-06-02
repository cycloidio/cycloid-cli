package oidc

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/beta/oidc/integration"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/beta/oidc/mappings"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/beta/oidc/settings"
)

// NewCommands returns the top-level `oidc` cobra command.
func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oidc",
		Short: "Manage OIDC group mappings and settings (beta)",
		Long: `Manage OIDC group-to-team mappings and per-organization reconciliation
settings. These commands are experimental; backwards-compatibility is not guaranteed.`,
	}

	cmd.AddCommand(
		integration.NewCommands(),
		mappings.NewCommands(),
		settings.NewCommands(),
	)

	return cmd
}
