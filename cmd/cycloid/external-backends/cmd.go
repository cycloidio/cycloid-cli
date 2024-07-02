package externalBackends

import (
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

	cmd.AddCommand(NewGetCommand(),
		NewDeleteCommand(),
		NewListCommand(),
		NewCreateCommand())

	return cmd
}
