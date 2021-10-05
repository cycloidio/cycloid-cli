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

func (m *middleware) GetStackConfig(org, ref string) (interface{}, error) {

	params := service_catalogs.NewGetServiceCatalogConfigParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogRef(ref)

	resp, err := m.api.ServiceCatalogs.GetServiceCatalogConfig(params, common.ClientCredentials(&org))
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

// convertFormFile takes a models.FormsFile and converts its variables
// from map[interface{}]interface{} to map[string]interface{}, allowing
// to use those properly with the API - as JSON cannot marshal/unmarshal
// map[interface{}]interface{}
func convertFormFile(mff models.FormsFile) models.FormsFile {
	for _, useCases := range mff {
		for _, groups := range useCases {
			for _, entities := range groups {
				for i, entity := range entities {
					entities[i].Default = ConvertMapInterfaceToMapString(entity.Default)
					for ni, v := range entity.Values {
						entities[i].Values[ni] = ConvertMapInterfaceToMapString(v)
					}
				}
			}
		}
	}
	return mff
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
	formsfile = convertFormFile(formsfile)

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
