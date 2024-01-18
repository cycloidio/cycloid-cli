package middleware

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/client/client/organization_invitations"
	"github.com/cycloidio/cycloid-cli/client/client/organization_members"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
)

func (m *middleware) ListMembers(org string) ([]*models.MemberOrg, error) {
	params := organization_members.NewGetOrgMembersParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationMembers.GetOrgMembers(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()

	err = p.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	d := p.Data

	return d, nil
}

func (m *middleware) ListInvites(org string) ([]*models.Invitation, error) {
	params := organization_invitations.NewGetInvitationsParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationInvitations.GetInvitations(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()

	d := p.Data

	return d, nil
}

func (m *middleware) GetMember(org string, name string) (*models.MemberOrg, error) {
	params := organization_members.NewGetOrgMemberParams()
	params.SetOrganizationCanonical(org)
	params.SetUsername(name)

	resp, err := m.api.OrganizationMembers.GetOrgMember(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()

	err = p.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	d := p.Data

	return d, nil
}

func (m *middleware) DeleteMember(org string, name string) error {
	params := organization_members.NewRemoveOrgMemberParams()
	params.SetOrganizationCanonical(org)
	params.SetUsername(name)

	_, err := m.api.OrganizationMembers.RemoveOrgMember(params, common.ClientCredentials(&org))
	if err != nil {
		return err
	}

	return nil
}

func (m *middleware) UpdateMembers(org, name, role string) (*models.MemberOrg, error) {
	params := organization_members.NewUpdateOrgMemberParams()

	params.SetOrganizationCanonical(org)
	params.SetUsername(name)

	body := &models.MemberAssignation{
		RoleCanonical: &role,
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	resp, err := m.api.OrganizationMembers.UpdateOrgMember(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	//TODO verify why getpayload no defined ?
	p := resp.GetPayload()

	err = p.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	d := p.Data

	return d, nil
}

func (m *middleware) InviteMember(org, email, role string) error {
	params := organization_members.NewInviteUserToOrgMemberParams()
	params.SetOrganizationCanonical(org)

	fmtEmail := strfmt.Email(email)
	body := &models.NewMemberInvitation{
		Email:         &fmtEmail,
		RoleCanonical: &role,
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	_, err = m.api.OrganizationMembers.InviteUserToOrgMember(params, common.ClientCredentials(&org))
	if err != nil {
		return NewApiError(err)
	}

	return nil
}

func (m *middleware) DeleteInvite(org string, invite string) error {
	params := organization_invitations.NewDeleteInvitationParams()
	params.SetOrganizationCanonical(org)

	i64, err := strconv.ParseInt(invite, 10, 32)
	if err != nil {
		return err
	}
	params.SetInvitationID(uint32(i64))

	_, err = m.api.OrganizationInvitations.DeleteInvitation(params, common.ClientCredentials(&org))
	if err != nil {
		return NewApiError(err)
	}

	return nil
}
