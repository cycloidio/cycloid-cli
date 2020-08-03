package pipelines

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewTriggerBuildCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "trigger-build",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}
	return cmd
}

// /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds
// post: createBuild
// Create a new build for the job
