package middleware

import (
	"net/http"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) GetLicence(org string) (*models.Licence, *http.Response, error) {
	var result *models.Licence
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "licence"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) ActivateLicence(org, licence string) (*http.Response, error) {
	body := map[string]string{"key": licence}
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "licence"},
		Body:         body,
	}, nil)
	return resp, err
}
