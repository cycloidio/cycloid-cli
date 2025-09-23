package middleware

import (
	"fmt"

	"github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/client/client/organization_licence"
	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) ActivateLicence(org, licence string) error {
	params := organization_licence.NewActivateLicenceParams()
	params.WithOrganizationCanonical(org)
	body := models.NewLicence{
		Key: &licence,
	}
	err := body.Validate(strfmt.Default)
	if err != nil {
		return fmt.Errorf("invalid body for activateLicence: %v", err)
	}

	params.WithBody(&body)
	_, err = m.api.OrganizationLicence.ActivateLicence(params)
	if err != nil {
		NewApiError(err)
	}

	return nil
}
