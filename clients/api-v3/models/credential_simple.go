// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CredentialSimple Credential Simple
//
// Represents the Credential without the raw and owner
// swagger:model CredentialSimple
type CredentialSimple struct {

	// description
	Description string `json:"description,omitempty"`

	// id
	// Required: true
	// Minimum: 1
	ID *uint32 `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// path
	// Required: true
	Path *string `json:"path"`

	// type
	// Required: true
	// Enum: [ssh aws custom azure azure_storage gcp basic_auth elasticsearch swift]
	Type *string `json:"type"`
}

// Validate validates this credential simple
func (m *CredentialSimple) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
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

func (m *CredentialSimple) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinimumInt("id", "body", int64(*m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *CredentialSimple) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CredentialSimple) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

var credentialSimpleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ssh","aws","custom","azure","azure_storage","gcp","basic_auth","elasticsearch","swift"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		credentialSimpleTypeTypePropEnum = append(credentialSimpleTypeTypePropEnum, v)
	}
}

const (

	// CredentialSimpleTypeSSH captures enum value "ssh"
	CredentialSimpleTypeSSH string = "ssh"

	// CredentialSimpleTypeAws captures enum value "aws"
	CredentialSimpleTypeAws string = "aws"

	// CredentialSimpleTypeCustom captures enum value "custom"
	CredentialSimpleTypeCustom string = "custom"

	// CredentialSimpleTypeAzure captures enum value "azure"
	CredentialSimpleTypeAzure string = "azure"

	// CredentialSimpleTypeAzureStorage captures enum value "azure_storage"
	CredentialSimpleTypeAzureStorage string = "azure_storage"

	// CredentialSimpleTypeGcp captures enum value "gcp"
	CredentialSimpleTypeGcp string = "gcp"

	// CredentialSimpleTypeBasicAuth captures enum value "basic_auth"
	CredentialSimpleTypeBasicAuth string = "basic_auth"

	// CredentialSimpleTypeElasticsearch captures enum value "elasticsearch"
	CredentialSimpleTypeElasticsearch string = "elasticsearch"

	// CredentialSimpleTypeSwift captures enum value "swift"
	CredentialSimpleTypeSwift string = "swift"
)

// prop value enum
func (m *CredentialSimple) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, credentialSimpleTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CredentialSimple) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CredentialSimple) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CredentialSimple) UnmarshalBinary(b []byte) error {
	var res CredentialSimple
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
