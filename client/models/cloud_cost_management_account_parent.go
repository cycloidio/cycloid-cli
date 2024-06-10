// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CloudCostManagementAccountParent CloudCostManagementAccountParent
//
// Object containing Cloud Cost Management parent account parameters. The difference
// between it and CloudCostManagementAccount is that parent has no EB/Credential
// displayed, as it's not necessary
//
// swagger:model CloudCostManagementAccountParent
type CloudCostManagementAccountParent struct {

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

	// A user-defined name for the account
	// Required: true
	Name *string `json:"name"`

	// status
	// Required: true
	// Enum: ["idle","error","import"]
	Status *string `json:"status"`

	// status message
	StatusMessage string `json:"status_message,omitempty"`

	// updated at
	// Required: true
	// Minimum: 0
	UpdatedAt *uint64 `json:"updated_at"`
}

// Validate validates this cloud cost management account parent
func (m *CloudCostManagementAccountParent) Validate(formats strfmt.Registry) error {
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

	if err := m.validateName(formats); err != nil {
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

func (m *CloudCostManagementAccountParent) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("account_id", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *CloudCostManagementAccountParent) validateCanonical(formats strfmt.Registry) error {

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

func (m *CloudCostManagementAccountParent) validateCloudProvider(formats strfmt.Registry) error {

	if err := validate.Required("cloud_provider", "body", m.CloudProvider); err != nil {
		return err
	}

	if m.CloudProvider != nil {
		if err := m.CloudProvider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloud_provider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloud_provider")
			}
			return err
		}
	}

	return nil
}

func (m *CloudCostManagementAccountParent) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.MinimumUint("created_at", "body", *m.CreatedAt, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *CloudCostManagementAccountParent) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *CloudCostManagementAccountParent) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinimumUint("id", "body", uint64(*m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *CloudCostManagementAccountParent) validateLastIngestionEndedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.LastIngestionEndedAt) { // not required
		return nil
	}

	if err := validate.MinimumUint("last_ingestion_ended_at", "body", *m.LastIngestionEndedAt, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *CloudCostManagementAccountParent) validateLastIngestionStartedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.LastIngestionStartedAt) { // not required
		return nil
	}

	if err := validate.MinimumUint("last_ingestion_started_at", "body", *m.LastIngestionStartedAt, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *CloudCostManagementAccountParent) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var cloudCostManagementAccountParentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["idle","error","import"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudCostManagementAccountParentTypeStatusPropEnum = append(cloudCostManagementAccountParentTypeStatusPropEnum, v)
	}
}

const (

	// CloudCostManagementAccountParentStatusIdle captures enum value "idle"
	CloudCostManagementAccountParentStatusIdle string = "idle"

	// CloudCostManagementAccountParentStatusError captures enum value "error"
	CloudCostManagementAccountParentStatusError string = "error"

	// CloudCostManagementAccountParentStatusImport captures enum value "import"
	CloudCostManagementAccountParentStatusImport string = "import"
)

// prop value enum
func (m *CloudCostManagementAccountParent) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cloudCostManagementAccountParentTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CloudCostManagementAccountParent) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *CloudCostManagementAccountParent) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.MinimumUint("updated_at", "body", *m.UpdatedAt, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cloud cost management account parent based on the context it is used
func (m *CloudCostManagementAccountParent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCloudProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudCostManagementAccountParent) contextValidateCloudProvider(ctx context.Context, formats strfmt.Registry) error {

	if m.CloudProvider != nil {

		if err := m.CloudProvider.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloud_provider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloud_provider")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudCostManagementAccountParent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudCostManagementAccountParent) UnmarshalBinary(b []byte) error {
	var res CloudCostManagementAccountParent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
