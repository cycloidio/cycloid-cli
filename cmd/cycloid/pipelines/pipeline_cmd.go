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
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(
		NewBuildsCommand(),
		NewJobsCommand(),
		NewPipelineClearTaskCacheCommand(),
		NewPipelineDiffCommand(),
		NewPipelineGetCommand(),
		NewPipelineGetCommand(),
		NewPipelineLastUsedCommand(),
		NewPipelineListCommand(),
		NewPipelinePauseCommand(),
		NewPipelineSyncedCommand(),
		NewPipelineUnpauseCommand(),
		NewPipelineUpdateCommand(),
	)

	return cmd
}
