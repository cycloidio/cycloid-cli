package apiclient

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/cycloidio/cycloid-cli/gen/models"
	"github.com/cycloidio/cycloid-cli/utils/ptr"
)

// InitFirstOrg will create the first user, org and inject the licence to the
// current console.
// If apiKeyCanonical != nil, will also create an api key admin and add it to a
// credential.
func (m *apiClient) InitFirstOrg(org, userName, fullName, email, password, licence string, apiKeyCanonical *string) (*FirstOrgData, *http.Response, error) {
	originalToken := m.api.Config.Token
	defer func() { m.api.Config.Token = originalToken }()

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

	refresh, _, err = m.RefreshToken(&org, nil, *refresh.Token)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to refresh token: %w", err)
	}
	m.api.Config.Token = *refresh.Token

	apiKeyToken, err := m.provisionAdminAPIKey(org, fullName, apiKeyCanonical)
	if isStaleTokenErr(err) {
		// Another test package's own InitFirstOrg call against this same shared org
		// (e.g. re-activating the licence) can invalidate this token mid-flight; one
		// refresh-and-retry clears the race
		//
		// RefreshToken is passed the same token that was just rejected: the rejection is
		// about the claims baked into it being stale (authenticate_user.go:269-282), not
		// about the token's signature, and RefreshToken re-derives fresh claims from it
		// rather than re-checking the old ones
		refresh, _, err = m.RefreshToken(&org, nil, *refresh.Token)
		if err != nil {
			return output, nil, fmt.Errorf("failed to refresh token: %w", err)
		}
		m.api.Config.Token = *refresh.Token
		apiKeyToken, err = m.provisionAdminAPIKey(org, fullName, apiKeyCanonical)
	}
	if err != nil {
		return output, nil, err
	}

	output.APIKey = apiKeyToken
	output.CredentialCanonical = apiKeyCanonical
	return output, nil, nil
}

// isStaleTokenErr reports whether err is the API's 403 "Need to refresh token"
// response, returned when the org/user state changed after the token was issued
func isStaleTokenErr(err error) bool {
	var apiErr *APIResponseError
	if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusForbidden || apiErr.Payload == nil {
		return false
	}
	for _, e := range apiErr.Payload.Errors {
		if e.Code != nil && *e.Code == "UnauthorizedRefreshToken" {
			return true
		}
	}
	return false
}

// provisionAdminAPIKey makes API key creation idempotent: it recreates the key
// only when the stored credential doesn't already match the current api-key
func (m *apiClient) provisionAdminAPIKey(org, fullName string, apiKeyCanonical *string) (*string, error) {
	cred, _, err := m.GetCredential(org, *apiKeyCanonical)
	var credErr *APIResponseError
	if errors.As(err, &credErr) && credErr.StatusCode != 404 {
		return nil, fmt.Errorf("api error: %w", err)
	}

	currentAPIKey, _, err := m.GetAPIKey(org, *apiKeyCanonical)
	var apiErr *APIResponseError
	if errors.As(err, &apiErr) && apiErr.StatusCode != 404 {
		return nil, fmt.Errorf("failed to fetch current apiKey %q: %w", *apiKeyCanonical, err)
	}

	var credAPIKey string
	if cred != nil {
		credAPIKey, _ = cred.Raw.Raw.(map[string]any)["key"].(string)
	}

	if cred != nil && currentAPIKey != nil && strings.HasSuffix(credAPIKey, *currentAPIKey.LastSeven) {
		return &credAPIKey, nil
	}

	if currentAPIKey != nil {
		_, err := m.DeleteAPIKey(org, *apiKeyCanonical)
		if err != nil {
			return nil, fmt.Errorf("failed to remove previous key %q: %w", *apiKeyCanonical, err)
		}
	}

	APIKey, _, err := m.CreateAPIKey(
		org, *apiKeyCanonical, "Initial api key admin", fullName, apiKeyCanonical,
		[]*models.NewRule{
			{Action: ptr.Ptr("organization:**"), Effect: ptr.Ptr("allow"), Resources: []string{}},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create api-key %q: %w", *apiKeyCanonical, err)
	}

	var credAPIErr *APIResponseError
	_, _, err = m.CreateCredential(org, *apiKeyCanonical, "custom",
		&models.CredentialRaw{Raw: map[string]string{"key": APIKey.Token}},
		"", *apiKeyCanonical, "First Admin API Key.",
	)
	credentialPreexisted := errors.As(err, &credAPIErr) && credAPIErr.StatusCode == 409
	if credentialPreexisted {
		_, _, err = m.UpdateCredential(org, *apiKeyCanonical, "custom",
			&models.CredentialRaw{Raw: map[string]string{"key": APIKey.Token}},
			"", *apiKeyCanonical, "First Admin API Key.",
		)
	}
	if err != nil {
		// Cleanup in case of failure. Only remove the credential when we created it
		// here: on the update branch it pre-existed, and a failed update leaves it
		// unchanged, so deleting it would leave the org worse off than before this call
		defer m.DeleteAPIKey(org, *apiKeyCanonical)
		if !credentialPreexisted {
			defer m.DeleteCredential(org, *apiKeyCanonical)
		}
		return nil, fmt.Errorf("failed to persist api key to credential %q: %w", *apiKeyCanonical, err)
	}

	return &APIKey.Token, nil
}
