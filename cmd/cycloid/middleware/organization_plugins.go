package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) ListPlugins(org string) ([]*models.Plugin, *http.Response, error) {
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

func (m *middleware) GetPlugin(org string, installID uint32) (*models.Plugin, *http.Response, error) {
	var result *models.Plugin
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "plugins", strconv.FormatUint(uint64(installID), 10)},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) UpdatePlugin(org string, installID, versionID uint32, configuration map[string]string) (*models.Plugin, *http.Response, error) {
	body := &models.UpdatePluginInstall{
		PluginVersionID: &versionID,
		Configuration:   configuration,
	}

	var result *models.Plugin
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "plugins", strconv.FormatUint(uint64(installID), 10)},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update plugin: %w", err)
	}
	return result, resp, nil
}

func (m *middleware) DeletePlugin(org string, installID uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "plugins", strconv.FormatUint(uint64(installID), 10)},
	}, nil)
	return resp, err
}

func (m *middleware) GetPluginLogs(org string, installID uint32) (*models.PluginInstallLog, *http.Response, error) {
	var result *models.PluginInstallLog
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "plugins", strconv.FormatUint(uint64(installID), 10), "logs"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) ListPluginWidgets(org, placement string) ([]*models.PluginWidget, *http.Response, error) {
	type query struct {
		Placement string `url:"placement,omitempty"`
	}

	var result []*models.PluginWidget
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "plugin_widgets"},
		Query:        query{Placement: placement},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
