package version

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewDeleteCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "delete <registry> <plugin> <version-id>",
		Args:              cobra.ExactArgs(3),
		ValidArgsFunction: cyargs.CompleteRegistryPluginID,
		Short:             "[beta] Delete a plugin version",
		Example: `
  cy beta plugin registry plugin version delete my-registry my-plugin 7
`,
		RunE: deleteVersion,
	}
}

func deleteVersion(cmd *cobra.Command, args []string) error {
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
		return err
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return err
	}

	_, err = m.DeletePluginVersion(org, registryID, pluginID, versionID)
	return printer.SmartPrint(p, nil, err, "unable to delete plugin version", printer.Options{}, cmd.OutOrStdout())
}
