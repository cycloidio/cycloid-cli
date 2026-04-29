package widget

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Args:  cobra.NoArgs,
		Short: "List org-level plugin widgets",
		Example: `
  cy plugin widget list
`,
		RunE: listPluginWidgets,
	}
}

func listPluginWidgets(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	result, _, err := m.ListPluginWidgets(org)
	return cyout.PrintWithOptions(cmd, result, err, "unable to list plugin widgets", printer.Options{})
}
