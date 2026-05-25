package cyargs

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

func AddCloudAccountFlag(cmd *cobra.Command) string {
	flagName := "cloud-account"
	cmd.Flags().StringP(flagName, "a", "", "cloud account canonical")
	_ = cmd.RegisterFlagCompletionFunc(flagName, CompleteCloudAccountCanonical)
	return flagName
}

func GetCloudAccount(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("cloud-account")
}

func AddExistingCredentialFlag(cmd *cobra.Command) string {
	flagName := "credential"
	cmd.Flags().String(flagName, "", "canonical of an existing credential to wrap")
	_ = cmd.RegisterFlagCompletionFunc(flagName, CompleteCredentialCanonical)
	return flagName
}

func GetExistingCredential(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("credential")
}

func AddNewCredentialTypeFlag(cmd *cobra.Command) string {
	flagName := "new-credential"
	cmd.Flags().String(flagName, "", "create a new credential inline (ssh|aws|azure|azure_storage|gcp|custom|basic_auth|elasticsearch|swift)")
	return flagName
}

func GetNewCredentialType(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("new-credential")
}

func AddEnvironmentTypeCanonicalFlag(cmd *cobra.Command) string {
	flagName := "environment-type"
	cmd.Flags().String(flagName, "", "environment type canonical")
	_ = cmd.RegisterFlagCompletionFunc(flagName, CompleteEnvironmentType)
	return flagName
}

func GetEnvironmentTypeCanonical(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("environment-type")
}

func CompleteCloudAccount(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org parameter for completion"),
			cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)
	accounts, _, err := m.ListCloudAccounts(org)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list cloud accounts: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	completions := make([]cobra.Completion, 0, len(accounts))
	for _, account := range accounts {
		if account.Canonical != nil && strings.HasPrefix(*account.Canonical, toComplete) {
			name := ""
			if account.Name != nil {
				name = *account.Name
			}
			completions = append(completions, cobra.CompletionWithDesc(*account.Canonical, name))
		}
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}
