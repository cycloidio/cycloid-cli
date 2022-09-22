// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"
)

// ExternalBackendConfiguration external backend configuration
// swagger:discriminator ExternalBackendConfiguration engine
type ExternalBackendConfiguration interface {
	runtime.Validatable

	// engine
	// Required: true
	Engine() string
	SetEngine(string)
}

type externalBackendConfiguration struct {
	engineField string
}

// Engine gets the engine of this polymorphic type
func (m *externalBackendConfiguration) Engine() string {
	return "ExternalBackendConfiguration"
}

// SetEngine sets the engine of this polymorphic type
func (m *externalBackendConfiguration) SetEngine(val string) {

}

// UnmarshalExternalBackendConfigurationSlice unmarshals polymorphic slices of ExternalBackendConfiguration
func UnmarshalExternalBackendConfigurationSlice(reader io.Reader, consumer runtime.Consumer) ([]ExternalBackendConfiguration, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []ExternalBackendConfiguration
	for _, element := range elements {
		obj, err := unmarshalExternalBackendConfiguration(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalExternalBackendConfiguration unmarshals polymorphic ExternalBackendConfiguration
func UnmarshalExternalBackendConfiguration(reader io.Reader, consumer runtime.Consumer) (ExternalBackendConfiguration, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalExternalBackendConfiguration(data, consumer)
}

func unmarshalExternalBackendConfiguration(data []byte, consumer runtime.Consumer) (ExternalBackendConfiguration, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the engine property.
	var getType struct {
		Engine string `json:"engine"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("engine", "body", getType.Engine); err != nil {
		return nil, err
	}

	// The value of engine is used to determine which type to create and unmarshal the data into
	switch getType.Engine {
	case "AWSCloudWatchLogs":
		var result AWSCloudWatchLogs
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "AWSRemoteTFState":
		var result AWSRemoteTFState
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "AWSStorage":
		var result AWSStorage
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "AzureRemoteTFState":
		var result AzureRemoteTFState
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "AzureStorage":
		var result AzureStorage
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "ElasticsearchLogs":
		var result ElasticsearchLogs
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "ExternalBackendConfiguration":
		var result externalBackendConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "GCPCostStorage":
		var result GCPCostStorage
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "GCPRemoteTFState":
		var result GCPRemoteTFState
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "GCPStorage":
		var result GCPStorage
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "SwiftRemoteTFState":
		var result SwiftRemoteTFState
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "SwiftStorage":
		var result SwiftStorage
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "VMwareVsphere":
		var result VMwareVsphere
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid engine value: %q", getType.Engine)

}

// Validate validates this external backend configuration
func (m *externalBackendConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}
