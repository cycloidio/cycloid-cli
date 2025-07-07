package apikey

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use: "api-key",
		Aliases: []string{
			"api-keys",
		},
		Example: `cy api-key [create|list|get|delete]`,
		Short:   "Manage organization API keys",
		Args:    cobra.NoArgs,
	}

	cmd.AddCommand(
		NewDeleteCommand(),
		NewGetCommand(),
		NewListCommand(),
		NewCreateCommand(),
	)

	return cmd
}
