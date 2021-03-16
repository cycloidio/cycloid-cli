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

// Rule Rule
//
// Rules define the specific access to the platform
// swagger:model Rule
type Rule struct {

	// It can be the normal Policy.Code or contain globs like `*` or `**`
	// Required: true
	Action *string `json:"action"`

	// effect
	// Required: true
	// Enum: [allow]
	Effect *string `json:"effect"`

	// This is the id of the row from the database, but for blocking organizations we generate rules that are not in the database. When this happens the id is allowed to be 0.
	// Required: true
	// Minimum: 0
	ID *uint32 `json:"id"`

	// It is the list of resources in which this Rule applies to, the format of it is the one on the Policy.Code but with the `canonical` of the entities like `organization:org-can:team:team-can` for an action of `organization:team:read`
	// Required: true
	Resources []string `json:"resources"`
}

// Validate validates this rule
func (m *Rule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
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

func (m *Rule) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

var ruleTypeEffectPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["allow"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ruleTypeEffectPropEnum = append(ruleTypeEffectPropEnum, v)
	}
}

const (

	// RuleEffectAllow captures enum value "allow"
	RuleEffectAllow string = "allow"
)

// prop value enum
func (m *Rule) validateEffectEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ruleTypeEffectPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Rule) validateEffect(formats strfmt.Registry) error {

	if err := validate.Required("effect", "body", m.Effect); err != nil {
		return err
	}

	// value enum
	if err := m.validateEffectEnum("effect", "body", *m.Effect); err != nil {
		return err
	}

	return nil
}

func (m *Rule) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinimumInt("id", "body", int64(*m.ID), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Rule) validateResources(formats strfmt.Registry) error {

	if err := validate.Required("resources", "body", m.Resources); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Rule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Rule) UnmarshalBinary(b []byte) error {
	var res Rule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
