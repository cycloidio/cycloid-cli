package middleware

import (
	"fmt"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/pkg/errors"

	"github.com/cycloidio/cycloid-cli/client/client/organization_service_catalog_sources"
	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) ListCatalogRepositories(org string) ([]*models.ServiceCatalogSource, error) {
	params := organization_service_catalog_sources.NewGetServiceCatalogSourcesParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationServiceCatalogSources.GetServiceCatalogSources(params, m.api.Credentials(&org))
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

func (m *middleware) GetCatalogRepository(org, catalogRepo string) (*models.ServiceCatalogSource, error) {
	params := organization_service_catalog_sources.NewGetServiceCatalogSourceParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogSourceCanonical(catalogRepo)

	resp, err := m.api.OrganizationServiceCatalogSources.GetServiceCatalogSource(params, m.api.Credentials(&org))
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

func (m *middleware) DeleteCatalogRepository(org, catalogRepo string) error {
	params := organization_service_catalog_sources.NewDeleteServiceCatalogSourceParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogSourceCanonical(catalogRepo)

	_, err := m.api.OrganizationServiceCatalogSources.DeleteServiceCatalogSource(params, m.api.Credentials(&org))
	if err != nil {
		return NewAPIError(err)
	}
	return nil
}

func (m *middleware) CreateCatalogRepository(org, name, url, branch, cred, visibility, teamCanonical string) (*models.ServiceCatalogSource, error) {
	params := organization_service_catalog_sources.NewCreateServiceCatalogSourceParams()
	params.SetOrganizationCanonical(org)

	var body *models.NewServiceCatalogSource

	if len(cred) != 0 {
		body = &models.NewServiceCatalogSource{
			Branch:              &branch,
			CredentialCanonical: cred,
			Name:                &name,
			URL:                 &url,
		}
	} else {
		body = &models.NewServiceCatalogSource{
			Branch: &branch,
			Name:   &name,
			URL:    &url,
		}
	}

	switch visibility {
	case "shared", "local", "hidden":
		body.Visibility = visibility
	case "":
		break
	default:
		return nil, errors.New("invalid visibility parameter for CreateCatalogRepository, accepted values are 'local', 'shared' or 'hidden'")
	}

	if teamCanonical != "" {
		body.TeamCanonical = teamCanonical
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	resp, err := m.api.OrganizationServiceCatalogSources.CreateServiceCatalogSource(params, m.api.Credentials(&org))
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

func (m *middleware) UpdateCatalogRepository(org, catalogRepo string, name, url, branch, cred string,
) (*models.ServiceCatalogSource, error) {
	params := organization_service_catalog_sources.NewUpdateServiceCatalogSourceParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogSourceCanonical(catalogRepo)

	body := &models.UpdateServiceCatalogSource{
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

	resp, err := m.api.OrganizationServiceCatalogSources.UpdateServiceCatalogSource(params, m.api.Credentials(&org))
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

func (m *middleware) RefreshCatalogRepository(org, catalogRepo string) (*models.ServiceCatalogChanges, error) {
	params := organization_service_catalog_sources.NewRefreshServiceCatalogSourceParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogSourceCanonical(catalogRepo)

	resp, err := m.api.OrganizationServiceCatalogSources.RefreshServiceCatalogSource(params, m.api.Credentials(&org))
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
