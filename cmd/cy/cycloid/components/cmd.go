package components

import "github.com/spf13/cobra"

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "components",
		Args: cobra.NoArgs,
		Aliases: []string{
			"comp",
			"component",
		},
		Short: "Manage components.",
	}

	cmd.AddCommand(
		NewGetComponentCommand(),
		NewGetComponentsCommand(),
		NewCreateComponentCommand(),
		NewUpdateComponentCommand(),
		NewDeleteComponentCommand(),
		NewMigrateCommand(),
		NewConfigCommands(),
	)

	return cmd
}
