package apiclient

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/gen/models"
)

// ListMembers lists organization members.
//
// Supported LHS filter attributes: user_canonical, user_full_name,
// invitation_state, invitation_created_at, role_name.
func (m *apiClient) ListMembers(org string, filters ...LHSFilter) ([]*models.MemberOrg, *http.Response, error) {
	var result []*models.MemberOrg
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "members"},
		LHSFilters:   filters,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) ListInvites(org string) ([]*models.MemberOrg, *http.Response, error) {
	query := url.Values{"invitation_state": []string{"pending"}}

	var result []*models.MemberOrg
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "members"},
		Query:        query,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) GetMember(org string, id uint32) (*models.MemberOrg, *http.Response, error) {
	var result *models.MemberOrg
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "members", strconv.FormatUint(uint64(id), 10)},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) DeleteMember(org string, id uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "members", strconv.FormatUint(uint64(id), 10)},
	}, nil)
	return resp, err
}

func (m *apiClient) UpdateMember(org string, id uint32, role string) (*models.MemberOrg, *http.Response, error) {
	body := &models.MemberAssignation{
		RoleCanonical: &role,
	}

	var result *models.MemberOrg
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "members", strconv.FormatUint(uint64(id), 10)},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) InviteMember(org, email, role string) (*models.MemberOrg, *http.Response, error) {
	fmtEmail := strfmt.Email(email)
	body := &models.NewMemberInvitation{
		Email:         &fmtEmail,
		RoleCanonical: &role,
	}

	var result *models.MemberOrg
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "members"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to invite member: %w", err)
	}
	return result, resp, nil
}
