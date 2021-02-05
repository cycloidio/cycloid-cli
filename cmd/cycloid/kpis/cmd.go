package kpis

import (
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
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
	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	cmd.AddCommand(NewCreateCommand(),
		NewDeleteCommand(),
		NewlistAvailableCommand(),
		NewListCommand())
	return cmd
}
