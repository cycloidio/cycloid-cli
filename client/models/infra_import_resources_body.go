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

// InfraImportResourcesBody Provider's Resources body
//
// # Entry that represents all the data needed for fetching resources
//
// swagger:model InfraImportResourcesBody
type InfraImportResourcesBody struct {
	configurationField CloudProviderConfiguration

	// Credential that will be used to import from the provider
	// Required: true
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	CredentialCanonical *string `json:"credential_canonical"`
}

// Configuration gets the configuration of this base type
func (m *InfraImportResourcesBody) Configuration() CloudProviderConfiguration {
	return m.configurationField
}

// SetConfiguration sets the configuration of this base type
func (m *InfraImportResourcesBody) SetConfiguration(val CloudProviderConfiguration) {
	m.configurationField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *InfraImportResourcesBody) UnmarshalJSON(raw []byte) error {
	var data struct {
		Configuration json.RawMessage `json:"configuration"`

		CredentialCanonical *string `json:"credential_canonical"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	propConfiguration, err := UnmarshalCloudProviderConfiguration(bytes.NewBuffer(data.Configuration), runtime.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	var result InfraImportResourcesBody

	// configuration
	result.configurationField = propConfiguration

	// credential_canonical
	result.CredentialCanonical = data.CredentialCanonical

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m InfraImportResourcesBody) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		CredentialCanonical *string `json:"credential_canonical"`
	}{

		CredentialCanonical: m.CredentialCanonical,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Configuration CloudProviderConfiguration `json:"configuration"`
	}{

		Configuration: m.configurationField,
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this infra import resources body
func (m *InfraImportResourcesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialCanonical(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InfraImportResourcesBody) validateConfiguration(formats strfmt.Registry) error {

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

func (m *InfraImportResourcesBody) validateCredentialCanonical(formats strfmt.Registry) error {

	if err := validate.Required("credential_canonical", "body", m.CredentialCanonical); err != nil {
		return err
	}

	if err := validate.MinLength("credential_canonical", "body", *m.CredentialCanonical, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("credential_canonical", "body", *m.CredentialCanonical, 100); err != nil {
		return err
	}

	if err := validate.Pattern("credential_canonical", "body", *m.CredentialCanonical, `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this infra import resources body based on the context it is used
func (m *InfraImportResourcesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InfraImportResourcesBody) contextValidateConfiguration(ctx context.Context, formats strfmt.Registry) error {

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
func (m *InfraImportResourcesBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfraImportResourcesBody) UnmarshalBinary(b []byte) error {
	var res InfraImportResourcesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
