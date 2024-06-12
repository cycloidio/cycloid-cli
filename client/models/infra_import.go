// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InfraImport Infra Import
//
// The entity which represents the information of for the import of a new Stack or Project.
//
// swagger:model InfraImport
type InfraImport struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// logs
	// Required: true
	Logs *string `json:"logs"`

	// project canonical
	// Max Length: 100
	// Min Length: 1
	// Pattern: (^[a-z0-9]+(([a-z0-9\-_]+)?[a-z0-9]+)?$)
	ProjectCanonical string `json:"project_canonical,omitempty"`

	// It's the ref of the Service Catalog, like 'cycloidio:stack-magento'
	// Required: true
	ServiceCatalogRef *string `json:"service_catalog_ref"`

	// The import process status.
	// Required: true
	// Enum: ["succeeded","failed","importing"]
	Status *string `json:"status"`
}

// Validate validates this infra import
func (m *InfraImport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceCatalogRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InfraImport) validateError(formats strfmt.Registry) error {

	if err := validate.Required("error", "body", m.Error); err != nil {
		return err
	}

	return nil
}

func (m *InfraImport) validateLogs(formats strfmt.Registry) error {

	if err := validate.Required("logs", "body", m.Logs); err != nil {
		return err
	}

	return nil
}

func (m *InfraImport) validateProjectCanonical(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectCanonical) { // not required
		return nil
	}

	if err := validate.MinLength("project_canonical", "body", m.ProjectCanonical, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("project_canonical", "body", m.ProjectCanonical, 100); err != nil {
		return err
	}

	if err := validate.Pattern("project_canonical", "body", m.ProjectCanonical, `(^[a-z0-9]+(([a-z0-9\-_]+)?[a-z0-9]+)?$)`); err != nil {
		return err
	}

	return nil
}

func (m *InfraImport) validateServiceCatalogRef(formats strfmt.Registry) error {

	if err := validate.Required("service_catalog_ref", "body", m.ServiceCatalogRef); err != nil {
		return err
	}

	return nil
}

var infraImportTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["succeeded","failed","importing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		infraImportTypeStatusPropEnum = append(infraImportTypeStatusPropEnum, v)
	}
}

const (

	// InfraImportStatusSucceeded captures enum value "succeeded"
	InfraImportStatusSucceeded string = "succeeded"

	// InfraImportStatusFailed captures enum value "failed"
	InfraImportStatusFailed string = "failed"

	// InfraImportStatusImporting captures enum value "importing"
	InfraImportStatusImporting string = "importing"
)

// prop value enum
func (m *InfraImport) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, infraImportTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InfraImport) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this infra import based on context it is used
func (m *InfraImport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InfraImport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfraImport) UnmarshalBinary(b []byte) error {
	var res InfraImport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
