package middleware

import (
	"net/http"

	"github.com/cycloidio/cycloid-cli/client/models"
)

// GetAppVersion returns the version of the running Cycloid server
func (m *middleware) GetAppVersion() (*models.AppVersion, *http.Response, error) {
	var result *models.AppVersion
	resp, err := m.GenericRequest(Request{
		Method: "GET",
		NoAuth: true,
		Route:  []string{"version"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// GetStatus returns the status of the various Cycloid services
func (m *middleware) GetStatus() (*models.GeneralStatus, *http.Response, error) {
	var result *models.GeneralStatus
	resp, err := m.GenericRequest(Request{
		Method: "GET",
		NoAuth: true,
		Route:  []string{"status"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
