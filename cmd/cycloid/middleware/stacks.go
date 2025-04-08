package middleware

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/pkg/errors"

	"github.com/cycloidio/cycloid-cli/client/client/service_catalogs"
	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) ListStacks(org string) ([]*models.ServiceCatalog, error) {
	params := service_catalogs.NewListServiceCatalogsParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.ServiceCatalogs.ListServiceCatalogs(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()

	d := p.Data

	return d, nil
}

func (m *middleware) GetStack(org, ref string) (*models.ServiceCatalog, error) {
	params := service_catalogs.NewGetServiceCatalogParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogRef(ref)

	resp, err := m.api.ServiceCatalogs.GetServiceCatalog(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()

	d := p.Data

	return d, nil
}

func (m *middleware) UpdateStack(
	org, ref, teamCanonical string,
	visibility *string,
) (*models.ServiceCatalog, error) {
	params := service_catalogs.NewUpdateServiceCatalogParams()
	params.WithOrganizationCanonical(org)
	params.WithServiceCatalogRef(ref)

	body := &models.UpdateServiceCatalog{
		TeamCanonical: teamCanonical,
		Visibility:    visibility,
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, errors.Wrap(err, "validation failed for updateServiceCatalog input")
	}

	params.WithBody(body)

	response, err := m.api.ServiceCatalogs.UpdateServiceCatalog(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := response.GetPayload()

	// TODO: This is a local fix for https://github.com/cycloidio/youdeploy-http-api/issues/5020
	// Remove this condition when backend will be fixed
	// If the team attribute is nil, this means that the backend did not found the maitainer canonical
	if teamCanonical != "" && payload.Data.Team == nil {
		return payload.Data, errors.Errorf(
			"maintainer with canonical '%s' may not exists, maintainer on stack ref '%s' has been removed, please check you team canonical argument and ensure that the team exists.",
			teamCanonical, ref,
		)
	}

	return payload.Data, nil
}

func (m *middleware) GetStackConfig(org, ref string) (interface{}, error) {
	params := service_catalogs.NewGetServiceCatalogConfigParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogRef(ref)

	resp, err := m.api.ServiceCatalogs.GetServiceCatalogConfig(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()

	d := p.Data

	return d, nil
}
