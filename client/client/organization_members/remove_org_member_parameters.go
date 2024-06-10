// Code generated by go-swagger; DO NOT EDIT.

package organization_members

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
)

// NewRemoveOrgMemberParams creates a new RemoveOrgMemberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveOrgMemberParams() *RemoveOrgMemberParams {
	return &RemoveOrgMemberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveOrgMemberParamsWithTimeout creates a new RemoveOrgMemberParams object
// with the ability to set a timeout on a request.
func NewRemoveOrgMemberParamsWithTimeout(timeout time.Duration) *RemoveOrgMemberParams {
	return &RemoveOrgMemberParams{
		timeout: timeout,
	}
}

// NewRemoveOrgMemberParamsWithContext creates a new RemoveOrgMemberParams object
// with the ability to set a context for a request.
func NewRemoveOrgMemberParamsWithContext(ctx context.Context) *RemoveOrgMemberParams {
	return &RemoveOrgMemberParams{
		Context: ctx,
	}
}

// NewRemoveOrgMemberParamsWithHTTPClient creates a new RemoveOrgMemberParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveOrgMemberParamsWithHTTPClient(client *http.Client) *RemoveOrgMemberParams {
	return &RemoveOrgMemberParams{
		HTTPClient: client,
	}
}

/*
RemoveOrgMemberParams contains all the parameters to send to the API endpoint

	for the remove org member operation.

	Typically these are written to a http.Request.
*/
type RemoveOrgMemberParams struct {

	/* OrganizationCanonical.

	   A canonical of an organization.
	*/
	OrganizationCanonical string

	/* Username.

	   A username
	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove org member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveOrgMemberParams) WithDefaults() *RemoveOrgMemberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove org member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveOrgMemberParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove org member params
func (o *RemoveOrgMemberParams) WithTimeout(timeout time.Duration) *RemoveOrgMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove org member params
func (o *RemoveOrgMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove org member params
func (o *RemoveOrgMemberParams) WithContext(ctx context.Context) *RemoveOrgMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove org member params
func (o *RemoveOrgMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove org member params
func (o *RemoveOrgMemberParams) WithHTTPClient(client *http.Client) *RemoveOrgMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove org member params
func (o *RemoveOrgMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationCanonical adds the organizationCanonical to the remove org member params
func (o *RemoveOrgMemberParams) WithOrganizationCanonical(organizationCanonical string) *RemoveOrgMemberParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the remove org member params
func (o *RemoveOrgMemberParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithUsername adds the username to the remove org member params
func (o *RemoveOrgMemberParams) WithUsername(username string) *RemoveOrgMemberParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the remove org member params
func (o *RemoveOrgMemberParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveOrgMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
