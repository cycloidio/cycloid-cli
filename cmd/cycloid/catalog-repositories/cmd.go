package catalogRepositories

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "catalog-repo",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
	}
	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	cmd.AddCommand(NewUpdateCommand(),
		NewGetCommand(),
		NewDeleteCommand())

	return cmd
}
