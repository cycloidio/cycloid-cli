package pipelines

import (
	"github.com/spf13/cobra"
)

func NewBuildsCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "builds",
		Aliases: []string{
			"b",
			"build",
		},
		Short: "manage pipeline builds",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(
		NewBuildCreateCommand(),
		NewBuildListCommand(),
		NewBuildGetCommand(),
	)
	return cmd
}
