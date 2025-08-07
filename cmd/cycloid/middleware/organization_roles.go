package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client/organization_roles"
	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) ListRoles(org string) ([]*models.Role, error) {
	params := organization_roles.NewGetRolesParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationRoles.GetRoles(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) GetRole(org, role string) (*models.Role, error) {
	params := organization_roles.NewGetRoleParams()
	params.SetOrganizationCanonical(org)
	params.SetRoleCanonical(role)

	resp, err := m.api.OrganizationRoles.GetRole(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) DeleteRole(org, role string) error {
	params := organization_roles.NewDeleteRoleParams()
	params.SetOrganizationCanonical(org)
	params.SetRoleCanonical(role)

	_, err := m.api.OrganizationRoles.DeleteRole(params, m.api.Credentials(&org))
	if err != nil {
		return NewApiError(err)
	}

	return nil
}
