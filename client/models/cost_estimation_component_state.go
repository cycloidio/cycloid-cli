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

// CostEstimationComponentState CostEstimationComponentState
//
// Either a Prior or Planned cost component state.
// swagger:model CostEstimationComponentState
type CostEstimationComponentState struct {

	// Cost of the component state in decimal form.
	// Required: true
	Cost *string `json:"cost"`

	// List of items on which the cost is dependent.
	// Required: true
	Details []string `json:"details"`

	// Quantity of units.
	// Required: true
	Quantity *uint32 `json:"quantity"`
}

// Validate validates this cost estimation component state
func (m *CostEstimationComponentState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetails(formats); err != nil {
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

func (m *CostEstimationComponentState) validateCost(formats strfmt.Registry) error {

	if err := validate.Required("cost", "body", m.Cost); err != nil {
		return err
	}

	return nil
}

func (m *CostEstimationComponentState) validateDetails(formats strfmt.Registry) error {

	if err := validate.Required("details", "body", m.Details); err != nil {
		return err
	}

	return nil
}

func (m *CostEstimationComponentState) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CostEstimationComponentState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CostEstimationComponentState) UnmarshalBinary(b []byte) error {
	var res CostEstimationComponentState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}