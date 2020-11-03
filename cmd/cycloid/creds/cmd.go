package creds

import (
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "credential",
		Aliases: []string{
			"creds",
			"cred",
			"credentials",
			"c",
		},
		Short: "Manage the credentials",
	}

	cmd.AddCommand(NewCreateCommand(),
		NewUpdateCommand(),
		NewDeleteCommand(),
		NewListCommand(),
		NewGetCommand())
	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	return cmd
}
