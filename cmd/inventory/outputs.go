package inventory

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/printer"
)

// outputsTableOptions is the shared table view for inventory outputs.
var outputsTableOptions = printer.Options{
	Columns:    []string{"Key", "Type", "Sensitive", "Pinned", "Description"},
	Identifier: "Key",
}

func newOutputsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "outputs",
		Aliases: []string{"output", "out"},
		Short:   "Manage inventory outputs",
		Long:    "List and fetch Terraform state outputs tracked by the Cycloid inventory.",
		Args:    cobra.NoArgs,
	}

	cmd.AddCommand(
		newOutputsListCmd(),
		newOutputsGetCmd(),
	)

	return cmd
}
