package plugins

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/plugins/component"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/plugins/manager"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/plugins/registry"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/plugins/widget"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "plugin",
		Aliases: []string{"plugins"},
		Short: "Manage plugins",
		Long:  "Manage plugins, plugin managers, and plugin registries.",
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
		widget.NewCommands(),
	)
	return cmd
}
