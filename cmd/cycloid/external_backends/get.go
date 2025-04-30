package externalBackends

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
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

// /organizations/{organization_canonical}/external_backends/{external_backend_id}
// get: getExternalBackend
// Get the external backend
