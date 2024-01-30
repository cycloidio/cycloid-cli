package middleware

import (
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

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, errors.Wrap(err, "unable to validate request body")
	}

	resp, err := m.api.Organizations.CreateOrg(params, m.api.Credentials(nil))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()

	d := p.Data
	return d, nil
}

func (m *middleware) GetOrganization(org string) (*models.Organization, error) {

	params := organizations.NewGetOrgParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.Organizations.GetOrg(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()
	//err = p.Validate(strfmt.Default)
	//if err != nil {
	//return nil, err
	//}

	d := p.Data
	return d, nil
}

func (m *middleware) ListOrganizationWorkers(org string) ([]*models.Worker, error) {

	params := organization_workers.NewGetWorkersParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationWorkers.GetWorkers(params, m.api.Credentials(&org))
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

func (m *middleware) ListOrganizations() ([]*models.Organization, error) {

	params := organizations.NewGetOrgsParams()

	resp, err := m.api.Organizations.GetOrgs(params, m.api.Credentials(nil))
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

func (m *middleware) ListOrganizationChildrens(org string) ([]*models.Organization, error) {

	params := organization_children.NewGetChildrenParams()
	orderBy := "organization_canonical:asc"
	params.SetOrderBy(&orderBy)
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationChildren.GetChildren(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()

	d := p.Data
	return d, nil
}

func (m *middleware) CreateOrganizationChild(org, porg string) (*models.Organization, error) {

	params := organization_children.NewCreateChildParams()
	//params.SetOrganizationCanonical(porg)

	body := &models.NewOrganization{
		Name: &org,
	}

	params.SetBody(body)
	params.OrganizationCanonical = porg
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, errors.Wrap(err, "unable to validate request body")
	}

	resp, err := m.api.OrganizationChildren.CreateChild(params, m.api.Credentials(&porg))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()

	d := p.Data
	return d, nil
}

func (m *middleware) DeleteOrganization(org string) error {
	params := organizations.NewDeleteOrgParams()
	params.SetOrganizationCanonical(org)

	_, err := m.api.Organizations.DeleteOrg(params, m.api.Credentials(&org))
	if err != nil {
		return NewApiError(err)
	}
	return nil
}
