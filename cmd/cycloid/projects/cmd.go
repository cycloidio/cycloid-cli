package projects

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "project",
		Aliases: []string{
			"p",
			"projects",
		},
		Short: "Manage the projects",
	}

	cmd.AddCommand(NewDeleteCommand(),
		NewCreateCommand(),
		NewListCommand(),
		NewGetCommand(),
		NewUpdateCommand(),
		NewListEnvCommand(),
	)

	return cmd
}
