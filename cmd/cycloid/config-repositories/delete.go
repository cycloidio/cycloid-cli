package configRepositories

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "delete a config repository",
		Example: `
	# delete a config repository with the canonical my-config-repo
	cy  --org my-org config-repository delete --canonical my-config-repo
`,
		RunE:    deleteConfigRepository,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredFlag(common.WithFlagCan, cmd)

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

	can, err := cmd.Flags().GetString("canonical")
	if err != nil {
		return err
	}

	if err := m.DeleteConfigRepository(org, can); err != nil {
		return errors.Wrap(err, "unable to delete config repository")
	}

	return nil
}
