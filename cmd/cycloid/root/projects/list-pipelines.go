package projects

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewListPipelinesCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-pipelines",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}
	return cmd
}

// /organizations/{organization_canonical}/projects/{project_canonical}/pipelines
// get: getProjectPipelines
// Get the pipelines that the authenticated user has access to.
