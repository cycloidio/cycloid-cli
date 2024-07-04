package kpis

import (
	"github.com/spf13/cobra"
)

func NewCommands() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "kpi",
		Aliases: []string{
			"kpis",
		},
		Short: "Manage the kpis",
	}

	cmd.AddCommand(NewCreateCommand(),
		NewDeleteCommand(),
		NewListCommand())
	return cmd
}
