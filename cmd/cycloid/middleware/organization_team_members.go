package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) ListTeamMembers(org string, team string) ([]*models.MemberTeam, *http.Response, error) {
	var result []*models.MemberTeam
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "teams", team, "members"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) GetTeamMember(org string, team string, memberID uint32) (*models.MemberTeam, *http.Response, error) {
	var result *models.MemberTeam
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "teams", team, "members", strconv.FormatUint(uint64(memberID), 10)},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// AssignMemberToTeam will assign a organization member to a team using either username or email
// one of them is required
func (m *middleware) AssignMemberToTeam(org, team string, username, email *string) (*models.MemberTeam, *http.Response, error) {
	if username == nil && email == nil {
		return nil, nil, fmt.Errorf("missing email or username for AssignMemberToTeam")
	}

	body := map[string]string{}
	if username != nil {
		body["username"] = *username
	} else {
		body["email"] = *email
	}

	var result *models.MemberTeam
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "teams", team, "members"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) UnAssignMemberFromTeam(org, team string, memberID uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "teams", team, "members", strconv.FormatUint(uint64(memberID), 10)},
	}, nil)
	return resp, err
}
