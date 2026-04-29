package widget

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewQueryCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "query <widget-id>",
		Args:  cobra.ExactArgs(1),
		Short: "Query data for an org-level plugin widget",
		Example: `
  cy plugin widget query 42
`,
		RunE: queryPluginWidget,
	}
}

func queryPluginWidget(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	n, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid widget-id %q: must be a positive integer", args[0])
	}
	widgetID := uint32(n)

	result, _, err := m.QueryPluginWidget(org, widgetID)
	return cyout.PrintWithOptions(cmd, result, err, "unable to query plugin widget", printer.Options{})
}
