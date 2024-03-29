// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Subscription Subscription
//
// It reflects the relation between an Organization and a Plan which
// could be the Free Trial or others, for more info check https://www.cycloid.io/pricing
//
// swagger:model Subscription
type Subscription struct {

	// current members
	// Minimum: 0
	CurrentMembers *uint64 `json:"current_members,omitempty"`

	// expires at
	// Required: true
	// Minimum: 0
	ExpiresAt *uint64 `json:"expires_at"`

	// members count
	// Minimum: 0
	MembersCount *uint64 `json:"members_count,omitempty"`

	// plan
	// Required: true
	Plan *SubscriptionPlan `json:"plan"`
}

// Validate validates this subscription
func (m *Subscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembersCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Subscription) validateCurrentMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentMembers) { // not required
		return nil
	}

	if err := validate.MinimumInt("current_members", "body", int64(*m.CurrentMembers), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateExpiresAt(formats strfmt.Registry) error {

	if err := validate.Required("expires_at", "body", m.ExpiresAt); err != nil {
		return err
	}

	if err := validate.MinimumInt("expires_at", "body", int64(*m.ExpiresAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateMembersCount(formats strfmt.Registry) error {

	if swag.IsZero(m.MembersCount) { // not required
		return nil
	}

	if err := validate.MinimumInt("members_count", "body", int64(*m.MembersCount), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validatePlan(formats strfmt.Registry) error {

	if err := validate.Required("plan", "body", m.Plan); err != nil {
		return err
	}

	if m.Plan != nil {
		if err := m.Plan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Subscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Subscription) UnmarshalBinary(b []byte) error {
	var res Subscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
