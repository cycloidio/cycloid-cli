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

// CloudProvider Cloud Provider
//
// CloudProvider represents a cloud provider. Those cloud providers are used to identify the scope of projects and/or stacks.
//
// swagger:model CloudProvider
type CloudProvider struct {

	// abbreviation
	// Max Length: 60
	// Min Length: 2
	Abbreviation string `json:"abbreviation,omitempty"`

	// canonical
	// Required: true
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	// Enum: ["aws","google","azurerm","flexibleengine","openstack","scaleway","vmware","ovh","alibaba","oracle","vsphere","kubernetes"]
	Canonical *string `json:"canonical"`

	// created at
	// Minimum: 0
	CreatedAt *uint64 `json:"created_at,omitempty"`

	// id
	// Minimum: 1
	ID uint32 `json:"id,omitempty"`

	// name
	// Required: true
	// Max Length: 60
	// Min Length: 2
	Name *string `json:"name"`

	// regions
	Regions []string `json:"regions"`

	// updated at
	// Minimum: 0
	UpdatedAt *uint64 `json:"updated_at,omitempty"`
}

// Validate validates this cloud provider
func (m *CloudProvider) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbbreviation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudProvider) validateAbbreviation(formats strfmt.Registry) error {
	if swag.IsZero(m.Abbreviation) { // not required
		return nil
	}

	if err := validate.MinLength("abbreviation", "body", m.Abbreviation, 2); err != nil {
		return err
	}

	if err := validate.MaxLength("abbreviation", "body", m.Abbreviation, 60); err != nil {
		return err
	}

	return nil
}

var cloudProviderTypeCanonicalPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aws","google","azurerm","flexibleengine","openstack","scaleway","vmware","ovh","alibaba","oracle","vsphere","kubernetes"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudProviderTypeCanonicalPropEnum = append(cloudProviderTypeCanonicalPropEnum, v)
	}
}

const (

	// CloudProviderCanonicalAws captures enum value "aws"
	CloudProviderCanonicalAws string = "aws"

	// CloudProviderCanonicalGoogle captures enum value "google"
	CloudProviderCanonicalGoogle string = "google"

	// CloudProviderCanonicalAzurerm captures enum value "azurerm"
	CloudProviderCanonicalAzurerm string = "azurerm"

	// CloudProviderCanonicalFlexibleengine captures enum value "flexibleengine"
	CloudProviderCanonicalFlexibleengine string = "flexibleengine"

	// CloudProviderCanonicalOpenstack captures enum value "openstack"
	CloudProviderCanonicalOpenstack string = "openstack"

	// CloudProviderCanonicalScaleway captures enum value "scaleway"
	CloudProviderCanonicalScaleway string = "scaleway"

	// CloudProviderCanonicalVmware captures enum value "vmware"
	CloudProviderCanonicalVmware string = "vmware"

	// CloudProviderCanonicalOvh captures enum value "ovh"
	CloudProviderCanonicalOvh string = "ovh"

	// CloudProviderCanonicalAlibaba captures enum value "alibaba"
	CloudProviderCanonicalAlibaba string = "alibaba"

	// CloudProviderCanonicalOracle captures enum value "oracle"
	CloudProviderCanonicalOracle string = "oracle"

	// CloudProviderCanonicalVsphere captures enum value "vsphere"
	CloudProviderCanonicalVsphere string = "vsphere"

	// CloudProviderCanonicalKubernetes captures enum value "kubernetes"
	CloudProviderCanonicalKubernetes string = "kubernetes"
)

// prop value enum
func (m *CloudProvider) validateCanonicalEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cloudProviderTypeCanonicalPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CloudProvider) validateCanonical(formats strfmt.Registry) error {

	if err := validate.Required("canonical", "body", m.Canonical); err != nil {
		return err
	}

	if err := validate.MinLength("canonical", "body", *m.Canonical, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("canonical", "body", *m.Canonical, 100); err != nil {
		return err
	}

	if err := validate.Pattern("canonical", "body", *m.Canonical, `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	// value enum
	if err := m.validateCanonicalEnum("canonical", "body", *m.Canonical); err != nil {
		return err
	}

	return nil
}

func (m *CloudProvider) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.MinimumUint("created_at", "body", *m.CreatedAt, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *CloudProvider) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinimumUint("id", "body", uint64(m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *CloudProvider) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 2); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 60); err != nil {
		return err
	}

	return nil
}

func (m *CloudProvider) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.MinimumUint("updated_at", "body", *m.UpdatedAt, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cloud provider based on context it is used
func (m *CloudProvider) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CloudProvider) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudProvider) UnmarshalBinary(b []byte) error {
	var res CloudProvider
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
