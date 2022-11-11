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

// ResourcePool Resource Pool
//
// A Resource Pool holds the information of all the Resources that have the same label. The Used is the amount used by Projects using Quotas and Allocated is the amount declared by Quotas
// swagger:model ResourcePool
type ResourcePool struct {

	// The amount of CPU that is allocated (quotas defined) in units
	// Required: true
	// Minimum: 0
	AllocatedCPU *uint64 `json:"allocated_cpu"`

	// The amount of memory that is allocated (quotas defined) in MB
	// Required: true
	// Minimum: 0
	AllocatedMemory *uint64 `json:"allocated_memory"`

	// The amount of Storage that is allocated (quotas defined) in MB
	// Required: true
	// Minimum: 0
	AllocatedStorage *uint64 `json:"allocated_storage"`

	// canonical
	// Required: true
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	Canonical *string `json:"canonical"`

	// The amount of cpu that it has in units
	// Required: true
	// Minimum: 0
	CPU *uint64 `json:"cpu"`

	// id
	// Minimum: 1
	ID uint32 `json:"id,omitempty"`

	// The label to which match the resources
	// Required: true
	Label *string `json:"label"`

	// The amount of memory that it has in MB
	// Required: true
	// Minimum: 0
	Memory *uint64 `json:"memory"`

	// name
	// Required: true
	Name *string `json:"name"`

	// The amount of storage that it has in MB
	// Required: true
	// Minimum: 0
	Storage *uint64 `json:"storage"`

	// The amount of CPU that is used in units
	// Required: true
	// Minimum: 0
	UsedCPU *uint64 `json:"used_cpu"`

	// The amount of memory that is used in MB
	// Required: true
	// Minimum: 0
	UsedMemory *uint64 `json:"used_memory"`

	// The amount of Storage that is used in MB
	// Required: true
	// Minimum: 0
	UsedStorage *uint64 `json:"used_storage"`
}

// Validate validates this resource pool
func (m *ResourcePool) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocatedCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAllocatedMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAllocatedStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsedCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsedMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsedStorage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourcePool) validateAllocatedCPU(formats strfmt.Registry) error {

	if err := validate.Required("allocated_cpu", "body", m.AllocatedCPU); err != nil {
		return err
	}

	if err := validate.MinimumInt("allocated_cpu", "body", int64(*m.AllocatedCPU), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ResourcePool) validateAllocatedMemory(formats strfmt.Registry) error {

	if err := validate.Required("allocated_memory", "body", m.AllocatedMemory); err != nil {
		return err
	}

	if err := validate.MinimumInt("allocated_memory", "body", int64(*m.AllocatedMemory), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ResourcePool) validateAllocatedStorage(formats strfmt.Registry) error {

	if err := validate.Required("allocated_storage", "body", m.AllocatedStorage); err != nil {
		return err
	}

	if err := validate.MinimumInt("allocated_storage", "body", int64(*m.AllocatedStorage), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ResourcePool) validateCanonical(formats strfmt.Registry) error {

	if err := validate.Required("canonical", "body", m.Canonical); err != nil {
		return err
	}

	if err := validate.MinLength("canonical", "body", string(*m.Canonical), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("canonical", "body", string(*m.Canonical), 100); err != nil {
		return err
	}

	if err := validate.Pattern("canonical", "body", string(*m.Canonical), `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *ResourcePool) validateCPU(formats strfmt.Registry) error {

	if err := validate.Required("cpu", "body", m.CPU); err != nil {
		return err
	}

	if err := validate.MinimumInt("cpu", "body", int64(*m.CPU), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ResourcePool) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinimumInt("id", "body", int64(m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *ResourcePool) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *ResourcePool) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	if err := validate.MinimumInt("memory", "body", int64(*m.Memory), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ResourcePool) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ResourcePool) validateStorage(formats strfmt.Registry) error {

	if err := validate.Required("storage", "body", m.Storage); err != nil {
		return err
	}

	if err := validate.MinimumInt("storage", "body", int64(*m.Storage), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ResourcePool) validateUsedCPU(formats strfmt.Registry) error {

	if err := validate.Required("used_cpu", "body", m.UsedCPU); err != nil {
		return err
	}

	if err := validate.MinimumInt("used_cpu", "body", int64(*m.UsedCPU), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ResourcePool) validateUsedMemory(formats strfmt.Registry) error {

	if err := validate.Required("used_memory", "body", m.UsedMemory); err != nil {
		return err
	}

	if err := validate.MinimumInt("used_memory", "body", int64(*m.UsedMemory), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ResourcePool) validateUsedStorage(formats strfmt.Registry) error {

	if err := validate.Required("used_storage", "body", m.UsedStorage); err != nil {
		return err
	}

	if err := validate.MinimumInt("used_storage", "body", int64(*m.UsedStorage), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourcePool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourcePool) UnmarshalBinary(b []byte) error {
	var res ResourcePool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}