package pluginregistries

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewVersionLogsCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "logs",
		Args:  cobra.NoArgs,
		Short: "get logs for a plugin version build",
		Example: `
	# Get logs for version 3 of plugin 2 in registry 1
	cy --org my-org plugin-registry version logs --registry-id 1 --plugin-id 2 --version-id 3
`,
		RunE: versionLogs,
	}

	cmd.MarkFlagRequired(cyargs.AddPluginRegistryIDFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddPluginIDFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddPluginVersionIDFlag(cmd))

	return cmd
}

func versionLogs(cmd *cobra.Command, args []string) error {
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

	result, _, err := m.GetPluginVersionLogs(org, registryID, pluginID, versionID)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to get plugin version logs", printer.Options{}, cmd.OutOrStderr())
	}
	return printer.SmartPrint(p, result, nil, "", printer.Options{}, cmd.OutOrStdout())
}
