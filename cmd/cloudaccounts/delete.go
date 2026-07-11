package cloudaccounts

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
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
	m := apiclient.NewAPIClient(api)

	for _, canonical := range args {
		if _, err := m.DeleteCloudAccount(org, canonical); err != nil {
			return cyout.Print(cmd, nil, err, "failed to delete cloud account "+canonical)
		}
		fmt.Fprintf(cmd.OutOrStdout(), "deleted %s\n", canonical)
	}
	return nil
}
