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
		NewListCommand(),
		NewGetPipelineCommand(),
		NewLastUsedCommand(),
		NewUpdateCommand(),
		NewJobsCommand(),
		// NewGetListBuildsCommand(),
		// NewDiffCommand(),
		// NewUnpauseCommand(),
		// NewPauseCommand(),
		// NewGetBuildCommand(),
		// NewTriggerBuildCommand(),
		// NewClearTaskCacheCommand(),
		// NewGetCommand(),
		// NewSyncedCommand(),
		// NewPauseJobCommand(),
	)

	return cmd
}
