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

// UpdateCredential Credential
//
// Represents the Credential
// swagger:model UpdateCredential
type UpdateCredential struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// name
	// Required: true
	Name *string `json:"name"`

	// User canonical that owns this credential. When a user is the owner of a credential he has
	// all the permissions on it.
	//
	// Required: true
	Owner *string `json:"owner"`

	// path
	// Required: true
	// Pattern: [a-zA-z0-9_\-./]
	Path *string `json:"path"`

	// raw
	// Required: true
	Raw *CredentialRaw `json:"raw"`

	// type
	// Required: true
	// Enum: [ssh aws custom azure azure_storage gcp basic_auth elasticsearch]
	Type *string `json:"type"`
}

// Validate validates this update credential
func (m *UpdateCredential) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateCredential) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *UpdateCredential) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *UpdateCredential) validateOwner(formats strfmt.Registry) error {

	if err := validate.Required("owner", "body", m.Owner); err != nil {
		return err
	}

	return nil
}

func (m *UpdateCredential) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	if err := validate.Pattern("path", "body", string(*m.Path), `[a-zA-z0-9_\-./]`); err != nil {
		return err
	}

	return nil
}

func (m *UpdateCredential) validateRaw(formats strfmt.Registry) error {

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

var updateCredentialTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ssh","aws","custom","azure","azure_storage","gcp","basic_auth","elasticsearch"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateCredentialTypeTypePropEnum = append(updateCredentialTypeTypePropEnum, v)
	}
}

const (

	// UpdateCredentialTypeSSH captures enum value "ssh"
	UpdateCredentialTypeSSH string = "ssh"

	// UpdateCredentialTypeAws captures enum value "aws"
	UpdateCredentialTypeAws string = "aws"

	// UpdateCredentialTypeCustom captures enum value "custom"
	UpdateCredentialTypeCustom string = "custom"

	// UpdateCredentialTypeAzure captures enum value "azure"
	UpdateCredentialTypeAzure string = "azure"

	// UpdateCredentialTypeAzureStorage captures enum value "azure_storage"
	UpdateCredentialTypeAzureStorage string = "azure_storage"

	// UpdateCredentialTypeGcp captures enum value "gcp"
	UpdateCredentialTypeGcp string = "gcp"

	// UpdateCredentialTypeBasicAuth captures enum value "basic_auth"
	UpdateCredentialTypeBasicAuth string = "basic_auth"

	// UpdateCredentialTypeElasticsearch captures enum value "elasticsearch"
	UpdateCredentialTypeElasticsearch string = "elasticsearch"
)

// prop value enum
func (m *UpdateCredential) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, updateCredentialTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UpdateCredential) validateType(formats strfmt.Registry) error {

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
func (m *UpdateCredential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateCredential) UnmarshalBinary(b []byte) error {
	var res UpdateCredential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}