package version

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
	return &cobra.Command{
		Use:               "logs <registry> <plugin> <version-id>",
		Args:              cobra.ExactArgs(3),
		ValidArgsFunction: cyargs.CompleteRegistryPluginID,
		Short:             "[beta] Show installation logs for a plugin version",
		Example: `
  cy beta plugin registry plugin version logs my-registry my-plugin 7
`,
		RunE: versionLogs,
	}
}

func versionLogs(cmd *cobra.Command, args []string) error {
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

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	result, _, err := m.ListPluginVersionLogs(org, registryID, pluginID, versionID)
	return printer.SmartPrint(p, result, err, "unable to get plugin version logs", printer.Options{}, cmd.OutOrStdout())
}
