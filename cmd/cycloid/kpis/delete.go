package kpis

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
		Short: "delete a kpi",
		Example: `
	# delete a kpi
	cy --org my-org kpi delete \
		--canonical kpi
`,
		RunE:    delete,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredFlag(common.WithFlagCan, cmd)

	return cmd
}

func delete(cmd *cobra.Command, args []string) error {
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

	if err = m.DeleteKpi(org, can); err != nil {
		return errors.Wrap(err, "unable to delete kpi")
	}
	return nil
}
