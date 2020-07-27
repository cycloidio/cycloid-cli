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

// PublicBuildOutput PublicBuildOutput
//
// Represents the information of a build output
// swagger:model PublicBuildOutput
type PublicBuildOutput struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// version
	// Required: true
	Version map[string]string `json:"version"`
}

// Validate validates this public build output
func (m *PublicBuildOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicBuildOutput) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PublicBuildOutput) validateVersion(formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PublicBuildOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicBuildOutput) UnmarshalBinary(b []byte) error {
	var res PublicBuildOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
