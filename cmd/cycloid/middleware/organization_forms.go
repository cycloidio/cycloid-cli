package middleware

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/client/client/organization_forms"
	"github.com/cycloidio/cycloid-cli/client/models"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// // from https://github.com/cycloidio/youdeploy-http-api/blob/develop/services/youdeploy/svccat/form/file.go#L12
// // modify Entity by interface and add Data from FileV1
// type FileForms struct {
// 	Version  *string                `yaml:"version" json:"version"`
// 	UseCases interface{}            `yaml:"use_cases" json:"use_cases"`
// 	Data     map[string]interface{} `yaml:",inline"`
// }

func (m *middleware) ValidateForm(org string, rawForms []byte) (*models.FormsValidationResult, error) {
	var formsfile any

	err := yaml.Unmarshal(rawForms, &formsfile)
	if err != nil {
		// return nil, err
		// Convert swagger validation error as FormsValidationResult
		// to keep the same display on validation error for the end user
		ve := &models.FormsValidationResult{
			Errors: []string{err.Error()},
			Forms:  nil,
		}
		return ve, nil
	}

	params := organization_forms.NewValidateFormsFileParams()
	params.SetOrganizationCanonical(org)

	body := &models.FormsValidation{
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
	resp, err := m.api.OrganizationForms.ValidateFormsFile(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	// Don't validate this payload.
	// Validation will silence an expected error related to SF validation

	return payload.Data, nil
}

func (m *middleware) InterpolateFormsConfig(org, project, env, component, serviceCatalogRef, useCase string, inputs *models.FormVariables) (*models.ServiceCatalogConfig, error) {
	if inputs == nil {
		return nil, errors.New("form inputs for interpolateFormsConfig must not be nil")
	}
	body := organization_forms.InterpolateFormsConfigBody{
		ServiceCatalogRef:  &serviceCatalogRef,
		ComponentCanonical: &component,
		UseCase:            &useCase,
		Vars:               *inputs,
	}

	params := organization_forms.NewInterpolateFormsConfigParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)
	params.WithEnvironmentCanonical(env)
	params.WithBody(body)

	if err := params.Body.Validate(strfmt.Default); err != nil {
		return nil, err
	}

	resp, err := m.api.OrganizationForms.InterpolateFormsConfig(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}
