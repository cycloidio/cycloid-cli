package version

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewRetryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "retry <version-id>",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginVersionID,
		Short:             "Retry a failed plugin version installation",
		Example: `
  cy plugin registry plugin version retry 7 --registry my-registry --plugin my-plugin
`,
		RunE: retryVersion,
	}
	_ = cmd.MarkFlagRequired(cyargs.AddRegistryFlag(cmd))
	_ = cmd.MarkFlagRequired(cyargs.AddPluginFlag(cmd))
	return cmd
}

func retryVersion(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	versionID, err := parseVersionID(args[0])
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	registryID, pluginID, err := resolveRegistryAndPlugin(org, cmd, m)
	if err != nil {
		return err
	}

	_, err = m.RetryPluginVersion(org, registryID, pluginID, versionID)
	return cyout.PrintWithOptions(cmd, nil, err, "unable to retry plugin version install", printer.Options{})
}
