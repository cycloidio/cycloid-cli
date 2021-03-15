// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewUserAccount Sign up
//
// Create a new user account.
// swagger:model NewUserAccount
type NewUserAccount struct {

	// Code of a country the user is from
	// Required: true
	// Pattern: ^[A-Z]{2}$
	CountryCode *string `json:"country_code"`

	// email
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

	// The field is used when a user signup from an invitation to an organization. Giving the token, the created user will be automatically added to the organization.
	// Min Length: 5
	InvitationToken string `json:"invitation_token,omitempty"`

	// User's preferred language
	// Required: true
	// Enum: [en fr es]
	Locale *string `json:"locale"`

	// password
	// Required: true
	// Min Length: 8
	// Format: password
	Password *strfmt.Password `json:"password"`

	// username
	// Required: true
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	Username *string `json:"username"`
}

// Validate validates this new user account
func (m *NewUserAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountryCode(formats); err != nil {
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

	if err := m.validateInvitationToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocale(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
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

func (m *NewUserAccount) validateCountryCode(formats strfmt.Registry) error {

	if err := validate.Required("country_code", "body", m.CountryCode); err != nil {
		return err
	}

	if err := validate.Pattern("country_code", "body", string(*m.CountryCode), `^[A-Z]{2}$`); err != nil {
		return err
	}

	return nil
}

func (m *NewUserAccount) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewUserAccount) validateFamilyName(formats strfmt.Registry) error {

	if err := validate.Required("family_name", "body", m.FamilyName); err != nil {
		return err
	}

	if err := validate.MinLength("family_name", "body", string(*m.FamilyName), 2); err != nil {
		return err
	}

	return nil
}

func (m *NewUserAccount) validateGivenName(formats strfmt.Registry) error {

	if err := validate.Required("given_name", "body", m.GivenName); err != nil {
		return err
	}

	if err := validate.MinLength("given_name", "body", string(*m.GivenName), 2); err != nil {
		return err
	}

	return nil
}

func (m *NewUserAccount) validateInvitationToken(formats strfmt.Registry) error {

	if swag.IsZero(m.InvitationToken) { // not required
		return nil
	}

	if err := validate.MinLength("invitation_token", "body", string(m.InvitationToken), 5); err != nil {
		return err
	}

	return nil
}

var newUserAccountTypeLocalePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["en","fr","es"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		newUserAccountTypeLocalePropEnum = append(newUserAccountTypeLocalePropEnum, v)
	}
}

const (

	// NewUserAccountLocaleEn captures enum value "en"
	NewUserAccountLocaleEn string = "en"

	// NewUserAccountLocaleFr captures enum value "fr"
	NewUserAccountLocaleFr string = "fr"

	// NewUserAccountLocaleEs captures enum value "es"
	NewUserAccountLocaleEs string = "es"
)

// prop value enum
func (m *NewUserAccount) validateLocaleEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, newUserAccountTypeLocalePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NewUserAccount) validateLocale(formats strfmt.Registry) error {

	if err := validate.Required("locale", "body", m.Locale); err != nil {
		return err
	}

	// value enum
	if err := m.validateLocaleEnum("locale", "body", *m.Locale); err != nil {
		return err
	}

	return nil
}

func (m *NewUserAccount) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.MinLength("password", "body", string(*m.Password), 8); err != nil {
		return err
	}

	if err := validate.FormatOf("password", "body", "password", m.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewUserAccount) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	if err := validate.MinLength("username", "body", string(*m.Username), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("username", "body", string(*m.Username), 100); err != nil {
		return err
	}

	if err := validate.Pattern("username", "body", string(*m.Username), `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewUserAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewUserAccount) UnmarshalBinary(b []byte) error {
	var res NewUserAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}