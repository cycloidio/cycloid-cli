package orgplugins

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewLogsCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "logs",
		Args:  cobra.NoArgs,
		Short: "get logs for an installed plugin",
		Example: `
	# Get logs for installed plugin with install ID 5
	cy --org my-org plugin logs --install-id 5
`,
		RunE: pluginLogs,
	}

	cmd.MarkFlagRequired(cyargs.AddPluginInstallIDFlag(cmd))

	return cmd
}

func pluginLogs(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}
	installID, err := cyargs.GetPluginInstallID(cmd)
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

	result, _, err := m.GetPluginLogs(org, installID)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to get plugin logs", printer.Options{}, cmd.OutOrStderr())
	}
	return printer.SmartPrint(p, result, nil, "", printer.Options{}, cmd.OutOrStdout())
}
