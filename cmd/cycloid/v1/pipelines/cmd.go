package pipelines

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "pipeline",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
	}

	cmd.AddCommand(NewUnpauseJobCommand(),
		NewListJobsCommand(),
		NewDiffCommand(),
		NewUnpauseCommand(),
		NewPauseCommand(),
		NewGetBuildCommand(),
		NewTriggerBuildCommand(),
		NewUpdateCommand(),
		NewGetCommand(),
		NewClearTaskCacheCommand(),
		NewListCommand(),
		NewGetJobCommand(),
		NewGetListBuildsCommand(),
		NewPauseJobCommand())

	return cmd
}
