package middleware

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/go-openapi/strfmt"
)

func (m *middleware) ListMembers(org string) ([]*models.MemberOrg, *http.Response, error) {
	var result []*models.MemberOrg
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "members"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) ListInvites(org string) ([]*models.MemberOrg, *http.Response, error) {
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

func (m *middleware) GetMember(org string, id uint32) (*models.MemberOrg, *http.Response, error) {
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

func (m *middleware) DeleteMember(org string, id uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "members", strconv.FormatUint(uint64(id), 10)},
	}, nil)
	return resp, err
}

func (m *middleware) UpdateMember(org string, id uint32, role string) (*models.MemberOrg, *http.Response, error) {
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

func (m *middleware) InviteMember(org, email, role string) (*models.MemberOrg, *http.Response, error) {
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
