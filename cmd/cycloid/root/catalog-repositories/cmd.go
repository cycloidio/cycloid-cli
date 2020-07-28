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

	cmd.AddCommand(NewUpdateCommand(),
		NewGetCommand(),
		NewDeleteCommand())

	return cmd
}
