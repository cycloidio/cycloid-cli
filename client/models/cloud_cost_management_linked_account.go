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

// CloudCostManagementLinkedAccount CloudCostManagementLinkedAccount
//
// An account that is linked to a Cloud Cost Management account in the cloud
// provider. It is not managed by Cloud Cost Management, but appears in the cost records.
//
// swagger:model CloudCostManagementLinkedAccount
type CloudCostManagementLinkedAccount struct {

	// ID of the account in the cloud provider
	// Required: true
	AccountID *string `json:"account_id"`

	// The cloud provider canonical that this project is using - between the
	// supported ones.
	//
	// Required: true
	CloudProvider *string `json:"cloud_provider"`

	// id
	// Required: true
	// Minimum: 1
	ID *uint32 `json:"id"`

	// A user-defined name for the account
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this cloud cost management linked account
func (m *CloudCostManagementLinkedAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudCostManagementLinkedAccount) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("account_id", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *CloudCostManagementLinkedAccount) validateCloudProvider(formats strfmt.Registry) error {

	if err := validate.Required("cloud_provider", "body", m.CloudProvider); err != nil {
		return err
	}

	return nil
}

func (m *CloudCostManagementLinkedAccount) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinimumUint("id", "body", uint64(*m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *CloudCostManagementLinkedAccount) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cloud cost management linked account based on context it is used
func (m *CloudCostManagementLinkedAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CloudCostManagementLinkedAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudCostManagementLinkedAccount) UnmarshalBinary(b []byte) error {
	var res CloudCostManagementLinkedAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
