package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewDeprecatedExampleCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "depcrecatedexample",
		Short: "...",
		Deprecated: "This command is deprecated, please use ...",
		Long:  `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}
	return cmd
}
