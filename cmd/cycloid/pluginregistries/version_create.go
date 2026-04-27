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

func NewVersionCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Args:  cobra.NoArgs,
		Short: "create a plugin version",
		Example: `
	# Create a version for plugin 2 in registry 1
	cy --org my-org plugin-registry version create --registry-id 1 --plugin-id 2 --url https://example.com/plugin.tar.gz
`,
		RunE: createVersion,
	}

	cmd.MarkFlagRequired(cyargs.AddPluginRegistryIDFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddPluginIDFlag(cmd))
	cmd.Flags().String("url", "", "URL of the plugin version archive")
	cmd.MarkFlagRequired("url")

	return cmd
}

func createVersion(cmd *cobra.Command, args []string) error {
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
	url, err := cmd.Flags().GetString("url")
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

	result, _, err := m.CreatePluginVersion(org, registryID, pluginID, url)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to create plugin version", printer.Options{}, cmd.OutOrStderr())
	}
	return printer.SmartPrint(p, result, nil, "", printer.Options{}, cmd.OutOrStdout())
}
