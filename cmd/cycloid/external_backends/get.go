package external_backends

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:    "get",
		Args:   cobra.NoArgs,
		Hidden: true,
		Short:  "not implemented yet",
		Long:   `not implemented yet`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("not implemented yet")
			return nil
		},
	}

	return cmd
}

// /organizations/{organization_canonical}/external_backends/{external_backend_id}
// get: getExternalBackend
// Get the external backend
