package pipelines

import (
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "pipeline",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
	}

	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	cmd.AddCommand(NewUnpauseJobCommand(),
		NewUpdateCommand(),
		NewGetJobCommand(),
		NewGetListBuildsCommand(),
		NewListJobsCommand(),
		NewDiffCommand(),
		NewUnpauseCommand(),
		NewPauseCommand(),
		// NewGetBuildCommand(),
		NewTriggerBuildCommand(),
		// NewGetCommand(),
		NewClearTaskCacheCommand(),
		// NewListCommand(),
		NewPauseJobCommand())

	return cmd
}
