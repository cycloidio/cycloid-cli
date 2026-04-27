package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) ListPluginVersions(org string, registryID, pluginID uint32) ([]*models.PluginVersion, *http.Response, error) {
	var result []*models.PluginVersion
	rid := strconv.FormatUint(uint64(registryID), 10)
	pid := strconv.FormatUint(uint64(pluginID), 10)
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_registries", rid, "plugins", pid, "versions"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) GetPluginVersion(org string, registryID, pluginID, versionID uint32) (*models.PluginVersion, *http.Response, error) {
	var result *models.PluginVersion
	rid := strconv.FormatUint(uint64(registryID), 10)
	pid := strconv.FormatUint(uint64(pluginID), 10)
	vid := strconv.FormatUint(uint64(versionID), 10)
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_registries", rid, "plugins", pid, "versions", vid},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) CreatePluginVersion(org string, registryID, pluginID uint32, url string) (*models.PluginVersion, *http.Response, error) {
	body := &models.NewPluginVersion{}
	body.URL = uriPtr(url)

	rid := strconv.FormatUint(uint64(registryID), 10)
	pid := strconv.FormatUint(uint64(pluginID), 10)

	var result *models.PluginVersion
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_registries", rid, "plugins", pid, "versions"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create plugin version: %w", err)
	}
	return result, resp, nil
}

func (m *middleware) DeletePluginVersion(org string, registryID, pluginID, versionID uint32) (*http.Response, error) {
	rid := strconv.FormatUint(uint64(registryID), 10)
	pid := strconv.FormatUint(uint64(pluginID), 10)
	vid := strconv.FormatUint(uint64(versionID), 10)
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_registries", rid, "plugins", pid, "versions", vid},
	}, nil)
	return resp, err
}

func (m *middleware) RetryPluginVersion(org string, registryID, pluginID, versionID uint32) (*models.PluginVersion, *http.Response, error) {
	rid := strconv.FormatUint(uint64(registryID), 10)
	pid := strconv.FormatUint(uint64(pluginID), 10)
	vid := strconv.FormatUint(uint64(versionID), 10)

	var result *models.PluginVersion
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_registries", rid, "plugins", pid, "versions", vid, "retry"},
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to retry plugin version: %w", err)
	}
	return result, resp, nil
}

func (m *middleware) GetPluginVersionLogs(org string, registryID, pluginID, versionID uint32) ([]*models.PluginVersionLog, *http.Response, error) {
	rid := strconv.FormatUint(uint64(registryID), 10)
	pid := strconv.FormatUint(uint64(pluginID), 10)
	vid := strconv.FormatUint(uint64(versionID), 10)

	var result []*models.PluginVersionLog
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_registries", rid, "plugins", pid, "versions", vid, "logs"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) InstallPluginVersion(org string, registryID, pluginID, versionID uint32, configuration map[string]string) (*models.PluginInstall, *http.Response, error) {
	rid := strconv.FormatUint(uint64(registryID), 10)
	pid := strconv.FormatUint(uint64(pluginID), 10)
	vid := strconv.FormatUint(uint64(versionID), 10)

	body := &models.NewPluginInstall{
		Configuration: configuration,
	}

	var result *models.PluginInstall
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_registries", rid, "plugins", pid, "versions", vid, "install"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to install plugin version: %w", err)
	}
	return result, resp, nil
}
