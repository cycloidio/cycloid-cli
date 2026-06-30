package inventory

import (
	"github.com/spf13/cobra"
)

// NewCommands returns the root "inventory" command with all subcommands attached.
func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "inventory",
		Aliases: []string{"inv"},
		Short:   "Manage the inventory",
		Long:    "Browse Terraform state outputs and resources tracked by the Cycloid inventory.",
	}

	cmd.AddCommand(
		newOutputsCmd(),
		newResourcesCmd(),
	)

	return cmd
}
