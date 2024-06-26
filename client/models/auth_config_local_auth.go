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

// AuthConfigLocalAuth AppConfigLocalAuth
//
// swagger:model AuthConfigLocalAuth
type AuthConfigLocalAuth struct {

	// Whether local authentication is enabled
	// Required: true
	Enabled *bool `json:"enabled"`
}

// Validate validates this auth config local auth
func (m *AuthConfigLocalAuth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthConfigLocalAuth) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this auth config local auth based on context it is used
func (m *AuthConfigLocalAuth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthConfigLocalAuth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthConfigLocalAuth) UnmarshalBinary(b []byte) error {
	var res AuthConfigLocalAuth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
