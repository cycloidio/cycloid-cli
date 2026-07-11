package apiclient

import (
	"fmt"
	"net/http"

	"github.com/cycloidio/cycloid-cli/gen/models"
	"github.com/cycloidio/cycloid-cli/utils/ptr"
)

// ListRoles lists roles for an organization.
//
// NOTE: the backend handler for this route does not call lhs.ParseQuery, so
// LHS filters are accepted by the apiClient but silently ignored server-side.
func (m *apiClient) ListRoles(org string, filters ...LHSFilter) ([]*models.Role, *http.Response, error) {
	var result []*models.Role
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "roles"},
		LHSFilters:   filters,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) GetRole(org, role string) (*models.Role, *http.Response, error) {
	var result *models.Role
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "roles", role},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) DeleteRole(org, role string) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "roles", role},
	}, nil)
	return resp, err
}

// CreateRole requires org, name or canonical and rules
func (m *apiClient) CreateRole(org string, name, canonical, description *string, rules []*models.NewRule) (*models.NewRole, *http.Response, error) {
	n, c, err := NameOrCanonical(name, canonical)
	if err != nil {
		return nil, nil, err
	}

	body := &models.NewRole{
		Name:        &n,
		Canonical:   c,
		Description: ptr.Value(description),
		Rules:       rules,
	}

	var result *models.NewRole
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "roles"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// UpdateRole updates an existing role (PUT). roleCanonical must match the resolved canonical from name/--role.
func (m *apiClient) UpdateRole(org, roleCanonical string, name, canonical, description *string, rules []*models.NewRule) (*models.Role, *http.Response, error) {
	n, c, err := NameOrCanonical(name, canonical)
	if err != nil {
		return nil, nil, err
	}
	if c != roleCanonical {
		return nil, nil, fmt.Errorf("role canonical %q must match update target %q", c, roleCanonical)
	}

	body := &models.NewRole{
		Name:        &n,
		Canonical:   c,
		Description: ptr.Value(description),
		Rules:       rules,
	}

	var result *models.Role
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "roles", roleCanonical},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
