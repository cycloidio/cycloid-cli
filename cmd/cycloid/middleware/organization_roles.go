package middleware

import (
	"net/http"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
)

func (m *middleware) ListRoles(org string) ([]*models.Role, *http.Response, error) {
	var result []*models.Role
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "roles"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) GetRole(org, role string) (*models.Role, *http.Response, error) {
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

func (m *middleware) DeleteRole(org, role string) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "roles", role},
	}, nil)
	return resp, err
}

// CreateRole requires org, name or canonical and rules
func (m *middleware) CreateRole(org string, name, canonical, description *string, rules []*models.NewRule) (*models.NewRole, *http.Response, error) {
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
