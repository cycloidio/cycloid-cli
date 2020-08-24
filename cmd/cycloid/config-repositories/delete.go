package configRepositories

import (
	"fmt"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_config_repositories"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/spf13/cobra"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  deleteConfigRepository,
	}

	common.RequiredFlag(common.WithFlagID, cmd)

	return cmd
}

// /organizations/{organization_canonical}/config_repositories/{config_repository_id}
// delete: deleteConfigRepository
// delete a Config Repositories

func deleteConfigRepository(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	id, err := cmd.Flags().GetUint32("id")
	if err != nil {
		return err
	}

	params := organization_config_repositories.NewDeleteConfigRepositoryParams()
	params.SetOrganizationCanonical(org)
	params.SetConfigRepositoryID(id)

	resp, err := api.OrganizationConfigRepositories.DeleteConfigRepository(params, root.ClientCredentials())
	if err != nil {
		return err
	}

	fmt.Println(resp)
	fmt.Printf("%+v\n", err)
	return nil
}
