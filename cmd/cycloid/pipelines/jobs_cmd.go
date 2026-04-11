package pipelines

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/printer"
)

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
		NewJobLogsCommand(),
	)

	return cmd
}

var jobTableOptions = printer.Options{
	Columns:    []string{"Name", "Paused", "PipelineName", "CurrentBuildID"},
	Identifier: "Name",
	Transform: func(obj interface{}) map[string]string {
		j, ok := obj.(*models.Job)
		if !ok {
			return map[string]string{}
		}
		name := ""
		if j.Name != nil {
			name = *j.Name
		}
		paused := strconv.FormatBool(j.Paused)
		currentBuildID := ""
		switch {
		case j.NextBuild != nil && j.NextBuild.ID != nil:
			currentBuildID = strconv.FormatUint(*j.NextBuild.ID, 10)
		case j.FinishedBuild != nil && j.FinishedBuild.ID != nil:
			currentBuildID = strconv.FormatUint(*j.FinishedBuild.ID, 10)
		}
		return map[string]string{
			"Name":           name,
			"Paused":         paused,
			"PipelineName":   j.PipelineName,
			"CurrentBuildID": currentBuildID,
		}
	},
}
