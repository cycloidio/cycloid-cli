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

// ComponentSimple ComponentSimple
//
// The entity which represents minimal information of a component.
//
// swagger:model ComponentSimple
type ComponentSimple struct {

	// canonical
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[\da-zA-Z]+(?:(?:[\da-zA-Z\-._]+)?[\da-zA-Z])?$
	Canonical *string `json:"canonical"`

	// The cloud provider object that this environment is using.
	// In the event where the cloud provider is not yet defined/supported
	// that field might be empty.
	//
	CloudProvider *CloudProvider `json:"cloud_provider,omitempty"`

	// created at
	// Required: true
	// Minimum: 0
	CreatedAt *uint64 `json:"created_at"`

	// description
	Description string `json:"description,omitempty"`

	// id
	// Required: true
	// Minimum: 1
	ID *uint32 `json:"id"`

	// name
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// The Service Catalog that was used to create project.
	// Required: true
	ServiceCatalog *ServiceCatalog `json:"service_catalog"`

	// updated at
	// Required: true
	// Minimum: 0
	UpdatedAt *uint64 `json:"updated_at"`

	// use case
	// Required: true
	UseCase *string `json:"use_case"`
}

// Validate validates this component simple
func (m *ComponentSimple) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceCatalog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseCase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComponentSimple) validateCanonical(formats strfmt.Registry) error {

	if err := validate.Required("canonical", "body", m.Canonical); err != nil {
		return err
	}

	if err := validate.MinLength("canonical", "body", *m.Canonical, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("canonical", "body", *m.Canonical, 100); err != nil {
		return err
	}

	if err := validate.Pattern("canonical", "body", *m.Canonical, `^[\da-zA-Z]+(?:(?:[\da-zA-Z\-._]+)?[\da-zA-Z])?$`); err != nil {
		return err
	}

	return nil
}

func (m *ComponentSimple) validateCloudProvider(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudProvider) { // not required
		return nil
	}

	if m.CloudProvider != nil {
		if err := m.CloudProvider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloud_provider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloud_provider")
			}
			return err
		}
	}

	return nil
}

func (m *ComponentSimple) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.MinimumUint("created_at", "body", *m.CreatedAt, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ComponentSimple) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinimumUint("id", "body", uint64(*m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *ComponentSimple) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	return nil
}

func (m *ComponentSimple) validateServiceCatalog(formats strfmt.Registry) error {

	if err := validate.Required("service_catalog", "body", m.ServiceCatalog); err != nil {
		return err
	}

	if m.ServiceCatalog != nil {
		if err := m.ServiceCatalog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_catalog")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service_catalog")
			}
			return err
		}
	}

	return nil
}

func (m *ComponentSimple) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.MinimumUint("updated_at", "body", *m.UpdatedAt, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ComponentSimple) validateUseCase(formats strfmt.Registry) error {

	if err := validate.Required("use_case", "body", m.UseCase); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this component simple based on the context it is used
func (m *ComponentSimple) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCloudProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceCatalog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComponentSimple) contextValidateCloudProvider(ctx context.Context, formats strfmt.Registry) error {

	if m.CloudProvider != nil {

		if swag.IsZero(m.CloudProvider) { // not required
			return nil
		}

		if err := m.CloudProvider.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloud_provider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloud_provider")
			}
			return err
		}
	}

	return nil
}

func (m *ComponentSimple) contextValidateServiceCatalog(ctx context.Context, formats strfmt.Registry) error {

	if m.ServiceCatalog != nil {

		if err := m.ServiceCatalog.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_catalog")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service_catalog")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComponentSimple) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComponentSimple) UnmarshalBinary(b []byte) error {
	var res ComponentSimple
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
