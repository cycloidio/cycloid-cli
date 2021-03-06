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

// InfraPoliciesValidationResultItem InfraPoliciesValidationResultItem represents a checked rule
// with a failed result generated from the validation process.
//
// swagger:model InfraPoliciesValidationResultItem
type InfraPoliciesValidationResultItem struct {

	// infra policy
	// Required: true
	InfraPolicy *InfraPolicy `json:"infra_policy"`

	// The messages about the reason of the validation failure
	// that's written in the InfraPolicy's Body of the failed rule.
	//
	Reasons []string `json:"reasons"`
}

// Validate validates this infra policies validation result item
func (m *InfraPoliciesValidationResultItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInfraPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InfraPoliciesValidationResultItem) validateInfraPolicy(formats strfmt.Registry) error {

	if err := validate.Required("infra_policy", "body", m.InfraPolicy); err != nil {
		return err
	}

	if m.InfraPolicy != nil {
		if err := m.InfraPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infra_policy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InfraPoliciesValidationResultItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfraPoliciesValidationResultItem) UnmarshalBinary(b []byte) error {
	var res InfraPoliciesValidationResultItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
