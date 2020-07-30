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

// ServiceCatalogChanges ServiceCatalogChanges
//
// Represents list of service catalogs changes during the refresh of a service catalog source.
// swagger:model ServiceCatalogChanges
type ServiceCatalogChanges struct {

	// created
	// Required: true
	Created []*ServiceCatalog `json:"created"`

	// deleted
	// Required: true
	Deleted []*ServiceCatalog `json:"deleted"`

	// updated
	// Required: true
	Updated []*ServiceCatalog `json:"updated"`
}

// Validate validates this service catalog changes
func (m *ServiceCatalogChanges) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceCatalogChanges) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	for i := 0; i < len(m.Created); i++ {
		if swag.IsZero(m.Created[i]) { // not required
			continue
		}

		if m.Created[i] != nil {
			if err := m.Created[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("created" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceCatalogChanges) validateDeleted(formats strfmt.Registry) error {

	if err := validate.Required("deleted", "body", m.Deleted); err != nil {
		return err
	}

	for i := 0; i < len(m.Deleted); i++ {
		if swag.IsZero(m.Deleted[i]) { // not required
			continue
		}

		if m.Deleted[i] != nil {
			if err := m.Deleted[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deleted" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceCatalogChanges) validateUpdated(formats strfmt.Registry) error {

	if err := validate.Required("updated", "body", m.Updated); err != nil {
		return err
	}

	for i := 0; i < len(m.Updated); i++ {
		if swag.IsZero(m.Updated[i]) { // not required
			continue
		}

		if m.Updated[i] != nil {
			if err := m.Updated[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updated" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceCatalogChanges) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCatalogChanges) UnmarshalBinary(b []byte) error {
	var res ServiceCatalogChanges
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}