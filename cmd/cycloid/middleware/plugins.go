package middleware

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func registryPluginsRoute(org string, registryID uint32, extra ...string) []string {
	base := []string{"organizations", org, "plugin_registries", fmt.Sprint(registryID), "plugins"}
	return append(base, extra...)
}

// --- Plugin Managers ---

func (m *middleware) ListPluginManagers(org string) ([]*models.PluginManager, *http.Response, error) {
	var result []*models.PluginManager
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_managers"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) GetPluginManager(org string, id uint32) (*models.PluginManager, *http.Response, error) {
	var result *models.PluginManager
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_managers", fmt.Sprint(id)},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) CreatePluginManager(org, name, url string) (*models.PluginManager, *http.Response, error) {
	u := strfmt.URI(url)
	body := &models.NewPluginManager{
		Name: &name,
		URL:  &u,
	}
	var result *models.PluginManager
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_managers"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) UpdatePluginManager(org string, id uint32, inviteStatus string) (*models.PluginManager, *http.Response, error) {
	body := &models.UpdatePluginManager{
		InviteStatus: &inviteStatus,
	}
	var result *models.PluginManager
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_managers", fmt.Sprint(id)},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) DeletePluginManager(org string, id uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_managers", fmt.Sprint(id)},
	}, nil)
	return resp, err
}

// --- Plugin Installs ---

func (m *middleware) ListPlugins(org string) ([]*models.PluginInstall, *http.Response, error) {
	var result []*models.PluginInstall
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "plugins"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) GetPlugin(org string, id uint32) (*models.PluginInstall, *http.Response, error) {
	var result *models.PluginInstall
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "plugins", fmt.Sprint(id)},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// createPluginBody is a superset of models.NewPluginInstall that optionally
// includes a version ID. The generated model is not used here so that we can
// add fields without editing auto-generated code.
type createPluginBody struct {
	Configuration   map[string]string `json:"configuration"`
	PluginVersionID *uint32           `json:"plugin_version_id,omitempty"`
}

func (m *middleware) CreatePlugin(org string, versionID *uint32, config map[string]string) (*models.PluginInstall, *http.Response, error) {
	body := &createPluginBody{
		Configuration:   config,
		PluginVersionID: versionID,
	}
	var result *models.PluginInstall
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "plugins"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) UpdatePlugin(org string, id, versionID uint32, config map[string]string) (*models.PluginInstall, *http.Response, error) {
	body := &models.UpdatePluginInstall{
		Configuration:   config,
		PluginVersionID: &versionID,
	}
	var result *models.PluginInstall
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "plugins", fmt.Sprint(id)},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) DeletePlugin(org string, id uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "plugins", fmt.Sprint(id)},
	}, nil)
	return resp, err
}

func (m *middleware) ListPluginLogs(org string, id uint32) (*models.PluginLogs, *http.Response, error) {
	var result *models.PluginLogs
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "plugins", fmt.Sprint(id), "logs"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// --- Plugin Registries ---

func (m *middleware) ListPluginRegistries(org string) ([]*models.PluginRegistry, *http.Response, error) {
	var result []*models.PluginRegistry
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_registries"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) GetPluginRegistry(org string, id uint32) (*models.PluginRegistry, *http.Response, error) {
	var result *models.PluginRegistry
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_registries", fmt.Sprint(id)},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) CreatePluginRegistry(org, name, url string) (*models.PluginRegistry, *http.Response, error) {
	u := strfmt.URI(url)
	body := &models.NewPluginRegistry{
		Name: &name,
		URL:  &u,
	}
	var result *models.PluginRegistry
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_registries"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) UpdatePluginRegistry(org string, id uint32, name string) (*models.PluginRegistry, *http.Response, error) {
	body := &models.UpdatePluginRegistry{Name: &name}
	var result *models.PluginRegistry
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_registries", fmt.Sprint(id)},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) DeletePluginRegistry(org string, id uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_registries", fmt.Sprint(id)},
	}, nil)
	return resp, err
}

// --- Registry Plugins ---

func (m *middleware) ListRegistryPlugins(org string, registryID uint32) ([]*models.Plugin, *http.Response, error) {
	var result []*models.Plugin
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        registryPluginsRoute(org, registryID),
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) GetRegistryPlugin(org string, registryID, pluginID uint32) (*models.Plugin, *http.Response, error) {
	var result *models.Plugin
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        registryPluginsRoute(org, registryID, fmt.Sprint(pluginID)),
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) CreateRegistryPlugin(org string, registryID uint32, name string) (*models.Plugin, *http.Response, error) {
	body := &models.NewPlugin{Name: &name}
	var result *models.Plugin
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        registryPluginsRoute(org, registryID),
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) UpdateRegistryPlugin(org string, registryID, pluginID uint32, name string) (*models.Plugin, *http.Response, error) {
	body := &models.UpdatePlugin{Name: &name}
	var result *models.Plugin
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        registryPluginsRoute(org, registryID, fmt.Sprint(pluginID)),
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) DeleteRegistryPlugin(org string, registryID, pluginID uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        registryPluginsRoute(org, registryID, fmt.Sprint(pluginID)),
	}, nil)
	return resp, err
}

// --- Registry Plugin Versions ---

func versionsRoute(org string, registryID, pluginID uint32, extra ...string) []string {
	base := registryPluginsRoute(org, registryID, fmt.Sprint(pluginID), "versions")
	return append(base, extra...)
}

func (m *middleware) ListPluginVersions(org string, registryID, pluginID uint32) ([]*models.PluginVersion, *http.Response, error) {
	var result []*models.PluginVersion
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        versionsRoute(org, registryID, pluginID),
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) GetPluginVersion(org string, registryID, pluginID, versionID uint32) (*models.PluginVersion, *http.Response, error) {
	var result *models.PluginVersion
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        versionsRoute(org, registryID, pluginID, fmt.Sprint(versionID)),
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) CreatePluginVersion(org string, registryID, pluginID uint32, url string) (*models.PluginVersion, *http.Response, error) {
	// Use a plain struct instead of models.NewPluginVersion to avoid strfmt.URI
	// coercion, which rejects Docker image references (scheme-less URIs).
	body := struct {
		URL string `json:"url"`
	}{URL: url}
	var result *models.PluginVersion
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        versionsRoute(org, registryID, pluginID),
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) DeletePluginVersion(org string, registryID, pluginID, versionID uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        versionsRoute(org, registryID, pluginID, fmt.Sprint(versionID)),
	}, nil)
	return resp, err
}

func (m *middleware) InstallPluginVersion(org string, registryID, pluginID, versionID uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        versionsRoute(org, registryID, pluginID, fmt.Sprint(versionID), "install"),
	}, nil)
	return resp, err
}

func (m *middleware) RetryPluginVersion(org string, registryID, pluginID, versionID uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        versionsRoute(org, registryID, pluginID, fmt.Sprint(versionID), "retry"),
	}, nil)
	return resp, err
}

func (m *middleware) ListPluginVersionLogs(org string, registryID, pluginID, versionID uint32) ([]*models.PluginVersionLog, *http.Response, error) {
	var result []*models.PluginVersionLog
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        versionsRoute(org, registryID, pluginID, fmt.Sprint(versionID), "logs"),
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// --- Component Plugins ---

func (m *middleware) ListComponentPlugins(org, project, env, component string) ([]*models.Plugin, *http.Response, error) {
	var result []*models.Plugin
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "plugins"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) SetComponentPluginRelation(org, project, env, component string, pluginInstallID uint32, enabled bool) (*models.PluginRelation, *http.Response, error) {
	body := &models.PluginRelation{
		Enabled:   &enabled,
		Relations: map[string]any{},
	}
	var result *models.PluginRelation
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "plugins", fmt.Sprint(pluginInstallID), "relation"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
