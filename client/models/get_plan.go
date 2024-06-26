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

// GetPlan GetPlan
//
// The plan to get before running another one.
//
// swagger:model GetPlan
type GetPlan struct {

	// name
	Name string `json:"name,omitempty"`

	// params
	Params map[string]interface{} `json:"params,omitempty"`

	// resource
	// Required: true
	Resource *string `json:"resource"`

	// source
	// Required: true
	Source map[string]interface{} `json:"source"`

	// tags
	Tags []string `json:"tags"`

	// type
	// Required: true
	Type *string `json:"type"`

	// version
	Version map[string]string `json:"version,omitempty"`

	// version from
	VersionFrom string `json:"version_from,omitempty"`

	// versioned resource types
	VersionedResourceTypes []*VersionedResourceType `json:"versioned_resource_types"`
}

// Validate validates this get plan
func (m *GetPlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionedResourceTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetPlan) validateResource(formats strfmt.Registry) error {

	if err := validate.Required("resource", "body", m.Resource); err != nil {
		return err
	}

	return nil
}

func (m *GetPlan) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	for k := range m.Source {

		if err := validate.Required("source"+"."+k, "body", m.Source[k]); err != nil {
			return err
		}

	}

	return nil
}

func (m *GetPlan) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *GetPlan) validateVersionedResourceTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.VersionedResourceTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.VersionedResourceTypes); i++ {
		if swag.IsZero(m.VersionedResourceTypes[i]) { // not required
			continue
		}

		if m.VersionedResourceTypes[i] != nil {
			if err := m.VersionedResourceTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("versioned_resource_types" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("versioned_resource_types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get plan based on the context it is used
func (m *GetPlan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVersionedResourceTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetPlan) contextValidateVersionedResourceTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VersionedResourceTypes); i++ {

		if m.VersionedResourceTypes[i] != nil {

			if swag.IsZero(m.VersionedResourceTypes[i]) { // not required
				return nil
			}

			if err := m.VersionedResourceTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("versioned_resource_types" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("versioned_resource_types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetPlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetPlan) UnmarshalBinary(b []byte) error {
	var res GetPlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
