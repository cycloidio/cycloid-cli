package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client/organization_roles"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"github.com/go-openapi/strfmt"
)

func (m *middleware) ListRoles(org string) ([]*models.Role, error) {
	params := organization_roles.NewGetRolesParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationRoles.GetRoles(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
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
		return nil, NewAPIError(err)
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
		return NewAPIError(err)
	}

	return nil
}

// CreateRole requires org, name or canonical and rules
func (m *middleware) CreateRole(org string, name, canonical, description *string, rules []*models.NewRule) (*models.NewRole, error) {
	params := organization_roles.NewCreateRoleParams()
	params.WithOrganizationCanonical(org)
	n, c, err := NameOrCanonical(name, canonical)
	if err != nil {
		return nil, err
	}

	body := &models.NewRole{
		Name:        &n,
		Canonical:   c,
		Description: ptr.Value(description),
		Rules:       rules,
	}

	params.WithBody(body)
	params.Body.Validate(strfmt.Default)

	resp, err := m.api.OrganizationRoles.CreateRole(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	return resp.Payload.Data, nil
}
