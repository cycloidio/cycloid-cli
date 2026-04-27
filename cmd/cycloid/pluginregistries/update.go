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

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "update",
		Args:  cobra.NoArgs,
		Short: "update a plugin registry",
		Example: `
	# Update a plugin registry name
	cy --org my-org plugin-registry update --registry-id 1 --name new-name
`,
		RunE: updateRegistry,
	}

	cmd.MarkFlagRequired(cyargs.AddPluginRegistryIDFlag(cmd))
	cmd.Flags().String("name", "", "new name for the plugin registry")
	cmd.MarkFlagRequired("name")

	return cmd
}

func updateRegistry(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}
	registryID, err := cyargs.GetPluginRegistryID(cmd)
	if err != nil {
		return err
	}
	name, err := cmd.Flags().GetString("name")
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

	result, _, err := m.UpdatePluginRegistry(org, registryID, name)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to update plugin registry", printer.Options{}, cmd.OutOrStderr())
	}
	return printer.SmartPrint(p, result, nil, "", printer.Options{}, cmd.OutOrStdout())
}
