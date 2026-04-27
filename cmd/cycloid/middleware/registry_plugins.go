package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) GetRegistryPlugin(org string, registryID, pluginID uint32) (*models.Plugin, *http.Response, error) {
	var result *models.Plugin
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_registries", strconv.FormatUint(uint64(registryID), 10), "plugins", strconv.FormatUint(uint64(pluginID), 10)},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) CreateRegistryPlugin(org string, registryID uint32, name string) (*models.Plugin, *http.Response, error) {
	body := &models.NewPlugin{
		Name: &name,
	}

	var result *models.Plugin
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_registries", strconv.FormatUint(uint64(registryID), 10), "plugins"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create registry plugin: %w", err)
	}
	return result, resp, nil
}

func (m *middleware) UpdateRegistryPlugin(org string, registryID, pluginID uint32, name string) (*models.Plugin, *http.Response, error) {
	body := &models.UpdatePlugin{
		Name: &name,
	}

	var result *models.Plugin
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_registries", strconv.FormatUint(uint64(registryID), 10), "plugins", strconv.FormatUint(uint64(pluginID), 10)},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update registry plugin: %w", err)
	}
	return result, resp, nil
}

func (m *middleware) DeleteRegistryPlugin(org string, registryID, pluginID uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_registries", strconv.FormatUint(uint64(registryID), 10), "plugins", strconv.FormatUint(uint64(pluginID), 10)},
	}, nil)
	return resp, err
}
