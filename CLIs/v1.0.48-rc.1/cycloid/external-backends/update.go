package externalBackends

import (
	"fmt"

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
	}

	return cmd
}

// Update will not be available on logs, only create/delete as it's too complexe

// /organizations/{organization_canonical}/external_backends/{external_backend_id}
// put: updateExternalBackend
// Update an External Backend
