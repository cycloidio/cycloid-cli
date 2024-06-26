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

// TerraformImage TerraformImage
//
// # The Image from the TF structure
//
// swagger:model TerraformImage
type TerraformImage struct {

	// image
	// Required: true
	// Format: byte
	Image *strfmt.Base64 `json:"image"`
}

// Validate validates this terraform image
func (m *TerraformImage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerraformImage) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this terraform image based on context it is used
func (m *TerraformImage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TerraformImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TerraformImage) UnmarshalBinary(b []byte) error {
	var res TerraformImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
