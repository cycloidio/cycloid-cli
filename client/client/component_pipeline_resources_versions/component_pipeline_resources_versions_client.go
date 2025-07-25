// Code generated by go-swagger; DO NOT EDIT.

package component_pipeline_resources_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new component pipeline resources versions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new component pipeline resources versions API client with basic auth credentials.
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

// New creates a new component pipeline resources versions API client with a bearer token for authentication.
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
Client for component pipeline resources versions API
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

// WithContentType allows the client to force the Content-Type header
// to negotiate a specific Consumer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithContentType(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ConsumesMediaTypes = []string{mime}
	}
}

// WithContentTypeApplicationJSON sets the Content-Type header to "application/json".
func WithContentTypeApplicationJSON(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/json"}
}

// WithContentTypeApplicationVndCycloidIoV1JSON sets the Content-Type header to "application/vnd.cycloid.io.v1+json".
func WithContentTypeApplicationVndCycloidIoV1JSON(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/vnd.cycloid.io.v1+json"}
}

// WithContentTypeApplicationxWwwFormUrlencoded sets the Content-Type header to "application/x-www-form-urlencoded".
func WithContentTypeApplicationxWwwFormUrlencoded(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/x-www-form-urlencoded"}
}

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
	DisableResourceVersion(params *DisableResourceVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DisableResourceVersionNoContent, error)

	EnableResourceVersion(params *EnableResourceVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EnableResourceVersionNoContent, error)

	GetBuildsWithVersionAsInput(params *GetBuildsWithVersionAsInputParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBuildsWithVersionAsInputOK, error)

	GetBuildsWithVersionAsOutput(params *GetBuildsWithVersionAsOutputParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBuildsWithVersionAsOutputOK, error)

	GetResourceVersions(params *GetResourceVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetResourceVersionsOK, error)

	PinResourceVersion(params *PinResourceVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PinResourceVersionNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DisableResourceVersion Disable a specified version of a resource
*/
func (a *Client) DisableResourceVersion(params *DisableResourceVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DisableResourceVersionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisableResourceVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "disableResourceVersion",
		Method:             "PUT",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/versions/{resource_version_id}/disable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DisableResourceVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*DisableResourceVersionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DisableResourceVersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EnableResourceVersion Enable a specified version of a resource
*/
func (a *Client) EnableResourceVersion(params *EnableResourceVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EnableResourceVersionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnableResourceVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "enableResourceVersion",
		Method:             "PUT",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/versions/{resource_version_id}/enable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EnableResourceVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*EnableResourceVersionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EnableResourceVersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetBuildsWithVersionAsInput Get builds that used a specified resource version as input.
*/
func (a *Client) GetBuildsWithVersionAsInput(params *GetBuildsWithVersionAsInputParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBuildsWithVersionAsInputOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBuildsWithVersionAsInputParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBuildsWithVersionAsInput",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/versions/{resource_version_id}/input_to",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBuildsWithVersionAsInputReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetBuildsWithVersionAsInputOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBuildsWithVersionAsInputDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetBuildsWithVersionAsOutput Get builds that used a specified resource version as output.
*/
func (a *Client) GetBuildsWithVersionAsOutput(params *GetBuildsWithVersionAsOutputParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBuildsWithVersionAsOutputOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBuildsWithVersionAsOutputParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBuildsWithVersionAsOutput",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/versions/{resource_version_id}/output_of",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBuildsWithVersionAsOutputReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetBuildsWithVersionAsOutputOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBuildsWithVersionAsOutputDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetResourceVersions Get versions of a pipeline's resource
*/
func (a *Client) GetResourceVersions(params *GetResourceVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetResourceVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetResourceVersionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getResourceVersions",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetResourceVersionsReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GetResourceVersionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetResourceVersionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PinResourceVersion Pin a specified version of a resource
*/
func (a *Client) PinResourceVersion(params *PinResourceVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PinResourceVersionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPinResourceVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pinResourceVersion",
		Method:             "PUT",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/versions/{resource_version_id}/pin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PinResourceVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*PinResourceVersionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PinResourceVersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
