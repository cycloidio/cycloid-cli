// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CredentialInUse Credential in use
//
// InUse represents the resources that are using provided credential.
//
// swagger:model CredentialInUse
type CredentialInUse struct {

	// config repositories
	ConfigRepositories []*InUseConfigRepository `json:"config_repositories"`

	// external backends
	ExternalBackends []*InUseExternalBackend `json:"external_backends"`

	// service catalog sources
	ServiceCatalogSources []*InUseServiceCatalogSource `json:"service_catalog_sources"`
}

// Validate validates this credential in use
func (m *CredentialInUse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigRepositories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalBackends(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceCatalogSources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CredentialInUse) validateConfigRepositories(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigRepositories) { // not required
		return nil
	}

	for i := 0; i < len(m.ConfigRepositories); i++ {
		if swag.IsZero(m.ConfigRepositories[i]) { // not required
			continue
		}

		if m.ConfigRepositories[i] != nil {
			if err := m.ConfigRepositories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("config_repositories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("config_repositories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CredentialInUse) validateExternalBackends(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalBackends) { // not required
		return nil
	}

	for i := 0; i < len(m.ExternalBackends); i++ {
		if swag.IsZero(m.ExternalBackends[i]) { // not required
			continue
		}

		if m.ExternalBackends[i] != nil {
			if err := m.ExternalBackends[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("external_backends" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("external_backends" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CredentialInUse) validateServiceCatalogSources(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceCatalogSources) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceCatalogSources); i++ {
		if swag.IsZero(m.ServiceCatalogSources[i]) { // not required
			continue
		}

		if m.ServiceCatalogSources[i] != nil {
			if err := m.ServiceCatalogSources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("service_catalog_sources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("service_catalog_sources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this credential in use based on the context it is used
func (m *CredentialInUse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfigRepositories(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalBackends(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceCatalogSources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CredentialInUse) contextValidateConfigRepositories(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConfigRepositories); i++ {

		if m.ConfigRepositories[i] != nil {

			if swag.IsZero(m.ConfigRepositories[i]) { // not required
				return nil
			}

			if err := m.ConfigRepositories[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("config_repositories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("config_repositories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CredentialInUse) contextValidateExternalBackends(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExternalBackends); i++ {

		if m.ExternalBackends[i] != nil {

			if swag.IsZero(m.ExternalBackends[i]) { // not required
				return nil
			}

			if err := m.ExternalBackends[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("external_backends" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("external_backends" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CredentialInUse) contextValidateServiceCatalogSources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServiceCatalogSources); i++ {

		if m.ServiceCatalogSources[i] != nil {

			if swag.IsZero(m.ServiceCatalogSources[i]) { // not required
				return nil
			}

			if err := m.ServiceCatalogSources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("service_catalog_sources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("service_catalog_sources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CredentialInUse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CredentialInUse) UnmarshalBinary(b []byte) error {
	var res CredentialInUse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
