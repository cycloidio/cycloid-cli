package pipelines

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewGetJobCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-job",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}
	return cmd
}

// /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}
// get: getJob
// Get the information of the job.
