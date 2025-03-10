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

// NewServiceCatalogSource NewServiceCatalogSource
//
// swagger:model NewServiceCatalogSource
type NewServiceCatalogSource struct {

	// branch
	// Required: true
	Branch *string `json:"branch"`

	// canonical
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	Canonical string `json:"canonical,omitempty"`

	// credential canonical
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	CredentialCanonical string `json:"credential_canonical,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// User canonical that owns this service catalog source. If omitted then the person
	// creating this service catalog source will be assigned as owner. When a user is the
	// owner of a service catalog source they has all the permissions on it.
	//
	Owner string `json:"owner,omitempty"`

	// Team responsible for the maintenance of the underlying service catalogs
	//
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	TeamCanonical string `json:"team_canonical,omitempty"`

	// url
	// Required: true
	// Pattern: ^((/|~)[^/]*)+.(\.git)|(([\w\]+@[\w\.]+))(:(//)?)([\w\.@\:/\-~]+)(/)?
	URL *string `json:"url"`

	// The visibility setting allows to specify which visibility will be applied to stacks in this catalog repository.
	// This option is only applied during initial catalog repository creation, not for subsequent updates.
	//
	Visibility string `json:"visibility,omitempty"`
}

// Validate validates this new service catalog source
func (m *NewServiceCatalogSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBranch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamCanonical(formats); err != nil {
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

func (m *NewServiceCatalogSource) validateBranch(formats strfmt.Registry) error {

	if err := validate.Required("branch", "body", m.Branch); err != nil {
		return err
	}

	return nil
}

func (m *NewServiceCatalogSource) validateCanonical(formats strfmt.Registry) error {
	if swag.IsZero(m.Canonical) { // not required
		return nil
	}

	if err := validate.MinLength("canonical", "body", m.Canonical, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("canonical", "body", m.Canonical, 100); err != nil {
		return err
	}

	if err := validate.Pattern("canonical", "body", m.Canonical, `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *NewServiceCatalogSource) validateCredentialCanonical(formats strfmt.Registry) error {
	if swag.IsZero(m.CredentialCanonical) { // not required
		return nil
	}

	if err := validate.MinLength("credential_canonical", "body", m.CredentialCanonical, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("credential_canonical", "body", m.CredentialCanonical, 100); err != nil {
		return err
	}

	if err := validate.Pattern("credential_canonical", "body", m.CredentialCanonical, `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *NewServiceCatalogSource) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NewServiceCatalogSource) validateTeamCanonical(formats strfmt.Registry) error {
	if swag.IsZero(m.TeamCanonical) { // not required
		return nil
	}

	if err := validate.MinLength("team_canonical", "body", m.TeamCanonical, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("team_canonical", "body", m.TeamCanonical, 100); err != nil {
		return err
	}

	if err := validate.Pattern("team_canonical", "body", m.TeamCanonical, `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *NewServiceCatalogSource) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	if err := validate.Pattern("url", "body", *m.URL, `^((/|~)[^/]*)+.(\.git)|(([\w\]+@[\w\.]+))(:(//)?)([\w\.@\:/\-~]+)(/)?`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this new service catalog source based on context it is used
func (m *NewServiceCatalogSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewServiceCatalogSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewServiceCatalogSource) UnmarshalBinary(b []byte) error {
	var res NewServiceCatalogSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
