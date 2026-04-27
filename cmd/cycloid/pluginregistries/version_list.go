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

func NewVersionListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Args:  cobra.NoArgs,
		Short: "list plugin versions",
		Example: `
	# List versions of plugin 2 in registry 1
	cy --org my-org plugin-registry version list --registry-id 1 --plugin-id 2
`,
		RunE: listVersions,
	}

	cmd.MarkFlagRequired(cyargs.AddPluginRegistryIDFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddPluginIDFlag(cmd))

	return cmd
}

func listVersions(cmd *cobra.Command, args []string) error {
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

	result, _, err := m.ListPluginVersions(org, registryID, pluginID)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to list plugin versions", printer.Options{}, cmd.OutOrStderr())
	}
	return printer.SmartPrint(p, result, nil, "", printer.Options{}, cmd.OutOrStdout())
}
