package middleware

import (
	"fmt"
	"net/http"
)

// NavItem is a single sidebar entry with an explicit position.
//
// This type lives in the middleware package (not client/models) following the
// same convention as OIDCOrganizationSettings: the endpoint is simple enough
// (two fields, a GET/PUT pair) that a hand-defined type plus GenericRequest is
// more direct than swagger-driven codegen for it.
type NavItem struct {
	// Type is either "native" (a built-in section, identified by name, e.g.
	// "dashboard") or "plugin_widget" (identified by the widget's plugin ID as
	// a string).
	Type string `json:"type"`
	Key  string `json:"key"`
	// Position is 1-indexed; positions must be unique within a NavConfig.
	Position uint32 `json:"position"`
}

// NavConfig is the per-organization sidebar nav ordering configuration.
// Items is empty when no ordering has been saved — the console falls back to
// its default ordering in that case.
type NavConfig struct {
	Items []*NavItem `json:"items"`
}

// GetOrgNav returns the org's sidebar nav ordering config.
func (m *middleware) GetOrgNav(org string) (*NavConfig, *http.Response, error) {
	var result *NavConfig
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "nav"},
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get organization nav ordering: %w", err)
	}
	return result, resp, nil
}

// UpdateOrgNav creates or replaces the org's sidebar nav ordering config.
// Passing an empty (or nil) items slice resets the ordering to defaults.
func (m *middleware) UpdateOrgNav(org string, items []*NavItem) (*NavConfig, *http.Response, error) {
	body := &NavConfig{Items: items}
	if body.Items == nil {
		body.Items = []*NavItem{}
	}

	var result *NavConfig
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "nav"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update organization nav ordering: %w", err)
	}
	return result, resp, nil
}
