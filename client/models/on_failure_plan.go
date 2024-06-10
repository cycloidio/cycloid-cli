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

// OnFailurePlan OnFailurePlan
//
// The plan definition when the action has failed.
//
// swagger:model OnFailurePlan
type OnFailurePlan struct {

	// next
	// Required: true
	Next *Plan `json:"next"`

	// step
	// Required: true
	Step *Plan `json:"step"`
}

// Validate validates this on failure plan
func (m *OnFailurePlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStep(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OnFailurePlan) validateNext(formats strfmt.Registry) error {

	if err := validate.Required("next", "body", m.Next); err != nil {
		return err
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

func (m *OnFailurePlan) validateStep(formats strfmt.Registry) error {

	if err := validate.Required("step", "body", m.Step); err != nil {
		return err
	}

	if m.Step != nil {
		if err := m.Step.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("step")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("step")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this on failure plan based on the context it is used
func (m *OnFailurePlan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStep(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OnFailurePlan) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {

		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

func (m *OnFailurePlan) contextValidateStep(ctx context.Context, formats strfmt.Registry) error {

	if m.Step != nil {

		if err := m.Step.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("step")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("step")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OnFailurePlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OnFailurePlan) UnmarshalBinary(b []byte) error {
	var res OnFailurePlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
