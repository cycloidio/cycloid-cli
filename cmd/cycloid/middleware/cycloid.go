package middleware

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/client/client/cycloid"
	"github.com/cycloidio/cycloid-cli/client/models"
	strfmt "github.com/go-openapi/strfmt"
)

// GetAppVersion returns the version of the running Cycloid server
func (m *middleware) GetAppVersion() (*models.AppVersion, error) {
	params := cycloid.NewGetAppVersionParams()

	resp, err := m.api.Cycloid.GetAppVersion(params)
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}

// GetStatus returns the status of the various Cycloid services
func (m *middleware) GetStatus() (*models.GeneralStatus, error) {
	resp, err := m.api.Cycloid.GetStatus(nil)
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, NewAPIError(err)
}
