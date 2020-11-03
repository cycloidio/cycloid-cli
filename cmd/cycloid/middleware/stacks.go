package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client/organization_forms"
	"github.com/cycloidio/cycloid-cli/client/client/service_catalogs"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	strfmt "github.com/go-openapi/strfmt"
	"gopkg.in/yaml.v2"
)

func (m *middleware) ListStacks(org string) ([]*models.ServiceCatalog, error) {

	params := service_catalogs.NewGetServiceCatalogsParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.ServiceCatalogs.GetServiceCatalogs(params, common.ClientCredentials(&org))
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

func (m *middleware) GetStack(org, ref string) (*models.ServiceCatalog, error) {

	params := service_catalogs.NewGetServiceCatalogParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogRef(ref)

	resp, err := m.api.ServiceCatalogs.GetServiceCatalog(params, common.ClientCredentials(&org))
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

func (m *middleware) ValidateForm(org string, rawForms []byte) (*models.FormsValidationResult, error) {
	var body *models.FormsValidation
	var formsfile models.FormsFile

	err := yaml.Unmarshal(rawForms, &formsfile)
	if err != nil {
		// return nil, err
		// Convert swagger validation error as FormsValidationResult
		// to keep the same display on validation error for the end user
		ve := &models.FormsValidationResult{
			Errors: []string{err.Error()},
		}
		return ve, nil
	}

	params := organization_forms.NewValidateFormsFileParams()
	params.SetOrganizationCanonical(org)

	body = &models.FormsValidation{
		FormFile: formsfile,
	}
	err = body.Validate(strfmt.Default)
	if err != nil {
		// return nil, err
		// Convert swagger validation error as FormsValidationResult
		// to keep the same display on validation error for the end user
		ve := &models.FormsValidationResult{
			Errors: []string{err.Error()},
		}
		return ve, nil
	}

	params.SetBody(body)

	resp, err := m.api.OrganizationForms.ValidateFormsFile(params, common.ClientCredentials(&org))
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
