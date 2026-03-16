package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListConfigRepositories(t *testing.T) {
	m := config.Middleware

	repos, _, err := m.ListConfigRepositories(config.Org)
	require.NoError(t, err)
	assert.NotEmpty(t, repos)
}

func TestGetConfigRepository(t *testing.T) {
	m := config.Middleware

	repo, _, err := m.GetConfigRepository(config.Org, *config.ConfigRepo.Canonical)
	require.NoError(t, err)
	assert.Equal(t, *config.ConfigRepo.Canonical, *repo.Canonical)
}
