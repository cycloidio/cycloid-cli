package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func (m *middleware) ActivateLicence(org, licence string) error {
	// Request built by hand due to invalid api spec
	body := fmt.Sprintf(`{"key": "%s"}`, licence)
	url := fmt.Sprintf("%s/organizations/%s/licence", m.api.Config.URL, org)
	req, err := http.NewRequest(
		http.MethodPost,
		url,
		strings.NewReader(body),
	)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+m.api.GetToken(&org))

	client := http.DefaultClient
	resp, httpErr := client.Do(req)
	if httpErr != nil || resp.StatusCode != 204 {
		return &APIError{
			HTTPMethod: http.MethodPost,
			HTTPCode:   resp.Status,
			URL:        url,
			APIAction:  "activateLicence",
			Payload:    nil,
		}
	}
	defer resp.Body.Close()
	return nil
}
