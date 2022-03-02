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

// Summary Summary of the organization
// swagger:model Summary
type Summary struct {

	// config repositories
	// Required: true
	ConfigRepositories *uint64 `json:"config_repositories"`

	// credentials
	// Required: true
	Credentials *uint64 `json:"credentials"`

	// pipelines
	// Required: true
	Pipelines *uint64 `json:"pipelines"`

	// projects
	// Required: true
	Projects *uint64 `json:"projects"`

	// roles
	// Required: true
	Roles *uint64 `json:"roles"`

	// service catalog sources
	// Required: true
	ServiceCatalogSources *uint64 `json:"service_catalog_sources"`

	// service catalogs
	// Required: true
	ServiceCatalogs *uint64 `json:"service_catalogs"`

	// teams
	// Required: true
	Teams *uint64 `json:"teams"`

	// users
	// Required: true
	Users *uint64 `json:"users"`
}

// Validate validates this summary
func (m *Summary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigRepositories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePipelines(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceCatalogSources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceCatalogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Summary) validateConfigRepositories(formats strfmt.Registry) error {

	if err := validate.Required("config_repositories", "body", m.ConfigRepositories); err != nil {
		return err
	}

	return nil
}

func (m *Summary) validateCredentials(formats strfmt.Registry) error {

	if err := validate.Required("credentials", "body", m.Credentials); err != nil {
		return err
	}

	return nil
}

func (m *Summary) validatePipelines(formats strfmt.Registry) error {

	if err := validate.Required("pipelines", "body", m.Pipelines); err != nil {
		return err
	}

	return nil
}

func (m *Summary) validateProjects(formats strfmt.Registry) error {

	if err := validate.Required("projects", "body", m.Projects); err != nil {
		return err
	}

	return nil
}

func (m *Summary) validateRoles(formats strfmt.Registry) error {

	if err := validate.Required("roles", "body", m.Roles); err != nil {
		return err
	}

	return nil
}

func (m *Summary) validateServiceCatalogSources(formats strfmt.Registry) error {

	if err := validate.Required("service_catalog_sources", "body", m.ServiceCatalogSources); err != nil {
		return err
	}

	return nil
}

func (m *Summary) validateServiceCatalogs(formats strfmt.Registry) error {

	if err := validate.Required("service_catalogs", "body", m.ServiceCatalogs); err != nil {
		return err
	}

	return nil
}

func (m *Summary) validateTeams(formats strfmt.Registry) error {

	if err := validate.Required("teams", "body", m.Teams); err != nil {
		return err
	}

	return nil
}

func (m *Summary) validateUsers(formats strfmt.Registry) error {

	if err := validate.Required("users", "body", m.Users); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Summary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Summary) UnmarshalBinary(b []byte) error {
	var res Summary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
