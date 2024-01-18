package middleware

import (
	"github.com/pkg/errors"

	"github.com/cycloidio/cycloid-cli/client/client/cycloid"
	"github.com/cycloidio/cycloid-cli/client/models"
)

// GetAppVersion returns the version of the running Cycloid server
func (m *middleware) GetAppVersion() (*models.AppVersion, error) {

	params := cycloid.NewGetAppVersionParams()

	resp, err := m.api.Cycloid.GetAppVersion(params)
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()

	d := p.Data

	return d, nil
}

// GetStatus returns the status of the various Cycloid services
func (m *middleware) GetStatus() (*models.GeneralStatus, error) {
	resp, err := m.api.Cycloid.GetStatus(nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get answer from the API")
	}

	p := resp.GetPayload()

	return p.Data, NewApiError(err)
}
