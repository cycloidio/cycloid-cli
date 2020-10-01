package middleware

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/client/client/organization_members"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
)

func (m *middleware) ListMembers(org string) ([]*models.MemberOrg, error) {
	params := organization_members.NewGetOrgMembersParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationMembers.GetOrgMembers(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()

	// TODO this validate have been removed https://github.com/cycloidio/youdeploy-http-api/issues/2262
	// err = p.Validate(strfmt.Default)
	// if err != nil {
	// 	return err
	// }

	d := p.Data

	return d, err
}

func (m *middleware) GetMember(org string, name string) (*models.MemberOrg, error) {
	params := organization_members.NewGetOrgMemberParams()
	params.SetOrganizationCanonical(org)
	params.SetUsername(name)

	resp, err := m.api.OrganizationMembers.GetOrgMember(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()

	// TODO this validate have been removed https://github.com/cycloidio/youdeploy-http-api/issues/2262
	// err = p.Validate(strfmt.Default)
	// if err != nil {
	// 	return err
	// }

	d := p.Data

	return d, err
}

func (m *middleware) DeleteMember(org string, name string) error {
	params := organization_members.NewRemoveOrgMemberParams()
	params.SetOrganizationCanonical(org)
	params.SetUsername(name)

	_, err := m.api.OrganizationMembers.RemoveOrgMember(params, common.ClientCredentials(&org))

	return err
}

func (m *middleware) UpdateMembers(org string, name string, roleID uint32) (*models.MemberOrg, error) {
	params := organization_members.NewUpdateOrgMemberParams()

	params.SetOrganizationCanonical(org)
	params.SetUsername(name)

	body := &models.MemberAssignation{
		RoleID: &roleID,
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	resp, err := m.api.OrganizationMembers.UpdateOrgMember(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, err
	}

	//TODO verify why getpayload no defined ?
	p := resp.GetPayload()

	err = p.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	d := p.Data

	return d, err
}

func (m *middleware) InviteMember(org string, email string, roleID uint32) error {
	params := organization_members.NewInviteUserToOrgMemberParams()
	params.SetOrganizationCanonical(org)

	fmtEmail := strfmt.Email(email)
	body := &models.NewMemberInvitation{
		Email:  &fmtEmail,
		RoleID: &roleID,
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	_, err = m.api.OrganizationMembers.InviteUserToOrgMember(params, common.ClientCredentials(&org))

	return err
}
