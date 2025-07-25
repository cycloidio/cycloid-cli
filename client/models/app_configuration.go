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

// AppConfiguration AppConfiguration
//
// # Contains application-wide configuration
//
// swagger:model AppConfiguration
type AppConfiguration struct {

	// Indicates if Sentry is enabled.
	// Required: true
	SentryEnabled *bool `json:"sentry_enabled"`
}

// Validate validates this app configuration
func (m *AppConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSentryEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppConfiguration) validateSentryEnabled(formats strfmt.Registry) error {

	if err := validate.Required("sentry_enabled", "body", m.SentryEnabled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this app configuration based on context it is used
func (m *AppConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppConfiguration) UnmarshalBinary(b []byte) error {
	var res AppConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
