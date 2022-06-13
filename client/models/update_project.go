// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateProject Update Project
//
// The entity which represents the information of the project to be updated.
// swagger:model UpdateProject
type UpdateProject struct {

	// The cloud provider canonical that this project is using - between the
	// supported ones.
	//
	// Enum: [aws google azurerm flexibleengine openstack]
	CloudProvider string `json:"cloud_provider,omitempty"`

	// The config_repository_canonical points to new Config Repository the project
	// will be using. If this value is filled and it's different from the
	// current one, the whole project will be migrated to new CR, meaning
	// configuration files will also be moved.
	// If the project didn't have config_repository_canonical set, this action will
	// only attach the project to the CR, it won't create/move any files.
	// In order to be sure everything works, make sure the
	// config_repository_canonical is pointing at the CR with the same git
	// repository that was used during project creation.
	// Although the config_repository_canonical is not marked as required,
	// it's actually required for projects that are already using CR. This
	// field not being required is only for compatibility with older projects,
	// which are not having CR yet.
	//
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	ConfigRepositoryCanonical string `json:"config_repository_canonical,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// environments
	// Min Items: 1
	Environments []*NewEnvironment `json:"environments"`

	// The variables set within a form with the corresponding environment
	// canonical and use case
	//
	Inputs []*FormInput `json:"inputs"`

	// name
	// Required: true
	// Min Length: 3
	Name *string `json:"name"`

	// User canonical that owns this project. Only the owner or an
	// organization admin can update such a field. When a user is the owner
	// of a project it has all the permission on it.
	//
	Owner string `json:"owner,omitempty"`

	// It's the ref of the Service Catalog, like 'cycloidio:stack-magento'
	// Required: true
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+:[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	ServiceCatalogRef *string `json:"service_catalog_ref"`
}

// Validate validates this update project
func (m *UpdateProject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigRepositoryCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInputs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceCatalogRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateProjectTypeCloudProviderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aws","google","azurerm","flexibleengine","openstack"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateProjectTypeCloudProviderPropEnum = append(updateProjectTypeCloudProviderPropEnum, v)
	}
}

const (

	// UpdateProjectCloudProviderAws captures enum value "aws"
	UpdateProjectCloudProviderAws string = "aws"

	// UpdateProjectCloudProviderGoogle captures enum value "google"
	UpdateProjectCloudProviderGoogle string = "google"

	// UpdateProjectCloudProviderAzurerm captures enum value "azurerm"
	UpdateProjectCloudProviderAzurerm string = "azurerm"

	// UpdateProjectCloudProviderFlexibleengine captures enum value "flexibleengine"
	UpdateProjectCloudProviderFlexibleengine string = "flexibleengine"

	// UpdateProjectCloudProviderOpenstack captures enum value "openstack"
	UpdateProjectCloudProviderOpenstack string = "openstack"
)

// prop value enum
func (m *UpdateProject) validateCloudProviderEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, updateProjectTypeCloudProviderPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UpdateProject) validateCloudProvider(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudProvider) { // not required
		return nil
	}

	// value enum
	if err := m.validateCloudProviderEnum("cloud_provider", "body", m.CloudProvider); err != nil {
		return err
	}

	return nil
}

func (m *UpdateProject) validateConfigRepositoryCanonical(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigRepositoryCanonical) { // not required
		return nil
	}

	if err := validate.MinLength("config_repository_canonical", "body", string(m.ConfigRepositoryCanonical), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("config_repository_canonical", "body", string(m.ConfigRepositoryCanonical), 100); err != nil {
		return err
	}

	if err := validate.Pattern("config_repository_canonical", "body", string(m.ConfigRepositoryCanonical), `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *UpdateProject) validateEnvironments(formats strfmt.Registry) error {

	if swag.IsZero(m.Environments) { // not required
		return nil
	}

	iEnvironmentsSize := int64(len(m.Environments))

	if err := validate.MinItems("environments", "body", iEnvironmentsSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Environments); i++ {
		if swag.IsZero(m.Environments[i]) { // not required
			continue
		}

		if m.Environments[i] != nil {
			if err := m.Environments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("environments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateProject) validateInputs(formats strfmt.Registry) error {

	if swag.IsZero(m.Inputs) { // not required
		return nil
	}

	for i := 0; i < len(m.Inputs); i++ {
		if swag.IsZero(m.Inputs[i]) { // not required
			continue
		}

		if m.Inputs[i] != nil {
			if err := m.Inputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateProject) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 3); err != nil {
		return err
	}

	return nil
}

func (m *UpdateProject) validateServiceCatalogRef(formats strfmt.Registry) error {

	if err := validate.Required("service_catalog_ref", "body", m.ServiceCatalogRef); err != nil {
		return err
	}

	if err := validate.Pattern("service_catalog_ref", "body", string(*m.ServiceCatalogRef), `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+:[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateProject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateProject) UnmarshalBinary(b []byte) error {
	var res UpdateProject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
