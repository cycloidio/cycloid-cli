package component

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/plugins/component/widget"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "component",
		Aliases: []string{"components"},
		Short:   "Manage component-plugin relations",
	}

	cmd.AddCommand(
		NewListCommand(),
		NewRelationSetCommand(),
		widget.NewCommands(),
	)
	return cmd
}
