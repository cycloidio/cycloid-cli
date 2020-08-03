package pipelines

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewListJobsCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-jobs",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}
	return cmd
}

// /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs
// get: getJobs
// Get the jobs of the pipeline that the authenticated user has access to.
