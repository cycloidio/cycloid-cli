package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
)

// InitFirstOrg will create the first user, org and inject the licence to the
// current console.
// If apiKeyCanonical != nil, will also create an api key admin and add it to a
// credential.
func (m *middleware) InitFirstOrg(org, userName, fullName, email, password, licence string, apiKeyCanonical *string) (*FirstOrgData, *http.Response, error) {
	_, err := m.UserSignup(userName, email, password, fullName)
	var signupErr *APIResponseError
	if errors.As(err, &signupErr) {
		if signupErr.StatusCode != 409 && err != nil {
			return nil, nil, fmt.Errorf("failed to signup first user: %w", err)
		}
	}

	login, _, err := m.UserLogin(&org, &email, password)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to login with admin user: %w", err)
	}
	m.api.Config.Token = *login.Token

	_, _, err = m.CreateOrganization(org)
	var orgErr *APIResponseError
	if errors.As(err, &orgErr) {
		if orgErr.StatusCode != 409 && err != nil {
			return nil, nil, fmt.Errorf("failed to create first org: %w", err)
		}
	}

	refresh, _, err := m.RefreshToken(&org, nil, *login.Token)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to refresh token: %w", err)
	}
	m.api.Config.Token = *refresh.Token

	_, err = m.ActivateLicence(org, licence)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to activate Licence: %w", err)
	}

	output := &FirstOrgData{
		Org:      org,
		Username: userName,
		FullName: fullName,
		Email:    email,
		Password: password,
		Token:    *refresh.Token,
	}

	if apiKeyCanonical == nil {
		return output, nil, nil
	}

	refresh, _, err = m.RefreshToken(&org, nil, *login.Token)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to refresh token: %w", err)
	}
	m.api.Config.Token = *refresh.Token

	// to make API key creation idempotent, let's recreate it everytime this runs
	cred, _, err := m.GetCredential(org, *apiKeyCanonical)
	var credErr *APIResponseError
	if errors.As(err, &credErr) && credErr.StatusCode != 404 {
		return output, nil, fmt.Errorf("api error: %w", err)
	}

	currentAPIKey, _, err := m.GetAPIKey(org, *apiKeyCanonical)
	var apiErr *APIResponseError
	if errors.As(err, &apiErr) && apiErr.StatusCode != 404 {
		return output, nil, fmt.Errorf("failed to fetch current apiKey %q: %w", *apiKeyCanonical, err)
	}

	var credAPIKey string
	if cred != nil {
		credAPIKey, _ = cred.Raw.Raw.(map[string]any)["key"].(string)
	}

	if cred == nil || !strings.HasSuffix(credAPIKey, *currentAPIKey.LastSeven) {
		if currentAPIKey != nil {
			_, err := m.DeleteAPIKey(org, *apiKeyCanonical)
			if err != nil {
				return output, nil, fmt.Errorf("failed to remove previous key %q: %w", *apiKeyCanonical, err)
			}
		}

		APIKey, _, err := m.CreateAPIKey(
			org, *apiKeyCanonical, "Initial api key admin", fullName, apiKeyCanonical,
			[]*models.NewRule{
				{Action: ptr.Ptr("organization:**"), Effect: ptr.Ptr("allow"), Resources: []string{}},
			},
		)
		if err != nil {
			return output, nil, fmt.Errorf("failed to create api-key %q: %w", *apiKeyCanonical, err)
		}

		output.APIKey = &APIKey.Token

		var credAPIErr *APIResponseError
		_, _, err = m.CreateCredential(org, *apiKeyCanonical, "custom",
			&models.CredentialRaw{Raw: map[string]string{"key": APIKey.Token}},
			"", *apiKeyCanonical, "First Admin API Key.",
		)
		if errors.As(err, &credAPIErr) && credAPIErr.StatusCode == 409 {
			_, _, err = m.UpdateCredential(org, *apiKeyCanonical, "custom",
				&models.CredentialRaw{Raw: map[string]string{"key": APIKey.Token}},
				"", *apiKeyCanonical, "First Admin API Key.",
			)
		}
		if err != nil {
			// Cleanup in case of failure
			defer m.DeleteAPIKey(org, *apiKeyCanonical)
			defer m.DeleteCredential(org, *apiKeyCanonical)
			return output, nil, fmt.Errorf("failed to persist api key to credential %q: %w", *apiKeyCanonical, err)
		}
	} else {
		output.APIKey = &credAPIKey
	}

	output.CredentialCanonical = apiKeyCanonical
	return output, nil, nil
}
