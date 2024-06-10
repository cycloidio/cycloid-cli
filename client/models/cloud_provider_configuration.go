// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// CloudProviderConfiguration cloud provider configuration
//
// swagger:discriminator CloudProviderConfiguration type
type CloudProviderConfiguration interface {
	runtime.Validatable
	runtime.ContextValidatable

	// type
	// Required: true
	Type() string
	SetType(string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type cloudProviderConfiguration struct {
	typeField string
}

// Type gets the type of this polymorphic type
func (m *cloudProviderConfiguration) Type() string {
	return "CloudProviderConfiguration"
}

// SetType sets the type of this polymorphic type
func (m *cloudProviderConfiguration) SetType(val string) {
}

// UnmarshalCloudProviderConfigurationSlice unmarshals polymorphic slices of CloudProviderConfiguration
func UnmarshalCloudProviderConfigurationSlice(reader io.Reader, consumer runtime.Consumer) ([]CloudProviderConfiguration, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []CloudProviderConfiguration
	for _, element := range elements {
		obj, err := unmarshalCloudProviderConfiguration(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalCloudProviderConfiguration unmarshals polymorphic CloudProviderConfiguration
func UnmarshalCloudProviderConfiguration(reader io.Reader, consumer runtime.Consumer) (CloudProviderConfiguration, error) {
	// we need to read this twice, so first into a buffer
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalCloudProviderConfiguration(data, consumer)
}

func unmarshalCloudProviderConfiguration(data []byte, consumer runtime.Consumer) (CloudProviderConfiguration, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var getType struct {
		Type string `json:"type"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("type", "body", getType.Type); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch getType.Type {
	case "CloudProviderAWSConfiguration":
		var result CloudProviderAWSConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "CloudProviderAzureConfiguration":
		var result CloudProviderAzureConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "CloudProviderConfiguration":
		var result cloudProviderConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "CloudProviderGCPConfiguration":
		var result CloudProviderGCPConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "CloudProviderVMWareVSphereConfiguration":
		var result CloudProviderVMWareVSphereConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid type value: %q", getType.Type)
}

// Validate validates this cloud provider configuration
func (m *cloudProviderConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cloud provider configuration based on context it is used
func (m *cloudProviderConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
