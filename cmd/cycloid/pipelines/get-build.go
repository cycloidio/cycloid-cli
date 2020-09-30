package pipelines

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/internal"
	"github.com/spf13/cobra"
)

func NewGetBuildCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "get-build",
		Hidden:  true,
		Short:   "not implemented yet",
		Long:    `not implemented yet`,
		PreRunE: internal.CheckAPIAndCLIVersion,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("not implemented yet")
			return nil
		},
	}
	return cmd
}

// /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}
// get: getBuild
// Get the information of the build.
