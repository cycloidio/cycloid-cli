package configRepositories

import (
	"fmt"
	"time"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_config_repositories"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
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
	var err error
	var org string

	org, err = cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	params := organization_config_repositories.NewGetConfigRepositoriesParams()
	params.SetOrganizationCanonical(org)

	resp, err := api.OrganizationConfigRepositories.GetConfigRepositories(params, root.ClientCredentials())
	if err != nil {
		return err
	}
	fmt.Println(resp)
	fmt.Println("...")

	crs := resp.GetPayload()

	for _, cr := range crs.Data {
		fmt.Printf("id: %d    name: %s    url: %s    branch: %s    default: %t    credential_id: %d\n", *cr.ID, *cr.Name, *cr.URL, cr.Branch, *cr.Default, cr.CredentialID)
		fmt.Printf("created_at: %v    updated_at: %v\n", time.Unix(*cr.CreatedAt, 0), time.Unix(*cr.UpdatedAt, 0))
	}
	fmt.Println(resp)
	fmt.Printf("%+v\n", err)

	return nil
}
