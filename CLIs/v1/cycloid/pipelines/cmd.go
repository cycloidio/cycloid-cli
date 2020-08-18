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
