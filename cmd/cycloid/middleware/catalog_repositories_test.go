package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListCatalogRepositories(t *testing.T) {
	m := config.Middleware

	repos, _, err := m.ListCatalogRepositories(config.Org)
	require.NoError(t, err)
	assert.NotEmpty(t, repos)
}

func TestGetCatalogRepository(t *testing.T) {
	m := config.Middleware

	repo, _, err := m.GetCatalogRepository(config.Org, *config.CatalogRepo.Canonical)
	require.NoError(t, err)
	assert.Equal(t, *config.CatalogRepo.Canonical, *repo.Canonical)
}

func TestRefreshCatalogRepository(t *testing.T) {
	m := config.Middleware

	changes, _, err := m.RefreshCatalogRepository(config.Org, *config.CatalogRepo.Canonical)
	require.NoError(t, err)
	assert.NotNil(t, changes)
}
