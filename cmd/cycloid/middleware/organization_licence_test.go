package middleware_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetLicence(t *testing.T) {
	m := config.Middleware

	licence, _, err := m.GetLicence(config.Org)
	require.NoError(t, err)
	require.NotNil(t, licence)
	require.NotNil(t, licence.Key)
	assert.NotEmpty(t, *licence.Key)
}

func TestActivateLicenceOverwrite(t *testing.T) {
	m := config.Middleware

	licenceKey, ok := os.LookupEnv("API_LICENCE_KEY")
	require.True(t, ok, "API_LICENCE_KEY must be set for licence tests")

	_, err := m.ActivateLicence(config.Org, licenceKey)
	require.NoError(t, err)

	got, _, err := m.GetLicence(config.Org)
	require.NoError(t, err)
	require.NotNil(t, got.Key)
	assert.Equal(t, licenceKey, *got.Key)

	_, err = m.ActivateLicence(config.Org, licenceKey)
	require.NoError(t, err)

	gotAgain, _, err := m.GetLicence(config.Org)
	require.NoError(t, err)
	require.NotNil(t, gotAgain.Key)
	assert.Equal(t, licenceKey, *gotAgain.Key)
}
