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
		Short:   "not implemented yet",
		Long:    `not implemented yet`,
		PreRunE: internal.CheckAPIAndCLIVersion,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("not implemented yet")
			return nil
		},
	}
	return cmd
}
