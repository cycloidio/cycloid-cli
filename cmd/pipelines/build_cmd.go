package pipelines

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/printer"
)

func NewBuildsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "builds",
		Aliases: []string{
			"b",
			"build",
		},
		Short: "manage pipeline builds",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(
		NewBuildCreateCommand(),
		NewBuildListCommand(),
		NewBuildGetCommand(),
		NewBuildLogsCommand(),
	)
	return cmd
}

var buildTableOptions = printer.Options{
	Columns:    []string{"ID", "Name", "Status", "JobName", "StartTime"},
	Identifier: "ID",
}
