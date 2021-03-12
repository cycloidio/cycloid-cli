package pipelines

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "get",
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
