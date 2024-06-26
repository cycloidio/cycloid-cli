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

// UpdateCloudCostManagementAccount Update CloudCostManagementAccount
//
// Update a Cloud Cost Management account to connect CP.
//
// swagger:model UpdateCloudCostManagementAccount
type UpdateCloudCostManagementAccount struct {

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// external backend
	// Required: true
	ExternalBackend *UpdateExternalBackend `json:"external_backend"`

	// A user-defined name for the account
	Name string `json:"name,omitempty"`
}

// Validate validates this update cloud cost management account
func (m *UpdateCloudCostManagementAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
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

func (m *UpdateCloudCostManagementAccount) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *UpdateCloudCostManagementAccount) validateExternalBackend(formats strfmt.Registry) error {

	if err := validate.Required("external_backend", "body", m.ExternalBackend); err != nil {
		return err
	}

	if m.ExternalBackend != nil {
		if err := m.ExternalBackend.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external_backend")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("external_backend")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update cloud cost management account based on the context it is used
func (m *UpdateCloudCostManagementAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExternalBackend(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateCloudCostManagementAccount) contextValidateExternalBackend(ctx context.Context, formats strfmt.Registry) error {

	if m.ExternalBackend != nil {

		if err := m.ExternalBackend.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external_backend")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("external_backend")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateCloudCostManagementAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateCloudCostManagementAccount) UnmarshalBinary(b []byte) error {
	var res UpdateCloudCostManagementAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
