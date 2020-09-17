package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewValidateFormCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "validate-form",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}
	return cmd
}

// /organizations/{organization_canonical}/forms/validate
// post: validateFormsFile
// Validate a forms file definition
