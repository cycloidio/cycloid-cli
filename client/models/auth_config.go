// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AuthConfig AuthConfig
//
// swagger:model AuthConfig
type AuthConfig struct {

	// Local authentication method configuration.
	// Required: true
	Local *AuthConfigLocalAuth `json:"local"`

	oauthField []AuthConfigOAuth

	// List of SAML2 providers.
	// Required: true
	Saml2 []*AuthConfigSAML `json:"saml2"`
}

// Oauth gets the oauth of this base type
func (m *AuthConfig) Oauth() []AuthConfigOAuth {
	return m.oauthField
}

// SetOauth sets the oauth of this base type
func (m *AuthConfig) SetOauth(val []AuthConfigOAuth) {
	m.oauthField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AuthConfig) UnmarshalJSON(raw []byte) error {
	var data struct {
		Local *AuthConfigLocalAuth `json:"local"`

		Oauth json.RawMessage `json:"oauth"`

		Saml2 []*AuthConfigSAML `json:"saml2"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	propOauth, err := UnmarshalAuthConfigOAuthSlice(bytes.NewBuffer(data.Oauth), runtime.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	var result AuthConfig

	// local
	result.Local = data.Local

	// oauth
	result.oauthField = propOauth

	// saml2
	result.Saml2 = data.Saml2

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AuthConfig) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		Local *AuthConfigLocalAuth `json:"local"`

		Saml2 []*AuthConfigSAML `json:"saml2"`
	}{

		Local: m.Local,

		Saml2: m.Saml2,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Oauth []AuthConfigOAuth `json:"oauth"`
	}{

		Oauth: m.oauthField,
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("local")
			}
			return err
		}
	}

	return nil
}

func (m *AuthConfig) validateOauth(formats strfmt.Registry) error {

	if err := validate.Required("oauth", "body", m.Oauth()); err != nil {
		return err
	}

	for i := 0; i < len(m.Oauth()); i++ {

		if err := m.oauthField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oauth" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oauth" + "." + strconv.Itoa(i))
			}
			return err
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("saml2" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this auth config based on the context it is used
func (m *AuthConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOauth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSaml2(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthConfig) contextValidateLocal(ctx context.Context, formats strfmt.Registry) error {

	if m.Local != nil {

		if err := m.Local.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("local")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("local")
			}
			return err
		}
	}

	return nil
}

func (m *AuthConfig) contextValidateOauth(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Oauth()); i++ {

		if swag.IsZero(m.oauthField[i]) { // not required
			return nil
		}

		if err := m.oauthField[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oauth" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oauth" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AuthConfig) contextValidateSaml2(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Saml2); i++ {

		if m.Saml2[i] != nil {

			if swag.IsZero(m.Saml2[i]) { // not required
				return nil
			}

			if err := m.Saml2[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("saml2" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("saml2" + "." + strconv.Itoa(i))
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
