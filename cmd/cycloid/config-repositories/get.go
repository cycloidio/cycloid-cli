package configRepositories

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}

	return cmd
}

// /organizations/{organization_canonical}/config_repositories/{config_repository_id}
// get: getConfigRepository
// Return the Config Repository
