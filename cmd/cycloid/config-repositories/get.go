package configRepositories

import (
	"fmt"
	"time"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_config_repositories"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  getConfigRepository,
	}

	common.RequiredFlag(common.WithFlagID, cmd)

	return cmd
}

// /organizations/{organization_canonical}/config_repositories/{config_repository_id}
// get: getConfigRepository
// Return the Config Repository

func getConfigRepository(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	id, err := cmd.Flags().GetUint32("id")
	if err != nil {
		return err
	}

	params := organization_config_repositories.NewGetConfigRepositoryParams()
	params.SetOrganizationCanonical(org)
	params.SetConfigRepositoryID(id)

	resp, err := api.OrganizationConfigRepositories.GetConfigRepository(params, root.ClientCredentials())
	if err != nil {
		return err
	}

	cr := resp.GetPayload()

	fmt.Printf("id: %d    name: %s    url: %s    branch: %s    default: %t    credential_id: %d\n", *cr.Data.ID, *cr.Data.Name, *cr.Data.URL, cr.Data.Branch, *cr.Data.Default, cr.Data.CredentialID)
	fmt.Printf("created_at: %v    updated_at: %v\n", time.Unix(*cr.Data.CreatedAt, 0), time.Unix(*cr.Data.UpdatedAt, 0))

	fmt.Println(resp)
	fmt.Printf("%+v\n", err)

	return nil
}
