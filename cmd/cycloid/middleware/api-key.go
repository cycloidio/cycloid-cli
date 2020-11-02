package middleware

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/client/client/organization_api_keys"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
)

// CreateAPIKey will request API to generate and return an API key
func (m *middleware) CreateAPIKey(org, name, canonical, description string, roleID uint32) (*models.APIKey, error) {
	params := organization_api_keys.NewCreateAPIKeyParams()

	body := &models.NewAPIKey{
		Canonical:   &canonical,
		Description: description,
		Name:        &name,
		RoleID:      &roleID,
	}

	params.SetBody(body)
	params.SetOrganizationCanonical(org)

	res, err := m.api.OrganizationAPIKeys.CreateAPIKey(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, fmt.Errorf("unable to create API key: %w", err)
	}

	return res.GetPayload().Data, nil
}

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
