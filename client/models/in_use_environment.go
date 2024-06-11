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

// InUseEnvironment InUseEnvironment
//
// # Represents an environment that's using credential
//
// swagger:model InUseEnvironment
type InUseEnvironment struct {

	// canonical
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: (^[a-z0-9]+(([a-z0-9\-_]+)?[a-z0-9]+)?$)
	Canonical *string `json:"canonical"`
}

// Validate validates this in use environment
func (m *InUseEnvironment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanonical(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InUseEnvironment) validateCanonical(formats strfmt.Registry) error {

	if err := validate.Required("canonical", "body", m.Canonical); err != nil {
		return err
	}

	if err := validate.MinLength("canonical", "body", *m.Canonical, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("canonical", "body", *m.Canonical, 100); err != nil {
		return err
	}

	if err := validate.Pattern("canonical", "body", *m.Canonical, `(^[a-z0-9]+(([a-z0-9\-_]+)?[a-z0-9]+)?$)`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this in use environment based on context it is used
func (m *InUseEnvironment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InUseEnvironment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InUseEnvironment) UnmarshalBinary(b []byte) error {
	var res InUseEnvironment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
