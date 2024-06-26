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

// NewServiceCatalogFromTemplate Service Catalog
//
// # Represents the Service Catalog item
//
// swagger:model NewServiceCatalogFromTemplate
type NewServiceCatalogFromTemplate struct {

	// author
	Author string `json:"author,omitempty"`

	// canonical
	// Required: true
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	Canonical *string `json:"canonical"`

	// created at
	// Minimum: 0
	CreatedAt *uint64 `json:"created_at,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// service catalog source canonical
	// Required: true
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	ServiceCatalogSourceCanonical *string `json:"service_catalog_source_canonical"`

	// updated at
	// Minimum: 0
	UpdatedAt *uint64 `json:"updated_at,omitempty"`

	// use case
	// Required: true
	UseCase *string `json:"use_case"`
}

// Validate validates this new service catalog from template
func (m *NewServiceCatalogFromTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceCatalogSourceCanonical(formats); err != nil {
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

func (m *NewServiceCatalogFromTemplate) validateCanonical(formats strfmt.Registry) error {

	if err := validate.Required("canonical", "body", m.Canonical); err != nil {
		return err
	}

	if err := validate.MinLength("canonical", "body", *m.Canonical, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("canonical", "body", *m.Canonical, 100); err != nil {
		return err
	}

	if err := validate.Pattern("canonical", "body", *m.Canonical, `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *NewServiceCatalogFromTemplate) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.MinimumUint("created_at", "body", *m.CreatedAt, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *NewServiceCatalogFromTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NewServiceCatalogFromTemplate) validateServiceCatalogSourceCanonical(formats strfmt.Registry) error {

	if err := validate.Required("service_catalog_source_canonical", "body", m.ServiceCatalogSourceCanonical); err != nil {
		return err
	}

	if err := validate.MinLength("service_catalog_source_canonical", "body", *m.ServiceCatalogSourceCanonical, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("service_catalog_source_canonical", "body", *m.ServiceCatalogSourceCanonical, 100); err != nil {
		return err
	}

	if err := validate.Pattern("service_catalog_source_canonical", "body", *m.ServiceCatalogSourceCanonical, `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *NewServiceCatalogFromTemplate) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.MinimumUint("updated_at", "body", *m.UpdatedAt, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *NewServiceCatalogFromTemplate) validateUseCase(formats strfmt.Registry) error {

	if err := validate.Required("use_case", "body", m.UseCase); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this new service catalog from template based on context it is used
func (m *NewServiceCatalogFromTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewServiceCatalogFromTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewServiceCatalogFromTemplate) UnmarshalBinary(b []byte) error {
	var res NewServiceCatalogFromTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
