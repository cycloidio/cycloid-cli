package pipelines

import "github.com/spf13/cobra"

func NewJobsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "job",
		Aliases: []string{
			"jobs",
			"j",
		},
		Short: "Manage pipeline jobs",
	}

	cmd.AddCommand(
		NewJobsGetCommand(),
		NewJobsListCommand(),
		NewJobsPauseCommand(),
		NewJobsUnpauseCommand(),
	)

	return cmd
}
