// Code generated by go-swagger; DO NOT EDIT.

package organization_teams

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

// NewGetTeamParams creates a new GetTeamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTeamParams() *GetTeamParams {
	return &GetTeamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTeamParamsWithTimeout creates a new GetTeamParams object
// with the ability to set a timeout on a request.
func NewGetTeamParamsWithTimeout(timeout time.Duration) *GetTeamParams {
	return &GetTeamParams{
		timeout: timeout,
	}
}

// NewGetTeamParamsWithContext creates a new GetTeamParams object
// with the ability to set a context for a request.
func NewGetTeamParamsWithContext(ctx context.Context) *GetTeamParams {
	return &GetTeamParams{
		Context: ctx,
	}
}

// NewGetTeamParamsWithHTTPClient creates a new GetTeamParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTeamParamsWithHTTPClient(client *http.Client) *GetTeamParams {
	return &GetTeamParams{
		HTTPClient: client,
	}
}

/*
GetTeamParams contains all the parameters to send to the API endpoint

	for the get team operation.

	Typically these are written to a http.Request.
*/
type GetTeamParams struct {

	/* OrganizationCanonical.

	   A canonical of an organization.
	*/
	OrganizationCanonical string

	/* TeamCanonical.

	   A canonical of a team.
	*/
	TeamCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTeamParams) WithDefaults() *GetTeamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTeamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get team params
func (o *GetTeamParams) WithTimeout(timeout time.Duration) *GetTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get team params
func (o *GetTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get team params
func (o *GetTeamParams) WithContext(ctx context.Context) *GetTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get team params
func (o *GetTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get team params
func (o *GetTeamParams) WithHTTPClient(client *http.Client) *GetTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get team params
func (o *GetTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationCanonical adds the organizationCanonical to the get team params
func (o *GetTeamParams) WithOrganizationCanonical(organizationCanonical string) *GetTeamParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get team params
func (o *GetTeamParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithTeamCanonical adds the teamCanonical to the get team params
func (o *GetTeamParams) WithTeamCanonical(teamCanonical string) *GetTeamParams {
	o.SetTeamCanonical(teamCanonical)
	return o
}

// SetTeamCanonical adds the teamCanonical to the get team params
func (o *GetTeamParams) SetTeamCanonical(teamCanonical string) {
	o.TeamCanonical = teamCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *GetTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	// path param team_canonical
	if err := r.SetPathParam("team_canonical", o.TeamCanonical); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
