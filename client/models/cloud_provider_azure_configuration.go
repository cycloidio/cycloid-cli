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

// CloudProviderAzureConfiguration Representation of Azure configuration
//
// swagger:model CloudProviderAzureConfiguration
type CloudProviderAzureConfiguration struct {

	// The Azure environment of the configuration
	//
	// Required: true
	Environment *string `json:"environment"`

	// The Azure resource group name of the configuration
	//
	// Required: true
	ResourceGroupName *string `json:"resource_group_name"`
}

// Type gets the type of this subtype
func (m *CloudProviderAzureConfiguration) Type() string {
	return "CloudProviderAzureConfiguration"
}

// SetType sets the type of this subtype
func (m *CloudProviderAzureConfiguration) SetType(val string) {

}

// Environment gets the environment of this subtype

// ResourceGroupName gets the resource group name of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *CloudProviderAzureConfiguration) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The Azure environment of the configuration
		//
		// Required: true
		Environment *string `json:"environment"`

		// The Azure resource group name of the configuration
		//
		// Required: true
		ResourceGroupName *string `json:"resource_group_name"`
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

	var result CloudProviderAzureConfiguration

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Environment = data.Environment

	result.ResourceGroupName = data.ResourceGroupName

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m CloudProviderAzureConfiguration) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The Azure environment of the configuration
		//
		// Required: true
		Environment *string `json:"environment"`

		// The Azure resource group name of the configuration
		//
		// Required: true
		ResourceGroupName *string `json:"resource_group_name"`
	}{

		Environment: m.Environment,

		ResourceGroupName: m.ResourceGroupName,
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

// Validate validates this cloud provider azure configuration
func (m *CloudProviderAzureConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceGroupName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudProviderAzureConfiguration) validateEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("environment", "body", m.Environment); err != nil {
		return err
	}

	return nil
}

func (m *CloudProviderAzureConfiguration) validateResourceGroupName(formats strfmt.Registry) error {

	if err := validate.Required("resource_group_name", "body", m.ResourceGroupName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudProviderAzureConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudProviderAzureConfiguration) UnmarshalBinary(b []byte) error {
	var res CloudProviderAzureConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
