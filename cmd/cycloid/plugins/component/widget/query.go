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
	cmd := &cobra.Command{
		Use:   "query <widget-id>",
		Args:  cobra.ExactArgs(1),
		Short: "Query data for a component-level plugin widget",
		Example: `
  cy plugin component widget query 42 --project my-project --env production --component my-component
`,
		RunE: queryComponentPluginWidget,
	}

	cyargs.AddProjectFlag(cmd)
	cyargs.AddEnvFlag(cmd)
	cyargs.AddComponentFlag(cmd)
	return cmd
}

func queryComponentPluginWidget(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	project, err := cyargs.GetProject(cmd)
	if err != nil {
		return err
	}

	env, err := cyargs.GetEnv(cmd)
	if err != nil {
		return err
	}

	component, err := cyargs.GetComponent(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	n, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid widget-id %q: must be a positive integer", args[0])
	}
	widgetID := uint32(n)

	result, _, err := m.QueryComponentPluginWidget(org, project, env, component, widgetID)
	return cyout.PrintWithOptions(cmd, result, err, "unable to query component plugin widget", printer.Options{})
}
