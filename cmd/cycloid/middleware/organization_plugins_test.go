package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListPlugins(t *testing.T) {
	m := config.Middleware

	plugins, _, err := m.ListPlugins(config.Org)
	require.NoError(t, err)
	assert.NotNil(t, plugins)
}

func TestGetPlugin(t *testing.T) {
	m := config.Middleware

	// No plugins installed, expect 404
	_, _, err := m.GetPlugin(config.Org, 999999)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "404")
}

func TestGetPluginLogs(t *testing.T) {
	m := config.Middleware

	// No plugins installed, expect 404
	_, _, err := m.GetPluginLogs(config.Org, 999999)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "404")
}

func TestListPluginWidgets(t *testing.T) {
	m := config.Middleware

	widgets, _, err := m.ListPluginWidgets(config.Org, "dashboard")
	require.NoError(t, err)
	assert.NotNil(t, widgets)
}
