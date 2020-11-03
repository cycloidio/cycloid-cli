package pipelines

import (
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
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

	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

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
		NewPauseJobCommand())

	return cmd
}
