// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Pipeline Pipeline
//
// The entity which represents a pipeline in the application.
//
// swagger:model Pipeline
type Pipeline struct {

	// archived
	Archived bool `json:"archived,omitempty"`

	// created at
	// Required: true
	// Minimum: 0
	CreatedAt *uint64 `json:"created_at"`

	// environment
	// Required: true
	// Pattern: ^[\da-zA-Z]+(?:(?:[\da-zA-Z\-._]+)?[\da-zA-Z])?$
	Environment *string `json:"environment"`

	// groups
	Groups []*GroupConfig `json:"groups"`

	// id
	// Required: true
	ID *uint64 `json:"id"`

	// jobs
	// Required: true
	Jobs []*Job `json:"jobs"`

	// name
	// Required: true
	Name *string `json:"name"`

	// paused
	// Required: true
	Paused *bool `json:"paused"`

	// project
	// Required: true
	Project *Project `json:"project"`

	// public
	// Required: true
	Public *bool `json:"public"`

	// started
	Started bool `json:"started,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// team name
	// Required: true
	TeamName *string `json:"team_name"`

	// updated at
	// Required: true
	// Minimum: 0
	UpdatedAt *uint64 `json:"updated_at"`

	// use case
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: (^[a-z0-9]+(([a-z0-9\-_]+)?[a-z0-9]+)?$)
	UseCase *string `json:"use_case"`
}

// Validate validates this pipeline
func (m *Pipeline) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaused(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseCase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Pipeline) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.MinimumInt("created_at", "body", int64(*m.CreatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Pipeline) validateEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("environment", "body", m.Environment); err != nil {
		return err
	}

	if err := validate.Pattern("environment", "body", string(*m.Environment), `^[\da-zA-Z]+(?:(?:[\da-zA-Z\-._]+)?[\da-zA-Z])?$`); err != nil {
		return err
	}

	return nil
}

func (m *Pipeline) validateGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.Groups) { // not required
		return nil
	}

	for i := 0; i < len(m.Groups); i++ {
		if swag.IsZero(m.Groups[i]) { // not required
			continue
		}

		if m.Groups[i] != nil {
			if err := m.Groups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Pipeline) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Pipeline) validateJobs(formats strfmt.Registry) error {

	if err := validate.Required("jobs", "body", m.Jobs); err != nil {
		return err
	}

	for i := 0; i < len(m.Jobs); i++ {
		if swag.IsZero(m.Jobs[i]) { // not required
			continue
		}

		if m.Jobs[i] != nil {
			if err := m.Jobs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("jobs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Pipeline) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Pipeline) validatePaused(formats strfmt.Registry) error {

	if err := validate.Required("paused", "body", m.Paused); err != nil {
		return err
	}

	return nil
}

func (m *Pipeline) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("project", "body", m.Project); err != nil {
		return err
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

func (m *Pipeline) validatePublic(formats strfmt.Registry) error {

	if err := validate.Required("public", "body", m.Public); err != nil {
		return err
	}

	return nil
}

func (m *Pipeline) validateTeamName(formats strfmt.Registry) error {

	if err := validate.Required("team_name", "body", m.TeamName); err != nil {
		return err
	}

	return nil
}

func (m *Pipeline) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.MinimumInt("updated_at", "body", int64(*m.UpdatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Pipeline) validateUseCase(formats strfmt.Registry) error {

	if err := validate.Required("use_case", "body", m.UseCase); err != nil {
		return err
	}

	if err := validate.MinLength("use_case", "body", string(*m.UseCase), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("use_case", "body", string(*m.UseCase), 100); err != nil {
		return err
	}

	if err := validate.Pattern("use_case", "body", string(*m.UseCase), `(^[a-z0-9]+(([a-z0-9\-_]+)?[a-z0-9]+)?$)`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Pipeline) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Pipeline) UnmarshalBinary(b []byte) error {
	var res Pipeline
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
