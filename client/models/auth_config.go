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

// AuthConfig AuthConfig
// swagger:model AuthConfig
type AuthConfig struct {

	// Local authentication method configuration.
	// Required: true
	Local *AuthConfigLocalAuth `json:"local"`

	// List of OAuth providers.
	// Required: true
	Oauth []*AuthConfigOAuth `json:"oauth"`

	// List of SAML2 providers.
	// Required: true
	Saml2 []*AuthConfigSAML `json:"saml2"`
}

// Validate validates this auth config
func (m *AuthConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOauth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSaml2(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthConfig) validateLocal(formats strfmt.Registry) error {

	if err := validate.Required("local", "body", m.Local); err != nil {
		return err
	}

	if m.Local != nil {
		if err := m.Local.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("local")
			}
			return err
		}
	}

	return nil
}

func (m *AuthConfig) validateOauth(formats strfmt.Registry) error {

	if err := validate.Required("oauth", "body", m.Oauth); err != nil {
		return err
	}

	for i := 0; i < len(m.Oauth); i++ {
		if swag.IsZero(m.Oauth[i]) { // not required
			continue
		}

		if m.Oauth[i] != nil {
			if err := m.Oauth[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("oauth" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuthConfig) validateSaml2(formats strfmt.Registry) error {

	if err := validate.Required("saml2", "body", m.Saml2); err != nil {
		return err
	}

	for i := 0; i < len(m.Saml2); i++ {
		if swag.IsZero(m.Saml2[i]) { // not required
			continue
		}

		if m.Saml2[i] != nil {
			if err := m.Saml2[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("saml2" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthConfig) UnmarshalBinary(b []byte) error {
	var res AuthConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
