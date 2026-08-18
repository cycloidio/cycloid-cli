package apiclient

import (
	"fmt"
	"net/http"

	"github.com/cycloidio/cycloid-cli/gen/models"
)

// ListAPIKeys lists API keys for an organization.
//
// Supported LHS filter attributes: organization_canonical, user_canonical.
func (m *apiClient) ListAPIKeys(org string, filters ...LHSFilter) ([]*models.APIKey, *http.Response, error) {
	var result []*models.APIKey
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "api_keys"},
		LHSFilters:   filters,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("unable to list API keys: %w", err)
	}
	return result, resp, nil
}

// GetAPIKey will request API to get a specified generated API key by its canonical
func (m *apiClient) GetAPIKey(org, canonical string) (*models.APIKey, *http.Response, error) {
	var result *models.APIKey
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "api_keys", canonical},
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("unable to get API key: %w", err)
	}
	return result, resp, nil
}

// CreateAPIKey will request API to create an API key
func (m *apiClient) CreateAPIKey(org, canonical, description, owner string, name *string, rules []*models.NewRule) (*models.APIKey, *http.Response, error) {
	body := models.NewAPIKey{
		Canonical:   canonical,
		Name:        name,
		Description: description,
		Rules:       rules,
	}

	if owner != "" {
		body.Owner = owner
	}

	var result *models.APIKey
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "api_keys"},
		Body:         &body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// DeleteAPIKey will request API to delete a specified generated API key
func (m *apiClient) DeleteAPIKey(org, canonical string) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "api_keys", canonical},
	}, nil)
	if err != nil {
		return resp, fmt.Errorf("unable to delete API key: %w", err)
	}
	return resp, nil
}
