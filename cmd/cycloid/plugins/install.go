package plugins

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewInstallCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Args:  cobra.NoArgs,
		Short: "Install a plugin",
		Example: `
  # Install a plugin with inline configuration
  cy plugin install --config host=localhost --config port=8080

  # Install a plugin from a JSON config file
  cy plugin install --config-file ./plugin-config.json

  # Install a specific plugin version (pin for CI/CD)
  cy plugin install --version-id 82 --config CY_API_KEY=secret
`,
		RunE: installPlugin,
	}

	cyargs.AddPluginConfigFlags(cmd)
	cyargs.AddPluginVersionIDFlag(cmd)
	return cmd
}

func installPlugin(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	versionID, err := cyargs.GetPluginVersionID(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get --version-id flag")
	}

	config, err := cyargs.GetPluginConfig(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get plugin configuration")
	}

	result, _, err := m.CreatePlugin(org, versionID, config)
	return cyout.PrintWithOptions(cmd, result, err, "unable to install plugin", printer.Options{})
}
