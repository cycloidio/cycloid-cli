// Code generated by go-swagger; DO NOT EDIT.

package organization_quotas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new organization quotas API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new organization quotas API client with basic auth credentials.
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

// New creates a new organization quotas API client with a bearer token for authentication.
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
Client for organization quotas API
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
	CreateQuota(params *CreateQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateQuotaOK, error)

	DeleteQuota(params *DeleteQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteQuotaNoContent, error)

	GetQuota(params *GetQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetQuotaOK, error)

	ListQuotaConsumptions(params *ListQuotaConsumptionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListQuotaConsumptionsOK, error)

	ListQuotas(params *ListQuotasParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListQuotasOK, error)

	UpdateQuota(params *UpdateQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateQuotaOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateQuota Create a new quota available in the organization.
*/
func (a *Client) CreateQuota(params *CreateQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateQuotaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateQuotaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createQuota",
		Method:             "POST",
		PathPattern:        "/organizations/{organization_canonical}/quotas",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateQuotaReader{formats: a.formats},
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
	success, ok := result.(*CreateQuotaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateQuotaDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteQuota Delete an existing quota in the organization.
*/
func (a *Client) DeleteQuota(params *DeleteQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteQuotaNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteQuotaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteQuota",
		Method:             "DELETE",
		PathPattern:        "/organizations/{organization_canonical}/quotas/{quota_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteQuotaReader{formats: a.formats},
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
	success, ok := result.(*DeleteQuotaNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteQuotaDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetQuota Get the quota available in the organization with an canonical
*/
func (a *Client) GetQuota(params *GetQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetQuotaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetQuotaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getQuota",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/quotas/{quota_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetQuotaReader{formats: a.formats},
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
	success, ok := result.(*GetQuotaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetQuotaDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListQuotaConsumptions Get the quota consumptions by project and environment
*/
func (a *Client) ListQuotaConsumptions(params *ListQuotaConsumptionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListQuotaConsumptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListQuotaConsumptionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listQuotaConsumptions",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/quotas/{quota_id}/consumptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListQuotaConsumptionsReader{formats: a.formats},
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
	success, ok := result.(*ListQuotaConsumptionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListQuotaConsumptionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListQuotas List of quotas available in the organization.
*/
func (a *Client) ListQuotas(params *ListQuotasParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListQuotasOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListQuotasParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listQuotas",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/quotas",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListQuotasReader{formats: a.formats},
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
	success, ok := result.(*ListQuotasOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListQuotasDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateQuota Update an existing quota in the organization.
*/
func (a *Client) UpdateQuota(params *UpdateQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateQuotaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateQuotaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateQuota",
		Method:             "PUT",
		PathPattern:        "/organizations/{organization_canonical}/quotas/{quota_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateQuotaReader{formats: a.formats},
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
	success, ok := result.(*UpdateQuotaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateQuotaDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
