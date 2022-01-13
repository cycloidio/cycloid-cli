// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FormUseCase Forms File Use case
// swagger:model FormUseCase
type FormUseCase struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// sections
	// Required: true
	Sections []*FormSection `json:"sections"`
}

// Validate validates this form use case
func (m *FormUseCase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSections(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FormUseCase) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *FormUseCase) validateSections(formats strfmt.Registry) error {

	if err := validate.Required("sections", "body", m.Sections); err != nil {
		return err
	}

	for i := 0; i < len(m.Sections); i++ {
		if swag.IsZero(m.Sections[i]) { // not required
			continue
		}

		if m.Sections[i] != nil {
			if err := m.Sections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FormUseCase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FormUseCase) UnmarshalBinary(b []byte) error {
	var res FormUseCase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
