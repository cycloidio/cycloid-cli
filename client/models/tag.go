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

// Tag Key and value pair
//
// Key and value pair defined with the widely adopted name, tag.
//
// swagger:model Tag
type Tag struct {

	// key
	// Required: true
	// Max Length: 254
	// Min Length: 1
	Key *string `json:"key"`

	// value
	// Required: true
	// Max Length: 254
	// Pattern: ^(?:[\w\-+=.:/@ ]*)$
	Value *string `json:"value"`
}

// Validate validates this tag
func (m *Tag) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tag) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	if err := validate.MinLength("key", "body", *m.Key, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("key", "body", *m.Key, 254); err != nil {
		return err
	}

	return nil
}

func (m *Tag) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	if err := validate.MaxLength("value", "body", *m.Value, 254); err != nil {
		return err
	}

	if err := validate.Pattern("value", "body", *m.Value, `^(?:[\w\-+=.:/@ ]*)$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this tag based on context it is used
func (m *Tag) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Tag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tag) UnmarshalBinary(b []byte) error {
	var res Tag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
