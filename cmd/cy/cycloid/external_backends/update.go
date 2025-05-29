package external_backends

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cy/cycloid/internal"
)

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "update",
		Args:    cobra.NoArgs,
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

// Update will not be available on logs, only create/delete as it's too complexe

// /organizations/{organization_canonical}/external_backends/{external_backend_id}
// put: updateExternalBackend
// Update an External Backend
