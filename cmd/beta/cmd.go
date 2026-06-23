package beta

import (
	"github.com/spf13/cobra"

	bootstrapfirstorg "github.com/cycloidio/cycloid-cli/cmd/beta/bootstrap_first_org"
	"github.com/cycloidio/cycloid-cli/cmd/beta/config"
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
		bootstrapfirstorg.NewCommands(),
	)
	return cmd
}
