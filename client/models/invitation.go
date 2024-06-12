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

// Invitation Invitation
//
// It represents an Invitation to join an Organization.
//
// swagger:model Invitation
type Invitation struct {

	// created at
	// Required: true
	// Minimum: 0
	CreatedAt *uint64 `json:"created_at"`

	// email
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// id
	// Required: true
	// Minimum: 1
	ID *uint32 `json:"id"`

	// invited by
	InvitedBy *MemberOrg `json:"invited_by,omitempty"`

	// invitee
	Invitee *UserAccount `json:"invitee,omitempty"`

	// resent at
	// Minimum: 0
	ResentAt *uint64 `json:"resent_at,omitempty"`

	// role
	// Required: true
	Role *Role `json:"role"`

	// state
	// Required: true
	// Enum: ["pending","accepted","declined"]
	State *string `json:"state"`
}

// Validate validates this invitation
func (m *Invitation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvitedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvitee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResentAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Invitation) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.MinimumUint("created_at", "body", *m.CreatedAt, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Invitation) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Invitation) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinimumUint("id", "body", uint64(*m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Invitation) validateInvitedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.InvitedBy) { // not required
		return nil
	}

	if m.InvitedBy != nil {
		if err := m.InvitedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("invited_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("invited_by")
			}
			return err
		}
	}

	return nil
}

func (m *Invitation) validateInvitee(formats strfmt.Registry) error {
	if swag.IsZero(m.Invitee) { // not required
		return nil
	}

	if m.Invitee != nil {
		if err := m.Invitee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("invitee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("invitee")
			}
			return err
		}
	}

	return nil
}

func (m *Invitation) validateResentAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ResentAt) { // not required
		return nil
	}

	if err := validate.MinimumUint("resent_at", "body", *m.ResentAt, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Invitation) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

var invitationTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending","accepted","declined"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invitationTypeStatePropEnum = append(invitationTypeStatePropEnum, v)
	}
}

const (

	// InvitationStatePending captures enum value "pending"
	InvitationStatePending string = "pending"

	// InvitationStateAccepted captures enum value "accepted"
	InvitationStateAccepted string = "accepted"

	// InvitationStateDeclined captures enum value "declined"
	InvitationStateDeclined string = "declined"
)

// prop value enum
func (m *Invitation) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, invitationTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Invitation) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this invitation based on the context it is used
func (m *Invitation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInvitedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInvitee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Invitation) contextValidateInvitedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.InvitedBy != nil {

		if swag.IsZero(m.InvitedBy) { // not required
			return nil
		}

		if err := m.InvitedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("invited_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("invited_by")
			}
			return err
		}
	}

	return nil
}

func (m *Invitation) contextValidateInvitee(ctx context.Context, formats strfmt.Registry) error {

	if m.Invitee != nil {

		if swag.IsZero(m.Invitee) { // not required
			return nil
		}

		if err := m.Invitee.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("invitee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("invitee")
			}
			return err
		}
	}

	return nil
}

func (m *Invitation) contextValidateRole(ctx context.Context, formats strfmt.Registry) error {

	if m.Role != nil {

		if err := m.Role.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Invitation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Invitation) UnmarshalBinary(b []byte) error {
	var res Invitation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
