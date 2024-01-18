package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client/organization_roles"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	strfmt "github.com/go-openapi/strfmt"
)

func (m *middleware) ListRoles(org string) ([]*models.Role, error) {
	params := organization_roles.NewGetRolesParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationRoles.GetRoles(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()

	err = p.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	d := p.Data

	return d, err
}

func (m *middleware) GetRole(org, role string) (*models.Role, error) {
	params := organization_roles.NewGetRoleParams()
	params.SetOrganizationCanonical(org)
	params.SetRoleCanonical(role)

	resp, err := m.api.OrganizationRoles.GetRole(params, common.ClientCredentials(&org))
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

func (m *middleware) DeleteRole(org, role string) error {
	params := organization_roles.NewDeleteRoleParams()
	params.SetOrganizationCanonical(org)
	params.SetRoleCanonical(role)

	_, err := m.api.OrganizationRoles.DeleteRole(params, common.ClientCredentials(&org))
	if err != nil {
		return NewApiError(err)
	}

	return nil
}
