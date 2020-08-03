package organizations

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "org",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
	}
	cmd.AddCommand(
		NewListCommand(),
		NewListWorkersCommand(),
		NewGetCommand())
	return cmd
}
