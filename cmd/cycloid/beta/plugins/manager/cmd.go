package manager

import "github.com/spf13/cobra"

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "manager",
		Aliases: []string{"managers"},
		Short:   "[beta] Manage plugin managers",
	}

	cmd.AddCommand(
		NewListCommand(),
		NewGetCommand(),
		NewCreateCommand(),
		NewAcceptCommand(),
		NewRejectCommand(),
		NewDeleteCommand(),
	)
	return cmd
}
