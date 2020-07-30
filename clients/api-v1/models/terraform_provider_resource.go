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

// TerraformProviderResource Resource
//
// A Resource of a Provider
// swagger:model TerraformProviderResource
type TerraformProviderResource struct {

	// canonical
	// Required: true
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	Canonical *string `json:"canonical"`

	// category
	// Required: true
	Category *string `json:"category"`

	// description
	// Required: true
	Description *string `json:"description"`

	// image
	// Required: true
	// Format: uri
	Image *strfmt.URI `json:"image"`

	// is edge
	// Required: true
	IsEdge *bool `json:"is_edge"`

	// is node
	// Required: true
	IsNode *bool `json:"is_node"`

	// keywords
	// Required: true
	Keywords []string `json:"keywords"`

	// schema
	// Required: true
	Schema interface{} `json:"schema"`

	// short description
	// Required: true
	ShortDescription *string `json:"short_description"`
}

// Validate validates this terraform provider resource
func (m *TerraformProviderResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsEdge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeywords(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShortDescription(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerraformProviderResource) validateCanonical(formats strfmt.Registry) error {

	if err := validate.Required("canonical", "body", m.Canonical); err != nil {
		return err
	}

	if err := validate.MinLength("canonical", "body", string(*m.Canonical), 3); err != nil {
		return err
	}

	if err := validate.Pattern("canonical", "body", string(*m.Canonical), `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *TerraformProviderResource) validateCategory(formats strfmt.Registry) error {

	if err := validate.Required("category", "body", m.Category); err != nil {
		return err
	}

	return nil
}

func (m *TerraformProviderResource) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *TerraformProviderResource) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	if err := validate.FormatOf("image", "body", "uri", m.Image.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TerraformProviderResource) validateIsEdge(formats strfmt.Registry) error {

	if err := validate.Required("is_edge", "body", m.IsEdge); err != nil {
		return err
	}

	return nil
}

func (m *TerraformProviderResource) validateIsNode(formats strfmt.Registry) error {

	if err := validate.Required("is_node", "body", m.IsNode); err != nil {
		return err
	}

	return nil
}

func (m *TerraformProviderResource) validateKeywords(formats strfmt.Registry) error {

	if err := validate.Required("keywords", "body", m.Keywords); err != nil {
		return err
	}

	return nil
}

func (m *TerraformProviderResource) validateSchema(formats strfmt.Registry) error {

	if err := validate.Required("schema", "body", m.Schema); err != nil {
		return err
	}

	return nil
}

func (m *TerraformProviderResource) validateShortDescription(formats strfmt.Registry) error {

	if err := validate.Required("short_description", "body", m.ShortDescription); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TerraformProviderResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TerraformProviderResource) UnmarshalBinary(b []byte) error {
	var res TerraformProviderResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}