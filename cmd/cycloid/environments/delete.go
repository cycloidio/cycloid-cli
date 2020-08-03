package environments

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}

	return cmd
}

// /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}
// delete: deleteProjectEnvironment
// Delete a project environment of the organization, and the project itself if it's the last environment.
