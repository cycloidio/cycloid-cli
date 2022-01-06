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

// Credential Credential
//
// Represents the Credential
// swagger:model Credential
type Credential struct {

	// canonical
	// Required: true
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	Canonical *string `json:"canonical"`

	// created at
	// Minimum: 0
	CreatedAt *uint64 `json:"created_at,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	// Required: true
	// Minimum: 1
	ID *uint32 `json:"id"`

	// in use
	InUse *CredentialInUse `json:"in_use,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// Organization member that owns this credential. When a user is the owner of a
	// credential he has all the permissions on it.
	// In the event where the user has been deleted that field might be empty.
	//
	Owner *User `json:"owner,omitempty"`

	// path
	// Required: true
	Path *string `json:"path"`

	// raw
	// Required: true
	Raw *CredentialRaw `json:"raw"`

	// type
	// Required: true
	// Enum: [ssh aws custom azure azure_storage gcp basic_auth elasticsearch swift]
	Type *string `json:"type"`

	// updated at
	// Minimum: 0
	UpdatedAt *uint64 `json:"updated_at,omitempty"`
}

// Validate validates this credential
func (m *Credential) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInUse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRaw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
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

func (m *Credential) validateCanonical(formats strfmt.Registry) error {

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

func (m *Credential) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.MinimumInt("created_at", "body", int64(*m.CreatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Credential) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinimumInt("id", "body", int64(*m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Credential) validateInUse(formats strfmt.Registry) error {

	if swag.IsZero(m.InUse) { // not required
		return nil
	}

	if m.InUse != nil {
		if err := m.InUse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("in_use")
			}
			return err
		}
	}

	return nil
}

func (m *Credential) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Credential) validateOwner(formats strfmt.Registry) error {

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

func (m *Credential) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *Credential) validateRaw(formats strfmt.Registry) error {

	if err := validate.Required("raw", "body", m.Raw); err != nil {
		return err
	}

	if m.Raw != nil {
		if err := m.Raw.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("raw")
			}
			return err
		}
	}

	return nil
}

var credentialTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ssh","aws","custom","azure","azure_storage","gcp","basic_auth","elasticsearch","swift"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		credentialTypeTypePropEnum = append(credentialTypeTypePropEnum, v)
	}
}

const (

	// CredentialTypeSSH captures enum value "ssh"
	CredentialTypeSSH string = "ssh"

	// CredentialTypeAws captures enum value "aws"
	CredentialTypeAws string = "aws"

	// CredentialTypeCustom captures enum value "custom"
	CredentialTypeCustom string = "custom"

	// CredentialTypeAzure captures enum value "azure"
	CredentialTypeAzure string = "azure"

	// CredentialTypeAzureStorage captures enum value "azure_storage"
	CredentialTypeAzureStorage string = "azure_storage"

	// CredentialTypeGcp captures enum value "gcp"
	CredentialTypeGcp string = "gcp"

	// CredentialTypeBasicAuth captures enum value "basic_auth"
	CredentialTypeBasicAuth string = "basic_auth"

	// CredentialTypeElasticsearch captures enum value "elasticsearch"
	CredentialTypeElasticsearch string = "elasticsearch"

	// CredentialTypeSwift captures enum value "swift"
	CredentialTypeSwift string = "swift"
)

// prop value enum
func (m *Credential) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, credentialTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Credential) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Credential) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.MinimumInt("updated_at", "body", int64(*m.UpdatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Credential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Credential) UnmarshalBinary(b []byte) error {
	var res Credential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CredentialInUse InUse represents the resources that are using provided credential.
//
// swagger:model CredentialInUse
type CredentialInUse struct {

	// config repositories
	ConfigRepositories []*InUseConfigRepository `json:"config_repositories"`

	// external backends
	ExternalBackends []*InUseExternalBackend `json:"external_backends"`

	// service catalog sources
	ServiceCatalogSources []*InUseServiceCatalogSource `json:"service_catalog_sources"`
}

// Validate validates this credential in use
func (m *CredentialInUse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigRepositories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalBackends(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceCatalogSources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CredentialInUse) validateConfigRepositories(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigRepositories) { // not required
		return nil
	}

	for i := 0; i < len(m.ConfigRepositories); i++ {
		if swag.IsZero(m.ConfigRepositories[i]) { // not required
			continue
		}

		if m.ConfigRepositories[i] != nil {
			if err := m.ConfigRepositories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("in_use" + "." + "config_repositories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CredentialInUse) validateExternalBackends(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalBackends) { // not required
		return nil
	}

	for i := 0; i < len(m.ExternalBackends); i++ {
		if swag.IsZero(m.ExternalBackends[i]) { // not required
			continue
		}

		if m.ExternalBackends[i] != nil {
			if err := m.ExternalBackends[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("in_use" + "." + "external_backends" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CredentialInUse) validateServiceCatalogSources(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceCatalogSources) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceCatalogSources); i++ {
		if swag.IsZero(m.ServiceCatalogSources[i]) { // not required
			continue
		}

		if m.ServiceCatalogSources[i] != nil {
			if err := m.ServiceCatalogSources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("in_use" + "." + "service_catalog_sources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CredentialInUse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CredentialInUse) UnmarshalBinary(b []byte) error {
	var res CredentialInUse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
