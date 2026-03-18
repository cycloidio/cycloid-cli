package roles

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var (
		example = `
	# Manage roles of my-org organization
	cy --org my-org roles [list|get|create|update|delete]
	`
		short = "Manage roles from the organization"
		long  = short
	)

	var cmd = &cobra.Command{
		Use:     "roles",
		Aliases: []string{"role"},
		Example: example,
		Short:   short,
		Long:    long,
	}

	cmd.AddCommand(
		NewListCommand(),
		NewGetCommand(),
		NewCreateCommand(),
		NewUpdateCommand(),
		NewDeleteCommand(),
	)

	return cmd
}
