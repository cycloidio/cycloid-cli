package environments

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use: "environment",
		Aliases: []string{
			"e",
			"env",
			"environments",
		},
		Short: "Manage organization environments.",
	}

	cmd.AddCommand(
		NewDeleteCommand(),
		NewCreateCommand(),
		NewListCommand(),
		NewGetCommand(),
		NewUpdateCommand(),
		NewLinkCommand(),
		NewUnlinkCommand(),
	)

	return cmd
}
