package roles

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

func NewDeleteCommand() *cobra.Command {
	var (
		example = `
	# Remove a role from my-org organization
	cy --org my-org roles delete --id my-role-id
	`
		short = "Remove a user from the organization"
		long  = short
	)

	var cmd = &cobra.Command{
		Use:     "delete",
		Example: example,
		Short:   short,
		Long:    long,
		RunE:    deleteRole,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredFlag(common.WithFlagID, cmd)

	return cmd
}

func deleteRole(cmd *cobra.Command, args []string) error {
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

	err = m.DeleteRole(org, id)
	if err != nil {
		return errors.Wrapf(err, "unable to remove role: %d", id)
	}
	return nil
}
