package projects

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
)

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:    "update",
		Short:  "...",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("not implemented yet")
		},
		PreRunE: internal.CheckAPIAndCLIVersion,
	}
	return cmd
}
