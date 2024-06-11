// Code generated by go-swagger; DO NOT EDIT.

package organization_invitations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteInvitationParams creates a new DeleteInvitationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteInvitationParams() *DeleteInvitationParams {
	return &DeleteInvitationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInvitationParamsWithTimeout creates a new DeleteInvitationParams object
// with the ability to set a timeout on a request.
func NewDeleteInvitationParamsWithTimeout(timeout time.Duration) *DeleteInvitationParams {
	return &DeleteInvitationParams{
		timeout: timeout,
	}
}

// NewDeleteInvitationParamsWithContext creates a new DeleteInvitationParams object
// with the ability to set a context for a request.
func NewDeleteInvitationParamsWithContext(ctx context.Context) *DeleteInvitationParams {
	return &DeleteInvitationParams{
		Context: ctx,
	}
}

// NewDeleteInvitationParamsWithHTTPClient creates a new DeleteInvitationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteInvitationParamsWithHTTPClient(client *http.Client) *DeleteInvitationParams {
	return &DeleteInvitationParams{
		HTTPClient: client,
	}
}

/*
DeleteInvitationParams contains all the parameters to send to the API endpoint

	for the delete invitation operation.

	Typically these are written to a http.Request.
*/
type DeleteInvitationParams struct {

	/* InvitationID.

	   Organization Invitation id.

	   Format: uint32
	*/
	InvitationID uint32

	/* OrganizationCanonical.

	   A canonical of an organization.
	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteInvitationParams) WithDefaults() *DeleteInvitationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteInvitationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete invitation params
func (o *DeleteInvitationParams) WithTimeout(timeout time.Duration) *DeleteInvitationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete invitation params
func (o *DeleteInvitationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete invitation params
func (o *DeleteInvitationParams) WithContext(ctx context.Context) *DeleteInvitationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete invitation params
func (o *DeleteInvitationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete invitation params
func (o *DeleteInvitationParams) WithHTTPClient(client *http.Client) *DeleteInvitationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete invitation params
func (o *DeleteInvitationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvitationID adds the invitationID to the delete invitation params
func (o *DeleteInvitationParams) WithInvitationID(invitationID uint32) *DeleteInvitationParams {
	o.SetInvitationID(invitationID)
	return o
}

// SetInvitationID adds the invitationId to the delete invitation params
func (o *DeleteInvitationParams) SetInvitationID(invitationID uint32) {
	o.InvitationID = invitationID
}

// WithOrganizationCanonical adds the organizationCanonical to the delete invitation params
func (o *DeleteInvitationParams) WithOrganizationCanonical(organizationCanonical string) *DeleteInvitationParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the delete invitation params
func (o *DeleteInvitationParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInvitationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param invitation_id
	if err := r.SetPathParam("invitation_id", swag.FormatUint32(o.InvitationID)); err != nil {
		return err
	}

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
