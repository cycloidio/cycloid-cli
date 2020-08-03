package pipelines

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewDiffCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "diff",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}
	return cmd
}

// /organizations/{organization_canonical}/pipelines/{inpath_pipeline_name}/diff
// put: diffPipeline
// The diff between the provided pipeline configuration and the pipeline from the given name.
