// Code generated by go-swagger; DO NOT EDIT.

package component_pipeline_resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new component pipeline resources API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new component pipeline resources API client with basic auth credentials.
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

// New creates a new component pipeline resources API client with a bearer token for authentication.
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
Client for component pipeline resources API
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
	CheckResource(params *CheckResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CheckResourceOK, error)

	GetPipelineResources(params *GetPipelineResourcesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPipelineResourcesOK, error)

	GetResource(params *GetResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetResourceOK, error)

	ResourceSetPinComment(params *ResourceSetPinCommentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResourceSetPinCommentNoContent, error)

	UnpinResource(params *UnpinResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UnpinResourceNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CheckResource Trigger a check for a specified resource.
*/
func (a *Client) CheckResource(params *CheckResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CheckResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckResourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "checkResource",
		Method:             "POST",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/check",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CheckResourceReader{formats: a.formats},
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
	success, ok := result.(*CheckResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CheckResourceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetPipelineResources Get the resources of the pipeline.
*/
func (a *Client) GetPipelineResources(params *GetPipelineResourcesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPipelineResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPipelineResourcesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPipelineResources",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPipelineResourcesReader{formats: a.formats},
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
	success, ok := result.(*GetPipelineResourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPipelineResourcesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetResource Get a specific resource of the pipeline.
*/
func (a *Client) GetResource(params *GetResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetResourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getResource",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetResourceReader{formats: a.formats},
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
	success, ok := result.(*GetResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetResourceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ResourceSetPinComment Set pin comment on a specified resource
*/
func (a *Client) ResourceSetPinComment(params *ResourceSetPinCommentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResourceSetPinCommentNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResourceSetPinCommentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "resourceSetPinComment",
		Method:             "PUT",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/pin_comment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResourceSetPinCommentReader{formats: a.formats},
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
	success, ok := result.(*ResourceSetPinCommentNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ResourceSetPinCommentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UnpinResource Unpin a specified resource
*/
func (a *Client) UnpinResource(params *UnpinResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UnpinResourceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnpinResourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "unpinResource",
		Method:             "PUT",
		PathPattern:        "/organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources/{resource_name}/unpin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UnpinResourceReader{formats: a.formats},
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
	success, ok := result.(*UnpinResourceNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UnpinResourceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
