package projects

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "project",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
	}
	cmd.AddCommand(NewUpdateCommand(),
		NewDeleteCommand(),
		NewCreateCommand(),
		NewListCommand(),
		NewGetCommand())

	return cmd
}
