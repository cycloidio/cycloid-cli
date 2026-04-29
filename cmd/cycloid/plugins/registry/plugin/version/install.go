package version

import (
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
		Short:             "Trigger installation of a plugin version",
		Example: `
  cy plugin registry plugin version install my-registry my-plugin 7
  cy plugin registry plugin version install my-registry my-plugin 7 --config key=value
`,
		RunE: installVersion,
	}
	cyargs.AddPluginConfigFlags(cmd)
	return cmd
}

func installVersion(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

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

	_, err = m.InstallPluginVersion(org, registryID, pluginID, versionID, configuration)
	return cyout.PrintWithOptions(cmd, nil, err, "unable to trigger plugin version install", printer.Options{})
}
