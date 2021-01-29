// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewRule NewRule
//
// NewRule represents an existing or new permission or constraint to access to an entity of the system. A Rule is aggregated into roles in order to be applied.
// swagger:model NewRule
type NewRule struct {

	// It can be the normal Policy.Code or contain globs like `*` or `**`
	// Required: true
	Action *string `json:"action"`

	// effect
	// Required: true
	// Enum: [allow]
	Effect *string `json:"effect"`

	// It is the list of resources in which this Rule applies to, the format of it is the one on the Policy.Code but with the `canonical` of the entities like `organization:org-can:team:team-can` for an action of `organization:team:read`
	// Required: true
	Resources []string `json:"resources"`
}

// Validate validates this new rule
func (m *NewRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewRule) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

var newRuleTypeEffectPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["allow"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		newRuleTypeEffectPropEnum = append(newRuleTypeEffectPropEnum, v)
	}
}

const (

	// NewRuleEffectAllow captures enum value "allow"
	NewRuleEffectAllow string = "allow"
)

// prop value enum
func (m *NewRule) validateEffectEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, newRuleTypeEffectPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NewRule) validateEffect(formats strfmt.Registry) error {

	if err := validate.Required("effect", "body", m.Effect); err != nil {
		return err
	}

	// value enum
	if err := m.validateEffectEnum("effect", "body", *m.Effect); err != nil {
		return err
	}

	return nil
}

func (m *NewRule) validateResources(formats strfmt.Registry) error {

	if err := validate.Required("resources", "body", m.Resources); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewRule) UnmarshalBinary(b []byte) error {
	var res NewRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}