package cy_args

import (
	"fmt"
	"slices"

	"github.com/cycloidio/cycloid-cli/client/models"
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

	return catalogRepository, err
}

func AddConfigRepositoryFlag(cmd *cobra.Command) {
	cmd.Flags().String("config-repository", "", "canonical of a config repository, if empty will use the default one in the current org.")
}

// GetConfigRepository return the config repository flag, if empty, will try to return
// the current org default config repository
func GetConfigRepository(cmd *cobra.Command, org string, m middleware.Middleware) (string, error) {
	configRepository, err := cmd.Flags().GetString("config-repository")
	if err != nil {
		return "", err
	}

	if configRepository == "" {
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

	return configRepository, err
}
