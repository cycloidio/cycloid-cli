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
		NewCreateCommand(),
		NewListCommand(),
		NewListWorkersCommand(),
		NewDeleteCommand(),
		NewListChildrensCommand(),
		NewCreateChildCommand(),
		NewGetCommand())
	return cmd
}
