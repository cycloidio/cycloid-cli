package plugins

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewLogsCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "logs <id-or-name>",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginInstallID,
		Short:             "[beta] Show deployment logs for a plugin install",
		Example: `
  cy beta plugin logs 42
  cy beta plugin logs my-plugin
`,
		RunE: pluginLogs,
	}
}

func pluginLogs(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	id, err := cyargs.ResolvePluginInstallID(org, args[0], m)
	if err != nil {
		return err
	}

	result, _, err := m.ListPluginLogs(org, id)
	return cyout.PrintWithOptions(cmd, result, err, "unable to get plugin logs", printer.Options{})
}
