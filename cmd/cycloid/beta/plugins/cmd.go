package plugins

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/beta/plugins/component"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/beta/plugins/manager"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/beta/plugins/registry"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "plugin",
		Aliases: []string{"plugins"},
		Short:   "[beta] Manage plugins",
		Long: `Manage plugins, plugin managers, and plugin registries.

This command group is in beta. Backwards compatibility is not guaranteed.`,
	}

	cmd.AddCommand(
		NewListCommand(),
		NewGetCommand(),
		NewInstallCommand(),
		NewUpgradeCommand(),
		NewUninstallCommand(),
		NewLogsCommand(),
		component.NewCommands(),
		manager.NewCommands(),
		registry.NewCommands(),
	)
	return cmd
}
