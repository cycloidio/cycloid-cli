package pipelines

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/internal"
	"github.com/spf13/cobra"
)

func NewGetBuildCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:    "get-build",
		Hidden: true,
		Short:  "...",
		Long:   `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
		PreRunE: internal.CheckAPIAndCLIVersion,
	}
	return cmd
}

// /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}
// get: getBuild
// Get the information of the build.
