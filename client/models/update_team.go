// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateTeam Update Team
//
// The entity which represents the information of the team to be updated.
//
// swagger:model UpdateTeam
type UpdateTeam struct {

	// canonical
	// Required: true
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	Canonical *string `json:"canonical"`

	// name
	// Required: true
	// Min Length: 3
	Name *string `json:"name"`

	// User canonical that owns this team. Only the owner or an
	// organization admin can update this field. When a user is the owner
	// of a team it has all the permission on it.
	//
	Owner string `json:"owner,omitempty"`

	// The roles to be re-assigned to a team.
	// Required: true
	RolesCanonical []string `json:"roles_canonical"`
}

// Validate validates this update team
func (m *UpdateTeam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRolesCanonical(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateTeam) validateCanonical(formats strfmt.Registry) error {

	if err := validate.Required("canonical", "body", m.Canonical); err != nil {
		return err
	}

	if err := validate.MinLength("canonical", "body", *m.Canonical, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("canonical", "body", *m.Canonical, 100); err != nil {
		return err
	}

	if err := validate.Pattern("canonical", "body", *m.Canonical, `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *UpdateTeam) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 3); err != nil {
		return err
	}

	return nil
}

func (m *UpdateTeam) validateRolesCanonical(formats strfmt.Registry) error {

	if err := validate.Required("roles_canonical", "body", m.RolesCanonical); err != nil {
		return err
	}

	for i := 0; i < len(m.RolesCanonical); i++ {

		if err := validate.MinLength("roles_canonical"+"."+strconv.Itoa(i), "body", m.RolesCanonical[i], 3); err != nil {
			return err
		}

		if err := validate.MaxLength("roles_canonical"+"."+strconv.Itoa(i), "body", m.RolesCanonical[i], 100); err != nil {
			return err
		}

		if err := validate.Pattern("roles_canonical"+"."+strconv.Itoa(i), "body", m.RolesCanonical[i], `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this update team based on context it is used
func (m *UpdateTeam) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateTeam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateTeam) UnmarshalBinary(b []byte) error {
	var res UpdateTeam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
