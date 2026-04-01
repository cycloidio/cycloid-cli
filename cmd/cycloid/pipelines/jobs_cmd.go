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
		Example: `# list jobs from a pipeline
cy pipeline job list --project my-project --env my-env --component my-component --pipeline my-pipeline

# get a specific job from a pipeline
cy pipeline job get --project my-project --env my-env --component my-component --pipeline my-pipeline --job my-job`,
		Args: cobra.NoArgs,
	}

	cmd.AddCommand(
		NewJobsGetCommand(),
		NewJobsListCommand(),
		NewJobsPauseCommand(),
		NewJobsUnpauseCommand(),
	)

	return cmd
}
