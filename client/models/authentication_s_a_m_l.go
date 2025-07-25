// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AuthenticationSAML AuthenticationSAML
//
// # SAML Authentication configuration
//
// swagger:model AuthenticationSAML
type AuthenticationSAML struct {
	enabledField *bool

	// Entity ID of the SAML2 identity provider.
	Provider string `json:"provider,omitempty"`

	// SSO URL to which the user should be redirected in order to authenticate with the Identity Provider.
	// Format: uri
	SsoURL strfmt.URI `json:"sso_url,omitempty"`
}

// Enabled gets the enabled of this subtype
func (m *AuthenticationSAML) Enabled() *bool {
	return m.enabledField
}

// SetEnabled sets the enabled of this subtype
func (m *AuthenticationSAML) SetEnabled(val *bool) {
	m.enabledField = val
}

// Type gets the type of this subtype
func (m *AuthenticationSAML) Type() string {
	return "AuthenticationSAML"
}

// SetType sets the type of this subtype
func (m *AuthenticationSAML) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AuthenticationSAML) UnmarshalJSON(raw []byte) error {
	var data struct {

		// Entity ID of the SAML2 identity provider.
		Provider string `json:"provider,omitempty"`

		// SSO URL to which the user should be redirected in order to authenticate with the Identity Provider.
		// Format: uri
		SsoURL strfmt.URI `json:"sso_url,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Enabled *bool `json:"enabled"`

		Type string `json:"type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result AuthenticationSAML

	result.enabledField = base.Enabled

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Provider = data.Provider
	result.SsoURL = data.SsoURL

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AuthenticationSAML) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// Entity ID of the SAML2 identity provider.
		Provider string `json:"provider,omitempty"`

		// SSO URL to which the user should be redirected in order to authenticate with the Identity Provider.
		// Format: uri
		SsoURL strfmt.URI `json:"sso_url,omitempty"`
	}{

		Provider: m.Provider,

		SsoURL: m.SsoURL,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Enabled *bool `json:"enabled"`

		Type string `json:"type"`
	}{

		Enabled: m.Enabled(),

		Type: m.Type(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this authentication s a m l
func (m *AuthenticationSAML) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsoURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthenticationSAML) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled()); err != nil {
		return err
	}

	return nil
}

func (m *AuthenticationSAML) validateSsoURL(formats strfmt.Registry) error {

	if swag.IsZero(m.SsoURL) { // not required
		return nil
	}

	if err := validate.FormatOf("sso_url", "body", "uri", m.SsoURL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this authentication s a m l based on the context it is used
func (m *AuthenticationSAML) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AuthenticationSAML) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthenticationSAML) UnmarshalBinary(b []byte) error {
	var res AuthenticationSAML
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
