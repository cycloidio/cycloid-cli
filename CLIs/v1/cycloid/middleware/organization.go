package middleware

import (
	"github.com/cycloidio/youdeploy-cli/client/client/organization_workers"
	"github.com/cycloidio/youdeploy-cli/client/client/organizations"
	"github.com/cycloidio/youdeploy-cli/client/models"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
)

func (m *middleware) GetOrganization(org string) (*models.Organization, error) {

	params := organizations.NewGetOrgParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.Organizations.GetOrg(params, common.ClientCredentials())

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

	resp, err := m.api.OrganizationWorkers.GetWorkers(params, common.ClientCredentials())

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

func (m *middleware) ListOrganizations() ([]*models.OrganizationBasicInfo, error) {

	params := organizations.NewGetOrgsParams()

	resp, err := m.api.Organizations.GetOrgs(params, common.ClientCredentials())
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
