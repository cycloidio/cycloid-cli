// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateCloudCostManagementTagMapping Create or Update CloudCostManagementTagMapping
//
// # Create or Update a Cloud Cost Management tag mapping for projects and environments
//
// swagger:model UpdateCloudCostManagementTagMapping
type UpdateCloudCostManagementTagMapping struct {

	// environment regex
	EnvironmentRegex string `json:"environment_regex,omitempty"`

	// environment tags
	EnvironmentTags []string `json:"environment_tags"`

	// project regex
	ProjectRegex string `json:"project_regex,omitempty"`

	// project tags
	ProjectTags []string `json:"project_tags"`
}

// Validate validates this update cloud cost management tag mapping
func (m *UpdateCloudCostManagementTagMapping) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update cloud cost management tag mapping based on context it is used
func (m *UpdateCloudCostManagementTagMapping) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateCloudCostManagementTagMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateCloudCostManagementTagMapping) UnmarshalBinary(b []byte) error {
	var res UpdateCloudCostManagementTagMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
