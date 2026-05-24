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
		Short:   "Manage plugins",
		Long: `Manage plugins in your organization.

Concepts and dependency order:

  Registry   An external source of plugins (Docker-registry-compatible).
             Must exist before any plugin can be installed.
             → cy plugin registry

  Plugin     A plugin definition published to a registry. Has a name,
             description, and one or more versions.
             → cy plugin registry plugin

  Version    A specific release of a plugin. Bundles a Docker image and
             a manifest that declares: required configuration fields,
             widgets (UI panels), and data queries.
             → cy plugin registry plugin version

  Manager    The agent that deploys and runs plugin containers in your
             infrastructure. Required before the first install.
             → cy plugin manager

  Install    Running instance of a plugin version, deployed by the manager
             with org-specific configuration (API keys, endpoints, etc.).
             This is what appears in 'cy plugin list'.
             → cy plugin registry plugin version install

  Widget     A UI panel backed by a running plugin. Each widget wraps one
             pre-configured query from the plugin manifest. Widgets are
             created automatically from the manifest on install.
             → cy plugin widget

To install a plugin for the first time:

  # 1. Add a registry
  cy plugin registry add --url https://registry.example.com

  # 2. Browse available plugins and versions
  cy plugin registry plugin list my-registry
  cy plugin registry plugin version list my-registry my-plugin

  # 3. Install (configuration keys are declared in the version manifest)
  cy plugin registry plugin version install my-registry my-plugin <version-id> \
    --config api_url=https://api.example.com \
    --config api_token=secret

Once installed, use 'cy plugin upgrade' to change version or reconfigure,
and 'cy plugin uninstall' to remove.`,
	}

	cmd.AddCommand(
		NewListCommand(),
		NewGetCommand(),
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
