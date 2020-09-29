package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewStatusCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:    "status",
		Hidden: true,
		Short:  "...",
		Long:   `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("not implemented yet")
		},
	}
	return cmd
}
