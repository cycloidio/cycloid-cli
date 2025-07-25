// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new organizations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new organizations API client with basic auth credentials.
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

// New creates a new organizations API client with a bearer token for authentication.
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
Client for organizations API
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
	CanDo(params *CanDoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CanDoOK, error)

	CreateOrg(params *CreateOrgParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOrgOK, error)

	DeleteOrg(params *DeleteOrgParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrgNoContent, error)

	GetAncestors(params *GetAncestorsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAncestorsOK, error)

	GetEvents(params *GetEventsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEventsOK, error)

	GetEventsCount(params *GetEventsCountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEventsCountOK, error)

	GetEventsTags(params *GetEventsTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEventsTagsOK, error)

	GetOrg(params *GetOrgParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrgOK, error)

	GetOrgs(params *GetOrgsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrgsOK, error)

	GetRepoBranches(params *GetRepoBranchesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRepoBranchesOK, error)

	GetSummary(params *GetSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSummaryOK, error)

	SendEvent(params *SendEventParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SendEventOK, error)

	UpdateOrg(params *UpdateOrgParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOrgOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CanDo Checks if the JWT can do the action
*/
func (a *Client) CanDo(params *CanDoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CanDoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCanDoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "canDo",
		Method:             "POST",
		PathPattern:        "/organizations/{organization_canonical}/can_do",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CanDoReader{formats: a.formats},
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
	success, ok := result.(*CanDoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CanDoDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateOrg Create a new organization, making the authenticated user the owner of it.
*/
func (a *Client) CreateOrg(params *CreateOrgParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOrgOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOrgParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createOrg",
		Method:             "POST",
		PathPattern:        "/organizations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateOrgReader{formats: a.formats},
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
	success, ok := result.(*CreateOrgOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateOrgDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteOrg Delete the organization.
*/
func (a *Client) DeleteOrg(params *DeleteOrgParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrgNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrgParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteOrg",
		Method:             "DELETE",
		PathPattern:        "/organizations/{organization_canonical}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOrgReader{formats: a.formats},
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
	success, ok := result.(*DeleteOrgNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteOrgDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetAncestors Get all the ancestors between the Organization and the User with the shortest path.
*/
func (a *Client) GetAncestors(params *GetAncestorsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAncestorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAncestorsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAncestors",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/ancestors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAncestorsReader{formats: a.formats},
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
	success, ok := result.(*GetAncestorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAncestorsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
		GetEvents Retrieve the list of events which has been registered on the organization. The events to request can be filtered using Unix timestamps in milliseconds (begin and end timestamps range), the event type and severity; when more than one are applied then they are applied with a logical AND.
	  - The Unix timestamps must always be specified, the rest of the filters
	    are not mandatory.
*/
func (a *Client) GetEvents(params *GetEventsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEventsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEvents",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEventsReader{formats: a.formats},
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
	success, ok := result.(*GetEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetEventsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetEventsCount Retrieve a count of events for given parameters
*/
func (a *Client) GetEventsCount(params *GetEventsCountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEventsCountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEventsCountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEventsCount",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/events/count",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEventsCountReader{formats: a.formats},
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
	success, ok := result.(*GetEventsCountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetEventsCountDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetEventsTags Retrieve the list of tags and set of values for all the events of the organization.
*/
func (a *Client) GetEventsTags(params *GetEventsTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEventsTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEventsTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEventsTags",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/events/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEventsTagsReader{formats: a.formats},
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
	success, ok := result.(*GetEventsTagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetEventsTagsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetOrg Get the information of the organization.
*/
func (a *Client) GetOrg(params *GetOrgParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrgOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrgParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrg",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrgReader{formats: a.formats},
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
	success, ok := result.(*GetOrgOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetOrgDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetOrgs Get the organizations that the authenticated user has access.
*/
func (a *Client) GetOrgs(params *GetOrgsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrgsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrgsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrgs",
		Method:             "GET",
		PathPattern:        "/organizations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrgsReader{formats: a.formats},
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
	success, ok := result.(*GetOrgsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetOrgsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	GetRepoBranches Return all the branches of repository. If the repository is empty then an

empty list will be returned.
*/
func (a *Client) GetRepoBranches(params *GetRepoBranchesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRepoBranchesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRepoBranchesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRepoBranches",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/branches",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRepoBranchesReader{formats: a.formats},
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
	success, ok := result.(*GetRepoBranchesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRepoBranchesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetSummary Get the summary of the organization
*/
func (a *Client) GetSummary(params *GetSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSummaryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSummary",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/summary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSummaryReader{formats: a.formats},
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
	success, ok := result.(*GetSummaryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSummary: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SendEvent Send a event on the organization to be registered.
*/
func (a *Client) SendEvent(params *SendEventParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SendEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendEventParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "sendEvent",
		Method:             "POST",
		PathPattern:        "/organizations/{organization_canonical}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SendEventReader{formats: a.formats},
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
	success, ok := result.(*SendEventOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sendEvent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateOrg Update the information of the organization.
*/
func (a *Client) UpdateOrg(params *UpdateOrgParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOrgOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrgParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateOrg",
		Method:             "PUT",
		PathPattern:        "/organizations/{organization_canonical}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateOrgReader{formats: a.formats},
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
	success, ok := result.(*UpdateOrgOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateOrgDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
