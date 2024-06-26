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

// CloudCostManagementProviderDetails CloudCostManagementProviderDetails
//
// # Description of the costs of a specific provider
//
// swagger:model CloudCostManagementProviderDetails
type CloudCostManagementProviderDetails struct {

	// cost histogram
	// Required: true
	CostHistogram *CloudCostManagementHistogram `json:"cost_histogram"`

	// filter values
	FilterValues *CloudCostManagementFilterValues `json:"filter_values,omitempty"`
}

// Validate validates this cloud cost management provider details
func (m *CloudCostManagementProviderDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCostHistogram(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilterValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudCostManagementProviderDetails) validateCostHistogram(formats strfmt.Registry) error {

	if err := validate.Required("cost_histogram", "body", m.CostHistogram); err != nil {
		return err
	}

	if m.CostHistogram != nil {
		if err := m.CostHistogram.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cost_histogram")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cost_histogram")
			}
			return err
		}
	}

	return nil
}

func (m *CloudCostManagementProviderDetails) validateFilterValues(formats strfmt.Registry) error {
	if swag.IsZero(m.FilterValues) { // not required
		return nil
	}

	if m.FilterValues != nil {
		if err := m.FilterValues.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter_values")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filter_values")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cloud cost management provider details based on the context it is used
func (m *CloudCostManagementProviderDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCostHistogram(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilterValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudCostManagementProviderDetails) contextValidateCostHistogram(ctx context.Context, formats strfmt.Registry) error {

	if m.CostHistogram != nil {

		if err := m.CostHistogram.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cost_histogram")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cost_histogram")
			}
			return err
		}
	}

	return nil
}

func (m *CloudCostManagementProviderDetails) contextValidateFilterValues(ctx context.Context, formats strfmt.Registry) error {

	if m.FilterValues != nil {

		if swag.IsZero(m.FilterValues) { // not required
			return nil
		}

		if err := m.FilterValues.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter_values")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filter_values")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudCostManagementProviderDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudCostManagementProviderDetails) UnmarshalBinary(b []byte) error {
	var res CloudCostManagementProviderDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
