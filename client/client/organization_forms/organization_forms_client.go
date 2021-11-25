// Code generated by go-swagger; DO NOT EDIT.

package organization_forms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new organization forms API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for organization forms API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateFormsConfig Generate a set of configs based on the forms inputs
*/
func (a *Client) CreateFormsConfig(params *CreateFormsConfigParams, authInfo runtime.ClientAuthInfoWriter) (*CreateFormsConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateFormsConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createFormsConfig",
		Method:             "POST",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/forms/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateFormsConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateFormsConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateFormsConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ValidateFormsFile Validate a forms file definition
*/
func (a *Client) ValidateFormsFile(params *ValidateFormsFileParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateFormsFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateFormsFileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateFormsFile",
		Method:             "POST",
		PathPattern:        "/organizations/{organization_canonical}/forms/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateFormsFileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ValidateFormsFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ValidateFormsFileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
