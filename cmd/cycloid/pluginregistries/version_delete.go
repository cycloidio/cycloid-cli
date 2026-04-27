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

func NewVersionDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Args:  cobra.NoArgs,
		Short: "delete a plugin version",
		Example: `
	# Delete version 3 of plugin 2 in registry 1
	cy --org my-org plugin-registry version delete --registry-id 1 --plugin-id 2 --version-id 3
`,
		RunE: deleteVersion,
	}

	cmd.MarkFlagRequired(cyargs.AddPluginRegistryIDFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddPluginIDFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddPluginVersionIDFlag(cmd))

	return cmd
}

func deleteVersion(cmd *cobra.Command, args []string) error {
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
	versionID, err := cyargs.GetPluginVersionID(cmd)
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

	_, err = m.DeletePluginVersion(org, registryID, pluginID, versionID)
	return printer.SmartPrint(p, nil, err, "unable to delete plugin version", printer.Options{}, cmd.OutOrStdout())
}
