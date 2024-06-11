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

// CostGroupDefinitions CostGroupDefinitions
//
// This object contains the items describe in https://docs.aws.amazon.com/sdk-for-go/api/service/costexplorer/#GroupDefinition It also grouping the costs based on different elements: az, services, tags, etc
//
// swagger:model CostGroupDefinitions
type CostGroupDefinitions struct {

	// group
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[a-zA-Z]+$
	Group *string `json:"group"`

	// key
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[a-zA-Z]+$
	Key *string `json:"key"`
}

// Validate validates this cost group definitions
func (m *CostGroupDefinitions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CostGroupDefinitions) validateGroup(formats strfmt.Registry) error {

	if err := validate.Required("group", "body", m.Group); err != nil {
		return err
	}

	if err := validate.MinLength("group", "body", *m.Group, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("group", "body", *m.Group, 100); err != nil {
		return err
	}

	if err := validate.Pattern("group", "body", *m.Group, `^[a-zA-Z]+$`); err != nil {
		return err
	}

	return nil
}

func (m *CostGroupDefinitions) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	if err := validate.MinLength("key", "body", *m.Key, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("key", "body", *m.Key, 100); err != nil {
		return err
	}

	if err := validate.Pattern("key", "body", *m.Key, `^[a-zA-Z]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cost group definitions based on context it is used
func (m *CostGroupDefinitions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CostGroupDefinitions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CostGroupDefinitions) UnmarshalBinary(b []byte) error {
	var res CostGroupDefinitions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
