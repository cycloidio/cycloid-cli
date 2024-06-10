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

// UserLogin Log in
//
// Validate the user to access to the application. The user can login with the primary email address or with username.
//
// MinProperties: 2
// MaxProperties: 2
//
// swagger:model UserLogin
type UserLogin struct {

	// email
	// Format: email
	Email strfmt.Email `json:"email,omitempty"`

	// organization canonical
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	OrganizationCanonical string `json:"organization_canonical,omitempty"`

	// password
	// Required: true
	// Min Length: 8
	// Format: password
	Password *strfmt.Password `json:"password"`

	// username
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	Username string `json:"username,omitempty"`

	// user login additional properties
	UserLoginAdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *UserLogin) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// email
		// Format: email
		Email strfmt.Email `json:"email,omitempty"`

		// organization canonical
		// Max Length: 100
		// Min Length: 3
		// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
		OrganizationCanonical string `json:"organization_canonical,omitempty"`

		// password
		// Required: true
		// Min Length: 8
		// Format: password
		Password *strfmt.Password `json:"password"`

		// username
		// Max Length: 100
		// Min Length: 3
		// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
		Username string `json:"username,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv UserLogin

	rcv.Email = stage1.Email
	rcv.OrganizationCanonical = stage1.OrganizationCanonical
	rcv.Password = stage1.Password
	rcv.Username = stage1.Username
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "email")
	delete(stage2, "organization_canonical")
	delete(stage2, "password")
	delete(stage2, "username")
	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.UserLoginAdditionalProperties = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m UserLogin) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// email
		// Format: email
		Email strfmt.Email `json:"email,omitempty"`

		// organization canonical
		// Max Length: 100
		// Min Length: 3
		// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
		OrganizationCanonical string `json:"organization_canonical,omitempty"`

		// password
		// Required: true
		// Min Length: 8
		// Format: password
		Password *strfmt.Password `json:"password"`

		// username
		// Max Length: 100
		// Min Length: 3
		// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
		Username string `json:"username,omitempty"`
	}

	stage1.Email = m.Email
	stage1.OrganizationCanonical = m.OrganizationCanonical
	stage1.Password = m.Password
	stage1.Username = m.Username

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.UserLoginAdditionalProperties) == 0 { // no additional properties
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.UserLoginAdditionalProperties)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 { // "{}": only additional properties
		return additional, nil
	}

	// concatenate the 2 objects
	return swag.ConcatJSON(props, additional), nil
}

// Validate validates this user login
func (m *UserLogin) Validate(formats strfmt.Registry) error {
	var res []error

	// short circuits minProperties > 0
	if m == nil {
		return errors.TooFewProperties("", "body", 2)
	}

	props := make(map[string]json.RawMessage, 4+10)
	j, err := swag.WriteJSON(m)
	if err != nil {
		return err
	}

	if err = swag.ReadJSON(j, &props); err != nil {
		return err
	}

	nprops := len(props)

	// minProperties: 2
	if nprops < 2 {
		return errors.TooFewProperties("", "body", 2)
	}

	// maxProperties: 2
	if nprops > 2 {
		return errors.TooManyProperties("", "body", 2)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationCanonical(formats); err != nil {
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

func (m *UserLogin) validateEmail(formats strfmt.Registry) error {
	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserLogin) validateOrganizationCanonical(formats strfmt.Registry) error {
	if swag.IsZero(m.OrganizationCanonical) { // not required
		return nil
	}

	if err := validate.MinLength("organization_canonical", "body", m.OrganizationCanonical, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("organization_canonical", "body", m.OrganizationCanonical, 100); err != nil {
		return err
	}

	if err := validate.Pattern("organization_canonical", "body", m.OrganizationCanonical, `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *UserLogin) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.MinLength("password", "body", m.Password.String(), 8); err != nil {
		return err
	}

	if err := validate.FormatOf("password", "body", "password", m.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserLogin) validateUsername(formats strfmt.Registry) error {
	if swag.IsZero(m.Username) { // not required
		return nil
	}

	if err := validate.MinLength("username", "body", m.Username, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("username", "body", m.Username, 100); err != nil {
		return err
	}

	if err := validate.Pattern("username", "body", m.Username, `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user login based on context it is used
func (m *UserLogin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserLogin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserLogin) UnmarshalBinary(b []byte) error {
	var res UserLogin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
