package cloudaccounts

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List cloud accounts",
		RunE:  list,
		Args:  cobra.NoArgs,
	}

	cyout.RegisterModel(cmd, models.CloudAccountDetail{})
	return cmd
}

func list(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	accounts, _, err := m.ListCloudAccounts(org)
	return cyout.PrintWithOptions(cmd, accounts, err, "failed to list cloud accounts", cloudAccountTableOptions)
}
