package middleware

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
)

// InitFirstOrg will create the first user, org and inject the licence to the
// current console.
// If apiKeyCanonical != nil, will also create an api key admin and add it to a
// credential.
func (m *middleware) InitFirstOrg(org, userName, givenName, famillyName, email, password, licence string, apiKeyCanonical *string) (*FirstOrgData, error) {
	err := m.UserSignup(userName, email, password, givenName, famillyName)
	var signupErr *APIError
	if errors.As(err, &signupErr) {
		if signupErr.HTTPCode != "409" && err != nil {
			return nil, fmt.Errorf("failed to signup first user: %w", err)
		}
	}

	login, err := m.UserLogin(&org, &email, &userName, password)
	if err != nil {
		return nil, fmt.Errorf("failed to login with admin user: %w", err)
	}
	m.api.Config.Token = *login.Token

	_, err = m.CreateOrganization(org)
	var orgErr *APIError
	if errors.As(err, &orgErr) {
		if orgErr.HTTPCode != "409" && err != nil {
			return nil, fmt.Errorf("failed to create first org: %w", err)
		}
	}

	refresh, err := m.RefreshToken(&org, nil, *login.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}
	m.api.Config.Token = *refresh.Token

	err = m.ActivateLicence(org, licence)
	if err != nil {
		return nil, fmt.Errorf("failed to activate Licence: %w", err)
	}

	output := &FirstOrgData{
		Org:         org,
		UserName:    userName,
		FamillyName: famillyName,
		GivenName:   givenName,
		Email:       email,
		Password:    password,
		Token:       *refresh.Token,
	}

	if apiKeyCanonical == nil {
		return output, nil
	}

	refresh, err = m.RefreshToken(&org, nil, *login.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}
	m.api.Config.Token = *refresh.Token

	// to make API key creation idempotent, let's recreate it everytime this runs
	cred, err := m.GetCredential(org, *apiKeyCanonical)
	var credErr *APIError
	if errors.As(err, &credErr) && credErr.HTTPCode != "404" {
		return output, fmt.Errorf("api error: %w", err)
	}

	currentAPIKey, err := m.GetAPIKey(org, *apiKeyCanonical)
	var apiErr *APIError
	if errors.As(err, &apiErr) && apiErr.HTTPCode != "404" {
		return output, fmt.Errorf("failed to fetch current apiKey %q: %w", *apiKeyCanonical, err)
	}

	var credAPIKey string
	if cred != nil {
		credAPIKey, _ = cred.Raw.Raw.(map[string]any)["key"].(string)
	}

	if cred == nil || !strings.HasSuffix(credAPIKey, *currentAPIKey.LastSeven) {
		if currentAPIKey != nil {
			err := m.DeleteAPIKey(org, *apiKeyCanonical)
			if err != nil {
				return output, fmt.Errorf("failed to remove previous key %q: %w", *apiKeyCanonical, err)
			}
		}

		APIKey, err := m.CreateAPIKey(
			org, *apiKeyCanonical, "Initial api key admin", userName, apiKeyCanonical,
			[]*models.NewRule{
				{Action: ptr.Ptr("organization:**"), Effect: ptr.Ptr("allow"), Resources: []string{}},
			},
		)
		if err != nil {
			return output, fmt.Errorf("failed to create api-key %q: %w", *apiKeyCanonical, err)
		}

		output.APIKey = &APIKey.Token

		var credAPIErr *APIError
		_, err = m.CreateCredential(org, *apiKeyCanonical, "custom",
			&models.CredentialRaw{Raw: map[string]string{"key": APIKey.Token}},
			"", *apiKeyCanonical, "First Admin API Key.",
		)
		if errors.As(err, &credAPIErr) && credAPIErr.HTTPCode == "409" {
			_, err = m.UpdateCredential(org, *apiKeyCanonical, "custom",
				&models.CredentialRaw{Raw: map[string]string{"key": APIKey.Token}},
				"", *apiKeyCanonical, "First Admin API Key.",
			)
		}
		if err != nil {
			// Cleanup in case of failure
			defer m.DeleteAPIKey(org, *apiKeyCanonical)
			defer m.DeleteCredential(org, *apiKeyCanonical)
			return output, fmt.Errorf("failed to persist api key to credential %q: %w", *apiKeyCanonical, err)
		}
	} else {
		output.APIKey = &credAPIKey
	}

	output.CredentialCanonical = apiKeyCanonical
	return output, nil
}
