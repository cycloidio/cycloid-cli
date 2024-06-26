// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Job Job
//
// The entity which represents a job output in the application.
//
// swagger:model Job
type Job struct {

	// disable manual trigger
	DisableManualTrigger bool `json:"disable_manual_trigger,omitempty"`

	// finished build
	FinishedBuild *Build `json:"finished_build,omitempty"`

	// first logged build id
	FirstLoggedBuildID uint64 `json:"first_logged_build_id,omitempty"`

	// groups
	Groups []string `json:"groups"`

	// has new inputs
	HasNewInputs bool `json:"has_new_inputs,omitempty"`

	// id
	// Required: true
	ID *uint64 `json:"id"`

	// inputs
	Inputs []*JobInput `json:"inputs"`

	// name
	// Required: true
	Name *string `json:"name"`

	// next build
	NextBuild *Build `json:"next_build,omitempty"`

	// outputs
	Outputs []*JobOutput `json:"outputs"`

	// paused
	Paused bool `json:"paused,omitempty"`

	// pipeline id
	PipelineID uint64 `json:"pipeline_id,omitempty"`

	// pipeline name
	PipelineName string `json:"pipeline_name,omitempty"`

	// team name
	TeamName string `json:"team_name,omitempty"`

	// transition build
	TransitionBuild *Build `json:"transition_build,omitempty"`
}

// Validate validates this job
func (m *Job) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFinishedBuild(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInputs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextBuild(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutputs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransitionBuild(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Job) validateFinishedBuild(formats strfmt.Registry) error {
	if swag.IsZero(m.FinishedBuild) { // not required
		return nil
	}

	if m.FinishedBuild != nil {
		if err := m.FinishedBuild.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("finished_build")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("finished_build")
			}
			return err
		}
	}

	return nil
}

func (m *Job) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Job) validateInputs(formats strfmt.Registry) error {
	if swag.IsZero(m.Inputs) { // not required
		return nil
	}

	for i := 0; i < len(m.Inputs); i++ {
		if swag.IsZero(m.Inputs[i]) { // not required
			continue
		}

		if m.Inputs[i] != nil {
			if err := m.Inputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Job) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Job) validateNextBuild(formats strfmt.Registry) error {
	if swag.IsZero(m.NextBuild) { // not required
		return nil
	}

	if m.NextBuild != nil {
		if err := m.NextBuild.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next_build")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next_build")
			}
			return err
		}
	}

	return nil
}

func (m *Job) validateOutputs(formats strfmt.Registry) error {
	if swag.IsZero(m.Outputs) { // not required
		return nil
	}

	for i := 0; i < len(m.Outputs); i++ {
		if swag.IsZero(m.Outputs[i]) { // not required
			continue
		}

		if m.Outputs[i] != nil {
			if err := m.Outputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("outputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Job) validateTransitionBuild(formats strfmt.Registry) error {
	if swag.IsZero(m.TransitionBuild) { // not required
		return nil
	}

	if m.TransitionBuild != nil {
		if err := m.TransitionBuild.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transition_build")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("transition_build")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this job based on the context it is used
func (m *Job) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFinishedBuild(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInputs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNextBuild(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOutputs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransitionBuild(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Job) contextValidateFinishedBuild(ctx context.Context, formats strfmt.Registry) error {

	if m.FinishedBuild != nil {

		if swag.IsZero(m.FinishedBuild) { // not required
			return nil
		}

		if err := m.FinishedBuild.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("finished_build")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("finished_build")
			}
			return err
		}
	}

	return nil
}

func (m *Job) contextValidateInputs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Inputs); i++ {

		if m.Inputs[i] != nil {

			if swag.IsZero(m.Inputs[i]) { // not required
				return nil
			}

			if err := m.Inputs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Job) contextValidateNextBuild(ctx context.Context, formats strfmt.Registry) error {

	if m.NextBuild != nil {

		if swag.IsZero(m.NextBuild) { // not required
			return nil
		}

		if err := m.NextBuild.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next_build")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next_build")
			}
			return err
		}
	}

	return nil
}

func (m *Job) contextValidateOutputs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Outputs); i++ {

		if m.Outputs[i] != nil {

			if swag.IsZero(m.Outputs[i]) { // not required
				return nil
			}

			if err := m.Outputs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("outputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Job) contextValidateTransitionBuild(ctx context.Context, formats strfmt.Registry) error {

	if m.TransitionBuild != nil {

		if swag.IsZero(m.TransitionBuild) { // not required
			return nil
		}

		if err := m.TransitionBuild.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transition_build")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("transition_build")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Job) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Job) UnmarshalBinary(b []byte) error {
	var res Job
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
