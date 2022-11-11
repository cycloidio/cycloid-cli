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

// NewQuota New Quota
//
// The Quota defines the basic needs to create a quota
// swagger:model NewQuota
type NewQuota struct {

	// The amount of cpu that it has in units
	// Required: true
	// Minimum: 0
	CPU *uint64 `json:"cpu"`

	// The amount of memory that it has in MB
	// Required: true
	// Minimum: 0
	Memory *uint64 `json:"memory"`

	// The Resource Pool it'll be used
	// Required: true
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	ResourcePoolCanonical *string `json:"resource_pool_canonical"`

	// The amount of storage that it has in MB
	// Required: true
	// Minimum: 0
	Storage *uint64 `json:"storage"`

	// The Team it'll be limiting
	// Required: true
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	TeamCanonical *string `json:"team_canonical"`
}

// Validate validates this new quota
func (m *NewQuota) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourcePoolCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamCanonical(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewQuota) validateCPU(formats strfmt.Registry) error {

	if err := validate.Required("cpu", "body", m.CPU); err != nil {
		return err
	}

	if err := validate.MinimumInt("cpu", "body", int64(*m.CPU), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *NewQuota) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	if err := validate.MinimumInt("memory", "body", int64(*m.Memory), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *NewQuota) validateResourcePoolCanonical(formats strfmt.Registry) error {

	if err := validate.Required("resource_pool_canonical", "body", m.ResourcePoolCanonical); err != nil {
		return err
	}

	if err := validate.MinLength("resource_pool_canonical", "body", string(*m.ResourcePoolCanonical), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("resource_pool_canonical", "body", string(*m.ResourcePoolCanonical), 100); err != nil {
		return err
	}

	if err := validate.Pattern("resource_pool_canonical", "body", string(*m.ResourcePoolCanonical), `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *NewQuota) validateStorage(formats strfmt.Registry) error {

	if err := validate.Required("storage", "body", m.Storage); err != nil {
		return err
	}

	if err := validate.MinimumInt("storage", "body", int64(*m.Storage), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *NewQuota) validateTeamCanonical(formats strfmt.Registry) error {

	if err := validate.Required("team_canonical", "body", m.TeamCanonical); err != nil {
		return err
	}

	if err := validate.MinLength("team_canonical", "body", string(*m.TeamCanonical), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("team_canonical", "body", string(*m.TeamCanonical), 100); err != nil {
		return err
	}

	if err := validate.Pattern("team_canonical", "body", string(*m.TeamCanonical), `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewQuota) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewQuota) UnmarshalBinary(b []byte) error {
	var res NewQuota
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}