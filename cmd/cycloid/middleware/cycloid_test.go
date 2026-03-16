package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAppVersion(t *testing.T) {
	m := config.Middleware

	version, _, err := m.GetAppVersion()
	require.NoError(t, err)
	assert.NotNil(t, version)
}

func TestGetStatus(t *testing.T) {
	m := config.Middleware

	status, _, err := m.GetStatus()
	require.NoError(t, err)
	assert.NotNil(t, status)
}
