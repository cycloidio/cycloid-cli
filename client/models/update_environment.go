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

// UpdateEnvironment Update Environment
//
// # Represent an entity necessary for environment update
//
// swagger:model UpdateEnvironment
type UpdateEnvironment struct {

	// cloud provider canonical
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	CloudProviderCanonical string `json:"cloud_provider_canonical,omitempty"`

	// color
	// Max Length: 64
	Color string `json:"color,omitempty"`

	// icon
	// Max Length: 64
	Icon string `json:"icon,omitempty"`

	// The variables set within a form with the corresponding environment
	// canonical and use case
	//
	Inputs []*FormInput `json:"inputs"`

	// use case
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	UseCase string `json:"use_case,omitempty"`
}

// Validate validates this update environment
func (m *UpdateEnvironment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudProviderCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIcon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInputs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseCase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateEnvironment) validateCloudProviderCanonical(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudProviderCanonical) { // not required
		return nil
	}

	if err := validate.MinLength("cloud_provider_canonical", "body", m.CloudProviderCanonical, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("cloud_provider_canonical", "body", m.CloudProviderCanonical, 100); err != nil {
		return err
	}

	if err := validate.Pattern("cloud_provider_canonical", "body", m.CloudProviderCanonical, `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *UpdateEnvironment) validateColor(formats strfmt.Registry) error {
	if swag.IsZero(m.Color) { // not required
		return nil
	}

	if err := validate.MaxLength("color", "body", m.Color, 64); err != nil {
		return err
	}

	return nil
}

func (m *UpdateEnvironment) validateIcon(formats strfmt.Registry) error {
	if swag.IsZero(m.Icon) { // not required
		return nil
	}

	if err := validate.MaxLength("icon", "body", m.Icon, 64); err != nil {
		return err
	}

	return nil
}

func (m *UpdateEnvironment) validateInputs(formats strfmt.Registry) error {
	if swag.IsZero(m.Inputs) { // not required
		return nil
	}

	for i := 0; i < len(m.Inputs); i++ {
		if swag.IsZero(m.Inputs[i]) { // not required
			continue
		}

		if m.Inputs[i] != nil {
			if err := m.Inputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateEnvironment) validateUseCase(formats strfmt.Registry) error {
	if swag.IsZero(m.UseCase) { // not required
		return nil
	}

	if err := validate.MinLength("use_case", "body", m.UseCase, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("use_case", "body", m.UseCase, 100); err != nil {
		return err
	}

	if err := validate.Pattern("use_case", "body", m.UseCase, `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update environment based on the context it is used
func (m *UpdateEnvironment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInputs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateEnvironment) contextValidateInputs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Inputs); i++ {

		if m.Inputs[i] != nil {

			if swag.IsZero(m.Inputs[i]) { // not required
				return nil
			}

			if err := m.Inputs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateEnvironment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateEnvironment) UnmarshalBinary(b []byte) error {
	var res UpdateEnvironment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
