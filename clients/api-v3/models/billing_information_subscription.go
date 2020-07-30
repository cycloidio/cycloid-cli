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

// BillingInformationSubscription Billing Information Subscription
//
// The information about the Subscription
// swagger:model BillingInformationSubscription
type BillingInformationSubscription struct {

	// amount
	// Required: true
	// Minimum: 0
	Amount *float64 `json:"amount"`

	// current period end
	// Required: true
	// Minimum: 0
	CurrentPeriodEnd *int64 `json:"current_period_end"`

	// quantity
	// Required: true
	// Minimum: 0
	Quantity *int64 `json:"quantity"`
}

// Validate validates this billing information subscription
func (m *BillingInformationSubscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentPeriodEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingInformationSubscription) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	if err := validate.Minimum("amount", "body", float64(*m.Amount), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *BillingInformationSubscription) validateCurrentPeriodEnd(formats strfmt.Registry) error {

	if err := validate.Required("current_period_end", "body", m.CurrentPeriodEnd); err != nil {
		return err
	}

	if err := validate.MinimumInt("current_period_end", "body", int64(*m.CurrentPeriodEnd), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *BillingInformationSubscription) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	if err := validate.MinimumInt("quantity", "body", int64(*m.Quantity), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BillingInformationSubscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingInformationSubscription) UnmarshalBinary(b []byte) error {
	var res BillingInformationSubscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}