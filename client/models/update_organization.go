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

// UpdateOrganization Update Organization
//
// The entity which represents the information of an organization to be updated.
// swagger:model UpdateOrganization
type UpdateOrganization struct {

	// name
	// Required: true
	// Min Length: 3
	Name *string `json:"name"`

	// quotas
	Quotas bool `json:"quotas,omitempty"`
}

// Validate validates this update organization
func (m *UpdateOrganization) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateOrganization) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 3); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateOrganization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateOrganization) UnmarshalBinary(b []byte) error {
	var res UpdateOrganization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
