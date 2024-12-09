package middleware

import (
	"fmt"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/cycloidio/cycloid-cli/client/client/organization_forms"
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
	org, ref, name, canonical, author, description, visibility, catalogRepoCanonical, teamCanonical string,
	image strfmt.URI,
	keywords []string,
	technologies []*models.ServiceCatalogTechnology,
	dependencies []*models.ServiceCatalogDependency,
) (*models.ServiceCatalog, error) {
	params := service_catalogs.NewUpdateServiceCatalogParams()
	params.WithOrganizationCanonical(org)
	params.WithServiceCatalogRef(ref)

	body := &models.UpdateServiceCatalog{
		Author:                        &author,
		Canonical:                     canonical,
		Dependencies:                  dependencies,
		Description:                   &description,
		Image:                         image,
		Keywords:                      keywords,
		Name:                          &name,
		ServiceCatalogSourceCanonical: &catalogRepoCanonical,
		TeamCanonical:                 teamCanonical,
		Technologies:                  technologies,
		Visibility:                    visibility,
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

// from https://github.com/cycloidio/youdeploy-http-api/blob/develop/services/youdeploy/svccat/form/file.go#L12
// modify Entity by interface and add Data from FileV1
type FileForms struct {
	Version  *string                `yaml:"version" json:"version"`
	UseCases interface{}            `yaml:"use_cases" json:"use_cases"`
	Data     map[string]interface{} `yaml:",inline"`
}

func (m *middleware) ValidateForm(org string, rawForms []byte) (*models.FormsValidationResult, error) {
	var body *models.FormsValidation
	var formsfile FileForms

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

	// We unmarchal stackforms yaml file in a generic structure FileForms.
	// Because the yaml file format could be v1 or v2 the FileForms is based on interfaces.
	// Unfortunately golang produce an error when you unmarchal yaml from interface, and marchal it later on into json
	// unable validate form: json: unsupported type: map[interface {}]interface {}

	var bodyFormFile interface{}
	if len(formsfile.Data) > 0 {
		// v1
		datas := map[string]interface{}{}
		for key, element := range formsfile.Data {
			datas[key] = ConvertMapInterfaceToMapString(element)
			// if element under Data, that means we use v1
		}
		bodyFormFile = datas
	} else {
		// v2
		formsfile.UseCases = ConvertMapInterfaceToMapString(formsfile.UseCases)
		bodyFormFile = formsfile
	}

	params := organization_forms.NewValidateFormsFileParams()
	params.SetOrganizationCanonical(org)

	body = &models.FormsValidation{
		FormFile: bodyFormFile,
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
	resp, err := m.api.OrganizationForms.ValidateFormsFile(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()

	d := p.Data
	return d, nil
}

func (m *middleware) CreateFormsConfig(org string, project string, serviceCatalogRef string, inputs []*models.FormInput) (map[string]any, error) {
	body := &models.FormInputs{
		ServiceCatalogRef: &serviceCatalogRef,
		Inputs:            inputs,
	}

	params := organization_forms.NewCreateFormsConfigParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)
	params.WithBody(body)

	resp, err := m.api.OrganizationForms.CreateFormsConfig(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	fmt.Println(resp.GetPayload().Data)
	return nil, nil
}
