// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TerraformProvider Provider
//
// # Provider of infrastructure
//
// swagger:model TerraformProvider
type TerraformProvider struct {

	// abbreviation
	// Required: true
	Abbreviation *string `json:"abbreviation"`

	// canonical
	// Required: true
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	Canonical *string `json:"canonical"`

	// cloud
	// Required: true
	Cloud *bool `json:"cloud"`

	// name
	// Required: true
	Name *string `json:"name"`

	// schema
	// Required: true
	Schema interface{} `json:"schema"`
}

// Validate validates this terraform provider
func (m *TerraformProvider) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbbreviation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloud(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerraformProvider) validateAbbreviation(formats strfmt.Registry) error {

	if err := validate.Required("abbreviation", "body", m.Abbreviation); err != nil {
		return err
	}

	return nil
}

func (m *TerraformProvider) validateCanonical(formats strfmt.Registry) error {

	if err := validate.Required("canonical", "body", m.Canonical); err != nil {
		return err
	}

	if err := validate.MinLength("canonical", "body", *m.Canonical, 3); err != nil {
		return err
	}

	if err := validate.Pattern("canonical", "body", *m.Canonical, `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *TerraformProvider) validateCloud(formats strfmt.Registry) error {

	if err := validate.Required("cloud", "body", m.Cloud); err != nil {
		return err
	}

	return nil
}

func (m *TerraformProvider) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *TerraformProvider) validateSchema(formats strfmt.Registry) error {

	if m.Schema == nil {
		return errors.Required("schema", "body", nil)
	}

	return nil
}

// ContextValidate validates this terraform provider based on context it is used
func (m *TerraformProvider) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TerraformProvider) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TerraformProvider) UnmarshalBinary(b []byte) error {
	var res TerraformProvider
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
