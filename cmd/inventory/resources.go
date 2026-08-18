package inventory

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/printer"
)

// resourcesTableOptions is the shared table view for inventory resources.
var resourcesTableOptions = printer.Options{
	Columns:    []string{"Name", "Provider", "Type", "Module", "Mode", "Label"},
	Identifier: "Name",
}

func newResourcesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "resources",
		Aliases: []string{"resource", "res"},
		Short:   "Manage inventory resources",
		Long:    "List Terraform resources tracked by the Cycloid inventory.",
		Args:    cobra.NoArgs,
	}

	cmd.AddCommand(
		newResourcesListCmd(),
	)

	return cmd
}
