package stacks

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "stack",
		Aliases: []string{
			"s",
			"stacks",
		},
		Short: "Manage the stacks",
	}

	cmd.AddCommand(
		NewListCommand(),
		NewGetCommand(),
		NewUpdateCommand(),
		NewCreateFromBlueprintCommand(),
		NewFormsCommands(),
		NewStacksGetComponentStackConfig(),
	)

	return cmd
}
