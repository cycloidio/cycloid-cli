// Code generated by go-swagger; DO NOT EDIT.

package organization_external_backends

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new organization external backends API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for organization external backends API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateExternalBackend Save information about the external backend
*/
func (a *Client) CreateExternalBackend(params *CreateExternalBackendParams, authInfo runtime.ClientAuthInfoWriter) (*CreateExternalBackendOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateExternalBackendParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createExternalBackend",
		Method:             "POST",
		PathPattern:        "/organizations/{organization_canonical}/external_backends",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateExternalBackendReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateExternalBackendOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createExternalBackend: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteExternalBackend delete an External Backend
*/
func (a *Client) DeleteExternalBackend(params *DeleteExternalBackendParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteExternalBackendNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteExternalBackendParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteExternalBackend",
		Method:             "DELETE",
		PathPattern:        "/organizations/{organization_canonical}/external_backends/{external_backend_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteExternalBackendReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteExternalBackendNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteExternalBackendDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetExternalBackend Get the external backend
*/
func (a *Client) GetExternalBackend(params *GetExternalBackendParams, authInfo runtime.ClientAuthInfoWriter) (*GetExternalBackendOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExternalBackendParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getExternalBackend",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/external_backends/{external_backend_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetExternalBackendReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetExternalBackendOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetExternalBackendDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetExternalBackends Get the list of organization external backends
*/
func (a *Client) GetExternalBackends(params *GetExternalBackendsParams, authInfo runtime.ClientAuthInfoWriter) (*GetExternalBackendsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExternalBackendsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getExternalBackends",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/external_backends",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetExternalBackendsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetExternalBackendsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetExternalBackendsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateExternalBackend Update an External Backend
*/
func (a *Client) UpdateExternalBackend(params *UpdateExternalBackendParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateExternalBackendOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateExternalBackendParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateExternalBackend",
		Method:             "PUT",
		PathPattern:        "/organizations/{organization_canonical}/external_backends/{external_backend_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateExternalBackendReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateExternalBackendOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateExternalBackendDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
