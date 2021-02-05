package middleware

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/pkg/errors"

	"github.com/cycloidio/cycloid-cli/client/client/organization_workers"
	"github.com/cycloidio/cycloid-cli/client/client/organizations"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
)

func (m *middleware) CreateOrganization(name string) (*models.Organization, error) {

	params := organizations.NewCreateOrgParams()

	body := &models.NewOrganization{
		Name:      &name,
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, errors.Wrap(err, "unable to validate request body")
	}

	resp, err := m.api.Organizations.CreateOrg(params, common.ClientCredentials(nil))

	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()

	d := p.Data
	return d, err
}

func (m *middleware) GetOrganization(org string) (*models.Organization, error) {

	params := organizations.NewGetOrgParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.Organizations.GetOrg(params, common.ClientCredentials(&org))

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

func (m *middleware) ListOrganizationWorkers(org string) ([]*models.Worker, error) {

	params := organization_workers.NewGetWorkersParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationWorkers.GetWorkers(params, common.ClientCredentials(&org))

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

func (m *middleware) ListOrganizations() ([]*models.Organization, error) {

	params := organizations.NewGetOrgsParams()

	resp, err := m.api.Organizations.GetOrgs(params, common.ClientCredentials(nil))
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

func (m *middleware) DeleteOrganization(org string) error {
	params := organizations.NewDeleteOrgParams()
	params.SetOrganizationCanonical(org)

	_, err := m.api.Organizations.DeleteOrg(params, common.ClientCredentials(&org))
	return err
}
