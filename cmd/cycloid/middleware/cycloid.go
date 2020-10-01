package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client/cycloid"
	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) GetAppVersion() (*models.AppVersion, error) {

	params := cycloid.NewGetAppVersionParams()

	resp, err := m.api.Cycloid.GetAppVersion(params)
	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()
	// TODO this validate have been removed https://github.com/cycloidio/youdeploy-http-api/issues/2262
	// err = p.Validate(strfmt.Default)
	// if err != nil {
	// 	return err
	// }

	d := p.Data

	return d, err
}
