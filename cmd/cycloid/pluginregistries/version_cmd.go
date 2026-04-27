package pluginregistries

import (
	"github.com/spf13/cobra"
)

func NewVersionCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "version",
		Aliases: []string{"versions"},
		Short:   "manage plugin versions in a registry",
	}

	cmd.AddCommand(
		NewVersionListCommand(),
		NewVersionGetCommand(),
		NewVersionCreateCommand(),
		NewVersionDeleteCommand(),
		NewVersionRetryCommand(),
		NewVersionLogsCommand(),
		NewVersionInstallCommand(),
	)

	return cmd
}
