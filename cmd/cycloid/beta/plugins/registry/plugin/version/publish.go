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

func NewPublishCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "publish <registry> <plugin>",
		Args:              cobra.ExactArgs(2),
		ValidArgsFunction: cyargs.CompleteRegistryPluginID,
		Short:             "[beta] Publish a new version of a plugin",
		Example: `
  cy beta plugin registry plugin version publish my-registry my-plugin --url https://example.com/plugin-v1.2.tar.gz
`,
		RunE: publishVersion,
	}

	cmd.Flags().String("url", "", "URL of the plugin version archive (required)")
	cmd.MarkFlagRequired("url")
	return cmd
}

func publishVersion(cmd *cobra.Command, args []string) error {
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

	url, err := cmd.Flags().GetString("url")
	if err != nil {
		return errors.Wrap(err, "unable to get --url flag")
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	result, _, err := m.CreatePluginVersion(org, registryID, pluginID, url)
	return printer.SmartPrint(p, result, err, "unable to publish plugin version", printer.Options{}, cmd.OutOrStdout())
}
