package stacks

import "github.com/spf13/cobra"

func NewFormsCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use: "forms",
		Aliases: []string{
			"form",
			"stackform",
			"stackforms",
		},
		Args:  cobra.NoArgs,
		Short: "use and manage Stackforms .forms.yml definition",
	}

	cmd.AddCommand(
		NewFormsValidateCommand(),
	)

	return cmd
}
