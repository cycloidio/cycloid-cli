package widget

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewQueryCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "query <widget-id>",
		Args:  cobra.ExactArgs(1),
		Short: "Execute the data query configured for a plugin widget",
		Long: `Execute the data query pre-configured for a plugin widget and return its results.

Each plugin widget is an instance of a query defined in the plugin manifest.
A single plugin may expose multiple queries; each becomes a separate widget with
its own ID. This command executes the query bound to the given widget and returns
whatever data the plugin produces — the shape is plugin-defined.

Only widgets of type "table" are supported. The associated plugin must be running.

Use 'cy plugin widget list' to discover widget IDs and their configured queries.`,
		Example: `
  # List available widgets to find the ID you want
  cy plugin widget list

  # Execute the query for widget 42 and print results as a table
  cy plugin widget query 42

  # Get the raw JSON output
  cy plugin widget query 42 --output json
`,
		RunE: queryPluginWidget,
	}
}

func queryPluginWidget(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	n, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid widget-id %q: must be a positive integer", args[0])
	}
	widgetID := uint32(n)

	result, _, err := m.QueryPluginWidget(org, widgetID)
	return cyout.PrintWithOptions(cmd, result, err, "unable to query plugin widget", printer.Options{})
}
