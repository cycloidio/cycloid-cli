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

// AuthConfigOAuth AppConfigAuthOAuth
// swagger:discriminator AuthConfigOAuth type
type AuthConfigOAuth interface {
	runtime.Validatable

	// ID of the OAuth client.
	// Required: true
	ClientID() *string
	SetClientID(*string)

	// Name of the OAuth identity provider.
	// Required: true
	Provider() *string
	SetProvider(*string)

	// Type is used as discriminator for different OAuth identity provider definitions.
	// Required: true
	Type() string
	SetType(string)
}

type authConfigOAuth struct {
	clientIdField *string

	providerField *string

	typeField string
}

// ClientID gets the client id of this polymorphic type
func (m *authConfigOAuth) ClientID() *string {
	return m.clientIdField
}

// SetClientID sets the client id of this polymorphic type
func (m *authConfigOAuth) SetClientID(val *string) {
	m.clientIdField = val
}

// Provider gets the provider of this polymorphic type
func (m *authConfigOAuth) Provider() *string {
	return m.providerField
}

// SetProvider sets the provider of this polymorphic type
func (m *authConfigOAuth) SetProvider(val *string) {
	m.providerField = val
}

// Type gets the type of this polymorphic type
func (m *authConfigOAuth) Type() string {
	return "AuthConfigOAuth"
}

// SetType sets the type of this polymorphic type
func (m *authConfigOAuth) SetType(val string) {

}

// UnmarshalAuthConfigOAuthSlice unmarshals polymorphic slices of AuthConfigOAuth
func UnmarshalAuthConfigOAuthSlice(reader io.Reader, consumer runtime.Consumer) ([]AuthConfigOAuth, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []AuthConfigOAuth
	for _, element := range elements {
		obj, err := unmarshalAuthConfigOAuth(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalAuthConfigOAuth unmarshals polymorphic AuthConfigOAuth
func UnmarshalAuthConfigOAuth(reader io.Reader, consumer runtime.Consumer) (AuthConfigOAuth, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalAuthConfigOAuth(data, consumer)
}

func unmarshalAuthConfigOAuth(data []byte, consumer runtime.Consumer) (AuthConfigOAuth, error) {
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
	case "AuthConfigOAuth":
		var result authConfigOAuth
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "AzureADAuthConfig":
		var result AzureADAuthConfig
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "GitHubOAuthConfig":
		var result GitHubOAuthConfig
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "GoogleOAuthConfig":
		var result GoogleOAuthConfig
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid type value: %q", getType.Type)

}

// Validate validates this auth config o auth
func (m *authConfigOAuth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *authConfigOAuth) validateClientID(formats strfmt.Registry) error {

	if err := validate.Required("client_id", "body", m.ClientID()); err != nil {
		return err
	}

	return nil
}

func (m *authConfigOAuth) validateProvider(formats strfmt.Registry) error {

	if err := validate.Required("provider", "body", m.Provider()); err != nil {
		return err
	}

	return nil
}
