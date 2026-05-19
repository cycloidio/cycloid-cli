package middleware_test

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/pkg/testcfg"
)

// pluginsMiddlewareEnabled reports whether the plugin middleware tests should run.
func pluginsMiddlewareEnabled() bool {
	return os.Getenv("CY_TEST_E2E_PLUGINS") == "1"
}

// uniquePluginImage generates a unique semver-tagged plugin image URL for use with
// CreatePluginVersion. The plugin-registry enforces a global unique constraint on
// version URLs, so each version creation requires a different docker image tag.
func uniquePluginImage(t *testing.T) string {
	t.Helper()
	tag := fmt.Sprintf("1.0.%d", rand.Intn(999999))
	local := fmt.Sprintf("localhost:5000/plugin-hello-world:%s", tag)
	api := fmt.Sprintf("docker-registry:5000/plugin-hello-world:%s", tag)

	run := func(name string, args ...string) {
		t.Helper()
		cmd := exec.Command(name, args...)
		if err := cmd.Run(); err != nil {
			t.Fatalf("uniquePluginImage: %q %v failed: %v", name, args, err)
		}
	}
	run("docker", "tag", "cycloid/plugin-hello-world:latest", local)
	run("docker", "push", local)
	return api
}

func TestPluginManagersCRUD(t *testing.T) {
	if !pluginsMiddlewareEnabled() {
		t.Skip("set CY_TEST_E2E_PLUGINS=1 to run plugin middleware tests")
	}

	m := config.Middleware
	org := config.Org

	t.Run("List", func(t *testing.T) {
		managers, resp, err := m.ListPluginManagers(org)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NotNil(t, managers)
	})

	t.Run("AcceptTestPluginManager", func(t *testing.T) {
		mgr, err := config.AcceptPluginManager(t)
		require.NoError(t, err)
		require.NotNil(t, mgr)
		// API returns "invite_accepted" once accepted (not "accepted").
		assert.Contains(t, []string{"accepted", "invite_accepted"}, *mgr.InviteStatus)
	})
}

func TestPluginRegistryCRUD(t *testing.T) {
	if !pluginsMiddlewareEnabled() {
		t.Skip("set CY_TEST_E2E_PLUGINS=1 to run plugin middleware tests")
	}

	m := config.Middleware
	org := config.Org
	name := testcfg.RandomCanonical("mw-registry")
	regURL := "http://plugin-registry:4000"

	reg, _, err := m.CreatePluginRegistry(org, name, regURL)
	require.NoError(t, err)
	require.NotNil(t, reg)
	assert.Equal(t, name, *reg.Name)

	defer func() {
		if _, delErr := m.DeletePluginRegistry(org, *reg.ID); delErr != nil {
			t.Logf("cleanup: delete registry %d: %v", *reg.ID, delErr)
		}
	}()

	t.Run("Get", func(t *testing.T) {
		got, resp, err := m.GetPluginRegistry(org, *reg.ID)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, name, *got.Name)
	})

	t.Run("List", func(t *testing.T) {
		regs, resp, err := m.ListPluginRegistries(org)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var found bool
		for _, r := range regs {
			if *r.ID == *reg.ID {
				found = true
				break
			}
		}
		assert.True(t, found, "created registry not found in list")
	})

	t.Run("Update", func(t *testing.T) {
		newName := name + "-upd"
		updated, resp, err := m.UpdatePluginRegistry(org, *reg.ID, newName)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, newName, *updated.Name)
	})
}

func TestRegistryPluginCRUD(t *testing.T) {
	if !pluginsMiddlewareEnabled() {
		t.Skip("set CY_TEST_E2E_PLUGINS=1 to run plugin middleware tests")
	}

	m := config.Middleware
	org := config.Org

	reg, err := config.NewTestPluginRegistry(t, testcfg.RandomCanonical("mw-reg-pl"), "http://plugin-registry:4000")
	require.NoError(t, err)

	pluginName := testcfg.RandomCanonical("mw-plugin")
	plugin, _, err := m.CreateRegistryPlugin(org, *reg.ID, pluginName)
	require.NoError(t, err)
	require.NotNil(t, plugin)

	defer func() {
		if _, delErr := m.DeleteRegistryPlugin(org, *reg.ID, *plugin.ID); delErr != nil {
			t.Logf("cleanup: delete plugin %d: %v", *plugin.ID, delErr)
		}
	}()

	t.Run("Get", func(t *testing.T) {
		got, resp, err := m.GetRegistryPlugin(org, *reg.ID, *plugin.ID)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, pluginName, *got.Name)
	})

	t.Run("List", func(t *testing.T) {
		plugins, resp, err := m.ListRegistryPlugins(org, *reg.ID)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var found bool
		for _, p := range plugins {
			if *p.ID == *plugin.ID {
				found = true
				break
			}
		}
		assert.True(t, found, "created plugin not found in list")
	})

	t.Run("Update", func(t *testing.T) {
		newName := pluginName + "-upd"
		updated, resp, err := m.UpdateRegistryPlugin(org, *reg.ID, *plugin.ID, newName)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, newName, *updated.Name)
	})
}

func TestPluginVersionCRUD(t *testing.T) {
	if !pluginsMiddlewareEnabled() {
		t.Skip("set CY_TEST_E2E_PLUGINS=1 to run plugin middleware tests")
	}

	m := config.Middleware
	org := config.Org

	reg, err := config.NewTestPluginRegistry(t, testcfg.RandomCanonical("mw-reg-ver"), "http://plugin-registry:4000")
	require.NoError(t, err)

	plugin, err := config.NewTestPluginInRegistry(t, *reg.ID, testcfg.RandomCanonical("mw-plver"))
	require.NoError(t, err)

	ver, _, err := m.CreatePluginVersion(org, *reg.ID, *plugin.ID, uniquePluginImage(t))
	require.NoError(t, err)
	require.NotNil(t, ver)

	defer func() {
		if _, delErr := m.DeletePluginVersion(org, *reg.ID, *plugin.ID, *ver.ID); delErr != nil {
			t.Logf("cleanup: delete version %d: %v", *ver.ID, delErr)
		}
	}()

	t.Run("Get", func(t *testing.T) {
		got, resp, err := m.GetPluginVersion(org, *reg.ID, *plugin.ID, *ver.ID)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, *ver.ID, *got.ID)
	})

	t.Run("List", func(t *testing.T) {
		versions, resp, err := m.ListPluginVersions(org, *reg.ID, *plugin.ID)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NotEmpty(t, versions)
	})

	t.Run("ListLogs", func(t *testing.T) {
		logs, resp, err := m.ListPluginVersionLogs(org, *reg.ID, *plugin.ID, *ver.ID)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NotNil(t, logs)
	})
}

func TestPluginInstallAndGet(t *testing.T) {
	if !pluginsMiddlewareEnabled() {
		t.Skip("set CY_TEST_E2E_PLUGINS=1 to run plugin middleware tests")
	}

	m := config.Middleware
	org := config.Org

	_, err := config.AcceptPluginManager(t)
	require.NoError(t, err)

	reg, err := config.NewTestPluginRegistry(t, testcfg.RandomCanonical("mw-reg-ins"), "http://plugin-registry:4000")
	require.NoError(t, err)

	plugin, err := config.NewTestPluginInRegistry(t, *reg.ID, testcfg.RandomCanonical("mw-plinst"))
	require.NoError(t, err)

	ver, err := config.NewTestPluginVersion(t, *reg.ID, *plugin.ID, uniquePluginImage(t))
	require.NoError(t, err)

	// Install returns 200 with PluginInstall body (swagger: "200 The Plugin Install data").
	resp, err := m.InstallPluginVersion(org, *reg.ID, *plugin.ID, *ver.ID, nil)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Resolve install ID from list — allow a moment for the record to appear.
	var installID uint32
	for i := 0; i < 10; i++ {
		plugins, _, _ := m.ListPlugins(org)
		for _, p := range plugins {
			if p.Name != nil && *p.Name == *plugin.Name && p.Install != nil {
				installID = *p.Install.ID
			}
		}
		if installID != 0 {
			break
		}
		time.Sleep(2 * time.Second)
	}
	require.NotZero(t, installID, "install record not found in ListPlugins")

	t.Cleanup(func() {
		if _, delErr := m.DeletePlugin(org, installID); delErr != nil {
			t.Logf("cleanup: delete plugin install %d: %v", installID, delErr)
		}
	})

	t.Run("GetInstall", func(t *testing.T) {
		install, resp, err := m.GetPlugin(org, installID)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		require.NotNil(t, install)
		// The plugin-manager deploys asynchronously; status may be pending/installed.
		assert.NotNil(t, install.ID)
	})

	t.Run("ListInstalls", func(t *testing.T) {
		plugins, resp, err := m.ListPlugins(org)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var found bool
		for _, p := range plugins {
			if p.Install != nil && *p.Install.ID == installID {
				found = true
				break
			}
		}
		assert.True(t, found, "install %d not found in ListPlugins", installID)
	})

	t.Run("RetryIdempotent", func(t *testing.T) {
		// Second install should return 409 Conflict; retry endpoint returns 204.
		retryResp, err := m.InstallPluginVersion(org, *reg.ID, *plugin.ID, *ver.ID, nil)
		if err == nil {
			// First install on a freshly reset DB — no conflict yet.
			assert.Contains(t, []int{http.StatusOK, http.StatusAccepted}, retryResp.StatusCode)
			return
		}
		if retryResp != nil && retryResp.StatusCode == http.StatusConflict {
			retryResp2, retryErr := m.RetryPluginVersion(org, *reg.ID, *plugin.ID, *ver.ID)
			require.NoError(t, retryErr)
			assert.Equal(t, http.StatusNoContent, retryResp2.StatusCode)
			return
		}
		t.Fatalf("unexpected error on second install: %v", err)
	})

	t.Run("Logs", func(t *testing.T) {
		logs, resp, err := m.ListPluginLogs(org, installID)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NotNil(t, logs)
	})

	t.Run("ListWidgets", func(t *testing.T) {
		widgets, resp, err := m.ListPluginWidgets(org, "sideMenuPage")
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NotNil(t, widgets)
	})
}
