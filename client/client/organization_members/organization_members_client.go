// Code generated by go-swagger; DO NOT EDIT.

package organization_members

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new organization members API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for organization members API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetOrgMember Get the information of a member of the organization.
*/
func (a *Client) GetOrgMember(params *GetOrgMemberParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrgMemberOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrgMemberParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrgMember",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/members/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrgMemberReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrgMemberOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetOrgMemberDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetOrgMembers Get the members of an organization.
*/
func (a *Client) GetOrgMembers(params *GetOrgMembersParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrgMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrgMembersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrgMembers",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrgMembersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrgMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetOrgMembersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
InviteUserToOrgMember Invite a user to be a member of the organization.
*/
func (a *Client) InviteUserToOrgMember(params *InviteUserToOrgMemberParams, authInfo runtime.ClientAuthInfoWriter) (*InviteUserToOrgMemberNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInviteUserToOrgMemberParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "inviteUserToOrgMember",
		Method:             "PUT",
		PathPattern:        "/organizations/{organization_canonical}/members-invitations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InviteUserToOrgMemberReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InviteUserToOrgMemberNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*InviteUserToOrgMemberDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RemoveOrgMember Remove a member of the organization.
*/
func (a *Client) RemoveOrgMember(params *RemoveOrgMemberParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveOrgMemberNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveOrgMemberParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeOrgMember",
		Method:             "DELETE",
		PathPattern:        "/organizations/{organization_canonical}/members/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RemoveOrgMemberReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveOrgMemberNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RemoveOrgMemberDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateOrgMember Update member of the organization.
*/
func (a *Client) UpdateOrgMember(params *UpdateOrgMemberParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOrgMemberOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrgMemberParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateOrgMember",
		Method:             "PUT",
		PathPattern:        "/organizations/{organization_canonical}/members/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateOrgMemberReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateOrgMemberOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateOrgMemberDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
