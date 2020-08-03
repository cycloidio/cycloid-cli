package environments

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "environments",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
	}
	cmd.AddCommand(NewDeleteCommand(),
		NewListCommand(),
		// TODO do internal update project automated to create env
	// NewCreateCommand(),
	// NewGetCommand()
	)

	return cmd
}
