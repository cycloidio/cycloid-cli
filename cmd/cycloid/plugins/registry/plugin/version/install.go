package version

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewInstallCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "install <registry> <plugin> <version-id>",
		Args:              cobra.ExactArgs(3),
		ValidArgsFunction: cyargs.CompleteRegistryPluginID,
		Short:             "Install a plugin version into your organization",
		Long: `Install a specific plugin version into your organization.

This deploys the plugin container via the plugin manager and registers the
plugin as a running install in your org. The version's manifest declares which
configuration fields are required — check them first with:

  cy plugin registry plugin version get <registry> <plugin> <version-id>

Configuration is supplied as key=value pairs (--config) or as a JSON file
(--config-file). Both can be combined: the file provides base values and
--config flags override individual keys. This is useful in CI/CD pipelines
where non-secret config lives in a committed file and secrets are injected
at runtime.

Once installed, the plugin appears in 'cy plugin list' and its widgets become
available under 'cy plugin widget list'. Use 'cy plugin upgrade' to change
version or update configuration. Use 'cy plugin uninstall' to remove it.

A plugin manager must be configured and accepted before running this command.
See 'cy plugin manager --help'.

Use --retry to make the install idempotent: if the plugin version is already
installed, the command retries the installation instead of failing. Safe to use
in CI/CD pipelines where the command may run more than once.`,
		Example: `
  # List versions to find the one you want
  cy plugin registry plugin version list my-registry my-plugin

  # Check what configuration fields a version requires
  cy plugin registry plugin version get my-registry my-plugin 7

  # Install with inline configuration
  cy plugin registry plugin version install my-registry my-plugin 7 \
    --config api_url=https://api.example.com \
    --config api_token=secret

  # Install with a config file (non-secret values committed to repo)
  cy plugin registry plugin version install my-registry my-plugin 7 \
    --config-file ./plugin-config.json \
    --config api_token=secret

  # Pin to a specific version in CI/CD (reproducible installs)
  cy plugin registry plugin version install my-registry my-plugin 7 \
    --config-file ./plugin-config.json

  # Idempotent install (safe to run in CI/CD pipelines)
  cy plugin registry plugin version install my-registry my-plugin 7 \
    --config-file ./plugin-config.json --retry
`,
		RunE: installVersion,
	}
	cyargs.AddPluginConfigFlags(cmd)
	cmd.Flags().Bool("retry", false, "If the plugin version is already installed, retry the installation instead of failing")
	return cmd
}

func installVersion(cmd *cobra.Command, args []string) error {
	retry, err := cmd.Flags().GetBool("retry")
	if err != nil {
		return err
	}

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	registryID, pluginID, err := resolveRegistryAndPlugin(org, args, m)
	if err != nil {
		return err
	}

	versionID, err := parseVersionID(args[2])
	if err != nil {
		return err
	}

	configuration, err := cyargs.GetPluginConfig(cmd)
	if err != nil {
		return err
	}

	_, installErr := m.InstallPluginVersion(org, registryID, pluginID, versionID, configuration)
	if installErr != nil && retry && strings.Contains(strings.ToLower(installErr.Error()), "already") {
		_, installErr = m.RetryPluginVersion(org, registryID, pluginID, versionID)
	}
	return cyout.PrintWithOptions(cmd, nil, installErr, "unable to trigger plugin version install", printer.Options{})
}
