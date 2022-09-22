// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateExternalBackend Update External backend
//
// An external backend contains the configuration needed in order to be plugged into the Cycloid system. A backend is a general purpose concept, but Cycloid specifies which ones are supported and the list of those which are supported for every concrete feature.
// swagger:model UpdateExternalBackend
type UpdateExternalBackend struct {
	configurationField ExternalBackendConfiguration

	// The type of the credential must be one of: ["aws", "azure_storage", "elasticsearch", "gcp", "swift", "vmware"]
	//
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	CredentialCanonical string `json:"credential_canonical,omitempty"`

	// Will mark this EB as default for the specific purpose
	Default bool `json:"default,omitempty"`

	// environment canonical
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[\da-zA-Z]+(?:(?:[\da-zA-Z\-._]+)?[\da-zA-Z])?$
	EnvironmentCanonical string `json:"environment_canonical,omitempty"`

	// id
	// Minimum: 1
	ID uint32 `json:"id,omitempty"`

	// project canonical
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	ProjectCanonical string `json:"project_canonical,omitempty"`

	// purpose
	// Required: true
	// Enum: [events logs remote_tfstate cost_explorer]
	Purpose *string `json:"purpose"`
}

// Configuration gets the configuration of this base type
func (m *UpdateExternalBackend) Configuration() ExternalBackendConfiguration {
	return m.configurationField
}

// SetConfiguration sets the configuration of this base type
func (m *UpdateExternalBackend) SetConfiguration(val ExternalBackendConfiguration) {
	m.configurationField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *UpdateExternalBackend) UnmarshalJSON(raw []byte) error {
	var data struct {
		Configuration json.RawMessage `json:"configuration"`

		CredentialCanonical string `json:"credential_canonical,omitempty"`

		Default bool `json:"default,omitempty"`

		EnvironmentCanonical string `json:"environment_canonical,omitempty"`

		ID uint32 `json:"id,omitempty"`

		ProjectCanonical string `json:"project_canonical,omitempty"`

		Purpose *string `json:"purpose"`
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

	var result UpdateExternalBackend

	// configuration
	result.configurationField = propConfiguration

	// credential_canonical
	result.CredentialCanonical = data.CredentialCanonical

	// default
	result.Default = data.Default

	// environment_canonical
	result.EnvironmentCanonical = data.EnvironmentCanonical

	// id
	result.ID = data.ID

	// project_canonical
	result.ProjectCanonical = data.ProjectCanonical

	// purpose
	result.Purpose = data.Purpose

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m UpdateExternalBackend) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		CredentialCanonical string `json:"credential_canonical,omitempty"`

		Default bool `json:"default,omitempty"`

		EnvironmentCanonical string `json:"environment_canonical,omitempty"`

		ID uint32 `json:"id,omitempty"`

		ProjectCanonical string `json:"project_canonical,omitempty"`

		Purpose *string `json:"purpose"`
	}{

		CredentialCanonical: m.CredentialCanonical,

		Default: m.Default,

		EnvironmentCanonical: m.EnvironmentCanonical,

		ID: m.ID,

		ProjectCanonical: m.ProjectCanonical,

		Purpose: m.Purpose,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Configuration ExternalBackendConfiguration `json:"configuration"`
	}{

		Configuration: m.configurationField,
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this update external backend
func (m *UpdateExternalBackend) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialCanonical(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateExternalBackend) validateConfiguration(formats strfmt.Registry) error {

	if err := validate.Required("configuration", "body", m.Configuration()); err != nil {
		return err
	}

	if err := m.Configuration().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("configuration")
		}
		return err
	}

	return nil
}

func (m *UpdateExternalBackend) validateCredentialCanonical(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialCanonical) { // not required
		return nil
	}

	if err := validate.MinLength("credential_canonical", "body", string(m.CredentialCanonical), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("credential_canonical", "body", string(m.CredentialCanonical), 100); err != nil {
		return err
	}

	if err := validate.Pattern("credential_canonical", "body", string(m.CredentialCanonical), `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *UpdateExternalBackend) validateEnvironmentCanonical(formats strfmt.Registry) error {

	if swag.IsZero(m.EnvironmentCanonical) { // not required
		return nil
	}

	if err := validate.MinLength("environment_canonical", "body", string(m.EnvironmentCanonical), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("environment_canonical", "body", string(m.EnvironmentCanonical), 100); err != nil {
		return err
	}

	if err := validate.Pattern("environment_canonical", "body", string(m.EnvironmentCanonical), `^[\da-zA-Z]+(?:(?:[\da-zA-Z\-._]+)?[\da-zA-Z])?$`); err != nil {
		return err
	}

	return nil
}

func (m *UpdateExternalBackend) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinimumInt("id", "body", int64(m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *UpdateExternalBackend) validateProjectCanonical(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectCanonical) { // not required
		return nil
	}

	if err := validate.MinLength("project_canonical", "body", string(m.ProjectCanonical), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("project_canonical", "body", string(m.ProjectCanonical), 100); err != nil {
		return err
	}

	if err := validate.Pattern("project_canonical", "body", string(m.ProjectCanonical), `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

var updateExternalBackendTypePurposePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["events","logs","remote_tfstate","cost_explorer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateExternalBackendTypePurposePropEnum = append(updateExternalBackendTypePurposePropEnum, v)
	}
}

const (

	// UpdateExternalBackendPurposeEvents captures enum value "events"
	UpdateExternalBackendPurposeEvents string = "events"

	// UpdateExternalBackendPurposeLogs captures enum value "logs"
	UpdateExternalBackendPurposeLogs string = "logs"

	// UpdateExternalBackendPurposeRemoteTfstate captures enum value "remote_tfstate"
	UpdateExternalBackendPurposeRemoteTfstate string = "remote_tfstate"

	// UpdateExternalBackendPurposeCostExplorer captures enum value "cost_explorer"
	UpdateExternalBackendPurposeCostExplorer string = "cost_explorer"
)

// prop value enum
func (m *UpdateExternalBackend) validatePurposeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, updateExternalBackendTypePurposePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UpdateExternalBackend) validatePurpose(formats strfmt.Registry) error {

	if err := validate.Required("purpose", "body", m.Purpose); err != nil {
		return err
	}

	// value enum
	if err := m.validatePurposeEnum("purpose", "body", *m.Purpose); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateExternalBackend) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateExternalBackend) UnmarshalBinary(b []byte) error {
	var res UpdateExternalBackend
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
