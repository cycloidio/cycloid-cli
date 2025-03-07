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

// NewInfraImport New Infra Import
//
// # Entry that represents all the data needed to import a stack
//
// swagger:model NewInfraImport
type NewInfraImport struct {

	// component
	Component *NewInfraImportComponent `json:"component,omitempty"`

	configurationField CloudProviderConfiguration

	// Credential that will be used to import from the provider
	// Required: true
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	CredentialCanonical *string `json:"credential_canonical"`

	// environment
	Environment *NewEnvironment `json:"environment,omitempty"`

	// external backend
	ExternalBackend *NewInfraImportExternalBackend `json:"external_backend,omitempty"`

	// List of resources to import, these names are the ones on TF (ex: aws_instance). If not set then it means that all the resources will be imported
	Include []string `json:"include"`

	// It's a KV where the key is the resource name and the value is the list (array) of attributes to include as part of the module
	ModuleVariables map[string][]string `json:"module_variables,omitempty"`

	// project
	Project *NewInfraImportProject `json:"project,omitempty"`

	// stack
	// Required: true
	Stack *NewServiceCatalog `json:"stack"`

	// List of tags to filter with format NAME:VALUE
	Tags []string `json:"tags"`

	// List of resources to import via ID, those IDs are the ones documented on Terraform that are needed to Import. The format is 'aws_instance.ID'
	Targets []string `json:"targets"`
}

// Configuration gets the configuration of this base type
func (m *NewInfraImport) Configuration() CloudProviderConfiguration {
	return m.configurationField
}

// SetConfiguration sets the configuration of this base type
func (m *NewInfraImport) SetConfiguration(val CloudProviderConfiguration) {
	m.configurationField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *NewInfraImport) UnmarshalJSON(raw []byte) error {
	var data struct {
		Component *NewInfraImportComponent `json:"component,omitempty"`

		Configuration json.RawMessage `json:"configuration"`

		CredentialCanonical *string `json:"credential_canonical"`

		Environment *NewEnvironment `json:"environment,omitempty"`

		ExternalBackend *NewInfraImportExternalBackend `json:"external_backend,omitempty"`

		Include []string `json:"include"`

		ModuleVariables map[string][]string `json:"module_variables,omitempty"`

		Project *NewInfraImportProject `json:"project,omitempty"`

		Stack *NewServiceCatalog `json:"stack"`

		Tags []string `json:"tags"`

		Targets []string `json:"targets"`
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

	var result NewInfraImport

	// component
	result.Component = data.Component

	// configuration
	result.configurationField = propConfiguration

	// credential_canonical
	result.CredentialCanonical = data.CredentialCanonical

	// environment
	result.Environment = data.Environment

	// external_backend
	result.ExternalBackend = data.ExternalBackend

	// include
	result.Include = data.Include

	// module_variables
	result.ModuleVariables = data.ModuleVariables

	// project
	result.Project = data.Project

	// stack
	result.Stack = data.Stack

	// tags
	result.Tags = data.Tags

	// targets
	result.Targets = data.Targets

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m NewInfraImport) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		Component *NewInfraImportComponent `json:"component,omitempty"`

		CredentialCanonical *string `json:"credential_canonical"`

		Environment *NewEnvironment `json:"environment,omitempty"`

		ExternalBackend *NewInfraImportExternalBackend `json:"external_backend,omitempty"`

		Include []string `json:"include"`

		ModuleVariables map[string][]string `json:"module_variables,omitempty"`

		Project *NewInfraImportProject `json:"project,omitempty"`

		Stack *NewServiceCatalog `json:"stack"`

		Tags []string `json:"tags"`

		Targets []string `json:"targets"`
	}{

		Component: m.Component,

		CredentialCanonical: m.CredentialCanonical,

		Environment: m.Environment,

		ExternalBackend: m.ExternalBackend,

		Include: m.Include,

		ModuleVariables: m.ModuleVariables,

		Project: m.Project,

		Stack: m.Stack,

		Tags: m.Tags,

		Targets: m.Targets,
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

// Validate validates this new infra import
func (m *NewInfraImport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalBackend(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStack(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewInfraImport) validateComponent(formats strfmt.Registry) error {
	if swag.IsZero(m.Component) { // not required
		return nil
	}

	if m.Component != nil {
		if err := m.Component.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

func (m *NewInfraImport) validateConfiguration(formats strfmt.Registry) error {

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

func (m *NewInfraImport) validateCredentialCanonical(formats strfmt.Registry) error {

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

func (m *NewInfraImport) validateEnvironment(formats strfmt.Registry) error {
	if swag.IsZero(m.Environment) { // not required
		return nil
	}

	if m.Environment != nil {
		if err := m.Environment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("environment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("environment")
			}
			return err
		}
	}

	return nil
}

func (m *NewInfraImport) validateExternalBackend(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalBackend) { // not required
		return nil
	}

	if m.ExternalBackend != nil {
		if err := m.ExternalBackend.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external_backend")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("external_backend")
			}
			return err
		}
	}

	return nil
}

func (m *NewInfraImport) validateProject(formats strfmt.Registry) error {
	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

func (m *NewInfraImport) validateStack(formats strfmt.Registry) error {

	if err := validate.Required("stack", "body", m.Stack); err != nil {
		return err
	}

	if m.Stack != nil {
		if err := m.Stack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stack")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this new infra import based on the context it is used
func (m *NewInfraImport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComponent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnvironment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalBackend(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewInfraImport) contextValidateComponent(ctx context.Context, formats strfmt.Registry) error {

	if m.Component != nil {

		if swag.IsZero(m.Component) { // not required
			return nil
		}

		if err := m.Component.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

func (m *NewInfraImport) contextValidateConfiguration(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NewInfraImport) contextValidateEnvironment(ctx context.Context, formats strfmt.Registry) error {

	if m.Environment != nil {

		if swag.IsZero(m.Environment) { // not required
			return nil
		}

		if err := m.Environment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("environment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("environment")
			}
			return err
		}
	}

	return nil
}

func (m *NewInfraImport) contextValidateExternalBackend(ctx context.Context, formats strfmt.Registry) error {

	if m.ExternalBackend != nil {

		if swag.IsZero(m.ExternalBackend) { // not required
			return nil
		}

		if err := m.ExternalBackend.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external_backend")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("external_backend")
			}
			return err
		}
	}

	return nil
}

func (m *NewInfraImport) contextValidateProject(ctx context.Context, formats strfmt.Registry) error {

	if m.Project != nil {

		if swag.IsZero(m.Project) { // not required
			return nil
		}

		if err := m.Project.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

func (m *NewInfraImport) contextValidateStack(ctx context.Context, formats strfmt.Registry) error {

	if m.Stack != nil {

		if err := m.Stack.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stack")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewInfraImport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewInfraImport) UnmarshalBinary(b []byte) error {
	var res NewInfraImport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
