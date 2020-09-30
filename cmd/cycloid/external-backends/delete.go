package externalBackends

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/internal"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "delete an external backend configuration",
		Example: `
	# delete an existing external backend with ID 123
	cy --org my-org eb delete --id 123
`,
		RunE:    del,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}
	common.RequiredFlag(common.WithFlagID, cmd)
	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	return cmd
}

func del(cmd *cobra.Command, args []string) error {
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

	if err := m.DeleteExternalBackend(org, id); err != nil {
		return errors.Wrap(err, "unable to delete external backend")
	}

	return nil
}
