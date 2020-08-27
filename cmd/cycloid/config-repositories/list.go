package configRepositories

import (
	"fmt"
	"time"

	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  listConfigRepositories,
	}

	return cmd
}

// /organizations/{organization_canonical}/config_repositories
// get: getConfigRepositories
// Return all the config repositories
func listConfigRepositories(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	crs, err := m.ListConfigRepositories(org)

	for _, cr := range crs {
		fmt.Printf("id: %d    name: %s    url: %s    branch: %s    default: %t    credential_id: %d\n", *cr.ID, *cr.Name, *cr.URL, cr.Branch, *cr.Default, cr.CredentialID)
		fmt.Printf("created_at: %v    updated_at: %v\n", time.Unix(*cr.CreatedAt, 0), time.Unix(*cr.UpdatedAt, 0))
	}
	fmt.Println(crs)
	fmt.Printf("%+v\n", err)

	return nil
}
