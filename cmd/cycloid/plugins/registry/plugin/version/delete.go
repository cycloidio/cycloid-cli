package version

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "delete <version-id>",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginVersionID,
		Short:             "Delete a plugin version",
		Example: `
  cy plugin registry plugin version delete 7 --registry my-registry --plugin my-plugin
`,
		RunE: deleteVersion,
	}
	_ = cmd.MarkFlagRequired(cyargs.AddRegistryFlag(cmd))
	_ = cmd.MarkFlagRequired(cyargs.AddPluginFlag(cmd))
	return cmd
}

func deleteVersion(cmd *cobra.Command, args []string) error {
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

	_, err = m.DeletePluginVersion(org, registryID, pluginID, versionID)
	return cyout.PrintWithOptions(cmd, nil, err, "unable to delete plugin version", printer.Options{})
}
