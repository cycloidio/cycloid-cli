// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateConfigRepository UpdateConfigRepository
//
// swagger:model UpdateConfigRepository
type UpdateConfigRepository struct {

	// branch
	// Required: true
	Branch *string `json:"branch"`

	// credential canonical
	// Required: true
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	CredentialCanonical *string `json:"credential_canonical"`

	// default
	// Required: true
	Default *bool `json:"default"`

	// name
	// Required: true
	Name *string `json:"name"`

	// url
	// Required: true
	// Pattern: ^((/|~)[^/]*)+.(\.git)|(([\w\]+@[\w\.]+))(:(//)?)([\w\.@\:/\-~]+)(/)?
	URL *string `json:"url"`
}

// Validate validates this update config repository
func (m *UpdateConfigRepository) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBranch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefault(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateConfigRepository) validateBranch(formats strfmt.Registry) error {

	if err := validate.Required("branch", "body", m.Branch); err != nil {
		return err
	}

	return nil
}

func (m *UpdateConfigRepository) validateCredentialCanonical(formats strfmt.Registry) error {

	if err := validate.Required("credential_canonical", "body", m.CredentialCanonical); err != nil {
		return err
	}

	if err := validate.MinLength("credential_canonical", "body", *m.CredentialCanonical, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("credential_canonical", "body", *m.CredentialCanonical, 100); err != nil {
		return err
	}

	if err := validate.Pattern("credential_canonical", "body", *m.CredentialCanonical, `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *UpdateConfigRepository) validateDefault(formats strfmt.Registry) error {

	if err := validate.Required("default", "body", m.Default); err != nil {
		return err
	}

	return nil
}

func (m *UpdateConfigRepository) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *UpdateConfigRepository) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	if err := validate.Pattern("url", "body", *m.URL, `^((/|~)[^/]*)+.(\.git)|(([\w\]+@[\w\.]+))(:(//)?)([\w\.@\:/\-~]+)(/)?`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update config repository based on context it is used
func (m *UpdateConfigRepository) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateConfigRepository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateConfigRepository) UnmarshalBinary(b []byte) error {
	var res UpdateConfigRepository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
