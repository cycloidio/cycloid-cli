// Code generated by go-swagger; DO NOT EDIT.

package organization_components

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

// NewGetComponentInfrastructureParams creates a new GetComponentInfrastructureParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetComponentInfrastructureParams() *GetComponentInfrastructureParams {
	return &GetComponentInfrastructureParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetComponentInfrastructureParamsWithTimeout creates a new GetComponentInfrastructureParams object
// with the ability to set a timeout on a request.
func NewGetComponentInfrastructureParamsWithTimeout(timeout time.Duration) *GetComponentInfrastructureParams {
	return &GetComponentInfrastructureParams{
		timeout: timeout,
	}
}

// NewGetComponentInfrastructureParamsWithContext creates a new GetComponentInfrastructureParams object
// with the ability to set a context for a request.
func NewGetComponentInfrastructureParamsWithContext(ctx context.Context) *GetComponentInfrastructureParams {
	return &GetComponentInfrastructureParams{
		Context: ctx,
	}
}

// NewGetComponentInfrastructureParamsWithHTTPClient creates a new GetComponentInfrastructureParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetComponentInfrastructureParamsWithHTTPClient(client *http.Client) *GetComponentInfrastructureParams {
	return &GetComponentInfrastructureParams{
		HTTPClient: client,
	}
}

/*
GetComponentInfrastructureParams contains all the parameters to send to the API endpoint

	for the get component infrastructure operation.

	Typically these are written to a http.Request.
*/
type GetComponentInfrastructureParams struct {

	/* ComponentCanonical.

	   A canonical of a component.
	*/
	ComponentCanonical string

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

// WithDefaults hydrates default values in the get component infrastructure params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComponentInfrastructureParams) WithDefaults() *GetComponentInfrastructureParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get component infrastructure params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComponentInfrastructureParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get component infrastructure params
func (o *GetComponentInfrastructureParams) WithTimeout(timeout time.Duration) *GetComponentInfrastructureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get component infrastructure params
func (o *GetComponentInfrastructureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get component infrastructure params
func (o *GetComponentInfrastructureParams) WithContext(ctx context.Context) *GetComponentInfrastructureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get component infrastructure params
func (o *GetComponentInfrastructureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get component infrastructure params
func (o *GetComponentInfrastructureParams) WithHTTPClient(client *http.Client) *GetComponentInfrastructureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get component infrastructure params
func (o *GetComponentInfrastructureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComponentCanonical adds the componentCanonical to the get component infrastructure params
func (o *GetComponentInfrastructureParams) WithComponentCanonical(componentCanonical string) *GetComponentInfrastructureParams {
	o.SetComponentCanonical(componentCanonical)
	return o
}

// SetComponentCanonical adds the componentCanonical to the get component infrastructure params
func (o *GetComponentInfrastructureParams) SetComponentCanonical(componentCanonical string) {
	o.ComponentCanonical = componentCanonical
}

// WithEnvironmentCanonical adds the environmentCanonical to the get component infrastructure params
func (o *GetComponentInfrastructureParams) WithEnvironmentCanonical(environmentCanonical string) *GetComponentInfrastructureParams {
	o.SetEnvironmentCanonical(environmentCanonical)
	return o
}

// SetEnvironmentCanonical adds the environmentCanonical to the get component infrastructure params
func (o *GetComponentInfrastructureParams) SetEnvironmentCanonical(environmentCanonical string) {
	o.EnvironmentCanonical = environmentCanonical
}

// WithOrganizationCanonical adds the organizationCanonical to the get component infrastructure params
func (o *GetComponentInfrastructureParams) WithOrganizationCanonical(organizationCanonical string) *GetComponentInfrastructureParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get component infrastructure params
func (o *GetComponentInfrastructureParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithProjectCanonical adds the projectCanonical to the get component infrastructure params
func (o *GetComponentInfrastructureParams) WithProjectCanonical(projectCanonical string) *GetComponentInfrastructureParams {
	o.SetProjectCanonical(projectCanonical)
	return o
}

// SetProjectCanonical adds the projectCanonical to the get component infrastructure params
func (o *GetComponentInfrastructureParams) SetProjectCanonical(projectCanonical string) {
	o.ProjectCanonical = projectCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *GetComponentInfrastructureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param component_canonical
	if err := r.SetPathParam("component_canonical", o.ComponentCanonical); err != nil {
		return err
	}

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
