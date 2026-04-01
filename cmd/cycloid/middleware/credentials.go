package middleware

import (
	"net/http"
	"net/url"
	"regexp"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) CreateCredential(org, name, credentialType string, rawCred *models.CredentialRaw, path, canonical, description string) (*models.Credential, *http.Response, error) {
	if path == "" {
		re := regexp.MustCompile(`[^a-zA-z0-9_\-./]`)
		path = re.ReplaceAllString(name, "-")
	}

	body := &models.NewCredential{
		Description: description,
		Name:        &name,
		Path:        &path,
		Raw:         rawCred,
		Type:        &credentialType,
		Canonical:   canonical,
	}

	var result *models.Credential
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "credentials"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) UpdateCredential(org, name, credentialType string, rawCred *models.CredentialRaw, path, canonical, description string) (*models.Credential, *http.Response, error) {
	if path == "" {
		re := regexp.MustCompile(`[^a-zA-z0-9_\-./]`)
		path = re.ReplaceAllString(name, "-")
	}

	body := &models.UpdateCredential{
		Description: description,
		Name:        &name,
		Path:        &path,
		Raw:         rawCred,
		Type:        &credentialType,
		Canonical:   &canonical,
	}

	var result *models.Credential
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "credentials", canonical},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) GetCredential(org, credential string) (*models.Credential, *http.Response, error) {
	var result *models.Credential
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "credentials", credential},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) DeleteCredential(org, credential string) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "credentials", credential},
	}, nil)
	return resp, err
}

func (m *middleware) ListCredentials(org, credentialType string) ([]*models.CredentialSimple, *http.Response, error) {
	// Explicit pagination matches swagger defaults (page_size 1000) so callers such as
	// `credential create --update` do not miss an existing credential on backends that
	// use a smaller default page size.
	query := url.Values{
		"page_index": []string{"1"},
		"page_size":  []string{"1000"},
	}
	if credentialType != "" {
		query.Set("credential_type", credentialType)
	}

	var result []*models.CredentialSimple
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "credentials"},
		Query:        query,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
