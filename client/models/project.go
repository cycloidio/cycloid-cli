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

// Project Project
//
// The entity which represents the information of a project.
// swagger:model Project
type Project struct {

	// canonical
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: (^[a-z0-9]+(([a-z0-9\-_]+)?[a-z0-9]+)?$)
	Canonical *string `json:"canonical"`

	// config repository canonical
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	ConfigRepositoryCanonical string `json:"config_repository_canonical,omitempty"`

	// created at
	// Required: true
	// Minimum: 0
	CreatedAt *uint64 `json:"created_at"`

	// description
	Description string `json:"description,omitempty"`

	// environments
	// Required: true
	Environments []*Environment `json:"environments"`

	// favorite
	Favorite bool `json:"favorite,omitempty"`

	// id
	// Required: true
	// Minimum: 1
	ID *uint32 `json:"id"`

	// The import process status.
	// Enum: [succeeded failed importing]
	ImportStatus string `json:"import_status,omitempty"`

	// name
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// Organization member that owns this project. When a user is the owner of a
	// project it has all the permissions on it.
	// In the event where the user has been deleted that field might be empty.
	//
	Owner *User `json:"owner,omitempty"`

	// The Service Catalog that was used to create project.
	ServiceCatalog *ServiceCatalog `json:"service_catalog,omitempty"`

	// The Team that was used to create project.
	Team *SimpleTeam `json:"team,omitempty"`

	// updated at
	// Required: true
	// Minimum: 0
	UpdatedAt *uint64 `json:"updated_at"`
}

// Validate validates this project
func (m *Project) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigRepositoryCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImportStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceCatalog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeam(formats); err != nil {
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

func (m *Project) validateCanonical(formats strfmt.Registry) error {

	if err := validate.Required("canonical", "body", m.Canonical); err != nil {
		return err
	}

	if err := validate.MinLength("canonical", "body", string(*m.Canonical), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("canonical", "body", string(*m.Canonical), 100); err != nil {
		return err
	}

	if err := validate.Pattern("canonical", "body", string(*m.Canonical), `(^[a-z0-9]+(([a-z0-9\-_]+)?[a-z0-9]+)?$)`); err != nil {
		return err
	}

	return nil
}

func (m *Project) validateConfigRepositoryCanonical(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigRepositoryCanonical) { // not required
		return nil
	}

	if err := validate.MinLength("config_repository_canonical", "body", string(m.ConfigRepositoryCanonical), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("config_repository_canonical", "body", string(m.ConfigRepositoryCanonical), 100); err != nil {
		return err
	}

	if err := validate.Pattern("config_repository_canonical", "body", string(m.ConfigRepositoryCanonical), `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Project) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.MinimumInt("created_at", "body", int64(*m.CreatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Project) validateEnvironments(formats strfmt.Registry) error {

	if err := validate.Required("environments", "body", m.Environments); err != nil {
		return err
	}

	for i := 0; i < len(m.Environments); i++ {
		if swag.IsZero(m.Environments[i]) { // not required
			continue
		}

		if m.Environments[i] != nil {
			if err := m.Environments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("environments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Project) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinimumInt("id", "body", int64(*m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

var projectTypeImportStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["succeeded","failed","importing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		projectTypeImportStatusPropEnum = append(projectTypeImportStatusPropEnum, v)
	}
}

const (

	// ProjectImportStatusSucceeded captures enum value "succeeded"
	ProjectImportStatusSucceeded string = "succeeded"

	// ProjectImportStatusFailed captures enum value "failed"
	ProjectImportStatusFailed string = "failed"

	// ProjectImportStatusImporting captures enum value "importing"
	ProjectImportStatusImporting string = "importing"
)

// prop value enum
func (m *Project) validateImportStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, projectTypeImportStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Project) validateImportStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ImportStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateImportStatusEnum("import_status", "body", m.ImportStatus); err != nil {
		return err
	}

	return nil
}

func (m *Project) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *Project) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *Project) validateServiceCatalog(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceCatalog) { // not required
		return nil
	}

	if m.ServiceCatalog != nil {
		if err := m.ServiceCatalog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_catalog")
			}
			return err
		}
	}

	return nil
}

func (m *Project) validateTeam(formats strfmt.Registry) error {

	if swag.IsZero(m.Team) { // not required
		return nil
	}

	if m.Team != nil {
		if err := m.Team.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team")
			}
			return err
		}
	}

	return nil
}

func (m *Project) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.MinimumInt("updated_at", "body", int64(*m.UpdatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Project) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Project) UnmarshalBinary(b []byte) error {
	var res Project
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
