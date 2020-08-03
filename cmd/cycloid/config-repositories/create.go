package configRepositories

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}

	return cmd
}

// /organizations/{organization_canonical}/config_repositories
// post: createConfigRepository
// Creates a config repository
