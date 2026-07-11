package environmenttypes

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/printer"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "environment-type",
		Aliases: []string{"environment-types", "env-type", "env-types"},
		Short:   "Manage environment types",
	}

	cmd.AddCommand(
		NewCreateCommand(),
		NewUpdateCommand(),
		NewGetCommand(),
		NewDeleteCommand(),
		NewListCommand(),
	)
	return cmd
}

var environmentTypeTableOptions = printer.Options{
	Columns:    []string{"Canonical", "Name", "Color"},
	Identifier: "Canonical",
}
