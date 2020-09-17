package middleware

import (
	"github.com/cycloidio/youdeploy-cli/client/client/organization_service_catalog_sources"
	"github.com/cycloidio/youdeploy-cli/client/models"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	strfmt "github.com/go-openapi/strfmt"
)

func (m *middleware) ListCatalogRepositories(org string) ([]*models.ServiceCatalogSource, error) {

	params := organization_service_catalog_sources.NewGetServiceCatalogSourcesParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationServiceCatalogSources.GetServiceCatalogSources(params, common.ClientCredentials(&org))
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

func (m *middleware) GetCatalogRepository(org string, catalogRepo uint32) (*models.ServiceCatalogSource, error) {

	params := organization_service_catalog_sources.NewGetServiceCatalogSourceParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogSourceID(catalogRepo)

	resp, err := m.api.OrganizationServiceCatalogSources.GetServiceCatalogSource(params, common.ClientCredentials(&org))
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

func (m *middleware) DeleteCatalogRepository(org string, catalogRepo uint32) error {
	params := organization_service_catalog_sources.NewDeleteServiceCatalogSourceParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogSourceID(catalogRepo)

	_, err := m.api.OrganizationServiceCatalogSources.DeleteServiceCatalogSource(params, common.ClientCredentials(&org))
	return err
}

func (m *middleware) CreateCatalogRepository(org, name, url, branch string, cred uint32) (*models.ServiceCatalogSource, error) {

	params := organization_service_catalog_sources.NewCreateServiceCatalogSourceParams()
	params.SetOrganizationCanonical(org)

	body := &models.CreateServiceCatalogSource{
		Branch:       &branch,
		CredentialID: cred,
		Name:         &name,
		URL:          &url,
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	resp, err := m.api.OrganizationServiceCatalogSources.CreateServiceCatalogSource(params, common.ClientCredentials(&org))
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

func (m *middleware) UpdateCatalogRepository(org string, catalogRepo uint32, name, url, branch string, cred uint32) (*models.ServiceCatalogSource, error) {

	params := organization_service_catalog_sources.NewUpdateServiceCatalogSourceParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogSourceID(catalogRepo)

	body := &models.UpdateServiceCatalogSource{
		Branch:       branch,
		CredentialID: cred,
		Name:         &name,
		URL:          &url,
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	resp, err := m.api.OrganizationServiceCatalogSources.UpdateServiceCatalogSource(params, common.ClientCredentials(&org))
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

func (m *middleware) RefreshCatalogRepository(org string, catalogRepo uint32) (*models.ServiceCatalogSource, error) {

	params := organization_service_catalog_sources.NewRefreshServiceCatalogSourceParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogSourceID(catalogRepo)

	resp, err := m.api.OrganizationServiceCatalogSources.RefreshServiceCatalogSource(params, common.ClientCredentials(&org))
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
