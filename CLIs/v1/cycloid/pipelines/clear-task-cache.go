package pipelines

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewClearTaskCacheCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "clear-task-cache",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}
	return cmd
}

// /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/tasks/{step_name}/cache
// delete: clearTaskCache
// Clear task cache
