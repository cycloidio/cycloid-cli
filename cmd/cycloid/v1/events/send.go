package events

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewSendCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "send",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}

	return cmd
}


// '/organizations/{organization_canonical}/events':
// post: sendOrgEvent
// Send a event on the organization to be registered.
