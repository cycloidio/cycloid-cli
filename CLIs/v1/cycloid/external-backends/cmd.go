package externalBackends

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "external-backends",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
	}

	cmd.AddCommand(NewUpdateCommand(),
		NewGetCommand(),
		NewDeleteCommand(),
		NewListCommand(),
		NewCreateCommand())

	return cmd
}
