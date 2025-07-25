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

// MigrateComponent Migrate Component
//
// The entity which represents the information of a component to migrate.
//
// swagger:model MigrateComponent
type MigrateComponent struct {

	// destination component canonical
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[\da-zA-Z]+(?:[\da-zA-Z\-._]+[\da-zA-Z]|[\da-zA-Z])$
	DestinationComponentCanonical string `json:"destination_component_canonical,omitempty"`

	// destination component name
	DestinationComponentName string `json:"destination_component_name,omitempty"`

	// destination environment canonical
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[\da-zA-Z]+(?:[\da-zA-Z\-._]+[\da-zA-Z]|[\da-zA-Z])$
	DestinationEnvironmentCanonical string `json:"destination_environment_canonical,omitempty"`

	// destination project canonical
	// Max Length: 100
	// Min Length: 1
	// Pattern: (^[a-z0-9]+(([a-z0-9\-_]+)?[a-z0-9]+)?$)
	DestinationProjectCanonical string `json:"destination_project_canonical,omitempty"`
}

// Validate validates this migrate component
func (m *MigrateComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationComponentCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationEnvironmentCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationProjectCanonical(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MigrateComponent) validateDestinationComponentCanonical(formats strfmt.Registry) error {
	if swag.IsZero(m.DestinationComponentCanonical) { // not required
		return nil
	}

	if err := validate.MinLength("destination_component_canonical", "body", m.DestinationComponentCanonical, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("destination_component_canonical", "body", m.DestinationComponentCanonical, 100); err != nil {
		return err
	}

	if err := validate.Pattern("destination_component_canonical", "body", m.DestinationComponentCanonical, `^[\da-zA-Z]+(?:[\da-zA-Z\-._]+[\da-zA-Z]|[\da-zA-Z])$`); err != nil {
		return err
	}

	return nil
}

func (m *MigrateComponent) validateDestinationEnvironmentCanonical(formats strfmt.Registry) error {
	if swag.IsZero(m.DestinationEnvironmentCanonical) { // not required
		return nil
	}

	if err := validate.MinLength("destination_environment_canonical", "body", m.DestinationEnvironmentCanonical, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("destination_environment_canonical", "body", m.DestinationEnvironmentCanonical, 100); err != nil {
		return err
	}

	if err := validate.Pattern("destination_environment_canonical", "body", m.DestinationEnvironmentCanonical, `^[\da-zA-Z]+(?:[\da-zA-Z\-._]+[\da-zA-Z]|[\da-zA-Z])$`); err != nil {
		return err
	}

	return nil
}

func (m *MigrateComponent) validateDestinationProjectCanonical(formats strfmt.Registry) error {
	if swag.IsZero(m.DestinationProjectCanonical) { // not required
		return nil
	}

	if err := validate.MinLength("destination_project_canonical", "body", m.DestinationProjectCanonical, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("destination_project_canonical", "body", m.DestinationProjectCanonical, 100); err != nil {
		return err
	}

	if err := validate.Pattern("destination_project_canonical", "body", m.DestinationProjectCanonical, `(^[a-z0-9]+(([a-z0-9\-_]+)?[a-z0-9]+)?$)`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this migrate component based on context it is used
func (m *MigrateComponent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MigrateComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MigrateComponent) UnmarshalBinary(b []byte) error {
	var res MigrateComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
