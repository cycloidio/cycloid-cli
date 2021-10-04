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

// Pagination Pagination
// swagger:model Pagination
type Pagination struct {

	// The index of the page sent (the first page is 1).
	// Required: true
	// Minimum: 1
	Index *uint64 `json:"index"`

	// The size of the page (the number of entities per page)
	// Required: true
	// Minimum: 0
	Size *uint64 `json:"size"`

	// The total number of items.
	// Required: true
	// Minimum: 0
	Total *uint64 `json:"total"`
}

// Validate validates this pagination
func (m *Pagination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Pagination) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	if err := validate.MinimumInt("index", "body", int64(*m.Index), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Pagination) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	if err := validate.MinimumInt("size", "body", int64(*m.Size), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Pagination) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("total", "body", m.Total); err != nil {
		return err
	}

	if err := validate.MinimumInt("total", "body", int64(*m.Total), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Pagination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Pagination) UnmarshalBinary(b []byte) error {
	var res Pagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
