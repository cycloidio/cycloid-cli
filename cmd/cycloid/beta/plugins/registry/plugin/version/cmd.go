package version

import "github.com/spf13/cobra"

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"versions"},
		Short:   "[beta] Manage plugin versions within a registry",
	}

	cmd.AddCommand(
		NewListCommand(),
		NewGetCommand(),
		NewPublishCommand(),
		NewDeleteCommand(),
		NewInstallCommand(),
		NewLogsCommand(),
		NewRetryCommand(),
	)
	return cmd
}
