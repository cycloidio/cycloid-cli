package configrepositories

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use: "config-repository",
		Aliases: []string{
			"config-repo",
			"config-repositories",
		},
		Short: "Manage the config repositories",
	}

	cmd.AddCommand(NewCreateCommand(),
		NewUpdateCommand(),
		NewDeleteCommand(),
		NewListCommand(),
		NewGetCommand())

	return cmd
}
