package root

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/internal"
	"github.com/spf13/cobra"
)

func NewValidateFormCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "validate-form",
		Short:   "...",
		Long:    `........ . . .... .. .. ....`,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}
	return cmd
}

// /organizations/{organization_canonical}/forms/validate
// post: validateFormsFile
// Validate a forms file definition
