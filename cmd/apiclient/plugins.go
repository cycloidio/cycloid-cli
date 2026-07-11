package apiclient

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/gen/models"
)

func registryPluginsRoute(org string, registryID uint32, extra ...string) []string {
	base := []string{"organizations", org, "plugin_registries", fmt.Sprint(registryID), "plugins"}
	return append(base, extra...)
}

// --- Plugin Managers ---

func (m *apiClient) ListPluginManagers(org string) ([]*models.PluginManager, *http.Response, error) {
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

func (m *apiClient) GetPluginManager(org string, id uint32) (*models.PluginManager, *http.Response, error) {
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

func (m *apiClient) CreatePluginManager(org, name, url string, autoRegister bool) (*models.PluginManager, *http.Response, error) {
	u := strfmt.URI(url)
	body := &models.NewPluginManager{
		AutoRegister: autoRegister,
		Name:         &name,
		URL:          &u,
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

func (m *apiClient) UpdatePluginManager(org string, id uint32, inviteStatus string) (*models.PluginManager, *http.Response, error) {
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

func (m *apiClient) DeletePluginManager(org string, id uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_managers", fmt.Sprint(id)},
	}, nil)
	return resp, err
}

// --- Plugin Installs ---

func (m *apiClient) ListPlugins(org string) ([]*models.Plugin, *http.Response, error) {
	var result []*models.Plugin
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

func (m *apiClient) GetPlugin(org string, id uint32) (*models.PluginInstall, *http.Response, error) {
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

func (m *apiClient) UpdatePlugin(org string, id, versionID uint32, config map[string]string) (*models.PluginInstall, *http.Response, error) {
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

func (m *apiClient) DeletePlugin(org string, id uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "plugins", fmt.Sprint(id)},
	}, nil)
	return resp, err
}

func (m *apiClient) ListPluginLogs(org string, id uint32) (*models.PluginLogs, *http.Response, error) {
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

func (m *apiClient) ListPluginRegistries(org string) ([]*models.PluginRegistry, *http.Response, error) {
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

// GetPluginRegistry retrieves a single plugin registry by ID.
// The API has no GET-by-ID endpoint for registries, so this lists all and filters.
func (m *apiClient) GetPluginRegistry(org string, id uint32) (*models.PluginRegistry, *http.Response, error) {
	regs, resp, err := m.ListPluginRegistries(org)
	if err != nil {
		return nil, resp, err
	}
	for _, r := range regs {
		if r.ID != nil && *r.ID == id {
			return r, resp, nil
		}
	}
	return nil, resp, fmt.Errorf("plugin registry %d not found", id)
}

func (m *apiClient) CreatePluginRegistry(org, name, url string) (*models.PluginRegistry, *http.Response, error) {
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

func (m *apiClient) UpdatePluginRegistry(org string, id uint32, name string) (*models.PluginRegistry, *http.Response, error) {
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

func (m *apiClient) DeletePluginRegistry(org string, id uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_registries", fmt.Sprint(id)},
	}, nil)
	return resp, err
}

// --- Registry Plugins ---

// ListRegistryPlugins lists plugins belonging to a specific registry.
// The API has no GET /plugin_registries/{id}/plugins endpoint (only POST for create),
// so we list all plugins and filter by registry ID.
func (m *apiClient) ListRegistryPlugins(org string, registryID uint32) ([]*models.Plugin, *http.Response, error) {
	all, resp, err := m.ListPlugins(org)
	if err != nil {
		return nil, resp, err
	}
	var result []*models.Plugin
	for _, p := range all {
		if p.Registry != nil && p.Registry.ID != nil && *p.Registry.ID == registryID {
			result = append(result, p)
		}
	}
	return result, resp, nil
}

func (m *apiClient) GetRegistryPlugin(org string, registryID, pluginID uint32) (*models.Plugin, *http.Response, error) {
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

func (m *apiClient) CreateRegistryPlugin(org string, registryID uint32, name string) (*models.Plugin, *http.Response, error) {
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

func (m *apiClient) UpdateRegistryPlugin(org string, registryID, pluginID uint32, name string) (*models.Plugin, *http.Response, error) {
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

func (m *apiClient) DeleteRegistryPlugin(org string, registryID, pluginID uint32) (*http.Response, error) {
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

func (m *apiClient) ListPluginVersions(org string, registryID, pluginID uint32) ([]*models.PluginVersion, *http.Response, error) {
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

func (m *apiClient) GetPluginVersion(org string, registryID, pluginID, versionID uint32) (*models.PluginVersion, *http.Response, error) {
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

func (m *apiClient) CreatePluginVersion(org string, registryID, pluginID uint32, url string) (*models.PluginVersion, *http.Response, error) {
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

func (m *apiClient) DeletePluginVersion(org string, registryID, pluginID, versionID uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        versionsRoute(org, registryID, pluginID, fmt.Sprint(versionID)),
	}, nil)
	return resp, err
}

func (m *apiClient) InstallPluginVersion(org string, registryID, pluginID, versionID uint32, configuration map[string]string) (*http.Response, error) {
	if configuration == nil {
		configuration = map[string]string{}
	}
	body := &models.NewPluginInstall{Configuration: configuration}
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        versionsRoute(org, registryID, pluginID, fmt.Sprint(versionID), "install"),
		Body:         body,
	}, nil)
	return resp, err
}

func (m *apiClient) RetryPluginVersion(org string, registryID, pluginID, versionID uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        versionsRoute(org, registryID, pluginID, fmt.Sprint(versionID), "retry"),
	}, nil)
	return resp, err
}

func (m *apiClient) ListPluginVersionLogs(org string, registryID, pluginID, versionID uint32) ([]*models.PluginVersionLog, *http.Response, error) {
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

func (m *apiClient) ListComponentPlugins(org, project, env, component string) ([]*models.Plugin, *http.Response, error) {
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

func (m *apiClient) SetComponentPluginRelation(org, project, env, component string, pluginInstallID uint32, enabled bool) (*models.PluginRelation, *http.Response, error) {
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

// --- Plugin Widgets (org) ---

type listPluginWidgetsQuery struct {
	Placement string `url:"placement"`
}

func (m *apiClient) ListPluginWidgets(org, placement string) ([]*models.PluginWidget, *http.Response, error) {
	var result []*models.PluginWidget
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_widgets"},
		Query:        listPluginWidgetsQuery{Placement: placement},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) QueryPluginWidget(org string, widgetID uint32) (*models.PluginWidgetData, *http.Response, error) {
	var result *models.PluginWidgetData
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_widgets", fmt.Sprint(widgetID), "query"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// --- Plugin Widgets (component) ---

func (m *apiClient) ListComponentPluginWidgets(org, project, env, component string) ([]*models.PluginWidget, *http.Response, error) {
	var result []*models.PluginWidget
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "plugin_widgets"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) QueryComponentPluginWidget(org, project, env, component string, widgetID uint32) (*models.PluginWidgetData, *http.Response, error) {
	var result *models.PluginWidgetData
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "components", component, "plugin_widgets", fmt.Sprint(widgetID), "query"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
