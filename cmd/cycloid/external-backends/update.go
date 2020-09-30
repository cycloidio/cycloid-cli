package externalBackends

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/internal"
	"github.com/spf13/cobra"
)

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:    "update",
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

// Update will not be available on logs, only create/delete as it's too complexe

// /organizations/{organization_canonical}/external_backends/{external_backend_id}
// put: updateExternalBackend
// Update an External Backend
