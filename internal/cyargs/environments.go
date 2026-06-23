package cyargs

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/gen/models"
	"github.com/cycloidio/cycloid-cli/utils/ptr"
)

func AddEnvironmentTypeFlag(cmd *cobra.Command) string {
	flagName := "type"
	cmd.Flags().String(flagName, "", "environment type canonical (e.g. production, staging, development)")
	_ = cmd.RegisterFlagCompletionFunc(flagName, CompleteEnvironmentType)
	return flagName
}

func GetEnvironmentType(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("type")
}

func CompleteEnvironmentType(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org parameter for completion"),
			cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)
	types, _, err := m.ListEnvironmentTypes(org)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list environment types: "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	completions := make([]cobra.Completion, 0, len(types))
	for _, envType := range types {
		if envType.Canonical != nil && strings.HasPrefix(*envType.Canonical, toComplete) {
			name := ""
			if envType.Name != nil {
				name = *envType.Name
			}
			completions = append(completions, cobra.CompletionWithDesc(*envType.Canonical, name))
		}
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}

func AddCloudAccountCanonicalsFlag(cmd *cobra.Command) string {
	flagName := "cloud-account"
	cmd.Flags().StringArray(flagName, nil, "cloud account canonical to link (repeatable); pass \"\" to unlink all")
	_ = cmd.RegisterFlagCompletionFunc(flagName, CompleteCloudAccountCanonical)
	return flagName
}

func GetCloudAccountCanonicals(cmd *cobra.Command) ([]string, error) {
	return cmd.Flags().GetStringArray("cloud-account")
}

func CompleteCloudAccountCanonical(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "missing org parameter for completion"),
			cobra.ShellCompDirectiveNoFileComp
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)
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

func AddEnvironmentVariablesFlag(cmd *cobra.Command) string {
	flagName := "variable"
	cmd.Flags().StringArray(flagName, nil, "environment variable as KEY=VALUE (repeatable)")
	return flagName
}

func AddEnvironmentVariablesFileFlag(cmd *cobra.Command) string {
	flagName := "variables-file"
	cmd.Flags().StringArray(flagName, nil, "path to a YAML file containing environment variables")
	cmd.MarkFlagFilename(flagName, "yaml", "yml")
	return flagName
}

func GetEnvironmentVariables(cmd *cobra.Command) ([]*models.EnvironmentVariableItem, error) {
	var items []*models.EnvironmentVariableItem

	files, err := cmd.Flags().GetStringArray("variables-file")
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		fileItems, err := parseEnvironmentVariablesFile(file)
		if err != nil {
			return nil, err
		}
		items = append(items, fileItems...)
	}

	variables, err := cmd.Flags().GetStringArray("variable")
	if err != nil {
		return nil, err
	}
	for _, variable := range variables {
		item, err := parseEnvironmentVariablePair(variable)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func parseEnvironmentVariablesFile(path string) ([]*models.EnvironmentVariableItem, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read variables file %q: %w", path, err)
	}

	var items []*models.EnvironmentVariableItem
	if err := yaml.Unmarshal(content, &items); err != nil {
		return nil, fmt.Errorf("failed to parse variables file %q: %w", path, err)
	}
	return items, nil
}

func parseEnvironmentVariablePair(pair string) (*models.EnvironmentVariableItem, error) {
	key, value, ok := strings.Cut(pair, "=")
	if !ok || key == "" {
		return nil, fmt.Errorf("invalid variable %q, expected KEY=VALUE", pair)
	}
	return &models.EnvironmentVariableItem{
		Key:   ptr.Ptr(key),
		Type:  ptr.Ptr(models.EnvironmentVariableItemTypeString),
		Value: value,
	}, nil
}

func AddEnvironmentOwnerFlag(cmd *cobra.Command) string {
	return AddOwnerFlag(cmd)
}

func GetEnvironmentOwner(cmd *cobra.Command) (string, error) {
	return GetOwner(cmd)
}
