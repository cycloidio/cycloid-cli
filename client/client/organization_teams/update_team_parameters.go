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

	"github.com/cycloidio/cycloid-cli/client/models"
)

// NewUpdateTeamParams creates a new UpdateTeamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateTeamParams() *UpdateTeamParams {
	return &UpdateTeamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTeamParamsWithTimeout creates a new UpdateTeamParams object
// with the ability to set a timeout on a request.
func NewUpdateTeamParamsWithTimeout(timeout time.Duration) *UpdateTeamParams {
	return &UpdateTeamParams{
		timeout: timeout,
	}
}

// NewUpdateTeamParamsWithContext creates a new UpdateTeamParams object
// with the ability to set a context for a request.
func NewUpdateTeamParamsWithContext(ctx context.Context) *UpdateTeamParams {
	return &UpdateTeamParams{
		Context: ctx,
	}
}

// NewUpdateTeamParamsWithHTTPClient creates a new UpdateTeamParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateTeamParamsWithHTTPClient(client *http.Client) *UpdateTeamParams {
	return &UpdateTeamParams{
		HTTPClient: client,
	}
}

/*
UpdateTeamParams contains all the parameters to send to the API endpoint

	for the update team operation.

	Typically these are written to a http.Request.
*/
type UpdateTeamParams struct {

	/* Body.

	   The information of the team to update.
	*/
	Body *models.UpdateTeam

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

// WithDefaults hydrates default values in the update team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTeamParams) WithDefaults() *UpdateTeamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update team params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTeamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update team params
func (o *UpdateTeamParams) WithTimeout(timeout time.Duration) *UpdateTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update team params
func (o *UpdateTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update team params
func (o *UpdateTeamParams) WithContext(ctx context.Context) *UpdateTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update team params
func (o *UpdateTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update team params
func (o *UpdateTeamParams) WithHTTPClient(client *http.Client) *UpdateTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update team params
func (o *UpdateTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update team params
func (o *UpdateTeamParams) WithBody(body *models.UpdateTeam) *UpdateTeamParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update team params
func (o *UpdateTeamParams) SetBody(body *models.UpdateTeam) {
	o.Body = body
}

// WithOrganizationCanonical adds the organizationCanonical to the update team params
func (o *UpdateTeamParams) WithOrganizationCanonical(organizationCanonical string) *UpdateTeamParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the update team params
func (o *UpdateTeamParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithTeamCanonical adds the teamCanonical to the update team params
func (o *UpdateTeamParams) WithTeamCanonical(teamCanonical string) *UpdateTeamParams {
	o.SetTeamCanonical(teamCanonical)
	return o
}

// SetTeamCanonical adds the teamCanonical to the update team params
func (o *UpdateTeamParams) SetTeamCanonical(teamCanonical string) {
	o.TeamCanonical = teamCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
