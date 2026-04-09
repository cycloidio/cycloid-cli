package registry

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/beta/plugins/registry/plugin"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "registry",
		Aliases: []string{"registries"},
		Short:   "[beta] Manage plugin registries",
	}

	cmd.AddCommand(
		NewListCommand(),
		NewGetCommand(),
		NewAddCommand(),
		NewDeleteCommand(),
		plugin.NewCommands(),
	)
	return cmd
}
