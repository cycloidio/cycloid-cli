package components

import "github.com/spf13/cobra"

func NewConfigCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "config",
		Args: cobra.NoArgs,
		Aliases: []string{
			"cfg",
		},
		Short: "Manage component config",
	}

	cmd.AddCommand(
		NewGetComponentConfigCommand(),
	)

	return cmd
}
