package components

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "components",
		Args: cobra.NoArgs,
		Aliases: []string{
			"component",
			"comp",
		},
		Short: "Manage components.",
	}

	cmd.AddCommand(
		NewGetComponentCommand(),
		NewGetComponentsCommand(),
		NewCreateComponentCommand(),
		NewUpdateComponentCommand(),
		NewDeleteComponentCommand(),
		NewMigrateCommand(),
		NewConfigCommands(),
	)

	return cmd
}

var componentTableOptions = printer.Options{
	Columns:    []string{"Canonical", "Name", "Description", "StackRef", "UseCase", "Version"},
	Identifier: "Canonical",
	Transform: func(obj interface{}) map[string]string {
		c, ok := obj.(*models.Component)
		if !ok {
			return map[string]string{}
		}
		canonical := ""
		if c.Canonical != nil {
			canonical = *c.Canonical
		}
		name := ""
		if c.Name != nil {
			name = *c.Name
		}
		stackRef := ""
		version := ""
		if c.ServiceCatalog != nil {
			if c.ServiceCatalog.Ref != nil {
				stackRef = *c.ServiceCatalog.Ref
			}
			if c.ServiceCatalog.Version != nil {
				version = *c.ServiceCatalog.Version
			}
		}
		useCase := c.UseCase
		return map[string]string{
			"Canonical":   canonical,
			"Name":        name,
			"Description": c.Description,
			"StackRef":    stackRef,
			"UseCase":     useCase,
			"Version":     version,
		}
	},
}
