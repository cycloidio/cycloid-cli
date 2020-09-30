package root

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/internal"
	"github.com/spf13/cobra"
)

func NewValidateFormCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "validate-form",
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

// /organizations/{organization_canonical}/forms/validate
// post: validateFormsFile
// Validate a forms file definition
