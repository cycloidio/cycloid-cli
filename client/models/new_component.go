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

// NewComponent New Component
//
// The entity which represents the information of a new component.
//
// swagger:model NewComponent
type NewComponent struct {

	// canonical
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[\da-zA-Z]+(?:[\da-zA-Z\-._]+[\da-zA-Z]|[\da-zA-Z])$
	Canonical string `json:"canonical,omitempty"`

	// cloud provider canonical
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	CloudProviderCanonical string `json:"cloud_provider_canonical,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// name
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// It's the ref of the Stack, like 'cycloidio:stack-magento'
	// Required: true
	ServiceCatalogRef *string `json:"service_catalog_ref"`

	// use case
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: (^[a-z0-9]+(([a-z0-9\-_]+)?[a-z0-9]+)?$)
	UseCase *string `json:"use_case"`

	// vars
	Vars FormVariables `json:"vars,omitempty"`
}

// Validate validates this new component
func (m *NewComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudProviderCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceCatalogRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseCase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVars(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewComponent) validateCanonical(formats strfmt.Registry) error {
	if swag.IsZero(m.Canonical) { // not required
		return nil
	}

	if err := validate.MinLength("canonical", "body", m.Canonical, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("canonical", "body", m.Canonical, 100); err != nil {
		return err
	}

	if err := validate.Pattern("canonical", "body", m.Canonical, `^[\da-zA-Z]+(?:[\da-zA-Z\-._]+[\da-zA-Z]|[\da-zA-Z])$`); err != nil {
		return err
	}

	return nil
}

func (m *NewComponent) validateCloudProviderCanonical(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudProviderCanonical) { // not required
		return nil
	}

	if err := validate.MinLength("cloud_provider_canonical", "body", m.CloudProviderCanonical, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("cloud_provider_canonical", "body", m.CloudProviderCanonical, 100); err != nil {
		return err
	}

	if err := validate.Pattern("cloud_provider_canonical", "body", m.CloudProviderCanonical, `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *NewComponent) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	return nil
}

func (m *NewComponent) validateServiceCatalogRef(formats strfmt.Registry) error {

	if err := validate.Required("service_catalog_ref", "body", m.ServiceCatalogRef); err != nil {
		return err
	}

	return nil
}

func (m *NewComponent) validateUseCase(formats strfmt.Registry) error {

	if err := validate.Required("use_case", "body", m.UseCase); err != nil {
		return err
	}

	if err := validate.MinLength("use_case", "body", *m.UseCase, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("use_case", "body", *m.UseCase, 100); err != nil {
		return err
	}

	if err := validate.Pattern("use_case", "body", *m.UseCase, `(^[a-z0-9]+(([a-z0-9\-_]+)?[a-z0-9]+)?$)`); err != nil {
		return err
	}

	return nil
}

func (m *NewComponent) validateVars(formats strfmt.Registry) error {
	if swag.IsZero(m.Vars) { // not required
		return nil
	}

	if m.Vars != nil {
		if err := m.Vars.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vars")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vars")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this new component based on the context it is used
func (m *NewComponent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVars(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewComponent) contextValidateVars(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Vars) { // not required
		return nil
	}

	if err := m.Vars.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("vars")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("vars")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewComponent) UnmarshalBinary(b []byte) error {
	var res NewComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
