package middleware

import (
	"errors"
	"net/http"

	"gopkg.in/yaml.v3"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) ValidateForm(org string, rawForms []byte) (*models.FormsValidationResult, *http.Response, error) {
	var formsfile any

	err := yaml.Unmarshal(rawForms, &formsfile)
	if err != nil {
		ve := &models.FormsValidationResult{
			Errors: []string{err.Error()},
			Forms:  nil,
		}
		return ve, nil, nil
	}

	body := &models.FormsValidation{
		FormFile: formsfile,
	}

	var result *models.FormsValidationResult
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "forms", "validate"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// interpolateFormsConfigBody mirrors the fields needed for the forms config interpolation
type interpolateFormsConfigBody struct {
	ServiceCatalogRef  *string              `json:"service_catalog_ref"`
	ComponentCanonical *string              `json:"component_canonical"`
	UseCase            *string              `json:"use_case"`
	Vars               models.FormVariables `json:"vars"`
}

func (m *middleware) InterpolateFormsConfig(org, project, env, component, serviceCatalogRef, useCase string, inputs models.FormVariables) (*models.ServiceCatalogConfig, *http.Response, error) {
	if inputs == nil {
		return nil, nil, errors.New("form inputs for interpolateFormsConfig must not be nil")
	}

	body := interpolateFormsConfigBody{
		ServiceCatalogRef:  &serviceCatalogRef,
		ComponentCanonical: &component,
		UseCase:            &useCase,
		Vars:               inputs,
	}

	var result *models.ServiceCatalogConfig
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "forms", "config"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
