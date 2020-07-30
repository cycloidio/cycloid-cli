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

// MemberTeam Member of a team
//
// Member is a user who is associated to a team.
// swagger:model MemberTeam
type MemberTeam struct {

	// When the user became a member.
	// Required: true
	// Minimum: 0
	CreatedAt *int64 `json:"created_at"`

	// user email
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// family name
	// Required: true
	// Min Length: 2
	FamilyName *string `json:"family_name"`

	// given name
	// Required: true
	// Min Length: 2
	GivenName *string `json:"given_name"`

	// id
	// Required: true
	// Minimum: 1
	ID *uint32 `json:"id"`

	// picture url
	// Format: uri
	PictureURL strfmt.URI `json:"picture_url,omitempty"`

	// When the user had the role modified.
	// Minimum: 0
	UpdatedAt *int64 `json:"updated_at,omitempty"`

	// username
	// Required: true
	// Max Length: 30
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	Username *string `json:"username"`
}

// Validate validates this member team
func (m *MemberTeam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFamilyName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGivenName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePictureURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MemberTeam) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.MinimumInt("created_at", "body", int64(*m.CreatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *MemberTeam) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MemberTeam) validateFamilyName(formats strfmt.Registry) error {

	if err := validate.Required("family_name", "body", m.FamilyName); err != nil {
		return err
	}

	if err := validate.MinLength("family_name", "body", string(*m.FamilyName), 2); err != nil {
		return err
	}

	return nil
}

func (m *MemberTeam) validateGivenName(formats strfmt.Registry) error {

	if err := validate.Required("given_name", "body", m.GivenName); err != nil {
		return err
	}

	if err := validate.MinLength("given_name", "body", string(*m.GivenName), 2); err != nil {
		return err
	}

	return nil
}

func (m *MemberTeam) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinimumInt("id", "body", int64(*m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *MemberTeam) validatePictureURL(formats strfmt.Registry) error {

	if swag.IsZero(m.PictureURL) { // not required
		return nil
	}

	if err := validate.FormatOf("picture_url", "body", "uri", m.PictureURL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MemberTeam) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.MinimumInt("updated_at", "body", int64(*m.UpdatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *MemberTeam) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	if err := validate.MinLength("username", "body", string(*m.Username), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("username", "body", string(*m.Username), 30); err != nil {
		return err
	}

	if err := validate.Pattern("username", "body", string(*m.Username), `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MemberTeam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MemberTeam) UnmarshalBinary(b []byte) error {
	var res MemberTeam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}