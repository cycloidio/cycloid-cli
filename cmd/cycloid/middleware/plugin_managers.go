package middleware

import (
	"fmt"
	"net/http"

	"github.com/cycloidio/cycloid-cli/client/models"
)

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

func (m *middleware) CreatePluginManager(org, name, url string) (*models.PluginManager, *http.Response, error) {
	// Use a map body because the generated model may not include auto_register
	body := map[string]any{
		"name":          name,
		"url":           url,
		"auto_register": true,
	}

	var result *models.PluginManager
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_managers"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create plugin manager: %w", err)
	}
	return result, resp, nil
}
