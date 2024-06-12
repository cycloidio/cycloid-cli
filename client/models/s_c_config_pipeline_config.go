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

// SCConfigPipelineConfig Pipeline Configuration
//
// # Represents the Service Catalog Configuration for a Pipeline of a given Use Case
//
// swagger:model SCConfigPipelineConfig
type SCConfigPipelineConfig struct {

	// pipeline
	// Required: true
	Pipeline *SCConfigPathConfig `json:"pipeline"`

	// variables
	// Required: true
	Variables *SCConfigPathDestConfig `json:"variables"`
}

// Validate validates this s c config pipeline config
func (m *SCConfigPipelineConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePipeline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SCConfigPipelineConfig) validatePipeline(formats strfmt.Registry) error {

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

func (m *SCConfigPipelineConfig) validateVariables(formats strfmt.Registry) error {

	if err := validate.Required("variables", "body", m.Variables); err != nil {
		return err
	}

	if m.Variables != nil {
		if err := m.Variables.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("variables")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("variables")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s c config pipeline config based on the context it is used
func (m *SCConfigPipelineConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePipeline(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVariables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SCConfigPipelineConfig) contextValidatePipeline(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SCConfigPipelineConfig) contextValidateVariables(ctx context.Context, formats strfmt.Registry) error {

	if m.Variables != nil {

		if err := m.Variables.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("variables")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("variables")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SCConfigPipelineConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SCConfigPipelineConfig) UnmarshalBinary(b []byte) error {
	var res SCConfigPipelineConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
