package externalBackends

import (
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "external-backend",
		Aliases: []string{
			"external-backends",
			"eb",
		},
		Short: "manage external backends",
	}
	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	cmd.AddCommand(NewGetCommand(),
		NewDeleteCommand(),
		NewListCommand(),
		NewCreateCommand())

	return cmd
}
