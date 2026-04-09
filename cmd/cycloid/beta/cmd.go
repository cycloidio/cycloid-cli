package beta

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/beta/config"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/beta/plugins"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "beta",
		Short: "Experimental commands.",
		Long: `Experimental Cycloid commands.
Those commands are feature in testing, retro-compatibility is not guaranteed.`,
	}

	cmd.AddCommand(
		config.NewCommands(),
		plugins.NewCommands(),
	)
	return cmd
}
