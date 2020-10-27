package organizations

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "organization",
		Aliases: []string{
			"org",
			"organizations",
			"o",
		},
		Short: "Manage the organizations",
	}
	cmd.AddCommand(
		NewListCommand(),
		NewListWorkersCommand(),
		NewDeleteCommand(),
		NewGetCommand())
	return cmd
}
