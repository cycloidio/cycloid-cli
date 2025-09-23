package cyargs

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

func AddAPIKeyDescriptionFlag(cmd *cobra.Command) string {
	flagName := "description"
	cmd.Flags().StringP(flagName, "d", "", "set the description of the apikey")
	return flagName
}

func GetAPIKeyDescription(cmd *cobra.Command) (string, error) {
	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return "", err
	}

	return description, nil
}

func AddAPIKeyNameFlag(cmd *cobra.Command) string {
	flagName := "name"
	cmd.Flags().StringP(flagName, "n", "", "set the name of the apikey, will default to canonical if unset.")
	return flagName
}

func GetAPIKeyName(cmd *cobra.Command) (string, error) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return "", err
	}

	return name, nil
}

func AddAPIKeyCanonicalFlag(cmd *cobra.Command) string {
	flagName := "canonical"
	cmd.Flags().String(flagName, "", "set the canonical of the API key, if unset, will use the sanitized name.")
	cmd.RegisterFlagCompletionFunc(flagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
			org, err := GetOrg(cmd)
			if err != nil {
				return cobra.AppendActiveHelp(nil, "missing org parameter for completion"),
					cobra.ShellCompDirectiveNoFileComp
			}

			api := common.NewAPI()
			m := middleware.NewMiddleware(api)
			APIKeys, err := m.ListAPIKeys(org)
			if err != nil {
				return cobra.AppendActiveHelp(nil, "failed to list API Keys: "+err.Error()),
					cobra.ShellCompDirectiveNoFileComp
			}

			var APIKeysComp = make([]cobra.Completion, len(APIKeys))
			for index, apiKey := range APIKeys {
				if apiKey.Canonical != nil && strings.HasPrefix(*apiKey.Canonical, toComplete) {
					APIKeysComp[index] = cobra.CompletionWithDesc(*apiKey.Canonical, fmt.Sprintf("%s from %s: %s", *apiKey.Name, *apiKey.Owner.Username, apiKey.Description))
				}
			}

			return APIKeysComp, cobra.ShellCompDirectiveNoFileComp
		},
	)
	return flagName
}

func GetAPIKeyCanonical(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("canonical")
}

func AddAPIKeyRulesFlag(cmd *cobra.Command) string {
	flagName := "rules"
	cmd.Flags().String(flagName, "", `Set the permissions of the api, must be a JSON array of rules formatted like: '[{"action": "organization:**", "effect": "allow", "resources": []}]'.\nSee API Docs https://docs.cycloid.io/api/#tag/Organization-API-keys/operation/createAPIKey`)
	cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
		if toComplete == "" {
			return []cobra.Completion{
				cobra.CompletionWithDesc(`[{"action":"organization:**","effect":"allow","resources":[]}]`, "Admin permissions"),
			}, cobra.ShellCompDirectiveNoFileComp
		}

		return []string{}, cobra.ShellCompDirectiveNoFileComp
	})
	return flagName
}

func GetAPIKeyRules(cmd *cobra.Command) (string, error) {
	rules, err := cmd.Flags().GetString("rules")
	if err != nil {
		return "", err
	}

	return rules, nil
}
