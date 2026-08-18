package licence

import "github.com/spf13/cobra"

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "licence",
		Aliases: []string{"license"},
		Short:   "Manage the organization's Cycloid licence",
		Args:    cobra.NoArgs,
	}
	cmd.AddCommand(
		NewActivateCommand(),
		NewGetCommand(),
	)
	return cmd
}
