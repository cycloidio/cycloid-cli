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
	cmd := &cobra.Command{
		Use:   "list",
		Args:  cobra.NoArgs,
		Short: "List org-level plugin widgets",
		Example: `
  cy plugin widget list --placement sideMenuPage
  cy plugin widget list --placement component
`,
		RunE: listPluginWidgets,
	}
	cmd.Flags().String("placement", "sideMenuPage", "filter widgets by placement type (e.g. sideMenuPage, component)")
	return cmd
}

func listPluginWidgets(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	placement, err := cmd.Flags().GetString("placement")
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	result, _, err := m.ListPluginWidgets(org, placement)
	return cyout.PrintWithOptions(cmd, result, err, "unable to list plugin widgets", printer.Options{})
}
