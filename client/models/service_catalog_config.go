// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceCatalogConfig Use Case Configuration
//
// # Represents the Service Catalog Configuration for a given Use Case
//
// swagger:model ServiceCatalogConfig
type ServiceCatalogConfig struct {

	// ansible
	Ansible SCConfigTechConfig `json:"ansible,omitempty"`

	// cloud provider
	// Required: true
	CloudProvider *string `json:"cloud_provider"`

	// description
	// Required: true
	Description *string `json:"description"`

	// forms
	Forms *FormUseCase `json:"forms,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// pipeline
	// Required: true
	Pipeline *SCConfigPipelineConfig `json:"pipeline"`

	// terraform
	Terraform SCConfigTechConfig `json:"terraform,omitempty"`
}

// Validate validates this service catalog config
func (m *ServiceCatalogConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnsible(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePipeline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerraform(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceCatalogConfig) validateAnsible(formats strfmt.Registry) error {
	if swag.IsZero(m.Ansible) { // not required
		return nil
	}

	if m.Ansible != nil {
		if err := m.Ansible.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ansible")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ansible")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceCatalogConfig) validateCloudProvider(formats strfmt.Registry) error {

	if err := validate.Required("cloud_provider", "body", m.CloudProvider); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalogConfig) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalogConfig) validateForms(formats strfmt.Registry) error {
	if swag.IsZero(m.Forms) { // not required
		return nil
	}

	if m.Forms != nil {
		if err := m.Forms.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("forms")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("forms")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceCatalogConfig) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCatalogConfig) validatePipeline(formats strfmt.Registry) error {

	if err := validate.Required("pipeline", "body", m.Pipeline); err != nil {
		return err
	}

	if m.Pipeline != nil {
		if err := m.Pipeline.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pipeline")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pipeline")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceCatalogConfig) validateTerraform(formats strfmt.Registry) error {
	if swag.IsZero(m.Terraform) { // not required
		return nil
	}

	if m.Terraform != nil {
		if err := m.Terraform.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terraform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terraform")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this service catalog config based on the context it is used
func (m *ServiceCatalogConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAnsible(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateForms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePipeline(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTerraform(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceCatalogConfig) contextValidateAnsible(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Ansible) { // not required
		return nil
	}

	if err := m.Ansible.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ansible")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ansible")
		}
		return err
	}

	return nil
}

func (m *ServiceCatalogConfig) contextValidateForms(ctx context.Context, formats strfmt.Registry) error {

	if m.Forms != nil {

		if swag.IsZero(m.Forms) { // not required
			return nil
		}

		if err := m.Forms.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("forms")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("forms")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceCatalogConfig) contextValidatePipeline(ctx context.Context, formats strfmt.Registry) error {

	if m.Pipeline != nil {

		if err := m.Pipeline.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pipeline")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pipeline")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceCatalogConfig) contextValidateTerraform(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Terraform) { // not required
		return nil
	}

	if err := m.Terraform.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("terraform")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("terraform")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceCatalogConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCatalogConfig) UnmarshalBinary(b []byte) error {
	var res ServiceCatalogConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
