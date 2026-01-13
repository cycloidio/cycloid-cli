package middleware

import (
	"fmt"

	"github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/client/client/organization_api_keys"
	"github.com/cycloidio/cycloid-cli/client/models"
)

// ListAPIKey will request API to list generated API keys
func (m *middleware) ListAPIKeys(org string) ([]*models.APIKey, error) {
	params := organization_api_keys.NewGetAPIKeysParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationAPIKeys.GetAPIKeys(params, m.api.Credentials(&org))
	if err != nil {
		return nil, fmt.Errorf("unable to list API keys: %w", NewAPIError(err))
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

// GetAPIKey will request API to get a specified generated API key by its canonical
func (m *middleware) GetAPIKey(org, canonical string) (*models.APIKey, error) {
	params := organization_api_keys.NewGetAPIKeyParams()
	params.SetOrganizationCanonical(org)
	params.SetAPIKeyCanonical(canonical)

	resp, err := m.api.OrganizationAPIKeys.GetAPIKey(params, m.api.Credentials(&org))
	if err != nil {
		return nil, fmt.Errorf("unable to get API key: %w", NewAPIError(err))
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

// GetAPIKey will request API to get a specified generated API key by its canonical
func (m *middleware) CreateAPIKey(org, canonical, description, owner string, name *string, rules []*models.NewRule) (*models.APIKey, error) {
	params := organization_api_keys.NewCreateAPIKeyParams()
	params.SetOrganizationCanonical(org)
	body := models.NewAPIKey{
		Canonical:   canonical,
		Name:        name,
		Description: description,
		Rules:       rules,
	}

	if owner != "" {
		body.Owner = owner
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, fmt.Errorf("invalid body for createAPIKey: %w", err)
	}

	params.SetBody(&body)

	resp, err := m.api.OrganizationAPIKeys.CreateAPIKey(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

// DeleteAPIKey will request API to delete a specified generated API key
func (m *middleware) DeleteAPIKey(org, canonical string) error {
	params := organization_api_keys.NewDeleteAPIKeyParams()
	params.SetOrganizationCanonical(org)
	params.SetAPIKeyCanonical(canonical)

	if _, err := m.api.OrganizationAPIKeys.DeleteAPIKey(params, m.api.Credentials(&org)); err != nil {
		return fmt.Errorf("unable to delete API key: %w", NewAPIError(err))
	}
	return nil
}
