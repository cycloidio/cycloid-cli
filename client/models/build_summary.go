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

// BuildSummary BuildSummary
//
// The information relative to a build summary.
//
// swagger:model BuildSummary
type BuildSummary struct {

	// end time
	EndTime uint64 `json:"end_time,omitempty"`

	// id
	// Required: true
	ID *uint64 `json:"id"`

	// job name
	JobName string `json:"job_name,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// pipeline id
	PipelineID uint64 `json:"pipeline_id,omitempty"`

	// pipeline name
	PipelineName string `json:"pipeline_name,omitempty"`

	// plan
	Plan interface{} `json:"plan,omitempty"`

	// start time
	StartTime uint64 `json:"start_time,omitempty"`

	// status
	// Required: true
	Status *string `json:"status"`

	// team name
	// Required: true
	TeamName *string `json:"team_name"`
}

// Validate validates this build summary
func (m *BuildSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildSummary) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *BuildSummary) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *BuildSummary) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *BuildSummary) validateTeamName(formats strfmt.Registry) error {

	if err := validate.Required("team_name", "body", m.TeamName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this build summary based on context it is used
func (m *BuildSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BuildSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildSummary) UnmarshalBinary(b []byte) error {
	var res BuildSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
