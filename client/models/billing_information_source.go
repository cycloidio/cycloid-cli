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

// BillingInformationSource Billing Information Source
//
// The information about the Source
// swagger:model BillingInformationSource
type BillingInformationSource struct {

	// brand
	// Required: true
	Brand *string `json:"brand"`

	// expiration month
	// Required: true
	// Minimum: 1
	ExpirationMonth *uint8 `json:"expiration_month"`

	// expiration year
	// Required: true
	ExpirationYear *uint8 `json:"expiration_year"`

	// last4
	// Required: true
	// Max Length: 4
	// Min Length: 4
	// Pattern: ^\d{4}$
	Last4 *string `json:"last4"`
}

// Validate validates this billing information source
func (m *BillingInformationSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBrand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationMonth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationYear(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLast4(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingInformationSource) validateBrand(formats strfmt.Registry) error {

	if err := validate.Required("brand", "body", m.Brand); err != nil {
		return err
	}

	return nil
}

func (m *BillingInformationSource) validateExpirationMonth(formats strfmt.Registry) error {

	if err := validate.Required("expiration_month", "body", m.ExpirationMonth); err != nil {
		return err
	}

	if err := validate.MinimumInt("expiration_month", "body", int64(*m.ExpirationMonth), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *BillingInformationSource) validateExpirationYear(formats strfmt.Registry) error {

	if err := validate.Required("expiration_year", "body", m.ExpirationYear); err != nil {
		return err
	}

	return nil
}

func (m *BillingInformationSource) validateLast4(formats strfmt.Registry) error {

	if err := validate.Required("last4", "body", m.Last4); err != nil {
		return err
	}

	if err := validate.MinLength("last4", "body", string(*m.Last4), 4); err != nil {
		return err
	}

	if err := validate.MaxLength("last4", "body", string(*m.Last4), 4); err != nil {
		return err
	}

	if err := validate.Pattern("last4", "body", string(*m.Last4), `^\d{4}$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BillingInformationSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingInformationSource) UnmarshalBinary(b []byte) error {
	var res BillingInformationSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
