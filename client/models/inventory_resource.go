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

// InventoryResource Inventory Resource
//
// # The Resource of the Inventory representing an element of your infrastructure
//
// swagger:model InventoryResource
type InventoryResource struct {

	// attributes
	Attributes interface{} `json:"attributes,omitempty"`

	// Category of the resource type
	Category string `json:"category,omitempty"`

	// The amount of cpu that it has in units
	// Minimum: 0
	CPU *uint64 `json:"cpu,omitempty"`

	// List of attributes of the Resource, can be anything
	CustomAttributes interface{} `json:"custom_attributes,omitempty"`

	// Full description of the resource type documentation
	Description string `json:"description,omitempty"`

	// Environment canonical in which this resource is used
	// Pattern: ^[\da-zA-Z]+(?:(?:[\da-zA-Z\-._]+)?[\da-zA-Z])?$
	EnvironmentCanonical string `json:"environment_canonical,omitempty"`

	// id
	// Minimum: 1
	ID uint32 `json:"id,omitempty"`

	// Image of the resource type
	// Format: uri
	Image strfmt.URI `json:"image,omitempty"`

	// Set of keywords to categorize the resource type
	Keywords []string `json:"keywords"`

	// A way to distinguish and categorize resources
	Label string `json:"label,omitempty"`

	// The amount of memory that it has in MB
	// Minimum: 0
	Memory *uint64 `json:"memory,omitempty"`

	// The way this resource is handled
	Mode string `json:"mode,omitempty"`

	// The module it belongs to
	Module string `json:"module,omitempty"`

	// The name of the resource
	// Required: true
	Name *string `json:"name"`

	// Project canonical in which this resource is used
	// Pattern: (^[a-z0-9]+(([a-z0-9\-_]+)?[a-z0-9]+)?$)
	ProjectCanonical string `json:"project_canonical,omitempty"`

	// The provider of the created Resource
	// Required: true
	Provider *string `json:"provider"`

	// Short description of the resource type documentation
	ShortDescription string `json:"short_description,omitempty"`

	// The amount of storage that it has in MB
	// Minimum: 0
	Storage *uint64 `json:"storage,omitempty"`

	// The type of the resource
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this inventory resource
func (m *InventoryResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InventoryResource) validateCPU(formats strfmt.Registry) error {
	if swag.IsZero(m.CPU) { // not required
		return nil
	}

	if err := validate.MinimumUint("cpu", "body", *m.CPU, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *InventoryResource) validateEnvironmentCanonical(formats strfmt.Registry) error {
	if swag.IsZero(m.EnvironmentCanonical) { // not required
		return nil
	}

	if err := validate.Pattern("environment_canonical", "body", m.EnvironmentCanonical, `^[\da-zA-Z]+(?:(?:[\da-zA-Z\-._]+)?[\da-zA-Z])?$`); err != nil {
		return err
	}

	return nil
}

func (m *InventoryResource) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinimumUint("id", "body", uint64(m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *InventoryResource) validateImage(formats strfmt.Registry) error {
	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if err := validate.FormatOf("image", "body", "uri", m.Image.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InventoryResource) validateMemory(formats strfmt.Registry) error {
	if swag.IsZero(m.Memory) { // not required
		return nil
	}

	if err := validate.MinimumUint("memory", "body", *m.Memory, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *InventoryResource) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *InventoryResource) validateProjectCanonical(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectCanonical) { // not required
		return nil
	}

	if err := validate.Pattern("project_canonical", "body", m.ProjectCanonical, `(^[a-z0-9]+(([a-z0-9\-_]+)?[a-z0-9]+)?$)`); err != nil {
		return err
	}

	return nil
}

func (m *InventoryResource) validateProvider(formats strfmt.Registry) error {

	if err := validate.Required("provider", "body", m.Provider); err != nil {
		return err
	}

	return nil
}

func (m *InventoryResource) validateStorage(formats strfmt.Registry) error {
	if swag.IsZero(m.Storage) { // not required
		return nil
	}

	if err := validate.MinimumUint("storage", "body", *m.Storage, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *InventoryResource) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this inventory resource based on context it is used
func (m *InventoryResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InventoryResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryResource) UnmarshalBinary(b []byte) error {
	var res InventoryResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
