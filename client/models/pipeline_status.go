// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PipelineStatus PipelineStatus
//
// Pipeline status returned upon pipelines comparison between the one locally stored in the database and its counter part on git.
//
// swagger:model PipelineStatus
type PipelineStatus struct {

	// diffs
	Diffs *PipelineDiffs `json:"diffs,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// The state can vary depending on how far the comparison process went. There are 4 possible states which are:
	//   - unknown: one of the pipeline (database/git) couldn't be retrieved
	//   - sycned: both database & git pipelines are identical
	//   - out_of_sync: database & git pipelines have some differences
	//   - errored: both pipelines got retrieved but the comparison triggered an error
	// Required: true
	// Enum: ["unknown","synced","out_of_sync","errored"]
	Synced *string `json:"synced"`
}

// Validate validates this pipeline status
func (m *PipelineStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiffs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSynced(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PipelineStatus) validateDiffs(formats strfmt.Registry) error {
	if swag.IsZero(m.Diffs) { // not required
		return nil
	}

	if m.Diffs != nil {
		if err := m.Diffs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("diffs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("diffs")
			}
			return err
		}
	}

	return nil
}

var pipelineStatusTypeSyncedPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unknown","synced","out_of_sync","errored"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pipelineStatusTypeSyncedPropEnum = append(pipelineStatusTypeSyncedPropEnum, v)
	}
}

const (

	// PipelineStatusSyncedUnknown captures enum value "unknown"
	PipelineStatusSyncedUnknown string = "unknown"

	// PipelineStatusSyncedSynced captures enum value "synced"
	PipelineStatusSyncedSynced string = "synced"

	// PipelineStatusSyncedOutOfSync captures enum value "out_of_sync"
	PipelineStatusSyncedOutOfSync string = "out_of_sync"

	// PipelineStatusSyncedErrored captures enum value "errored"
	PipelineStatusSyncedErrored string = "errored"
)

// prop value enum
func (m *PipelineStatus) validateSyncedEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pipelineStatusTypeSyncedPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PipelineStatus) validateSynced(formats strfmt.Registry) error {

	if err := validate.Required("synced", "body", m.Synced); err != nil {
		return err
	}

	// value enum
	if err := m.validateSyncedEnum("synced", "body", *m.Synced); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this pipeline status based on the context it is used
func (m *PipelineStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDiffs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PipelineStatus) contextValidateDiffs(ctx context.Context, formats strfmt.Registry) error {

	if m.Diffs != nil {

		if swag.IsZero(m.Diffs) { // not required
			return nil
		}

		if err := m.Diffs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("diffs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("diffs")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PipelineStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelineStatus) UnmarshalBinary(b []byte) error {
	var res PipelineStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
