package configRepositories

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "delete a config repository",
		Example: `
	# delete a catalog repository with the ID 123
	cy  --org my-org config-repository delete --id 123
`,
	}

	common.RequiredFlag(common.WithFlagID, cmd)

	return cmd
}

// /organizations/{organization_canonical}/config_repositories/{config_repository_id}
// delete: deleteConfigRepository
// delete a Config Repositories

func deleteConfigRepository(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	id, err := cmd.Flags().GetUint32("id")
	if err != nil {
		return err
	}

	if err := m.DeleteConfigRepository(org, id); err != nil {
		return errors.Wrap(err, "unable to delete config repository")
	}

	return nil
}
