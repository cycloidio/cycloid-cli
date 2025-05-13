package environments

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "environment",
		Aliases: []string{
			"e",
			"env",
			"environments",
		},
		Short: "Manage the environments of a project.",
	}

	cmd.AddCommand(
		NewDeleteCommand(),
		NewCreateCommand(),
		NewListCommand(),
		NewGetCommand(),
		NewUpdateCommand(),
	)

	return cmd
}
