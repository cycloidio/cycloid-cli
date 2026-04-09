package registry

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
		Use:               "delete <id-or-name-or-url>",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginRegistryID,
		Short:             "[beta] Delete a plugin registry",
		Example: `
  cy beta plugin registry delete 1
  cy beta plugin registry delete my-registry
`,
		RunE: deletePluginRegistry,
	}
}

func deletePluginRegistry(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	id, err := cyargs.ResolvePluginRegistryID(org, args[0], m)
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

	_, err = m.DeletePluginRegistry(org, id)
	return printer.SmartPrint(p, nil, err, "unable to delete plugin registry", printer.Options{}, cmd.OutOrStdout())
}
