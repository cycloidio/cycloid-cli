package version

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "get <version-id>",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginVersionID,
		Short:             "Get a specific plugin version",
		Example: `
  cy plugin registry plugin version get 7 --registry my-registry --plugin my-plugin
`,
		RunE: getVersion,
	}
	_ = cmd.MarkFlagRequired(cyargs.AddRegistryFlag(cmd))
	_ = cmd.MarkFlagRequired(cyargs.AddPluginFlag(cmd))
	return cmd
}

func getVersion(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	versionID, err := parseVersionID(args[0])
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	registryID, pluginID, err := resolveRegistryAndPlugin(org, cmd, m)
	if err != nil {
		return err
	}

	result, _, err := m.GetPluginVersion(org, registryID, pluginID, versionID)
	return cyout.PrintWithOptions(cmd, result, err, "unable to get plugin version", printer.Options{})
}
