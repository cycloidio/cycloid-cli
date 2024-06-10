// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CostEstimationResourceEstimate CostEstimationResourceEstimate
//
// Estimate for a single cloud resource.
//
// swagger:model CostEstimationResourceEstimate
type CostEstimationResourceEstimate struct {

	// Unique identifier of the resource.
	// Required: true
	Address *string `json:"address"`

	// List of resource cost components.
	// Required: true
	Components []*CostEstimationComponent `json:"components"`

	// Path to the image of the resource
	// Format: uri
	Image strfmt.URI `json:"image,omitempty"`

	// Planned monthly cost of the resource estimate in decimal form.
	PlannedCost string `json:"planned_cost,omitempty"`

	// Planned hourly cost of the resource estimate in decimal form.
	PlannedHourlyCost string `json:"planned_hourly_cost,omitempty"`

	// Prior monthly cost of the resource estimate in decimal form.
	PriorCost string `json:"prior_cost,omitempty"`

	// Prior hourly cost of the resource estimate in decimal form.
	PriorHourlyCost string `json:"prior_hourly_cost,omitempty"`

	// The resource's cloud provider.
	// Required: true
	Provider *string `json:"provider"`

	// Type of the resource.
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this cost estimation resource estimate
func (m *CostEstimationResourceEstimate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CostEstimationResourceEstimate) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *CostEstimationResourceEstimate) validateComponents(formats strfmt.Registry) error {

	if err := validate.Required("components", "body", m.Components); err != nil {
		return err
	}

	for i := 0; i < len(m.Components); i++ {
		if swag.IsZero(m.Components[i]) { // not required
			continue
		}

		if m.Components[i] != nil {
			if err := m.Components[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("components" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CostEstimationResourceEstimate) validateImage(formats strfmt.Registry) error {
	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if err := validate.FormatOf("image", "body", "uri", m.Image.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CostEstimationResourceEstimate) validateProvider(formats strfmt.Registry) error {

	if err := validate.Required("provider", "body", m.Provider); err != nil {
		return err
	}

	return nil
}

func (m *CostEstimationResourceEstimate) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cost estimation resource estimate based on the context it is used
func (m *CostEstimationResourceEstimate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CostEstimationResourceEstimate) contextValidateComponents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Components); i++ {

		if m.Components[i] != nil {

			if swag.IsZero(m.Components[i]) { // not required
				return nil
			}

			if err := m.Components[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("components" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CostEstimationResourceEstimate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CostEstimationResourceEstimate) UnmarshalBinary(b []byte) error {
	var res CostEstimationResourceEstimate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
