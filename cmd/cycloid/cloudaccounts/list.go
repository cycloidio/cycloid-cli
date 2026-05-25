package cloudaccounts

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List cloud accounts",
		RunE:    list,
		Args:    cobra.NoArgs,
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
	m := middleware.NewMiddleware(api)

	accounts, _, err := m.ListCloudAccounts(org)
	return cyout.PrintWithOptions(cmd, accounts, err, "failed to list cloud accounts", cloudAccountTableOptions)
}

func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "get [cloud_account_canonical...]",
		Short:             "Get cloud account details",
		RunE:              get,
		Args:              cyargs.RequireArgsOrFlag("cloud-account"),
		ValidArgsFunction: cyargs.CompleteCloudAccount,
	}

	cyargs.AddCloudAccountFlag(cmd)
	cyout.RegisterModel(cmd, models.CloudAccountDetail{})
	return cmd
}

func get(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	if flag, err := cyargs.GetCloudAccount(cmd); err != nil {
		return err
	} else if flag != "" {
		args = appendUniqueArg(args, flag)
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

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

func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "delete [cloud_account_canonical...]",
		Short:             "Delete cloud accounts",
		RunE:              deleteCloudAccount,
		Args:              cyargs.RequireArgsOrFlag("cloud-account"),
		ValidArgsFunction: cyargs.CompleteCloudAccount,
	}

	cyargs.AddCloudAccountFlag(cmd)
	return cmd
}

func deleteCloudAccount(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	if flag, err := cyargs.GetCloudAccount(cmd); err != nil {
		return err
	} else if flag != "" {
		args = appendUniqueArg(args, flag)
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	for _, canonical := range args {
		if _, err := m.DeleteCloudAccount(org, canonical); err != nil {
			return cyout.Print(cmd, nil, err, "failed to delete cloud account "+canonical)
		}
	}
	return nil
}

func appendUniqueArg(args []string, value string) []string {
	for _, existing := range args {
		if existing == value {
			return args
		}
	}
	return append(args, value)
}
