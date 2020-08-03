package pipelines

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewPauseCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "pause",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}
	return cmd
}

// /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/pause
// put: pausePipeline
// pause a pipeline
