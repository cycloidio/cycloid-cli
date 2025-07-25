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

// NewOAuthUser User's OAuth information
//
// # The User OAuth information
//
// swagger:model NewOAuthUser
type NewOAuthUser struct {

	// Code of a country the user is from
	// Pattern: ^[A-Z]{2}$
	CountryCode string `json:"country_code,omitempty"`

	// email
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// family name
	FamilyName string `json:"family_name,omitempty"`

	// given name
	// Required: true
	GivenName *string `json:"given_name"`

	// The field is used when a user signup from an invitation to an organization. Giving the token, the created user will be automatically added to the organization.
	// Min Length: 5
	InvitationToken string `json:"invitation_token,omitempty"`

	// User's preferred language
	// Enum: ["en","fr","es"]
	Locale string `json:"locale,omitempty"`

	// organization canonical
	OrganizationCanonical string `json:"organization_canonical,omitempty"`

	// picture url
	// Format: uri
	PictureURL strfmt.URI `json:"picture_url,omitempty"`

	// social id
	// Required: true
	SocialID *string `json:"social_id"`

	// username
	// Required: true
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	Username *string `json:"username"`
}

// Validate validates this new o auth user
func (m *NewOAuthUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountryCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
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

	if err := m.validatePictureURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSocialID(formats); err != nil {
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

func (m *NewOAuthUser) validateCountryCode(formats strfmt.Registry) error {
	if swag.IsZero(m.CountryCode) { // not required
		return nil
	}

	if err := validate.Pattern("country_code", "body", m.CountryCode, `^[A-Z]{2}$`); err != nil {
		return err
	}

	return nil
}

func (m *NewOAuthUser) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewOAuthUser) validateGivenName(formats strfmt.Registry) error {

	if err := validate.Required("given_name", "body", m.GivenName); err != nil {
		return err
	}

	return nil
}

func (m *NewOAuthUser) validateInvitationToken(formats strfmt.Registry) error {
	if swag.IsZero(m.InvitationToken) { // not required
		return nil
	}

	if err := validate.MinLength("invitation_token", "body", m.InvitationToken, 5); err != nil {
		return err
	}

	return nil
}

var newOAuthUserTypeLocalePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["en","fr","es"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		newOAuthUserTypeLocalePropEnum = append(newOAuthUserTypeLocalePropEnum, v)
	}
}

const (

	// NewOAuthUserLocaleEn captures enum value "en"
	NewOAuthUserLocaleEn string = "en"

	// NewOAuthUserLocaleFr captures enum value "fr"
	NewOAuthUserLocaleFr string = "fr"

	// NewOAuthUserLocaleEs captures enum value "es"
	NewOAuthUserLocaleEs string = "es"
)

// prop value enum
func (m *NewOAuthUser) validateLocaleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, newOAuthUserTypeLocalePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NewOAuthUser) validateLocale(formats strfmt.Registry) error {
	if swag.IsZero(m.Locale) { // not required
		return nil
	}

	// value enum
	if err := m.validateLocaleEnum("locale", "body", m.Locale); err != nil {
		return err
	}

	return nil
}

func (m *NewOAuthUser) validatePictureURL(formats strfmt.Registry) error {
	if swag.IsZero(m.PictureURL) { // not required
		return nil
	}

	if err := validate.FormatOf("picture_url", "body", "uri", m.PictureURL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewOAuthUser) validateSocialID(formats strfmt.Registry) error {

	if err := validate.Required("social_id", "body", m.SocialID); err != nil {
		return err
	}

	return nil
}

func (m *NewOAuthUser) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	if err := validate.MinLength("username", "body", *m.Username, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("username", "body", *m.Username, 100); err != nil {
		return err
	}

	if err := validate.Pattern("username", "body", *m.Username, `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this new o auth user based on context it is used
func (m *NewOAuthUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewOAuthUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewOAuthUser) UnmarshalBinary(b []byte) error {
	var res NewOAuthUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
