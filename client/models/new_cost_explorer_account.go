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

// NewCostExplorerAccount Create CostExplorerAccount
//
// Create a new Cost Explorer account to connect CP.
// If a canonical is missing, the one from credential will be applied
//
// swagger:model NewCostExplorerAccount
type NewCostExplorerAccount struct {

	// canonical
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	Canonical string `json:"canonical,omitempty"`

	// external backend
	// Required: true
	ExternalBackend *NewExternalBackend `json:"external_backend"`
}

// Validate validates this new cost explorer account
func (m *NewCostExplorerAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalBackend(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewCostExplorerAccount) validateCanonical(formats strfmt.Registry) error {

	if swag.IsZero(m.Canonical) { // not required
		return nil
	}

	if err := validate.MinLength("canonical", "body", string(m.Canonical), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("canonical", "body", string(m.Canonical), 100); err != nil {
		return err
	}

	if err := validate.Pattern("canonical", "body", string(m.Canonical), `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *NewCostExplorerAccount) validateExternalBackend(formats strfmt.Registry) error {

	if err := validate.Required("external_backend", "body", m.ExternalBackend); err != nil {
		return err
	}

	if m.ExternalBackend != nil {
		if err := m.ExternalBackend.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external_backend")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewCostExplorerAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewCostExplorerAccount) UnmarshalBinary(b []byte) error {
	var res NewCostExplorerAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
