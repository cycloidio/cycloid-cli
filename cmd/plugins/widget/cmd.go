package widget

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/plugins/widget/view"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "widget",
		Aliases: []string{"widgets"},
		Short:   "Manage org-level plugin widgets",
	}

	cmd.AddCommand(
		NewListCommand(),
		NewQueryCommand(),
		view.NewCommands(),
	)
	return cmd
}
