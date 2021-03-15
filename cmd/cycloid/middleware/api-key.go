package middleware

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/client/client/organization_api_keys"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
)

// ListAPIKey will request API to list generated API keys
func (m *middleware) ListAPIKey(org string) ([]*models.APIKey, error) {
	params := organization_api_keys.NewGetAPIKeysParams()
	params.SetOrganizationCanonical(org)

	res, err := m.api.OrganizationAPIKeys.GetAPIKeys(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, fmt.Errorf("unable to list API keys: %w", err)
	}

	return res.GetPayload().Data, nil
}

// GetAPIKey will request API to get a specified generated API key by its canonical
func (m *middleware) GetAPIKey(org, canonical string) (*models.APIKey, error) {
	params := organization_api_keys.NewGetAPIKeyParams()
	params.SetOrganizationCanonical(org)
	params.SetAPIKeyCanonical(canonical)

	res, err := m.api.OrganizationAPIKeys.GetAPIKey(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, fmt.Errorf("unable to get API key: %w", err)
	}

	return res.GetPayload().Data, nil
}

// DeleteAPIKey will request API to delete a specified generated API key
func (m *middleware) DeleteAPIKey(org, canonical string) error {
	params := organization_api_keys.NewDeleteAPIKeyParams()
	params.SetOrganizationCanonical(org)
	params.SetAPIKeyCanonical(canonical)

	if _, err := m.api.OrganizationAPIKeys.DeleteAPIKey(params, common.ClientCredentials(&org)); err != nil {
		return fmt.Errorf("unable to delete API key: %w", err)
	}
	return nil
}
