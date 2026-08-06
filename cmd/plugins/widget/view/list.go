package view

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "list <plugin-install-id-or-name>",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginInstallID,
		Short:             "List widget views for a plugin install",
		Example: `
  cy plugin widget view list my-plugin
  cy plugin widget view list 42
`,
		RunE: listWidgetViews,
	}
	return cmd
}

func listWidgetViews(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	id, err := cyargs.ResolvePluginInstallID(org, args[0], m)
	if err != nil {
		return err
	}

	result, _, err := m.ListPluginWidgetViews(org, id)
	return cyout.PrintWithOptions(cmd, result, err, "unable to list plugin widget views", printer.Options{})
}
