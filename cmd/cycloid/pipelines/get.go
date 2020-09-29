package pipelines

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:    "get",
		Short:  "...",
		Hidden: true,
		Long:   `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}
	return cmd
}

// /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}
// get: getPipeline
// Get the configuration of the pipeline.
