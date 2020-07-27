package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewStatusCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "status",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}
	return cmd
}

// /status
// get: getStatus
// Get the status of the Cycloid's services.
