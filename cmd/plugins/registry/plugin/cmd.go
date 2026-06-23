package plugin

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/plugins/registry/plugin/version"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "plugin",
		Aliases: []string{"plugins"},
		Short:   "Manage plugins within a registry",
	}

	cmd.AddCommand(
		NewListCommand(),
		NewGetCommand(),
		NewCreateCommand(),
		NewUpdateCommand(),
		NewDeleteCommand(),
		version.NewCommands(),
	)
	return cmd
}
