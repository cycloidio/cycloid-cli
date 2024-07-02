package pipelines

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "pipeline",
		Aliases: []string{
			"pp",
			"pipelines",
		},
		Short: "Manage the pipelines",
	}

	cmd.AddCommand(NewUnpauseJobCommand(),
		NewUpdateCommand(),
		NewGetJobCommand(),
		NewGetListBuildsCommand(),
		NewListJobsCommand(),
		NewDiffCommand(),
		NewUnpauseCommand(),
		NewPauseCommand(),
		NewGetBuildCommand(),
		NewTriggerBuildCommand(),
		NewGetCommand(),
		NewClearTaskCacheCommand(),
		NewListCommand(),
		NewGetCommand(),
		NewSyncedCommand(),
		NewPauseJobCommand())

	return cmd
}
