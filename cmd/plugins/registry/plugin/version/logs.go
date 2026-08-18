package version

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewLogsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "logs <version-id>",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginVersionID,
		Short:             "Show installation logs for a plugin version",
		Example: `
  cy plugin registry plugin version logs 7 --registry my-registry --plugin my-plugin
`,
		RunE: versionLogs,
	}
	_ = cmd.MarkFlagRequired(cyargs.AddRegistryFlag(cmd))
	_ = cmd.MarkFlagRequired(cyargs.AddPluginFlag(cmd))
	return cmd
}

func versionLogs(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	versionID, err := parseVersionID(args[0])
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	registryID, pluginID, err := resolveRegistryAndPlugin(org, cmd, m)
	if err != nil {
		return err
	}

	result, _, err := m.ListPluginVersionLogs(org, registryID, pluginID, versionID)
	return cyout.PrintWithOptions(cmd, result, err, "unable to get plugin version logs", printer.Options{})
}
