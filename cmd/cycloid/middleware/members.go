package middleware

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/client/client/organization_members"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
)

func (m *middleware) ListMembers(org string) ([]*models.MemberOrg, error) {
	params := organization_members.NewGetOrgMembersParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationMembers.GetOrgMembers(params, m.api.Credentials(&org))
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

func (m *middleware) ListInvites(org string) ([]*models.MemberOrg, error) {
	params := organization_members.NewGetOrgMembersParams()
	params.SetOrganizationCanonical(org)
	params.SetInvitationState(ptr.Ptr("pending"))

	resp, err := m.api.OrganizationMembers.GetOrgMembers(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()

	d := p.Data

	return d, nil
}

func (m *middleware) GetMember(org string, id uint32) (*models.MemberOrg, error) {
	params := organization_members.NewGetOrgMemberParams()
	params.SetOrganizationCanonical(org)
	params.SetMemberID(id)

	resp, err := m.api.OrganizationMembers.GetOrgMember(params, m.api.Credentials(&org))
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

func (m *middleware) DeleteMember(org string, id uint32) error {
	params := organization_members.NewRemoveOrgMemberParams()
	params.SetOrganizationCanonical(org)
	params.SetMemberID(id)

	_, err := m.api.OrganizationMembers.RemoveOrgMember(params, m.api.Credentials(&org))
	if err != nil {
		return err
	}

	return nil
}

func (m *middleware) UpdateMember(org string, id uint32, role string) (*models.MemberOrg, error) {
	params := organization_members.NewUpdateOrgMemberParams()

	params.SetOrganizationCanonical(org)
	params.SetMemberID(id)

	body := &models.MemberAssignation{
		RoleCanonical: &role,
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	resp, err := m.api.OrganizationMembers.UpdateOrgMember(params, m.api.Credentials(&org))
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

func (m *middleware) InviteMember(org, email, role string) (*models.MemberOrg, error) {
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
		return nil, err
	}

	resp, err := m.api.OrganizationMembers.InviteUserToOrgMember(params, m.api.Credentials(&org))
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
