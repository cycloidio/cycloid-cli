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

// UpdateWatchRule Update Watch rule
//
// # The entity which represents the information of a Watch rule
//
// swagger:model UpdateWatchRule
type UpdateWatchRule struct {

	// filters
	Filters []*NewWatchRuleFilter `json:"filters"`

	// muted
	// Required: true
	Muted *bool `json:"muted"`

	// name
	// Required: true
	// Min Length: 3
	Name *string `json:"name"`
}

// Validate validates this update watch rule
func (m *UpdateWatchRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMuted(formats); err != nil {
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

func (m *UpdateWatchRule) validateFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	for i := 0; i < len(m.Filters); i++ {
		if swag.IsZero(m.Filters[i]) { // not required
			continue
		}

		if m.Filters[i] != nil {
			if err := m.Filters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateWatchRule) validateMuted(formats strfmt.Registry) error {

	if err := validate.Required("muted", "body", m.Muted); err != nil {
		return err
	}

	return nil
}

func (m *UpdateWatchRule) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 3); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update watch rule based on the context it is used
func (m *UpdateWatchRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateWatchRule) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Filters); i++ {

		if m.Filters[i] != nil {

			if swag.IsZero(m.Filters[i]) { // not required
				return nil
			}

			if err := m.Filters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateWatchRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateWatchRule) UnmarshalBinary(b []byte) error {
	var res UpdateWatchRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
