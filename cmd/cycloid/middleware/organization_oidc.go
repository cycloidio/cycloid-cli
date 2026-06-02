package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cycloidio/cycloid-cli/client/models"
)

// OIDCGroupMapping maps an OIDC provider group-claim value to a Cycloid team.
// One mapping is a single (group_name, team) pair; a group that grants access
// to several teams is represented as several mappings.
//
// These types live in the middleware package (not client/models) because the
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
func (m *middleware) ListOIDCGroupMappings(org string, filters ...LHSFilter) ([]*OIDCGroupMapping, *http.Response, error) {
	var result []*OIDCGroupMapping
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "oidc-group-mappings"},
		LHSFilters:   filters,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// CreateOIDCGroupMapping creates a mapping from an OIDC group-claim value to a
// team. A group that should grant several teams needs one call per team.
func (m *middleware) CreateOIDCGroupMapping(org, groupName, teamCanonical string) (*OIDCGroupMapping, *http.Response, error) {
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
func (m *middleware) DeleteOIDCGroupMapping(org string, id uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "oidc-group-mappings", strconv.FormatUint(uint64(id), 10)},
	}, nil)
	return resp, err
}

// GetOIDCOrganizationSettings returns the per-org OIDC reconciliation settings.
func (m *middleware) GetOIDCOrganizationSettings(org string) (*OIDCOrganizationSettings, *http.Response, error) {
	var result *OIDCOrganizationSettings
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "oidc-organization-settings"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// UpdateOIDCOrganizationSettings upserts the per-org OIDC reconciliation
// settings.
func (m *middleware) UpdateOIDCOrganizationSettings(org string, settings UpdateOIDCOrganizationSettings) (*OIDCOrganizationSettings, *http.Response, error) {
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
