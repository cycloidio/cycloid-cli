package pipelines

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:    "list",
		Hidden: true,
		Short:  "...",
		Long:   `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}
	return cmd
}

// /organizations/{organization_canonical}/pipelines
// get: getPipelines
// Get all the pipelines that the authenticated user has access to.
