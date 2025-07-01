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

	cmd.AddCommand(
		NewListCommand(),
		NewGetPipelineCommand(),
		NewLastUsedCommand(),
		NewUpdateCommand(),
		NewJobsCommand(),
		NewBuildsCommand(),
		// NewGetListBuildsCommand(),
		// NewDiffCommand(),
		// NewUnpauseCommand(),
		// NewPauseCommand(),
		// NewGetBuildCommand(),
		// NewClearTaskCacheCommand(),
		// NewGetCommand(),
		// NewSyncedCommand(),
	)

	return cmd
}
