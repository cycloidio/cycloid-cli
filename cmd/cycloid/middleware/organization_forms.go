package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client/organization_forms"
	"github.com/cycloidio/cycloid-cli/client/models"
	strfmt "github.com/go-openapi/strfmt"
	"gopkg.in/yaml.v3"
)

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

func (m *middleware) InterpolateFormsConfig(org, project, component, serviceCatalogRef, useCase string, inputs *models.FormVariables) (*models.ServiceCatalogConfig, error) {
	body := organization_forms.InterpolateFormsConfigBody{
		ServiceCatalogRef:  &serviceCatalogRef,
		ComponentCanonical: &component,
		UseCase:            &useCase,
		Vars:               *inputs,
	}

	params := organization_forms.NewInterpolateFormsConfigParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)
	params.WithBody(body)

	if err := params.Body.Validate(strfmt.Default); err != nil {
		return nil, err
	}

	resp, err := m.api.OrganizationForms.InterpolateFormsConfig(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	return resp.GetPayload().Data, nil
}
