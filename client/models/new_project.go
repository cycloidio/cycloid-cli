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

// NewProject Create Project
//
// The entity which represents the information of a new project.
// swagger:model NewProject
type NewProject struct {

	// canonical
	// Max Length: 100
	// Min Length: 1
	// Pattern: (^[a-z0-9]+(([a-z0-9\-_]+)?[a-z0-9]+)?$)
	Canonical string `json:"canonical,omitempty"`

	// config repository canonical
	// Required: true
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	ConfigRepositoryCanonical *string `json:"config_repository_canonical"`

	// A description regarding the project to help identify/remember details,
	// implementation, purpose, etc.
	//
	Description string `json:"description,omitempty"`

	// The variables set within a form with the corresponding environment
	// canonical and use case
	//
	Inputs []*FormInput `json:"inputs"`

	// name
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// User canonical that owns this project. If omitted then the person
	// creating this project will be assigned as owner. When a user is the
	// owner of a project it has all the permissions on it.
	//
	Owner string `json:"owner,omitempty"`

	// Each instance should include passed_config if no inputs are sent on
	// project creation, otherwise it will be inferred internally.
	//
	// Required: true
	// Min Items: 1
	Pipelines []*NewPipeline `json:"pipelines"`

	// It's the ref of the Service Catalog, like 'cycloidio:stack-magento'
	// Required: true
	ServiceCatalogRef *string `json:"service_catalog_ref"`
}

// Validate validates this new project
func (m *NewProject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigRepositoryCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInputs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePipelines(formats); err != nil {
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

func (m *NewProject) validateCanonical(formats strfmt.Registry) error {

	if swag.IsZero(m.Canonical) { // not required
		return nil
	}

	if err := validate.MinLength("canonical", "body", string(m.Canonical), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("canonical", "body", string(m.Canonical), 100); err != nil {
		return err
	}

	if err := validate.Pattern("canonical", "body", string(m.Canonical), `(^[a-z0-9]+(([a-z0-9\-_]+)?[a-z0-9]+)?$)`); err != nil {
		return err
	}

	return nil
}

func (m *NewProject) validateConfigRepositoryCanonical(formats strfmt.Registry) error {

	if err := validate.Required("config_repository_canonical", "body", m.ConfigRepositoryCanonical); err != nil {
		return err
	}

	if err := validate.MinLength("config_repository_canonical", "body", string(*m.ConfigRepositoryCanonical), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("config_repository_canonical", "body", string(*m.ConfigRepositoryCanonical), 100); err != nil {
		return err
	}

	if err := validate.Pattern("config_repository_canonical", "body", string(*m.ConfigRepositoryCanonical), `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *NewProject) validateInputs(formats strfmt.Registry) error {

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

func (m *NewProject) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *NewProject) validatePipelines(formats strfmt.Registry) error {

	if err := validate.Required("pipelines", "body", m.Pipelines); err != nil {
		return err
	}

	iPipelinesSize := int64(len(m.Pipelines))

	if err := validate.MinItems("pipelines", "body", iPipelinesSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Pipelines); i++ {
		if swag.IsZero(m.Pipelines[i]) { // not required
			continue
		}

		if m.Pipelines[i] != nil {
			if err := m.Pipelines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pipelines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NewProject) validateServiceCatalogRef(formats strfmt.Registry) error {

	if err := validate.Required("service_catalog_ref", "body", m.ServiceCatalogRef); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewProject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewProject) UnmarshalBinary(b []byte) error {
	var res NewProject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
