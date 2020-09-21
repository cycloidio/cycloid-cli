package projects

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "update",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}
	return cmd
}

// /organizations/{organization_canonical}/projects/{project_canonical}
// put: updateProject
// Update the information of a project of the organization. If the project has some information on the fields which aren't required and they are not sent or set to their default vaules, which depend of their types, the information will be removed.
