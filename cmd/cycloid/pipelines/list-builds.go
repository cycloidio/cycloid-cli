package pipelines

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewGetListBuildsCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-builds",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}
	return cmd
}

// /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds
// get: getBuilds
// Get the pipeline job's builds that the authenticated user has access to.
