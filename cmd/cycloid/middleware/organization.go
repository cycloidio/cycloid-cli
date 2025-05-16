package middleware

import (
	"fmt"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/pkg/errors"

	"github.com/cycloidio/cycloid-cli/client/client/organization_children"
	"github.com/cycloidio/cycloid-cli/client/client/organization_workers"
	"github.com/cycloidio/cycloid-cli/client/client/organizations"
	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) CreateOrganization(name string) (*models.Organization, error) {
	params := organizations.NewCreateOrgParams()

	body := &models.NewOrganization{
		Name: &name,
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, errors.Wrap(err, "unable to validate request body")
	}
	params.SetBody(body)

	resp, err := m.api.Organizations.CreateOrg(params, m.api.Credentials(nil))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}

func (m *middleware) UpdateOrganization(can, name string) (*models.Organization, error) {
	params := organizations.NewUpdateOrgParams()
	params.SetOrganizationCanonical(can)

	body := &models.UpdateOrganization{
		Name: &name,
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, errors.Wrap(err, "unable to validate request body")
	}
	params.SetBody(body)

	resp, err := m.api.Organizations.UpdateOrg(params, m.api.Credentials(&can))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}

func (m *middleware) GetOrganization(org string) (*models.Organization, error) {
	params := organizations.NewGetOrgParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.Organizations.GetOrg(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}

func (m *middleware) ListOrganizationWorkers(org string) ([]*models.Worker, error) {
	params := organization_workers.NewGetWorkersParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationWorkers.GetWorkers(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}

func (m *middleware) ListOrganizations() ([]*models.Organization, error) {
	params := organizations.NewGetOrgsParams()

	resp, err := m.api.Organizations.GetOrgs(params, m.api.Credentials(nil))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}

func (m *middleware) ListOrganizationChildrens(org string) ([]*models.Organization, error) {
	params := organization_children.NewGetChildrenParams()
	orderBy := "organization_canonical:asc"
	params.SetOrderBy(&orderBy)
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationChildren.GetChildren(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}

func (m *middleware) CreateOrganizationChild(org, childOrg string, childOrgName *string) (*models.Organization, error) {
	if childOrgName == nil {
		childOrgName = &childOrg
	}

	params := organization_children.NewCreateChildParams()
	params.SetOrganizationCanonical(org)
	body := &models.NewOrganization{
		Name:      childOrgName,
		Canonical: childOrg,
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, errors.Wrap(err, "unable to validate request body")
	}

	params.SetBody(body)

	resp, err := m.api.OrganizationChildren.CreateChild(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}

func (m *middleware) DeleteOrganization(org string) error {
	params := organizations.NewDeleteOrgParams()
	params.SetOrganizationCanonical(org)

	_, err := m.api.Organizations.DeleteOrg(params, m.api.Credentials(&org))
	if err != nil {
		return NewAPIError(err)
	}
	return nil
}
