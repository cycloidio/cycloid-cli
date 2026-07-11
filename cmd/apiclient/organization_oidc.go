package apiclient

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cycloidio/cycloid-cli/gen/models"
)

// OIDCGroupMapping maps an OIDC provider group-claim value to a Cycloid team.
// One mapping is a single (group_name, team) pair; a group that grants access
// to several teams is represented as several mappings.
//
// These types live in the apiClient package (not client/models) because the
// OIDC group-mapping endpoints are not part of the published swagger spec yet.
// Reuse models.Role for the team's roles since that model is already generated.
type OIDCGroupMapping struct {
	ID        uint32                `json:"id"`
	GroupName string                `json:"group_name"`
	Team      *OIDCGroupMappingTeam `json:"team"`
	CreatedAt uint64                `json:"created_at,omitempty"`
	UpdatedAt uint64                `json:"updated_at,omitempty"`
}

// OIDCGroupMappingTeam is the team an OIDC group maps to, as returned by the API.
type OIDCGroupMappingTeam struct {
	ID        uint32         `json:"id"`
	Canonical string         `json:"canonical"`
	Name      string         `json:"name"`
	Roles     []*models.Role `json:"roles"`
}

// NewOIDCGroupMapping is the request body to create an OIDC group mapping.
type NewOIDCGroupMapping struct {
	GroupName     *string `json:"group_name"`
	TeamCanonical *string `json:"team_canonical"`
}

// OIDCOrganizationSettings holds the per-organization OIDC reconciliation
// settings returned by the API.
type OIDCOrganizationSettings struct {
	// DefaultRoleCanonical is the org role granted to OIDC-provisioned members.
	// Empty when no default role is configured.
	DefaultRoleCanonical string `json:"default_role_canonical,omitempty"`
	// OIDCManaged, when true, makes OIDC the source of truth for membership and
	// disables local member/team/invite mutations on the org.
	OIDCManaged bool `json:"oidc_managed"`
	// OIDCNoMatchPolicy is the action for a member whose OIDC groups map to no
	// team: "keep_membership" (default) or "eject".
	OIDCNoMatchPolicy string `json:"oidc_no_match_policy"`
}

// UpdateOIDCOrganizationSettings is the request body to upsert the per-org OIDC
// settings. oidc_no_match_policy="eject" requires oidc_managed=true (the API
// rejects the combination with HTTP 422). default_role_canonical is always sent
// (an empty string clears a previously-set default role).
type UpdateOIDCOrganizationSettings struct {
	OIDCManaged          bool   `json:"oidc_managed"`
	OIDCNoMatchPolicy    string `json:"oidc_no_match_policy"`
	DefaultRoleCanonical string `json:"default_role_canonical"`
}

// ListOIDCGroupMappings returns the OIDC group->team mappings configured on the
// organization.
func (m *apiClient) ListOIDCGroupMappings(org string, filters ...LHSFilter) ([]*OIDCGroupMapping, *http.Response, error) {
	var result []*OIDCGroupMapping
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "oidc-group-mappings"},
		LHSFilters:   filters,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list OIDC group mappings: %w", err)
	}
	return result, resp, nil
}

// CreateOIDCGroupMapping creates a mapping from an OIDC group-claim value to a
// team. A group that should grant several teams needs one call per team.
func (m *apiClient) CreateOIDCGroupMapping(org, groupName, teamCanonical string) (*OIDCGroupMapping, *http.Response, error) {
	body := &NewOIDCGroupMapping{
		GroupName:     &groupName,
		TeamCanonical: &teamCanonical,
	}

	var result *OIDCGroupMapping
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "oidc-group-mappings"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create OIDC group mapping: %w", err)
	}
	return result, resp, nil
}

// DeleteOIDCGroupMapping removes an OIDC group mapping by its ID.
func (m *apiClient) DeleteOIDCGroupMapping(org string, id uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "oidc-group-mappings", strconv.FormatUint(uint64(id), 10)},
	}, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to delete OIDC group mapping: %w", err)
	}
	return resp, nil
}

// GetOIDCOrganizationSettings returns the per-org OIDC reconciliation settings.
func (m *apiClient) GetOIDCOrganizationSettings(org string) (*OIDCOrganizationSettings, *http.Response, error) {
	var result *OIDCOrganizationSettings
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "oidc-organization-settings"},
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get OIDC organization settings: %w", err)
	}
	return result, resp, nil
}

// UpdateOIDCOrganizationSettings upserts the per-org OIDC reconciliation
// settings.
func (m *apiClient) UpdateOIDCOrganizationSettings(org string, settings UpdateOIDCOrganizationSettings) (*OIDCOrganizationSettings, *http.Response, error) {
	var result *OIDCOrganizationSettings
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "oidc-organization-settings"},
		Body:         &settings,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update OIDC organization settings: %w", err)
	}
	return result, resp, nil
}

// OIDCIntegration is the AuthenticationOIDC SSO integration config as returned by
// the API. Secrets are never returned: oidc_client_secret / oidc_ca_cert are
// replaced by the read-only HasSecret / HasCaCertificate presence booleans. These
// types live in the apiClient package because several fields (oidc_discovery_url,
// oidc_allow_insecure_discovery, oidc_groups_claim_name, oidc_session_ttl_seconds)
// are not in the published swagger / generated models.AuthenticationOIDC yet.
type OIDCIntegration struct {
	Type                       string  `json:"type,omitempty"`
	Enabled                    bool    `json:"enabled"`
	OidcDisplayName            string  `json:"oidc_display_name,omitempty"`
	OidcClientID               string  `json:"oidc_client_id,omitempty"`
	OidcIssuer                 string  `json:"oidc_issuer,omitempty"`
	OidcIcon                   string  `json:"oidc_icon,omitempty"`
	OidcClientSecretJwt        bool    `json:"oidc_client_secret_jwt,omitempty"`
	OidcUseCaCert              bool    `json:"oidc_use_ca_cert,omitempty"`
	OidcSkipTLSVerify          bool    `json:"oidc_skip_tls_verify,omitempty"`
	OidcDiscoveryURL           *string `json:"oidc_discovery_url,omitempty"`
	OidcAllowInsecureDiscovery *bool   `json:"oidc_allow_insecure_discovery,omitempty"`
	OidcAdoptManualMembers     *bool   `json:"oidc_adopt_manual_members,omitempty"`
	OidcGroupsClaimName        string  `json:"oidc_groups_claim_name,omitempty"`
	OidcSessionTTLSeconds      *int64  `json:"oidc_session_ttl_seconds,omitempty"`
	// Read-only presence flags — the API never returns the secret/cert values.
	HasSecret        *bool `json:"has_secret,omitempty"`
	HasCaCertificate *bool `json:"has_ca_certificate,omitempty"`
}

// authenticationEnvelope wraps an OIDCIntegration in the API's {"config": {...}}
// Authentication envelope used by the /authentications/{type} routes.
type authenticationEnvelope struct {
	Config *OIDCIntegration `json:"config"`
}

// GetOIDCIntegration returns the org's AuthenticationOIDC SSO integration config.
// The secret and CA cert are never included (HasSecret / HasCaCertificate report
// presence only).
func (m *apiClient) GetOIDCIntegration(org string) (*OIDCIntegration, *http.Response, error) {
	var result authenticationEnvelope
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "authentications", "AuthenticationOIDC"},
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get OIDC integration: %w", err)
	}
	return result.Config, resp, nil
}

// UpdateOIDCIntegration creates-or-updates the org's AuthenticationOIDC config.
// config carries the fields to apply; always include "type" ("AuthenticationOIDC")
// and "enabled". The backend MERGES: any key absent from config keeps its stored
// value, and an absent/empty oidc_client_secret or oidc_ca_cert preserves the
// stored secret. Pass only the keys you intend to change.
func (m *apiClient) UpdateOIDCIntegration(org string, config map[string]interface{}) (*OIDCIntegration, *http.Response, error) {
	var result authenticationEnvelope
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "authentications", "AuthenticationOIDC"},
		Body:         map[string]interface{}{"config": config},
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update OIDC integration: %w", err)
	}
	return result.Config, resp, nil
}
