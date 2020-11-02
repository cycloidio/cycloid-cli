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

// UserAccount User's account
//
// The user's account contains information related with the authenticated user.
// swagger:model UserAccount
type UserAccount struct {

	// created at
	// Required: true
	// Minimum: 0
	CreatedAt *uint64 `json:"created_at"`

	// emails
	// Required: true
	// Min Items: 1
	Emails []*UserAccountEmail `json:"emails"`

	// family name
	// Required: true
	// Min Length: 2
	FamilyName *string `json:"family_name"`

	// given name
	// Required: true
	// Min Length: 2
	GivenName *string `json:"given_name"`

	// guide
	Guide UserGuide `json:"guide,omitempty"`

	// last login
	// Required: true
	// Minimum: 0
	LastLogin *uint64 `json:"last_login"`

	// The local that the user prefer.
	// Required: true
	// Pattern: ^[a-z]{2}(?:-[a-z][a-z])?$
	Locale *string `json:"locale"`

	// picture url
	// Format: uri
	PictureURL strfmt.URI `json:"picture_url,omitempty"`

	// updated at
	// Required: true
	// Minimum: 0
	UpdatedAt *uint64 `json:"updated_at"`

	// username
	// Required: true
	// Max Length: 30
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	Username *string `json:"username"`
}

// Validate validates this user account
func (m *UserAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFamilyName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGivenName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastLogin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocale(formats); err != nil {
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

func (m *UserAccount) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.MinimumInt("created_at", "body", int64(*m.CreatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *UserAccount) validateEmails(formats strfmt.Registry) error {

	if err := validate.Required("emails", "body", m.Emails); err != nil {
		return err
	}

	iEmailsSize := int64(len(m.Emails))

	if err := validate.MinItems("emails", "body", iEmailsSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Emails); i++ {
		if swag.IsZero(m.Emails[i]) { // not required
			continue
		}

		if m.Emails[i] != nil {
			if err := m.Emails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserAccount) validateFamilyName(formats strfmt.Registry) error {

	if err := validate.Required("family_name", "body", m.FamilyName); err != nil {
		return err
	}

	if err := validate.MinLength("family_name", "body", string(*m.FamilyName), 2); err != nil {
		return err
	}

	return nil
}

func (m *UserAccount) validateGivenName(formats strfmt.Registry) error {

	if err := validate.Required("given_name", "body", m.GivenName); err != nil {
		return err
	}

	if err := validate.MinLength("given_name", "body", string(*m.GivenName), 2); err != nil {
		return err
	}

	return nil
}

func (m *UserAccount) validateLastLogin(formats strfmt.Registry) error {

	if err := validate.Required("last_login", "body", m.LastLogin); err != nil {
		return err
	}

	if err := validate.MinimumInt("last_login", "body", int64(*m.LastLogin), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *UserAccount) validateLocale(formats strfmt.Registry) error {

	if err := validate.Required("locale", "body", m.Locale); err != nil {
		return err
	}

	if err := validate.Pattern("locale", "body", string(*m.Locale), `^[a-z]{2}(?:-[a-z][a-z])?$`); err != nil {
		return err
	}

	return nil
}

func (m *UserAccount) validatePictureURL(formats strfmt.Registry) error {

	if swag.IsZero(m.PictureURL) { // not required
		return nil
	}

	if err := validate.FormatOf("picture_url", "body", "uri", m.PictureURL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserAccount) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.MinimumInt("updated_at", "body", int64(*m.UpdatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *UserAccount) validateUsername(formats strfmt.Registry) error {

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
func (m *UserAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserAccount) UnmarshalBinary(b []byte) error {
	var res UserAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
