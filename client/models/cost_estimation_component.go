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

// CostEstimationComponent CostEstimationComponent
//
// Cost component of a cloud resource estimate.
// swagger:model CostEstimationComponent
type CostEstimationComponent struct {

	// Human-readable label of the component.
	// Required: true
	Label *string `json:"label"`

	// planned
	Planned *CostEstimationComponentState `json:"planned,omitempty"`

	// prior
	Prior *CostEstimationComponentState `json:"prior,omitempty"`

	// Monthly rate per unit.
	// Required: true
	Rate *string `json:"rate"`

	// Unit of estimation.
	Unit string `json:"unit,omitempty"`
}

// Validate validates this cost estimation component
func (m *CostEstimationComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanned(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrior(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CostEstimationComponent) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *CostEstimationComponent) validatePlanned(formats strfmt.Registry) error {

	if swag.IsZero(m.Planned) { // not required
		return nil
	}

	if m.Planned != nil {
		if err := m.Planned.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("planned")
			}
			return err
		}
	}

	return nil
}

func (m *CostEstimationComponent) validatePrior(formats strfmt.Registry) error {

	if swag.IsZero(m.Prior) { // not required
		return nil
	}

	if m.Prior != nil {
		if err := m.Prior.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prior")
			}
			return err
		}
	}

	return nil
}

func (m *CostEstimationComponent) validateRate(formats strfmt.Registry) error {

	if err := validate.Required("rate", "body", m.Rate); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CostEstimationComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CostEstimationComponent) UnmarshalBinary(b []byte) error {
	var res CostEstimationComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}