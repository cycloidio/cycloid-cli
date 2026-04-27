package orgplugins

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

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "update",
		Args:  cobra.NoArgs,
		Short: "update an installed plugin",
		Example: `
	# Update plugin install 5 to version 10
	cy --org my-org plugin update --install-id 5 --version-id 10 --configuration '{"key":"value"}'
`,
		RunE: updatePlugin,
	}

	cmd.MarkFlagRequired(cyargs.AddPluginInstallIDFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddPluginVersionIDFlag(cmd))
	cmd.Flags().String("configuration", "{}", "JSON configuration for the plugin")

	return cmd
}

func updatePlugin(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}
	installID, err := cyargs.GetPluginInstallID(cmd)
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

	result, _, err := m.UpdatePlugin(org, installID, versionID, configuration)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to update plugin", printer.Options{}, cmd.OutOrStderr())
	}
	return printer.SmartPrint(p, result, nil, "", printer.Options{}, cmd.OutOrStdout())
}
