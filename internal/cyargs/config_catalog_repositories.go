package cyargs

import (
	"fmt"
	"slices"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/spf13/cobra"
)

func AddCatalogRepositoryFlag(cmd *cobra.Command) {
	cmd.Flags().String("catalog-repository", "", "canonical of a catalog repository")
}

func GetCatalogRepository(cmd *cobra.Command) (string, error) {
	catalogRepository, err := cmd.Flags().GetString("catalog-repository")
	if err != nil {
		return "", err
	}

	cmd.RegisterFlagCompletionFunc("catalog-repository", func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
		api := common.NewAPI()
		m := middleware.NewMiddleware(api)

		org, err := GetOrg(cmd)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "completion failed: "+err.Error()), cobra.ShellCompDirectiveError
		}

		stacks, err := m.ListCatalogRepositories(org)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "failed to list catalog repositories for completion in org '"+org+"': "+err.Error()),
				cobra.ShellCompDirectiveNoFileComp
		}

		var catalogRepositories = make([]cobra.Completion, len(stacks))
		for index, catalogRepository := range stacks {
			if catalogRepository.Canonical != nil {
				catalogRepositories[index] = cobra.CompletionWithDesc(*catalogRepository.Canonical, *catalogRepository.Name+" - branch: "+catalogRepository.Branch)
			}
		}

		return catalogRepositories, cobra.ShellCompDirectiveNoFileComp
	})

	return catalogRepository, err
}

func AddConfigRepositoryFlag(cmd *cobra.Command) string {
	flagName := "config-repository"
	cmd.Flags().String(flagName, "", "canonical of a config repository, if empty will use the default one in the current org.")
	cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
		api := common.NewAPI()
		m := middleware.NewMiddleware(api)

		org, err := GetOrg(cmd)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "completion failed: "+err.Error()), cobra.ShellCompDirectiveError
		}

		stacks, err := m.ListConfigRepositories(org)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "failed to list config repositories for completion in org '"+org+"' :"+err.Error()), cobra.ShellCompDirectiveError
		}

		var configRepositories = make([]string, len(stacks))
		for index, configRepository := range stacks {
			if configRepository.Canonical != nil {
				configRepositories[index] = cobra.CompletionWithDesc(*configRepository.Canonical, *configRepository.Name+" - branch: "+configRepository.Branch)
			}
		}

		return configRepositories, cobra.ShellCompDirectiveNoFileComp
	})
	return flagName
}

// GetConfigRepository return the config repository flag, if empty, will try to return
// the current org default config repository
func GetConfigRepository(cmd *cobra.Command) (string, error) {
	configRepository, err := cmd.Flags().GetString("config-repository")
	if err != nil {
		return "", err
	}

	return configRepository, err
}

func GetDefaultConfigRepository(cmd *cobra.Command) (string, error) {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := GetOrg(cmd)
	if err != nil {
		return "", fmt.Errorf("failed to get default config repository, missing org argument: %s", err)
	}

	configRepository, _ := GetConfigRepository(cmd)
	if configRepository != "" {
		return configRepository, nil
	}

	// TODO: This behavior will be pushed to backend
	// track issue: https://linear.app/cycloid/issue/BE-807/make-the-createproject-route-use-the-default-catalog-if
	catalogRepos, err := m.ListConfigRepositories(org)
	if err != nil {
		return "", fmt.Errorf("failed to get the default config repository: %s", err)
	}

	index := slices.IndexFunc(catalogRepos, func(c *models.ConfigRepository) bool {
		return *c.Default
	})
	if index == -1 {
		docURL := "https://docs.cycloid.io/reference/config-and-catalog-repository/"
		return "", fmt.Errorf("error: seems like your org '%s' does not have a default config repository, please add one using this doc: '%s'", org, docURL)
	}

	return *catalogRepos[index].Canonical, nil
}
