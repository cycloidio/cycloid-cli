package plugin

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "create <registry-id-or-name>",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginRegistryID,
		Short:             "[beta] Create a plugin in a registry",
		Example: `
  cy beta plugin registry plugin create my-registry --name my-plugin
`,
		RunE: createRegistryPlugin,
	}

	cmd.MarkFlagRequired(cyargs.AddNameFlag(cmd))
	return cmd
}

func createRegistryPlugin(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	registryID, err := cyargs.ResolvePluginRegistryID(org, args[0], m)
	if err != nil {
		return err
	}

	name, err := cyargs.GetName(cmd)
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

	result, _, err := m.CreateRegistryPlugin(org, registryID, name)
	return printer.SmartPrint(p, result, err, "unable to create registry plugin", printer.Options{}, cmd.OutOrStdout())
}
