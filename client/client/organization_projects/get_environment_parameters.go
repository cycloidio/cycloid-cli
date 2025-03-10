// Code generated by go-swagger; DO NOT EDIT.

package organization_projects

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

// NewGetEnvironmentParams creates a new GetEnvironmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEnvironmentParams() *GetEnvironmentParams {
	return &GetEnvironmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEnvironmentParamsWithTimeout creates a new GetEnvironmentParams object
// with the ability to set a timeout on a request.
func NewGetEnvironmentParamsWithTimeout(timeout time.Duration) *GetEnvironmentParams {
	return &GetEnvironmentParams{
		timeout: timeout,
	}
}

// NewGetEnvironmentParamsWithContext creates a new GetEnvironmentParams object
// with the ability to set a context for a request.
func NewGetEnvironmentParamsWithContext(ctx context.Context) *GetEnvironmentParams {
	return &GetEnvironmentParams{
		Context: ctx,
	}
}

// NewGetEnvironmentParamsWithHTTPClient creates a new GetEnvironmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEnvironmentParamsWithHTTPClient(client *http.Client) *GetEnvironmentParams {
	return &GetEnvironmentParams{
		HTTPClient: client,
	}
}

/*
GetEnvironmentParams contains all the parameters to send to the API endpoint

	for the get environment operation.

	Typically these are written to a http.Request.
*/
type GetEnvironmentParams struct {

	/* EnvironmentCanonical.

	   The environment canonical to use as part of a path
	*/
	EnvironmentCanonical string

	/* OrganizationCanonical.

	   A canonical of an organization.
	*/
	OrganizationCanonical string

	/* ProjectCanonical.

	   A canonical of a project.
	*/
	ProjectCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get environment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEnvironmentParams) WithDefaults() *GetEnvironmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get environment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEnvironmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get environment params
func (o *GetEnvironmentParams) WithTimeout(timeout time.Duration) *GetEnvironmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get environment params
func (o *GetEnvironmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get environment params
func (o *GetEnvironmentParams) WithContext(ctx context.Context) *GetEnvironmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get environment params
func (o *GetEnvironmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get environment params
func (o *GetEnvironmentParams) WithHTTPClient(client *http.Client) *GetEnvironmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get environment params
func (o *GetEnvironmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentCanonical adds the environmentCanonical to the get environment params
func (o *GetEnvironmentParams) WithEnvironmentCanonical(environmentCanonical string) *GetEnvironmentParams {
	o.SetEnvironmentCanonical(environmentCanonical)
	return o
}

// SetEnvironmentCanonical adds the environmentCanonical to the get environment params
func (o *GetEnvironmentParams) SetEnvironmentCanonical(environmentCanonical string) {
	o.EnvironmentCanonical = environmentCanonical
}

// WithOrganizationCanonical adds the organizationCanonical to the get environment params
func (o *GetEnvironmentParams) WithOrganizationCanonical(organizationCanonical string) *GetEnvironmentParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get environment params
func (o *GetEnvironmentParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithProjectCanonical adds the projectCanonical to the get environment params
func (o *GetEnvironmentParams) WithProjectCanonical(projectCanonical string) *GetEnvironmentParams {
	o.SetProjectCanonical(projectCanonical)
	return o
}

// SetProjectCanonical adds the projectCanonical to the get environment params
func (o *GetEnvironmentParams) SetProjectCanonical(projectCanonical string) {
	o.ProjectCanonical = projectCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *GetEnvironmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param environment_canonical
	if err := r.SetPathParam("environment_canonical", o.EnvironmentCanonical); err != nil {
		return err
	}

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	// path param project_canonical
	if err := r.SetPathParam("project_canonical", o.ProjectCanonical); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
