package catalogRepositories

import (
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
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
	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	cmd.AddCommand(NewCreateCommand(),
		NewUpdateCommand(),
		NewDeleteCommand(),
		NewListCommand(),
		NewGetCommand(),
		NewRefreshCommand())

	return cmd
}
