// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceCatalog Service Catalog
//
// Represents the Service Catalog item
// swagger:model ServiceCatalog
type ServiceCatalog struct {

	// author
	// Required: true
	Author *string `json:"author"`

	// canonical
	// Required: true
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	Canonical *string `json:"canonical"`

	// created at
	// Minimum: 0
	CreatedAt *uint64 `json:"created_at,omitempty"`

	// dependencies
	Dependencies []*ServiceCatalogDependency `json:"dependencies"`

	// description
	// Required: true
	Description *string `json:"description"`

	// Directory where the ServiceCatalog configuration is found.
	// Required: true
	Directory *string `json:"directory"`

	// Indicates if this stack can be configured with form's or not. Based on the presence or not of a valid .forms.yaml file since it's creation or last refresh.
	// Required: true
	FormEnabled *bool `json:"form_enabled"`

	// id
	// Required: true
	// Minimum: 1
	ID *uint32 `json:"id"`

	// image
	// Format: uri
	Image strfmt.URI `json:"image,omitempty"`

	// The import process status.
	// Enum: [succeeded failed importing]
	ImportStatus string `json:"import_status,omitempty"`

	// keywords
	// Required: true
	Keywords []string `json:"keywords"`

	// name
	// Required: true
	Name *string `json:"name"`

	// Readme of the stack
	Readme string `json:"readme,omitempty"`

	// ref
	// Required: true
	Ref *string `json:"ref"`

	// service catalog source canonical
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	ServiceCatalogSourceCanonical string `json:"service_catalog_source_canonical,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// technologies
	Technologies []*ServiceCatalogTechnology `json:"technologies"`

	// If 'true' the ServiceCatalog is from the main organization and can be trusted.
	// Required: true
	Trusted *bool `json:"trusted"`

	// updated at
	// Minimum: 0
	UpdatedAt *uint64 `json:"updated_at,omitempty"`
}

// Validate validates this service catalog
func (m *ServiceCatalog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDependencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImportStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeywords(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceCatalogSourceCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTechnologies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrusted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceCatalog) validateAuthor(formats strfmt.Registry) error {

	if err := validate.Required("author", "body", m.Author); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalog) validateCanonical(formats strfmt.Registry) error {

	if err := validate.Required("canonical", "body", m.Canonical); err != nil {
		return err
	}

	if err := validate.MinLength("canonical", "body", string(*m.Canonical), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("canonical", "body", string(*m.Canonical), 100); err != nil {
		return err
	}

	if err := validate.Pattern("canonical", "body", string(*m.Canonical), `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalog) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.MinimumInt("created_at", "body", int64(*m.CreatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalog) validateDependencies(formats strfmt.Registry) error {

	if swag.IsZero(m.Dependencies) { // not required
		return nil
	}

	for i := 0; i < len(m.Dependencies); i++ {
		if swag.IsZero(m.Dependencies[i]) { // not required
			continue
		}

		if m.Dependencies[i] != nil {
			if err := m.Dependencies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dependencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceCatalog) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalog) validateDirectory(formats strfmt.Registry) error {

	if err := validate.Required("directory", "body", m.Directory); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalog) validateFormEnabled(formats strfmt.Registry) error {

	if err := validate.Required("form_enabled", "body", m.FormEnabled); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalog) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinimumInt("id", "body", int64(*m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalog) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if err := validate.FormatOf("image", "body", "uri", m.Image.String(), formats); err != nil {
		return err
	}

	return nil
}

var serviceCatalogTypeImportStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["succeeded","failed","importing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceCatalogTypeImportStatusPropEnum = append(serviceCatalogTypeImportStatusPropEnum, v)
	}
}

const (

	// ServiceCatalogImportStatusSucceeded captures enum value "succeeded"
	ServiceCatalogImportStatusSucceeded string = "succeeded"

	// ServiceCatalogImportStatusFailed captures enum value "failed"
	ServiceCatalogImportStatusFailed string = "failed"

	// ServiceCatalogImportStatusImporting captures enum value "importing"
	ServiceCatalogImportStatusImporting string = "importing"
)

// prop value enum
func (m *ServiceCatalog) validateImportStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serviceCatalogTypeImportStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceCatalog) validateImportStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ImportStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateImportStatusEnum("import_status", "body", m.ImportStatus); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalog) validateKeywords(formats strfmt.Registry) error {

	if err := validate.Required("keywords", "body", m.Keywords); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalog) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalog) validateRef(formats strfmt.Registry) error {

	if err := validate.Required("ref", "body", m.Ref); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalog) validateServiceCatalogSourceCanonical(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceCatalogSourceCanonical) { // not required
		return nil
	}

	if err := validate.MinLength("service_catalog_source_canonical", "body", string(m.ServiceCatalogSourceCanonical), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("service_catalog_source_canonical", "body", string(m.ServiceCatalogSourceCanonical), 100); err != nil {
		return err
	}

	if err := validate.Pattern("service_catalog_source_canonical", "body", string(m.ServiceCatalogSourceCanonical), `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalog) validateTechnologies(formats strfmt.Registry) error {

	if swag.IsZero(m.Technologies) { // not required
		return nil
	}

	for i := 0; i < len(m.Technologies); i++ {
		if swag.IsZero(m.Technologies[i]) { // not required
			continue
		}

		if m.Technologies[i] != nil {
			if err := m.Technologies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("technologies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceCatalog) validateTrusted(formats strfmt.Registry) error {

	if err := validate.Required("trusted", "body", m.Trusted); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalog) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.MinimumInt("updated_at", "body", int64(*m.UpdatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceCatalog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCatalog) UnmarshalBinary(b []byte) error {
	var res ServiceCatalog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
