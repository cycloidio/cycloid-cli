package stacks

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}
	return cmd
}

// /organizations/{organization_canonical}/service_catalogs
// get: getServiceCatalogs
// Return all the service catalogs