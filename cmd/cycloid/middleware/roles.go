package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client/organization_roles"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
)

func (m *middleware) ListRoles(org string) ([]*models.Role, error) {
	params := organization_roles.NewGetOrgRolesParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationRoles.GetOrgRoles(params, common.ClientCredentials(&org))
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

func (m *middleware) GetRole(org, role string) (*models.Role, error) {
	params := organization_roles.NewGetOrgRoleParams()
	params.SetOrganizationCanonical(org)
	params.SetRoleCanonical(role)

	resp, err := m.api.OrganizationRoles.GetOrgRole(params, common.ClientCredentials(&org))
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

func (m *middleware) DeleteRole(org, role string) error {
	params := organization_roles.NewDeleteOrgRoleParams()
	params.SetOrganizationCanonical(org)
	params.SetRoleCanonical(role)

	_, err := m.api.OrganizationRoles.DeleteOrgRole(params, common.ClientCredentials(&org))

	return err
}
