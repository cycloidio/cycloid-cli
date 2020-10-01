package configRepositories

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
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
	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	cmd.AddCommand(NewCreateCommand(),
		NewUpdateCommand(),
		NewDeleteCommand(),
		NewListCommand(),
		NewGetCommand())

	return cmd
}
