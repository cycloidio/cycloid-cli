package e2e_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// pluginsEnabled reports whether the plugin e2e suite should run.
func pluginsEnabled() bool {
	return os.Getenv("CY_TEST_E2E_PLUGINS") == "1"
}

// pollPluginRunning polls GetPlugin until Status == "running" or timeout.
// Returns true if "running" was reached, false on timeout (plugin-manager may lack
// a container runtime in the compose test environment).
func pollPluginRunning(t *testing.T, org string, installID uint32) bool {
	t.Helper()
	deadline := time.Now().Add(90 * time.Second)
	for time.Now().Before(deadline) {
		install, _, err := config.Middleware.GetPlugin(org, installID)
		if err == nil && install != nil && install.Status != nil && *install.Status == "running" {
			return true
		}
		if err == nil && install != nil && install.Status != nil && *install.Status == "failed" {
			t.Logf("plugin install %d entered failed state (plugin-manager may lack container runtime)", installID)
			return false
		}
		time.Sleep(3 * time.Second)
	}
	t.Logf("plugin install %d did not reach running state within 90s (plugin-manager may lack container runtime)", installID)
	return false
}

func TestPluginManagers(t *testing.T) {
	if !pluginsEnabled() {
		t.Skip("set CY_TEST_E2E_PLUGINS=1 to run plugin e2e tests")
	}

	org := config.Org

	t.Run("List", func(t *testing.T) {
		out, err := executeCommand([]string{"plugin", "manager", "list", "--output", "json"})
		require.NoError(t, err, "plugin manager list failed: %s", out)
		var managers []map[string]any
		require.NoError(t, json.Unmarshal([]byte(out), &managers), "output not valid JSON")
	})

	t.Run("AcceptTestPluginManager", func(t *testing.T) {
		mgr, err := config.AcceptPluginManager(t)
		require.NoError(t, err, "failed to accept test-plugin-manager")
		require.NotNil(t, mgr)
		// API returns "invite_accepted" once accepted.
		assert.Contains(t, []string{"accepted", "invite_accepted"}, *mgr.InviteStatus)
	})

	t.Run("GetAfterAccept", func(t *testing.T) {
		managers, _, err := config.Middleware.ListPluginManagers(org)
		require.NoError(t, err)
		var found bool
		for _, m := range managers {
			if m.Name != nil && *m.Name == "test-plugin-manager" {
				found = true
				assert.Contains(t, []string{"accepted", "invite_accepted"}, *m.InviteStatus)
				break
			}
		}
		assert.True(t, found, "test-plugin-manager not found after accept")
	})
}

func TestPluginRegistries(t *testing.T) {
	if !pluginsEnabled() {
		t.Skip("set CY_TEST_E2E_PLUGINS=1 to run plugin e2e tests")
	}

	name := randomCanonical("e2e-registry")
	// Use a unique offline URL per test run to avoid the API's one-URL-per-registry constraint.
	regURL := fmt.Sprintf("http://test-reg-%s:9999", name)

	t.Run("Create", func(t *testing.T) {
		out, err := executeCommand([]string{
			"plugin", "registry", "create",
			"--name", name,
			"--url", regURL,
			"--output", "json",
		})
		require.NoError(t, err, "create failed: %s", out)
		var reg map[string]any
		require.NoError(t, json.Unmarshal([]byte(out), &reg))
		assert.Equal(t, name, reg["name"])
	})

	defer t.Run("Cleanup", func(t *testing.T) {
		_, err := executeCommand([]string{"plugin", "registry", "delete", name})
		assert.NoError(t, err, "cleanup delete failed")
	})

	t.Run("List", func(t *testing.T) {
		out, err := executeCommand([]string{"plugin", "registry", "list", "--output", "json"})
		require.NoError(t, err)
		names, err := JSONListExtractFields(out, "name", "", "")
		require.NoError(t, err)
		assert.Contains(t, names, name)
	})

	t.Run("Get", func(t *testing.T) {
		out, err := executeCommand([]string{"plugin", "registry", "get", name, "--output", "json"})
		require.NoError(t, err, "get failed: %s", out)
		var reg map[string]any
		require.NoError(t, json.Unmarshal([]byte(out), &reg))
		assert.Equal(t, name, reg["name"])
	})

	t.Run("UpdateName", func(t *testing.T) {
		renamed := name + "-r"
		out, err := executeCommand([]string{
			"plugin", "registry", "update", name,
			"--name", renamed,
			"--output", "json",
		})
		require.NoError(t, err, "update failed: %s", out)
		var reg map[string]any
		require.NoError(t, json.Unmarshal([]byte(out), &reg))
		assert.Equal(t, renamed, reg["name"])
		// Rename back so subsequent subtests can find by original name.
		_, err = executeCommand([]string{
			"plugin", "registry", "update", renamed,
			"--name", name,
		})
		require.NoError(t, err)
	})

	t.Run("CreateIdempotentWithUpdate", func(t *testing.T) {
		out, err := executeCommand([]string{
			"plugin", "registry", "create",
			"--name", name,
			"--url", regURL,
			"--update",
			"--output", "json",
		})
		require.NoError(t, err, "create --update failed: %s", out)
		var reg map[string]any
		require.NoError(t, json.Unmarshal([]byte(out), &reg))
		assert.Equal(t, name, reg["name"])
	})

	t.Run("CreateDuplicateFails", func(t *testing.T) {
		_, err := executeCommand([]string{
			"plugin", "registry", "create",
			"--name", name,
			"--url", regURL,
		})
		assert.Error(t, err, "expected error creating duplicate registry without --update")
	})
}

func TestRegistryPlugins(t *testing.T) {
	if !pluginsEnabled() {
		t.Skip("set CY_TEST_E2E_PLUGINS=1 to run plugin e2e tests")
	}

	regURL := "http://plugin-registry:4000"
	reg, err := config.NewTestPluginRegistry(t, randomCanonical("e2e-reg-pl"), regURL)
	require.NoError(t, err)
	regID := *reg.ID
	regName := *reg.Name

	pluginName := randomCanonical("e2e-plugin")

	t.Run("Create", func(t *testing.T) {
		out, err := executeCommand([]string{
			"plugin", "registry", "plugin", "create",
			regName,
			"--name", pluginName,
			"--output", "json",
		})
		require.NoError(t, err, "create plugin failed: %s", out)
		var p map[string]any
		require.NoError(t, json.Unmarshal([]byte(out), &p))
		assert.Equal(t, pluginName, p["name"])
	})

	defer t.Run("CleanupPlugin", func(t *testing.T) {
		_, _ = executeCommand([]string{
			"plugin", "registry", "plugin", "delete",
			regName,
			pluginName,
		})
	})

	t.Run("List", func(t *testing.T) {
		out, err := executeCommand([]string{
			"plugin", "registry", "plugin", "list",
			regName,
			"--output", "json",
		})
		require.NoError(t, err)
		names, err := JSONListExtractFields(out, "name", "", "")
		require.NoError(t, err)
		assert.Contains(t, names, pluginName)
	})

	t.Run("Get", func(t *testing.T) {
		out, err := executeCommand([]string{
			"plugin", "registry", "plugin", "get",
			pluginName,
			"--registry", regName,
			"--output", "json",
		})
		require.NoError(t, err, "get plugin failed: %s", out)
		var p map[string]any
		require.NoError(t, json.Unmarshal([]byte(out), &p))
		assert.Equal(t, pluginName, p["name"])
	})

	t.Run("Update", func(t *testing.T) {
		renamed := pluginName + "-r"
		out, err := executeCommand([]string{
			"plugin", "registry", "plugin", "update",
			regName,
			pluginName,
			"--name", renamed,
			"--output", "json",
		})
		require.NoError(t, err, "update plugin failed: %s", out)
		var p map[string]any
		require.NoError(t, json.Unmarshal([]byte(out), &p))
		assert.Equal(t, renamed, p["name"])
		// Rename back.
		_, err = executeCommand([]string{
			"plugin", "registry", "plugin", "update",
			regName,
			renamed,
			"--name", pluginName,
		})
		require.NoError(t, err)
	})

	_ = regID
}

func TestPluginVersions(t *testing.T) {
	if !pluginsEnabled() {
		t.Skip("set CY_TEST_E2E_PLUGINS=1 to run plugin e2e tests")
	}

	ensureHelloWorldImage(t)

	regURL := "http://plugin-registry:4000"
	reg, err := config.NewTestPluginRegistry(t, randomCanonical("e2e-reg-ver"), regURL)
	require.NoError(t, err)
	regName := *reg.Name

	plugin, err := config.NewTestPluginInRegistry(t, *reg.ID, randomCanonical("e2e-plver"))
	require.NoError(t, err)
	pluginID := *plugin.ID
	pluginName := *plugin.Name

	t.Run("Publish", func(t *testing.T) {
		out, err := executeCommand([]string{
			"plugin", "registry", "plugin", "version", "publish",
			"--registry", regName,
			"--plugin", pluginName,
			"--docker-image", uniqueHelloWorldVersion(t),
			"--output", "json",
		})
		require.NoError(t, err, "publish failed: %s", out)
		var v map[string]any
		require.NoError(t, json.Unmarshal([]byte(out), &v))
		assert.NotEmpty(t, v["id"])
	})

	versions, _, err := config.Middleware.ListPluginVersions(config.Org, *reg.ID, pluginID)
	require.NoError(t, err)
	require.NotEmpty(t, versions, "no versions after publish")
	versionID := *versions[0].ID

	t.Run("List", func(t *testing.T) {
		out, err := executeCommand([]string{
			"plugin", "registry", "plugin", "version", "list",
			"--registry", regName,
			"--plugin", pluginName,
			"--output", "json",
		})
		require.NoError(t, err)
		var vs []map[string]any
		require.NoError(t, json.Unmarshal([]byte(out), &vs))
		assert.NotEmpty(t, vs)
	})

	t.Run("Get", func(t *testing.T) {
		out, err := executeCommand([]string{
			"plugin", "registry", "plugin", "version", "get",
			fmt.Sprint(versionID),
			"--registry", regName,
			"--plugin", pluginName,
			"--output", "json",
		})
		require.NoError(t, err, "get version failed: %s", out)
		var v map[string]any
		require.NoError(t, json.Unmarshal([]byte(out), &v))
		assert.NotEmpty(t, v["id"])
	})

	t.Run("Logs", func(t *testing.T) {
		// Version logs endpoint should succeed even when empty.
		out, err := executeCommand([]string{
			"plugin", "registry", "plugin", "version", "logs",
			fmt.Sprint(versionID),
			"--registry", regName,
			"--plugin", pluginName,
			"--output", "json",
		})
		require.NoError(t, err, "version logs failed: %s", out)
	})

	defer t.Run("Delete", func(t *testing.T) {
		_, err := executeCommand([]string{
			"plugin", "registry", "plugin", "version", "delete",
			fmt.Sprint(versionID),
			"--registry", regName,
			"--plugin", pluginName,
		})
		assert.NoError(t, err)
	})
}

func TestPluginInstallLifecycle(t *testing.T) {
	if !pluginsEnabled() {
		t.Skip("set CY_TEST_E2E_PLUGINS=1 to run plugin e2e tests")
	}

	ensureHelloWorldImage(t)

	org := config.Org

	// Ensure test-plugin-manager is accepted.
	_, err := config.AcceptPluginManager(t)
	require.NoError(t, err, "failed to accept test-plugin-manager")

	regURL := "http://plugin-registry:4000"
	reg, err := config.NewTestPluginRegistry(t, randomCanonical("e2e-reg-inst"), regURL)
	require.NoError(t, err)
	regName := *reg.Name

	plugin, err := config.NewTestPluginInRegistry(t, *reg.ID, randomCanonical("e2e-plinst"))
	require.NoError(t, err)
	pluginName := *plugin.Name

	version, err := config.NewTestPluginVersion(t, *reg.ID, *plugin.ID, uniqueHelloWorldVersion(t))
	require.NoError(t, err)
	versionID := *version.ID

	var installID uint32

	t.Run("Install", func(t *testing.T) {
		_, err := executeCommand([]string{
			"plugin", "registry", "plugin", "version", "install",
			fmt.Sprint(versionID),
			"--registry", regName,
			"--plugin", pluginName,
		})
		require.NoError(t, err, "install failed")

		// Resolve the install ID from the plugin list.
		installs, _, err := config.Middleware.ListPlugins(org)
		require.NoError(t, err)
		for _, p := range installs {
			if p.Name != nil && *p.Name == pluginName && p.Install != nil {
				installID = *p.Install.ID
				break
			}
		}
		require.NotZero(t, installID, "install record not found after install")
	})

	var pluginRunning bool
	t.Run("PollRunning", func(t *testing.T) {
		if installID == 0 {
			t.Skip("install step did not complete")
		}
		pluginRunning = pollPluginRunning(t, org, installID)
		if !pluginRunning {
			t.Skip("plugin did not reach running state (plugin-manager may lack container runtime)")
		}
	})

	t.Run("Get", func(t *testing.T) {
		if installID == 0 {
			t.Skip("install step did not complete")
		}
		out, err := executeCommand([]string{
			"plugin", "get", fmt.Sprint(installID),
			"--output", "json",
		})
		require.NoError(t, err, "plugin get failed: %s", out)
		var install map[string]any
		require.NoError(t, json.Unmarshal([]byte(out), &install))
		assert.NotNil(t, install["id"], "install record should have an id")
	})

	t.Run("List", func(t *testing.T) {
		out, err := executeCommand([]string{"plugin", "list", "--output", "json"})
		require.NoError(t, err)
		var installs []map[string]any
		require.NoError(t, json.Unmarshal([]byte(out), &installs))
		assert.NotEmpty(t, installs)
	})

	t.Run("RetryIdempotent", func(t *testing.T) {
		// A second install with --retry should not fail with 409.
		_, err := executeCommand([]string{
			"plugin", "registry", "plugin", "version", "install",
			fmt.Sprint(versionID),
			"--registry", regName,
			"--plugin", pluginName,
			"--retry",
		})
		require.NoError(t, err, "install --retry on already-installed version should succeed")
	})

	defer t.Run("Uninstall", func(t *testing.T) {
		if installID == 0 {
			return
		}
		_, err := executeCommand([]string{"plugin", "uninstall", fmt.Sprint(installID)})
		assert.NoError(t, err, "uninstall failed")
	})
}

func TestPluginWidgets(t *testing.T) {
	if !pluginsEnabled() {
		t.Skip("set CY_TEST_E2E_PLUGINS=1 to run plugin e2e tests")
	}

	ensureHelloWorldImage(t)

	org := config.Org

	// Ensure test-plugin-manager is accepted.
	_, err := config.AcceptPluginManager(t)
	require.NoError(t, err, "failed to accept test-plugin-manager")

	regURL := "http://plugin-registry:4000"
	reg, err := config.NewTestPluginRegistry(t, randomCanonical("e2e-reg-wgt"), regURL)
	require.NoError(t, err)
	regName := *reg.Name

	plugin, err := config.NewTestPluginInRegistry(t, *reg.ID, randomCanonical("e2e-plwgt"))
	require.NoError(t, err)
	pluginName := *plugin.Name

	version, err := config.NewTestPluginVersion(t, *reg.ID, *plugin.ID, uniqueHelloWorldVersion(t))
	require.NoError(t, err)
	versionID := *version.ID

	// Install and wait for running.
	_, err = executeCommand([]string{
		"plugin", "registry", "plugin", "version", "install",
		fmt.Sprint(versionID),
		"--registry", regName,
		"--plugin", pluginName,
	})
	require.NoError(t, err)

	installs, _, err := config.Middleware.ListPlugins(org)
	require.NoError(t, err)
	var installID uint32
	for _, p := range installs {
		if p.Name != nil && *p.Name == pluginName && p.Install != nil {
			installID = *p.Install.ID
			break
		}
	}
	require.NotZero(t, installID)

	defer t.Run("CleanupInstall", func(t *testing.T) {
		_, _ = config.Middleware.DeletePlugin(org, installID)
	})

	running := pollPluginRunning(t, org, installID)

	t.Run("ListWidgets", func(t *testing.T) {
		if !running {
			t.Skip("plugin not running — skipping widget test (plugin-manager may lack container runtime)")
		}
		out, err := executeCommand([]string{
			"plugin", "widget", "list",
			"--placement", "sideMenuPage",
			"--output", "json",
		})
		require.NoError(t, err, "widget list failed: %s", out)
		var widgets []map[string]any
		require.NoError(t, json.Unmarshal([]byte(out), &widgets))
		assert.NotEmpty(t, widgets, "expected at least one widget after install")
	})

	t.Run("IframeServedContent", func(t *testing.T) {
		if !running {
			t.Skip("plugin not running — skipping iframe test (plugin-manager may lack container runtime)")
		}
		// Resolve the widget URL from the plugin widget list and verify it
		// serves HTML containing "Hello World".
		widgets, _, err := config.Middleware.ListPluginWidgets(org, "sideMenuPage")
		require.NoError(t, err)
		require.NotEmpty(t, widgets, "no widgets found for sideMenuPage placement")

		// Find the iframe widget belonging to our install.
		var iframeURL string
		for _, w := range widgets {
			if w.Type == nil || *w.Type != "iframe" {
				continue
			}
			// Widget field is a map with a "url" key.
			wmap, ok := w.Widget.(map[string]any)
			if !ok {
				continue
			}
			rawURL, ok := wmap["url"].(string)
			if !ok {
				continue
			}
			iframeURL = rawURL
			break
		}
		require.NotEmpty(t, iframeURL, "no iframe widget URL found")

		// plugin-manager is reachable on localhost:4000 (compose port mapping).
		hostURL := strings.ReplaceAll(iframeURL, "plugin-manager:4000", "localhost:4000")

		resp, err := http.Get(hostURL) //nolint:noctx
		require.NoError(t, err, "GET iframe URL failed: %s", hostURL)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Contains(t, strings.ToLower(string(body)), "hello world",
			"iframe URL %s did not return Hello World content", hostURL)
	})
}

func TestPluginCommandsErrorPaths(t *testing.T) {
	if !pluginsEnabled() {
		t.Skip("set CY_TEST_E2E_PLUGINS=1 to run plugin e2e tests")
	}

	t.Run("WatchIntervalTooShort", func(t *testing.T) {
		_, err := executeCommand([]string{
			"plugin", "logs", "999",
			"--watch-interval", "100ms",
		})
		assert.Error(t, err, "expected error for --watch-interval below 500ms")
		assert.Contains(t, err.Error(), "500ms")
	})

	t.Run("RegistryCreateMissingURL", func(t *testing.T) {
		_, err := executeCommand([]string{
			"plugin", "registry", "create",
			"--name", "missing-url-test",
		})
		assert.Error(t, err, "expected error when --url is absent")
	})

	t.Run("InstallRetryOnNonExistent", func(t *testing.T) {
		// Should fail cleanly, not panic.
		_, err := executeCommand([]string{
			"plugin", "registry", "plugin", "version", "install",
			"--registry", "nonexistent-registry-e2e",
			"--plugin", "nonexistent-plugin-e2e",
			"99999",
			"--retry",
		})
		assert.Error(t, err)
	})
}
