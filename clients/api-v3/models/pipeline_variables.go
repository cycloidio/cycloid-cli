// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PipelineVariables Pipeline variables
//
// The entity which contains pipeline's variables.
// swagger:model PipelineVariables
type PipelineVariables struct {

	// The has_saved_yaml_vars specifies whether the returned vars are from
	// the saved ones or the sample ones. If the has_saved_yaml_vars is true,
	// it means that the saved have been returned.
	//
	// Required: true
	HasSavedYamlVars *bool `json:"has_saved_yaml_vars"`

	// yaml vars
	// Required: true
	YamlVars *string `json:"yaml_vars"`
}

// Validate validates this pipeline variables
func (m *PipelineVariables) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHasSavedYamlVars(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYamlVars(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PipelineVariables) validateHasSavedYamlVars(formats strfmt.Registry) error {

	if err := validate.Required("has_saved_yaml_vars", "body", m.HasSavedYamlVars); err != nil {
		return err
	}

	return nil
}

func (m *PipelineVariables) validateYamlVars(formats strfmt.Registry) error {

	if err := validate.Required("yaml_vars", "body", m.YamlVars); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PipelineVariables) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelineVariables) UnmarshalBinary(b []byte) error {
	var res PipelineVariables
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
