package organizations

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewListWorkersCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-workers",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}
	return cmd
}

// /organizations/{organization_canonical}/workers
// get: getWorkers
// Get the workers that the authenticated user has access to.
