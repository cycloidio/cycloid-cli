// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExternalBackend External backend
//
// An external backend contains the configuration needed in order to be plugged into the Cycloid system. A backend is a general purpose concept, but Cycloid specifies which ones are supported and the list of those which are supported for every concrete feature.
//
// swagger:model ExternalBackend
type ExternalBackend struct {
	configurationField ExternalBackendConfiguration

	// created at
	// Minimum: 0
	CreatedAt *uint64 `json:"created_at,omitempty"`

	// credential canonical
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	CredentialCanonical string `json:"credential_canonical,omitempty"`

	// Will mark this EB as default for the specific purpose
	// Required: true
	Default *bool `json:"default"`

	// environment canonical
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[\da-zA-Z]+(?:(?:[\da-zA-Z\-._]+)?[\da-zA-Z])?$
	EnvironmentCanonical string `json:"environment_canonical,omitempty"`

	// id
	// Minimum: 1
	ID uint32 `json:"id,omitempty"`

	// JWT is a credential identifying this EB for a public interaction right now it's only filled when the Purpose == RemoteTFState as we'll use it for the Inventory.
	Jwt string `json:"jwt,omitempty"`

	// project canonical
	// Max Length: 100
	// Min Length: 1
	// Pattern: (^[a-z0-9]+(([a-z0-9\-_]+)?[a-z0-9]+)?$)
	ProjectCanonical string `json:"project_canonical,omitempty"`

	// purpose
	// Required: true
	Purpose *string `json:"purpose"`

	// updated at
	// Minimum: 0
	UpdatedAt *uint64 `json:"updated_at,omitempty"`
}

// Configuration gets the configuration of this base type
func (m *ExternalBackend) Configuration() ExternalBackendConfiguration {
	return m.configurationField
}

// SetConfiguration sets the configuration of this base type
func (m *ExternalBackend) SetConfiguration(val ExternalBackendConfiguration) {
	m.configurationField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *ExternalBackend) UnmarshalJSON(raw []byte) error {
	var data struct {
		Configuration json.RawMessage `json:"configuration"`

		CreatedAt *uint64 `json:"created_at,omitempty"`

		CredentialCanonical string `json:"credential_canonical,omitempty"`

		Default *bool `json:"default"`

		EnvironmentCanonical string `json:"environment_canonical,omitempty"`

		ID uint32 `json:"id,omitempty"`

		Jwt string `json:"jwt,omitempty"`

		ProjectCanonical string `json:"project_canonical,omitempty"`

		Purpose *string `json:"purpose"`

		UpdatedAt *uint64 `json:"updated_at,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	propConfiguration, err := UnmarshalExternalBackendConfiguration(bytes.NewBuffer(data.Configuration), runtime.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	var result ExternalBackend

	// configuration
	result.configurationField = propConfiguration

	// created_at
	result.CreatedAt = data.CreatedAt

	// credential_canonical
	result.CredentialCanonical = data.CredentialCanonical

	// default
	result.Default = data.Default

	// environment_canonical
	result.EnvironmentCanonical = data.EnvironmentCanonical

	// id
	result.ID = data.ID

	// jwt
	result.Jwt = data.Jwt

	// project_canonical
	result.ProjectCanonical = data.ProjectCanonical

	// purpose
	result.Purpose = data.Purpose

	// updated_at
	result.UpdatedAt = data.UpdatedAt

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m ExternalBackend) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		CreatedAt *uint64 `json:"created_at,omitempty"`

		CredentialCanonical string `json:"credential_canonical,omitempty"`

		Default *bool `json:"default"`

		EnvironmentCanonical string `json:"environment_canonical,omitempty"`

		ID uint32 `json:"id,omitempty"`

		Jwt string `json:"jwt,omitempty"`

		ProjectCanonical string `json:"project_canonical,omitempty"`

		Purpose *string `json:"purpose"`

		UpdatedAt *uint64 `json:"updated_at,omitempty"`
	}{

		CreatedAt: m.CreatedAt,

		CredentialCanonical: m.CredentialCanonical,

		Default: m.Default,

		EnvironmentCanonical: m.EnvironmentCanonical,

		ID: m.ID,

		Jwt: m.Jwt,

		ProjectCanonical: m.ProjectCanonical,

		Purpose: m.Purpose,

		UpdatedAt: m.UpdatedAt,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Configuration ExternalBackendConfiguration `json:"configuration"`
	}{

		Configuration: m.configurationField,
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this external backend
func (m *ExternalBackend) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefault(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurpose(formats); err != nil {
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

func (m *ExternalBackend) validateConfiguration(formats strfmt.Registry) error {

	if err := validate.Required("configuration", "body", m.Configuration()); err != nil {
		return err
	}

	if err := m.Configuration().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("configuration")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("configuration")
		}
		return err
	}

	return nil
}

func (m *ExternalBackend) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.MinimumUint("created_at", "body", *m.CreatedAt, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ExternalBackend) validateCredentialCanonical(formats strfmt.Registry) error {
	if swag.IsZero(m.CredentialCanonical) { // not required
		return nil
	}

	if err := validate.MinLength("credential_canonical", "body", m.CredentialCanonical, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("credential_canonical", "body", m.CredentialCanonical, 100); err != nil {
		return err
	}

	if err := validate.Pattern("credential_canonical", "body", m.CredentialCanonical, `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *ExternalBackend) validateDefault(formats strfmt.Registry) error {

	if err := validate.Required("default", "body", m.Default); err != nil {
		return err
	}

	return nil
}

func (m *ExternalBackend) validateEnvironmentCanonical(formats strfmt.Registry) error {
	if swag.IsZero(m.EnvironmentCanonical) { // not required
		return nil
	}

	if err := validate.MinLength("environment_canonical", "body", m.EnvironmentCanonical, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("environment_canonical", "body", m.EnvironmentCanonical, 100); err != nil {
		return err
	}

	if err := validate.Pattern("environment_canonical", "body", m.EnvironmentCanonical, `^[\da-zA-Z]+(?:(?:[\da-zA-Z\-._]+)?[\da-zA-Z])?$`); err != nil {
		return err
	}

	return nil
}

func (m *ExternalBackend) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinimumUint("id", "body", uint64(m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *ExternalBackend) validateProjectCanonical(formats strfmt.Registry) error {
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

func (m *ExternalBackend) validatePurpose(formats strfmt.Registry) error {

	if err := validate.Required("purpose", "body", m.Purpose); err != nil {
		return err
	}

	return nil
}

func (m *ExternalBackend) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.MinimumUint("updated_at", "body", *m.UpdatedAt, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this external backend based on the context it is used
func (m *ExternalBackend) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalBackend) contextValidateConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Configuration().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("configuration")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("configuration")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExternalBackend) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalBackend) UnmarshalBinary(b []byte) error {
	var res ExternalBackend
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
