// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CloudProviderGCPConfiguration Representation of Google configuration
//
// swagger:model CloudProviderGCPConfiguration
type CloudProviderGCPConfiguration struct {

	// The Google project where the resource exists
	//
	// Required: true
	Project *string `json:"project"`

	// The Google region where the resource exists
	//
	// Required: true
	Region *string `json:"region"`
}

// Type gets the type of this subtype
func (m *CloudProviderGCPConfiguration) Type() string {
	return "CloudProviderGCPConfiguration"
}

// SetType sets the type of this subtype
func (m *CloudProviderGCPConfiguration) SetType(val string) {

}

// Project gets the project of this subtype

// Region gets the region of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *CloudProviderGCPConfiguration) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The Google project where the resource exists
		//
		// Required: true
		Project *string `json:"project"`

		// The Google region where the resource exists
		//
		// Required: true
		Region *string `json:"region"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Type string `json:"type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result CloudProviderGCPConfiguration

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Project = data.Project

	result.Region = data.Region

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m CloudProviderGCPConfiguration) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The Google project where the resource exists
		//
		// Required: true
		Project *string `json:"project"`

		// The Google region where the resource exists
		//
		// Required: true
		Region *string `json:"region"`
	}{

		Project: m.Project,

		Region: m.Region,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Type string `json:"type"`
	}{

		Type: m.Type(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this cloud provider g c p configuration
func (m *CloudProviderGCPConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudProviderGCPConfiguration) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("project", "body", m.Project); err != nil {
		return err
	}

	return nil
}

func (m *CloudProviderGCPConfiguration) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudProviderGCPConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudProviderGCPConfiguration) UnmarshalBinary(b []byte) error {
	var res CloudProviderGCPConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}