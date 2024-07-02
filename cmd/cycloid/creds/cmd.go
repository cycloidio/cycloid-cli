package creds

import (
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

	return cmd
}
