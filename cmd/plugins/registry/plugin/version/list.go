package version

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "list",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cyargs.CompletePluginVersionID,
		Short:             "List versions of a plugin",
		Example: `
  cy plugin registry plugin version list --registry my-registry --plugin my-plugin
`,
		RunE: listVersions,
	}
	_ = cmd.MarkFlagRequired(cyargs.AddRegistryFlag(cmd))
	_ = cmd.MarkFlagRequired(cyargs.AddPluginFlag(cmd))
	return cmd
}

func listVersions(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	registryID, pluginID, err := resolveRegistryAndPlugin(org, cmd, m)
	if err != nil {
		return err
	}

	result, _, err := m.ListPluginVersions(org, registryID, pluginID)
	return cyout.PrintWithOptions(cmd, result, err, "unable to list plugin versions", printer.Options{})
}
