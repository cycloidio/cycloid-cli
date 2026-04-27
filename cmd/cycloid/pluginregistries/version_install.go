package pluginregistries

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewVersionInstallCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "install",
		Args:  cobra.NoArgs,
		Short: "install a plugin version",
		Example: `
	# Install version 3 of plugin 2 in registry 1
	cy --org my-org plugin-registry version install --registry-id 1 --plugin-id 2 --version-id 3 --configuration '{"key":"value"}'
`,
		RunE: installVersion,
	}

	cmd.MarkFlagRequired(cyargs.AddPluginRegistryIDFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddPluginIDFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddPluginVersionIDFlag(cmd))
	cmd.Flags().String("configuration", "{}", "JSON configuration for the plugin install")

	return cmd
}

func installVersion(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}
	registryID, err := cyargs.GetPluginRegistryID(cmd)
	if err != nil {
		return err
	}
	pluginID, err := cyargs.GetPluginID(cmd)
	if err != nil {
		return err
	}
	versionID, err := cyargs.GetPluginVersionID(cmd)
	if err != nil {
		return err
	}
	configStr, err := cmd.Flags().GetString("configuration")
	if err != nil {
		return err
	}

	var configuration map[string]string
	if err := json.Unmarshal([]byte(configStr), &configuration); err != nil {
		return errors.Wrap(err, "unable to parse configuration JSON")
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	result, _, err := m.InstallPluginVersion(org, registryID, pluginID, versionID, configuration)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to install plugin version", printer.Options{}, cmd.OutOrStderr())
	}
	return printer.SmartPrint(p, result, nil, "", printer.Options{}, cmd.OutOrStdout())
}
