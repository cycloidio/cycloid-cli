package config_repositories

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "config-repository",
		Aliases: []string{
			"config-repo",
			"config-repositories",
		},
		Short: "Manage the catalog repositories",
	}

	cmd.AddCommand(NewCreateCommand(),
		NewUpdateCommand(),
		NewDeleteCommand(),
		NewListCommand(),
		NewGetCommand())

	return cmd
}
