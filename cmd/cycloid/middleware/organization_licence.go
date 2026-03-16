package middleware

import "net/http"

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
