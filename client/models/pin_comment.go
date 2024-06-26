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

// PinComment PinComment
//
// # Represents a pin comment of a resource
//
// swagger:model PinComment
type PinComment struct {

	// pin comment
	// Required: true
	PinComment *string `json:"pin_comment"`
}

// Validate validates this pin comment
func (m *PinComment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePinComment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PinComment) validatePinComment(formats strfmt.Registry) error {

	if err := validate.Required("pin_comment", "body", m.PinComment); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this pin comment based on context it is used
func (m *PinComment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PinComment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PinComment) UnmarshalBinary(b []byte) error {
	var res PinComment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
