package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cycloidio/cycloid-cli/client/models"
)

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

func (m *middleware) CreatePluginRegistry(org, name, url string) (*models.PluginRegistry, *http.Response, error) {
	body := &models.NewPluginRegistry{
		Name: &name,
	}
	body.URL = uriPtr(url)

	var result *models.PluginRegistry
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_registries"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create plugin registry: %w", err)
	}
	return result, resp, nil
}

func (m *middleware) UpdatePluginRegistry(org string, registryID uint32, name string) (*models.PluginRegistry, *http.Response, error) {
	body := &models.UpdatePluginRegistry{
		Name: &name,
	}

	var result *models.PluginRegistry
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_registries", strconv.FormatUint(uint64(registryID), 10)},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update plugin registry: %w", err)
	}
	return result, resp, nil
}

func (m *middleware) DeletePluginRegistry(org string, registryID uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_registries", strconv.FormatUint(uint64(registryID), 10)},
	}, nil)
	return resp, err
}
