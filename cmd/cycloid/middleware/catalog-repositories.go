package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client/organization_catalog_repositories"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	strfmt "github.com/go-openapi/strfmt"
)

func (m *middleware) ListCatalogRepositories(org string) ([]*models.CatalogRepository, error) {

	params := organization_catalog_repositories.NewGetCatalogRepositoriesParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationCatalogRepositories.GetCatalogRepositories(params, common.ClientCredentials(&org))
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

func (m *middleware) GetCatalogRepository(org, catalogRepo string) (*models.CatalogRepository, error) {

	params := organization_catalog_repositories.NewGetCatalogRepositoryParams()
	params.SetOrganizationCanonical(org)
	params.SetCatalogRepositoryCanonical(catalogRepo)

	resp, err := m.api.OrganizationCatalogRepositories.GetCatalogRepository(params, common.ClientCredentials(&org))
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

func (m *middleware) DeleteCatalogRepository(org, catalogRepo string) error {
	params := organization_catalog_repositories.NewDeleteCatalogRepositoryParams()
	params.SetOrganizationCanonical(org)
	params.SetCatalogRepositoryCanonical(catalogRepo)

	_, err := m.api.OrganizationCatalogRepositories.DeleteCatalogRepository(params, common.ClientCredentials(&org))
	return err
}

func (m *middleware) CreateCatalogRepository(org, name, url, branch, cred string) (*models.CatalogRepository, error) {

	params := organization_catalog_repositories.NewCreateCatalogRepositoryParams()
	params.SetOrganizationCanonical(org)

	var body *models.NewCatalogRepository

	if len(cred) != 0 {
		body = &models.NewCatalogRepository{
			Branch:              &branch,
			CredentialCanonical: cred,
			Name:                &name,
			URL:                 &url,
		}
	} else {
		body = &models.NewCatalogRepository{
			Branch: &branch,
			Name:   &name,
			URL:    &url,
		}
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	resp, err := m.api.OrganizationCatalogRepositories.CreateCatalogRepository(params, common.ClientCredentials(&org))
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

func (m *middleware) UpdateCatalogRepository(org, catalogRepo string, name, url, branch, cred string) (*models.CatalogRepository, error) {

	params := organization_catalog_repositories.NewUpdateCatalogRepositoryParams()
	params.SetOrganizationCanonical(org)
	params.SetCatalogRepositoryCanonical(catalogRepo)

	body := &models.UpdateCatalogRepository{
		Branch:              branch,
		CredentialCanonical: cred,
		Name:                &name,
		URL:                 &url,
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	resp, err := m.api.OrganizationCatalogRepositories.UpdateCatalogRepository(params, common.ClientCredentials(&org))
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

func (m *middleware) RefreshCatalogRepository(org, catalogRepo string) (*models.StackChanges, error) {

	params := organization_catalog_repositories.NewRefreshCatalogRepositoryParams()
	params.SetOrganizationCanonical(org)
	params.SetCatalogRepositoryCanonical(catalogRepo)

	resp, err := m.api.OrganizationCatalogRepositories.RefreshCatalogRepository(params, common.ClientCredentials(&org))
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
