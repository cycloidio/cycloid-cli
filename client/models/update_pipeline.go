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

// UpdatePipeline Update Pipeline
//
// The entity which represents a new pipeline config to update in the application.
// swagger:model UpdatePipeline
type UpdatePipeline struct {

	// check credentials
	CheckCredentials bool `json:"check_credentials,omitempty"`

	// passed config
	// Required: true
	PassedConfig *string `json:"passed_config"`

	// yaml vars
	YamlVars string `json:"yaml_vars,omitempty"`
}

// Validate validates this update pipeline
func (m *UpdatePipeline) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassedConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdatePipeline) validatePassedConfig(formats strfmt.Registry) error {

	if err := validate.Required("passed_config", "body", m.PassedConfig); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdatePipeline) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdatePipeline) UnmarshalBinary(b []byte) error {
	var res UpdatePipeline
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
