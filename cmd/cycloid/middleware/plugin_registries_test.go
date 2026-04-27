package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListPluginRegistries(t *testing.T) {
	m := config.Middleware

	registries, _, err := m.ListPluginRegistries(config.Org)
	require.NoError(t, err)
	assert.NotNil(t, registries)
}

func TestPluginRegistryCRUD(t *testing.T) {
	m := config.Middleware

	// Create registry pointing to real plugin-registry service
	created, _, err := m.CreatePluginRegistry(config.Org, "test-registry", "http://192.168.10.13:4000")
	require.NoError(t, err)
	require.NotNil(t, created)
	assert.Equal(t, "test-registry", *created.Name)

	defer m.DeletePluginRegistry(config.Org, *created.ID)

	// Update
	updated, _, err := m.UpdatePluginRegistry(config.Org, *created.ID, "test-registry-updated")
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, "test-registry-updated", *updated.Name)

	// Create a plugin on the registry
	plugin, _, err := m.CreateRegistryPlugin(config.Org, *created.ID, "test-plugin")
	require.NoError(t, err)
	require.NotNil(t, plugin)
	require.NotNil(t, plugin.ID)
	assert.Equal(t, "test-plugin", *plugin.Name)

	defer m.DeleteRegistryPlugin(config.Org, *created.ID, *plugin.ID)

	// List versions (empty)
	versions, _, err := m.ListPluginVersions(config.Org, *created.ID, *plugin.ID)
	require.NoError(t, err)
	assert.NotNil(t, versions)

	// Delete plugin
	_, err = m.DeleteRegistryPlugin(config.Org, *created.ID, *plugin.ID)
	require.NoError(t, err)

	// Delete registry
	_, err = m.DeletePluginRegistry(config.Org, *created.ID)
	require.NoError(t, err)
}
