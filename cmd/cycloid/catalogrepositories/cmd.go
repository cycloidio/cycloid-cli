package catalogrepositories

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "catalog-repository",
		Aliases: []string{
			"catalog-repo",
			"catalog-repositories",
			"cr",
		},
		Short: "Manage the catalog repositories",
	}

	cmd.AddCommand(NewCreateCommand(),
		NewUpdateCommand(),
		NewDeleteCommand(),
		NewListCommand(),
		NewGetCommand(),
		NewRefreshCommand())

	return cmd
}
