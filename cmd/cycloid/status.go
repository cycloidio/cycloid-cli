package root

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/internal"
	"github.com/spf13/cobra"
)

func NewStatusCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "status",
		Hidden:  true,
		Short:   "...",
		Long:    `........ . . .... .. .. ....`,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("not implemented yet")
		},
	}
	return cmd
}
