package cloudaccounts

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "get <cloud_account_canonical> [<cloud_account_canonical>...]",
		Short:             "Get cloud account details",
		RunE:              get,
		Args:              cobra.MinimumNArgs(1),
		ValidArgsFunction: cyargs.CompleteCloudAccount,
	}

	cyout.RegisterModel(cmd, models.CloudAccountDetail{})
	return cmd
}

func get(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	if len(args) == 1 {
		account, _, err := m.GetCloudAccount(org, args[0])
		return cyout.PrintWithOptions(cmd, account, err, "failed to get cloud account", cloudAccountTableOptions)
	}

	results := make([]*models.CloudAccountDetail, 0, len(args))
	for _, canonical := range args {
		account, _, err := m.GetCloudAccount(org, canonical)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "failed to get cloud account "+canonical, cloudAccountTableOptions)
		}
		results = append(results, account)
	}
	return cyout.PrintWithOptions(cmd, results, nil, "", cloudAccountTableOptions)
}
