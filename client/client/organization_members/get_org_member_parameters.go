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

// NewGetOrgMemberParams creates a new GetOrgMemberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrgMemberParams() *GetOrgMemberParams {
	return &GetOrgMemberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrgMemberParamsWithTimeout creates a new GetOrgMemberParams object
// with the ability to set a timeout on a request.
func NewGetOrgMemberParamsWithTimeout(timeout time.Duration) *GetOrgMemberParams {
	return &GetOrgMemberParams{
		timeout: timeout,
	}
}

// NewGetOrgMemberParamsWithContext creates a new GetOrgMemberParams object
// with the ability to set a context for a request.
func NewGetOrgMemberParamsWithContext(ctx context.Context) *GetOrgMemberParams {
	return &GetOrgMemberParams{
		Context: ctx,
	}
}

// NewGetOrgMemberParamsWithHTTPClient creates a new GetOrgMemberParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrgMemberParamsWithHTTPClient(client *http.Client) *GetOrgMemberParams {
	return &GetOrgMemberParams{
		HTTPClient: client,
	}
}

/*
GetOrgMemberParams contains all the parameters to send to the API endpoint

	for the get org member operation.

	Typically these are written to a http.Request.
*/
type GetOrgMemberParams struct {

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

// WithDefaults hydrates default values in the get org member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgMemberParams) WithDefaults() *GetOrgMemberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get org member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgMemberParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get org member params
func (o *GetOrgMemberParams) WithTimeout(timeout time.Duration) *GetOrgMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get org member params
func (o *GetOrgMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get org member params
func (o *GetOrgMemberParams) WithContext(ctx context.Context) *GetOrgMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get org member params
func (o *GetOrgMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get org member params
func (o *GetOrgMemberParams) WithHTTPClient(client *http.Client) *GetOrgMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get org member params
func (o *GetOrgMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationCanonical adds the organizationCanonical to the get org member params
func (o *GetOrgMemberParams) WithOrganizationCanonical(organizationCanonical string) *GetOrgMemberParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get org member params
func (o *GetOrgMemberParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithUsername adds the username to the get org member params
func (o *GetOrgMemberParams) WithUsername(username string) *GetOrgMemberParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the get org member params
func (o *GetOrgMemberParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrgMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
