package externalBackends

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/internal"
	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:    "get",
		Short:  "...",
		Hidden: true,
		Long:   `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	return cmd
}

// /organizations/{organization_canonical}/external_backends/{external_backend_id}
// get: getExternalBackend
// Get the external backend
