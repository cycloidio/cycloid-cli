package registry

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/plugins/registry/plugin"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "registry",
		Aliases: []string{"registries"},
		Short:   "Manage plugin registries",
	}

	cmd.AddCommand(
		NewListCommand(),
		NewGetCommand(),
		NewCreateCommand(),
		NewUpdateCommand(),
		NewDeleteCommand(),
		plugin.NewCommands(),
	)
	return cmd
}
