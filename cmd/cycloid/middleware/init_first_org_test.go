package middleware_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitFirstOrgIdempotent(t *testing.T) {
	m := config.Middleware

	licenceKey, ok := os.LookupEnv("API_LICENCE_KEY")
	require.True(t, ok, "API_LICENCE_KEY must be set for bootstrap tests")

	apiKeyCanonical := "admin-init-first-org-test"
	result, _, err := m.InitFirstOrg(
		config.Org,
		"administrator",
		"administrator",
		"admin@cycloid.io",
		"cycloidadmin",
		licenceKey,
		&apiKeyCanonical,
	)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, config.Org, result.Org)
	assert.NotEmpty(t, result.Token)
	require.NotNil(t, result.APIKey)
	assert.NotEmpty(t, *result.APIKey)

	resultAgain, _, err := m.InitFirstOrg(
		config.Org,
		"administrator",
		"administrator",
		"admin@cycloid.io",
		"cycloidadmin",
		licenceKey,
		&apiKeyCanonical,
	)
	require.NoError(t, err)
	require.NotNil(t, resultAgain)
	assert.Equal(t, config.Org, resultAgain.Org)
	assert.NotEmpty(t, resultAgain.Token)
}
