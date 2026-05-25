package cloudaccounts

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "delete <cloud_account_canonical> [<cloud_account_canonical>...]",
		Short:             "Delete cloud accounts",
		RunE:              deleteCloudAccount,
		Args:              cobra.MinimumNArgs(1),
		ValidArgsFunction: cyargs.CompleteCloudAccount,
	}

	return cmd
}

func deleteCloudAccount(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	for _, canonical := range args {
		if _, err := m.DeleteCloudAccount(org, canonical); err != nil {
			return cyout.Print(cmd, nil, err, "failed to delete cloud account "+canonical)
		}
	}
	return cyout.Print(cmd, map[string]string{"status": "ok"}, nil, "")
}
