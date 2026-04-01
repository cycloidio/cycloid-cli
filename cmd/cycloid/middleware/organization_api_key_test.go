package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"github.com/cycloidio/cycloid-cli/internal/testcfg"
)

func TestAPIKeyCRUD(t *testing.T) {
	m := config.Middleware

	canonical := testcfg.RandomCanonical("test-apikey")
	name := ptr.Ptr("Test API Key")

	rule := &models.NewRule{Action: ptr.Ptr("*"), Effect: ptr.Ptr(models.NewRuleEffectAllow)}
	created, _, err := m.CreateAPIKey(config.Org, canonical, "test api key", "administrator", name, []*models.NewRule{rule})
	require.NoError(t, err, "CreateAPIKey should succeed")
	require.NotNil(t, created)

	defer func() {
		_, err := m.DeleteAPIKey(config.Org, canonical)
		require.NoError(t, err, "DeleteAPIKey should succeed")
	}()

	got, _, err := m.GetAPIKey(config.Org, canonical)
	require.NoError(t, err, "GetAPIKey should succeed")
	assert.Equal(t, canonical, *got.Canonical)

	list, _, err := m.ListAPIKeys(config.Org)
	require.NoError(t, err, "ListAPIKeys should succeed")
	assert.NotEmpty(t, list)
}
