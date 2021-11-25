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

// TerraformHCLConfig TerraformHCLConfig
//
// The HCL config for Terraform
// swagger:model TerraformHCLConfig
type TerraformHCLConfig struct {

	// config
	// Required: true
	Config *string `json:"config"`
}

// Validate validates this terraform h c l config
func (m *TerraformHCLConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerraformHCLConfig) validateConfig(formats strfmt.Registry) error {

	if err := validate.Required("config", "body", m.Config); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TerraformHCLConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TerraformHCLConfig) UnmarshalBinary(b []byte) error {
	var res TerraformHCLConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
