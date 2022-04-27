// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CostExplorerAccountParent CostExplorerAccountParent
//
// Object containing Cost Explorer parent account parameters. The difference
// between it and CostExplorerAccount is that parent has no EB/Credential
// displayed, as it's not necessary
//
// swagger:model CostExplorerAccountParent
type CostExplorerAccountParent struct {

	// The ID of an account on the CP
	// Required: true
	AccountID *string `json:"account_id"`

	// canonical
	// Required: true
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	Canonical *string `json:"canonical"`

	// cloud provider
	// Required: true
	CloudProvider *CloudProvider `json:"cloud_provider"`

	// created at
	// Required: true
	// Minimum: 0
	CreatedAt *uint64 `json:"created_at"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// id
	// Required: true
	// Minimum: 1
	ID *uint32 `json:"id"`

	// last ingestion ended at
	// Minimum: 0
	LastIngestionEndedAt *uint64 `json:"last_ingestion_ended_at,omitempty"`

	// last ingestion started at
	// Minimum: 0
	LastIngestionStartedAt *uint64 `json:"last_ingestion_started_at,omitempty"`

	// status
	// Required: true
	// Enum: [idle error import]
	Status *string `json:"status"`

	// status message
	StatusMessage string `json:"status_message,omitempty"`

	// updated at
	// Required: true
	// Minimum: 0
	UpdatedAt *uint64 `json:"updated_at"`
}

// Validate validates this cost explorer account parent
func (m *CostExplorerAccountParent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastIngestionEndedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastIngestionStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CostExplorerAccountParent) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("account_id", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *CostExplorerAccountParent) validateCanonical(formats strfmt.Registry) error {

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

func (m *CostExplorerAccountParent) validateCloudProvider(formats strfmt.Registry) error {

	if err := validate.Required("cloud_provider", "body", m.CloudProvider); err != nil {
		return err
	}

	if m.CloudProvider != nil {
		if err := m.CloudProvider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloud_provider")
			}
			return err
		}
	}

	return nil
}

func (m *CostExplorerAccountParent) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.MinimumInt("created_at", "body", int64(*m.CreatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *CostExplorerAccountParent) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *CostExplorerAccountParent) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinimumInt("id", "body", int64(*m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *CostExplorerAccountParent) validateLastIngestionEndedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.LastIngestionEndedAt) { // not required
		return nil
	}

	if err := validate.MinimumInt("last_ingestion_ended_at", "body", int64(*m.LastIngestionEndedAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *CostExplorerAccountParent) validateLastIngestionStartedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.LastIngestionStartedAt) { // not required
		return nil
	}

	if err := validate.MinimumInt("last_ingestion_started_at", "body", int64(*m.LastIngestionStartedAt), 0, false); err != nil {
		return err
	}

	return nil
}

var costExplorerAccountParentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["idle","error","import"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		costExplorerAccountParentTypeStatusPropEnum = append(costExplorerAccountParentTypeStatusPropEnum, v)
	}
}

const (

	// CostExplorerAccountParentStatusIdle captures enum value "idle"
	CostExplorerAccountParentStatusIdle string = "idle"

	// CostExplorerAccountParentStatusError captures enum value "error"
	CostExplorerAccountParentStatusError string = "error"

	// CostExplorerAccountParentStatusImport captures enum value "import"
	CostExplorerAccountParentStatusImport string = "import"
)

// prop value enum
func (m *CostExplorerAccountParent) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, costExplorerAccountParentTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CostExplorerAccountParent) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *CostExplorerAccountParent) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.MinimumInt("updated_at", "body", int64(*m.UpdatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CostExplorerAccountParent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CostExplorerAccountParent) UnmarshalBinary(b []byte) error {
	var res CostExplorerAccountParent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
