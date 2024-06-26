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

// UpdateQuota Update Quota
//
// # The Quota defines the basic needs to update a create
//
// swagger:model UpdateQuota
type UpdateQuota struct {

	// The amount of cpu that it has in units
	// Required: true
	// Minimum: 0
	CPU *uint64 `json:"cpu"`

	// The amount of memory that it has in MB
	// Required: true
	// Minimum: 0
	Memory *uint64 `json:"memory"`

	// The amount of storage that it has in MB
	// Required: true
	// Minimum: 0
	Storage *uint64 `json:"storage"`
}

// Validate validates this update quota
func (m *UpdateQuota) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateQuota) validateCPU(formats strfmt.Registry) error {

	if err := validate.Required("cpu", "body", m.CPU); err != nil {
		return err
	}

	if err := validate.MinimumUint("cpu", "body", *m.CPU, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *UpdateQuota) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	if err := validate.MinimumUint("memory", "body", *m.Memory, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *UpdateQuota) validateStorage(formats strfmt.Registry) error {

	if err := validate.Required("storage", "body", m.Storage); err != nil {
		return err
	}

	if err := validate.MinimumUint("storage", "body", *m.Storage, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update quota based on context it is used
func (m *UpdateQuota) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateQuota) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateQuota) UnmarshalBinary(b []byte) error {
	var res UpdateQuota
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
