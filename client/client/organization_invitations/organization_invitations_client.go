// Code generated by go-swagger; DO NOT EDIT.

package organization_invitations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new organization invitations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for organization invitations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteInvitation Delete an Organization's Invitation.
*/
func (a *Client) DeleteInvitation(params *DeleteInvitationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteInvitationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInvitationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteInvitation",
		Method:             "DELETE",
		PathPattern:        "/organizations/{organization_canonical}/invitations/{invitation_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteInvitationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteInvitationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteInvitationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetInvitations Get list of the Organization's Invitations.
*/
func (a *Client) GetInvitations(params *GetInvitationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetInvitationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInvitationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getInvitations",
		Method:             "GET",
		PathPattern:        "/organizations/{organization_canonical}/invitations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInvitationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInvitationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetInvitationsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetPendingInvitation Get the email address used for the pending invitation
*/
func (a *Client) GetPendingInvitation(params *GetPendingInvitationParams) (*GetPendingInvitationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPendingInvitationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPendingInvitation",
		Method:             "GET",
		PathPattern:        "/invitations/{verification_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPendingInvitationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPendingInvitationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPendingInvitationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ResendInvitation Resend the email containing the verification token to accept the Invitation.
*/
func (a *Client) ResendInvitation(params *ResendInvitationParams) (*ResendInvitationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResendInvitationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "resendInvitation",
		Method:             "PUT",
		PathPattern:        "/organizations/{organization_canonical}/invitations/{invitation_id}/resend",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResendInvitationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ResendInvitationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ResendInvitationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
