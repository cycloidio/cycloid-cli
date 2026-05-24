package testcfg

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/cycloidio/cycloid-cli/client/models"
)

// NewTestPluginRegistry creates a plugin registry for testing and registers cleanup.
// If a registry with the same URL already exists it is returned as-is (no cleanup
// registered for it, since we don't own it).
func (config *Config) NewTestPluginRegistry(t *testing.T, name, url string) (*models.PluginRegistry, error) {
	t.Helper()
	m := config.Middleware
	out, _, err := m.CreatePluginRegistry(config.Org, name, url)
	if err != nil {
		// 409 Conflict → registry with this URL already exists; find and return it.
		if strings.Contains(err.Error(), "409") || strings.Contains(err.Error(), "already exists") {
			regs, _, listErr := m.ListPluginRegistries(config.Org)
			if listErr != nil {
				return nil, fmt.Errorf("create failed (%v) and list fallback failed: %w", err, listErr)
			}
			for _, r := range regs {
				if r.URL != nil && string(*r.URL) == url {
					return r, nil
				}
			}
			return nil, fmt.Errorf("registry with url %q not found after 409: %w", url, err)
		}
		return nil, err
	}
	config.AppendCleanup(func() {
		if _, delErr := m.DeletePluginRegistry(config.Org, *out.ID); delErr != nil {
			log.Printf("failed to cleanup plugin registry %q: %v", name, delErr)
		}
	})
	return out, nil
}

// NewTestPluginInRegistry creates a plugin inside the given registry and registers cleanup.
func (config *Config) NewTestPluginInRegistry(t *testing.T, registryID uint32, name string) (*models.Plugin, error) {
	t.Helper()
	m := config.Middleware
	out, _, err := m.CreateRegistryPlugin(config.Org, registryID, name)
	if err != nil {
		return nil, err
	}
	config.AppendCleanup(func() {
		if _, err := m.DeleteRegistryPlugin(config.Org, registryID, *out.ID); err != nil {
			log.Printf("failed to cleanup registry plugin %q: %v", name, err)
		}
	})
	return out, nil
}

// NewTestPluginVersion publishes a plugin version, waits for it to reach "success"
// status (indicating the plugin-registry has validated the image), and registers cleanup.
func (config *Config) NewTestPluginVersion(t *testing.T, registryID, pluginID uint32, dockerImage string) (*models.PluginVersion, error) {
	t.Helper()
	m := config.Middleware
	out, _, err := m.CreatePluginVersion(config.Org, registryID, pluginID, dockerImage)
	if err != nil {
		return nil, err
	}
	config.AppendCleanup(func() {
		if _, delErr := m.DeletePluginVersion(config.Org, registryID, pluginID, *out.ID); delErr != nil {
			log.Printf("failed to cleanup plugin version %d: %v", *out.ID, delErr)
		}
	})

	// The plugin-registry validates the image asynchronously. Poll until success.
	deadline := time.Now().Add(60 * time.Second)
	for time.Now().Before(deadline) {
		ver, _, pollErr := m.GetPluginVersion(config.Org, registryID, pluginID, *out.ID)
		if pollErr == nil && ver.Status != nil {
			switch *ver.Status {
			case "success":
				return ver, nil
			case "failed":
				return nil, fmt.Errorf("plugin version %d validation failed", *out.ID)
			}
		}
		time.Sleep(3 * time.Second)
	}
	return nil, fmt.Errorf("plugin version %d did not reach success within 60s", *out.ID)
}

// AcceptPluginManager finds and accepts the auto-registered test-plugin-manager.
// The plugin-manager service in compose.yml auto-registers itself as "test-plugin-manager"
// on startup, but retries if the org doesn't exist yet. This helper polls until it
// appears (up to 90s) then accepts it.
func (config *Config) AcceptPluginManager(t *testing.T) (*models.PluginManager, error) {
	t.Helper()
	m := config.Middleware
	deadline := time.Now().Add(90 * time.Second)
	for {
		managers, _, listErr := m.ListPluginManagers(config.Org)
		if listErr == nil {
			for _, mgr := range managers {
				if mgr.Name != nil && *mgr.Name == "test-plugin-manager" {
					// "invite_accepted" is what the API returns once accepted.
					// Either form means already done.
					if mgr.InviteStatus != nil &&
						(*mgr.InviteStatus == "accepted" || *mgr.InviteStatus == "invite_accepted") {
						return mgr, nil
					}
					out, _, acceptErr := m.UpdatePluginManager(config.Org, *mgr.ID, "accepted")
					if acceptErr != nil {
						// Already accepted by a concurrent call — treat as success.
						if strings.Contains(acceptErr.Error(), "only pending") {
							return mgr, nil
						}
						return nil, acceptErr
					}
					return out, nil
				}
			}
		}
		// Any error (including JSON decode of numeric status on first registration)
		// means the manager is not ready yet — keep polling.
		if time.Now().After(deadline) {
			return nil, fmt.Errorf("test-plugin-manager did not register within 90s")
		}
		time.Sleep(5 * time.Second)
	}
}
