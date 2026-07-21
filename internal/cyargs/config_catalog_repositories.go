package cyargs

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
)

func AddCatalogRepositoryFlag(cmd *cobra.Command) string {
	flagName := "catalog-repository"
	cmd.Flags().String(flagName, "", "canonical of a catalog repository")
	cmd.RegisterFlagCompletionFunc("catalog-repository", CompleteCatalogRepository)
	return flagName
}

// AddCatalogRepoCanonicalFlag registers optional --canonical on catalog-repo create (identity for upsert).
func AddCatalogRepoCanonicalFlag(cmd *cobra.Command) string {
	const flagName = "canonical"
	cmd.Flags().String(flagName, "", "catalog repository canonical; if omitted, derived from --name")
	_ = cmd.RegisterFlagCompletionFunc(flagName, CompleteCatalogRepository)
	return flagName
}

// GetCatalogRepoCanonical returns the optional catalog repository canonical from create (empty if unset).
func GetCatalogRepoCanonical(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("canonical")
}

// GetConfigRepoCanonical returns the optional config repository canonical (empty if unset).
func GetConfigRepoCanonical(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("canonical")
}

func GetCatalogRepository(cmd *cobra.Command) (string, error) {
	catalogRepository, err := cmd.Flags().GetString("catalog-repository")
	if err != nil {
		return "", err
	}
	return catalogRepository, err
}

func CompleteCatalogRepository(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "completion failed: "+err.Error()), cobra.ShellCompDirectiveError
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	stacks, _, err := m.ListCatalogRepositories(org)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list catalog repositories for completion in org '"+org+"': "+err.Error()),
			cobra.ShellCompDirectiveNoFileComp
	}

	catalogRepositories := make([]cobra.Completion, len(stacks))
	for index, catalogRepository := range stacks {
		if catalogRepository.Canonical != nil {
			catalogRepositories[index] = cobra.CompletionWithDesc(*catalogRepository.Canonical, *catalogRepository.Name+" - branch: "+catalogRepository.Branch)
		}
	}

	return catalogRepositories, cobra.ShellCompDirectiveNoFileComp
}

func AddConfigRepositoryFlag(cmd *cobra.Command) string {
	flagName := "config-repository"
	cmd.Flags().String(flagName, "", "canonical of a config repository, if empty will use the default one in the current org.")
	cmd.RegisterFlagCompletionFunc(flagName, CompleteConfigRepository)
	return flagName
}

// AddConfigRepoCanonicalFlag registers --canonical for config-repository get/delete/update commands.
func AddConfigRepoCanonicalFlag(cmd *cobra.Command) string {
	const flagName = "canonical"
	cmd.Flags().String(flagName, "", "config repository canonical")
	_ = cmd.RegisterFlagCompletionFunc(flagName, CompleteConfigRepository)
	return flagName
}

func CompleteConfigRepository(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	org, err := GetOrg(cmd)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "completion failed: "+err.Error()), cobra.ShellCompDirectiveError
	}

	stacks, _, err := m.ListConfigRepositories(org)
	if err != nil {
		return cobra.AppendActiveHelp(nil, "failed to list config repositories for completion in org '"+org+"': "+err.Error()), cobra.ShellCompDirectiveError
	}

	configRepositories := make([]cobra.Completion, len(stacks))
	for index, configRepository := range stacks {
		if configRepository.Canonical != nil {
			configRepositories[index] = cobra.CompletionWithDesc(*configRepository.Canonical, *configRepository.Name+" - branch: "+configRepository.Branch)
		}
	}

	return configRepositories, cobra.ShellCompDirectiveNoFileComp
}

// GetConfigRepository returns the config-repository flag value
func GetConfigRepository(cmd *cobra.Command) (string, error) {
	configRepository, err := cmd.Flags().GetString("config-repository")
	if err != nil {
		return "", err
	}

	return configRepository, err
}

func GetDefaultConfigRepository(cmd *cobra.Command) (string, error) {
	configRepository, _ := GetConfigRepository(cmd)
	return configRepository, nil
}
