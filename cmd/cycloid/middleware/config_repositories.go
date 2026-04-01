package middleware

import (
	"net/http"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) ListConfigRepositories(org string) ([]*models.ConfigRepository, *http.Response, error) {
	var result []*models.ConfigRepository
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "config_repositories"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) GetConfigRepository(org, configRepo string) (*models.ConfigRepository, *http.Response, error) {
	var result *models.ConfigRepository
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "config_repositories", configRepo},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) DeleteConfigRepository(org, configRepo string) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "config_repositories", configRepo},
	}, nil)
	return resp, err
}

func (m *middleware) CreateConfigRepository(org, name, canonical, url, branch, cred string, setDefault bool) (*models.ConfigRepository, *http.Response, error) {
	if name == "" {
		name = canonical
	}

	body := &models.NewConfigRepository{
		Branch:    &branch,
		Canonical: canonical,
		Default:   &setDefault,
		Name:      &name,
		URL:       &url,
	}

	if cred != "" {
		body.CredentialCanonical = &cred
	}

	var result *models.ConfigRepository
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "config_repositories"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) UpdateConfigRepository(org, configRepo, cred, name, url, branch string, setDefault bool) (*models.ConfigRepository, *http.Response, error) {
	body := &models.UpdateConfigRepository{
		Branch:              &branch,
		CredentialCanonical: &cred,
		Default:             &setDefault,
		Name:                &name,
		URL:                 &url,
	}

	var result *models.ConfigRepository
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "config_repositories", configRepo},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
