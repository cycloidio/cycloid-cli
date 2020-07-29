package pipelines

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewPauseJobCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "pause-job",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}
	return cmd
}

// /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/pause
// put: pauseJob
// pause a job
