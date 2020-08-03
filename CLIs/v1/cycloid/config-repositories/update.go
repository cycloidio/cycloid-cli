package configRepositories

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "update",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}

	return cmd
}

// /organizations/{organization_canonical}/config_repositories/{config_repository_id}
// put: updateConfigRepository
// Update a config repository
