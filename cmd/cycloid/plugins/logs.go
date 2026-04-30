package plugins

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewLogsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "logs <id-or-name>",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginInstallID,
		Short:             "Show deployment logs for a plugin install",
		Long: `Show deployment and runtime logs for a plugin install.

Use --watch (-w) to poll continuously for new log entries. Useful for
following an ongoing deployment. Ctrl-C to stop.`,
		Example: `
  cy plugin logs 42
  cy plugin logs my-plugin

  # Follow logs as they arrive (poll every 5s)
  cy plugin logs my-plugin --watch

  # Custom polling interval
  cy plugin logs my-plugin --watch --watch-interval 10s
`,
		RunE: pluginLogs,
	}
	cmd.Flags().BoolP("watch", "w", false, "Poll for new log entries until interrupted")
	cmd.Flags().Duration("watch-interval", 5*time.Second, "Polling interval when --watch is set")
	return cmd
}

func pluginLogs(cmd *cobra.Command, args []string) error {
	watch, err := cmd.Flags().GetBool("watch")
	if err != nil {
		return err
	}
	interval, err := cmd.Flags().GetDuration("watch-interval")
	if err != nil {
		return err
	}

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	id, err := cyargs.ResolvePluginInstallID(org, args[0], m)
	if err != nil {
		return err
	}

	if !watch {
		result, _, err := m.ListPluginLogs(org, id)
		return cyout.PrintWithOptions(cmd, result, err, "unable to get plugin logs", printer.Options{})
	}

	// Watch mode: poll and print only new entries.
	var seenInstall, seenRuntime int
	for {
		result, _, err := m.ListPluginLogs(org, id)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "unable to get plugin logs", printer.Options{})
		}
		if result != nil {
			newInstall, newRuntime := logSlicesDelta(result, seenInstall, seenRuntime)
			if len(newInstall) > 0 || len(newRuntime) > 0 {
				partial := &models.PluginLogs{
					InstallLogs: newInstall,
					RuntimeLogs: newRuntime,
				}
				if printErr := cyout.PrintWithOptions(cmd, partial, nil, "", printer.Options{}); printErr != nil {
					return printErr
				}
				seenInstall += len(newInstall)
				seenRuntime += len(newRuntime)
			}
		}
		time.Sleep(interval)
	}
}

// logSlicesDelta returns the unseen portion of each log slice.
func logSlicesDelta(logs *models.PluginLogs, seenInstall, seenRuntime int) ([]*models.PluginInstallDeploymentLog, []*models.PluginInstallRuntimeLog) {
	var newInstall []*models.PluginInstallDeploymentLog
	var newRuntime []*models.PluginInstallRuntimeLog

	if seenInstall < len(logs.InstallLogs) {
		newInstall = logs.InstallLogs[seenInstall:]
	}
	if seenRuntime < len(logs.RuntimeLogs) {
		newRuntime = logs.RuntimeLogs[seenRuntime:]
	}
	return newInstall, newRuntime
}
