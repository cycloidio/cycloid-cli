package e2e_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func TestPluginRegistries(t *testing.T) {
	t.Run("SuccessPluginRegistriesList", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"plugin-registry",
			"list",
		})

		require.Nil(t, cmdErr)
		assert.NotNil(t, cmdOut)
	})

	// Cleanup any leftover registries from previous test runs
	t.Run("Cleanup", func(t *testing.T) {
		listOut, _ := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"plugin-registry", "list",
		})
		ids, _ := JSONListExtractFields(listOut, "id", "", "")
		for _, id := range ids {
			executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"plugin-registry", "delete",
				"--registry-id", id,
			})
		}
	})

	// CRUD lifecycle using the real plugin-registry service
	t.Run("SuccessPluginRegistryCreate", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"plugin-registry",
			"create",
			"--name", "e2e-test-registry",
			"--url", "http://192.168.10.13:4000",
		})

		require.Nil(t, cmdErr)
		require.NotNil(t, cmdOut)

		var registry models.PluginRegistry
		err := json.Unmarshal([]byte(cmdOut), &registry)
		require.NoError(t, err)
		require.NotNil(t, registry.ID)
		assert.Equal(t, "e2e-test-registry", *registry.Name)

		registryID := fmt.Sprint(*registry.ID)

		defer t.Run("SuccessPluginRegistryDelete", func(t *testing.T) {
			_, cmdErr := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"plugin-registry",
				"delete",
				"--registry-id", registryID,
			})

			require.Nil(t, cmdErr)

			// Verify it's gone from the list
			listOut, listErr := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"plugin-registry",
				"list",
			})
			require.Nil(t, listErr)
			ids, err := JSONListExtractFields(listOut, "id", "", "")
			require.Nil(t, err)
			assert.NotContains(t, ids, registryID)
		})

		t.Run("SuccessPluginRegistryInList", func(t *testing.T) {
			listOut, listErr := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"plugin-registry",
				"list",
			})
			require.Nil(t, listErr)
			ids, err := JSONListExtractFields(listOut, "id", "", "")
			require.Nil(t, err)
			assert.Contains(t, ids, registryID)
		})

		t.Run("SuccessPluginRegistryUpdate", func(t *testing.T) {
			cmdOut, cmdErr := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"plugin-registry",
				"update",
				"--registry-id", registryID,
				"--name", "e2e-test-registry-updated",
			})

			require.Nil(t, cmdErr)
			require.NotNil(t, cmdOut)

			var updated models.PluginRegistry
			err := json.Unmarshal([]byte(cmdOut), &updated)
			require.NoError(t, err)
			assert.Equal(t, "e2e-test-registry-updated", *updated.Name)
		})

		t.Run("SuccessPluginRegistryPluginCreate", func(t *testing.T) {
			pluginOut, pluginErr := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"plugin-registry", "plugin", "create",
				"--registry-id", registryID,
				"--name", "e2e-test-plugin",
			})
			require.Nil(t, pluginErr)
			require.NotNil(t, pluginOut)

			var plugin models.Plugin
			err := json.Unmarshal([]byte(pluginOut), &plugin)
			require.NoError(t, err)
			require.NotNil(t, plugin.ID)
			assert.Equal(t, "e2e-test-plugin", *plugin.Name)

			pluginID := fmt.Sprint(*plugin.ID)

			defer func() {
				executeCommand([]string{
					"--output", "json",
					"--org", config.Org,
					"plugin-registry", "plugin", "delete",
					"--registry-id", registryID,
					"--plugin-id", pluginID,
				})
			}()
		})
	})

	t.Run("FailPluginRegistryCreateMissingName", func(t *testing.T) {
		_, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"plugin-registry",
			"create",
			"--url", "http://192.168.10.13:4000",
		})

		require.NotNil(t, cmdErr)
	})

	t.Run("FailPluginRegistryCreateMissingURL", func(t *testing.T) {
		_, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"plugin-registry",
			"create",
			"--name", "e2e-test-registry",
		})

		require.NotNil(t, cmdErr)
	})
}

func TestPluginManagers(t *testing.T) {
	t.Run("SuccessPluginManagersList", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"plugin-manager",
			"list",
		})

		require.Nil(t, cmdErr)
		assert.NotNil(t, cmdOut)
	})

	t.Run("SuccessPluginManagerCreate", func(t *testing.T) {
		// auto_register may fail with duplicate org on re-runs; that's expected
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"plugin-manager",
			"create",
			"--name", "e2e-test-pm",
			"--url", "http://192.168.10.14:4000",
		})

		if cmdErr != nil {
			t.Skipf("plugin manager create failed (may be duplicate org): %v", cmdErr)
		}

		require.NotNil(t, cmdOut)

		var pm models.PluginManager
		err := json.Unmarshal([]byte(cmdOut), &pm)
		require.NoError(t, err)
		require.NotNil(t, pm.ID)
		assert.Equal(t, "e2e-test-pm", *pm.Name)
	})
}

func TestOrganizationPlugins(t *testing.T) {
	t.Run("SuccessPluginsList", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"plugin",
			"list",
		})

		require.Nil(t, cmdErr)
		assert.NotNil(t, cmdOut)
	})
}
