package middleware

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/client/client/organization_team_members"
	"github.com/cycloidio/cycloid-cli/client/models"
	strfmt "github.com/go-openapi/strfmt"
)

func (m *middleware) ListTeamMembers(org string, team string) ([]*models.MemberTeam, error) {
	params := organization_team_members.NewGetTeamMembersParams()
	params.SetOrganizationCanonical(org)
	params.SetTeamCanonical(team)

	resp, err := m.api.OrganizationTeamMembers.GetTeamMembers(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) GetTeamMembers(org string, team string) (*models.MemberTeam, error) {
	params := organization_team_members.NewGetTeamMemberParams()
	params.SetOrganizationCanonical(org)
	params.SetTeamCanonical(team)

	resp, err := m.api.OrganizationTeamMembers.GetTeamMember(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

// AssignMemberToTeam will assign a organization member to a team using either username or email
// one of them is required
func (m *middleware) AssignMemberToTeam(org, team string, username, email *string) (*models.MemberTeam, error) {
	params := organization_team_members.NewAssignMemberToTeamParams()
	params.SetOrganizationCanonical(org)
	body := &models.NewTeamMemberAssignation{}

	if username == nil && email == nil {
		return nil, fmt.Errorf("missing email or username for AssignMemberToTeam")
	}

	if username != nil {
		body.Username = *username
	}

	if email != nil {
		body.Email = strfmt.Email(*email)
	}
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	params.SetBody(body)
	resp, err := m.api.OrganizationTeamMembers.AssignMemberToTeam(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}
