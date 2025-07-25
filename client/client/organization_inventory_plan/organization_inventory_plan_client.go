// Code generated by go-swagger; DO NOT EDIT.

package organization_inventory_plan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new organization inventory plan API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new organization inventory plan API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new organization inventory plan API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for organization inventory plan API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// This client is generated with a few options you might find useful for your swagger spec.
//
// Feel free to add you own set of options.

// WithAccept allows the client to force the Accept header
// to negotiate a specific Producer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithAccept(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ProducesMediaTypes = []string{mime}
	}
}

// WithAcceptApplicationJSON sets the Accept header to "application/json".
func WithAcceptApplicationJSON(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/json"}
}

// WithAcceptApplicationVndCycloidIoV1JSON sets the Accept header to "application/vnd.cycloid.io.v1+json".
func WithAcceptApplicationVndCycloidIoV1JSON(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/vnd.cycloid.io.v1+json"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetInventoryPlan(params *GetInventoryPlanParams, opts ...ClientOption) (*GetInventoryPlanOK, error)

	SetInventoryPlan(params *SetInventoryPlanParams, opts ...ClientOption) (*SetInventoryPlanNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetInventoryPlan Get the terraform plan of the inventory. It uses '-' as the Concourse resource that uses it automatically adds '-plan' on it's implementation to push the plan. So to simplify the configuration we use '-' also.
*/
func (a *Client) GetInventoryPlan(params *GetInventoryPlanParams, opts ...ClientOption) (*GetInventoryPlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryPlanParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getInventoryPlan",
		Method:             "GET",
		PathPattern:        "/inventory-plan",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryPlanReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInventoryPlanOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetInventoryPlanDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SetInventoryPlan Create or replace an Inventory plan
*/
func (a *Client) SetInventoryPlan(params *SetInventoryPlanParams, opts ...ClientOption) (*SetInventoryPlanNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetInventoryPlanParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "setInventoryPlan",
		Method:             "POST",
		PathPattern:        "/inventory-plan",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetInventoryPlanReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetInventoryPlanNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SetInventoryPlanDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
