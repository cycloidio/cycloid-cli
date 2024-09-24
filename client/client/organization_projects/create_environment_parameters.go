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

	"github.com/cycloidio/cycloid-cli/client/models"
)

// NewCreateEnvironmentParams creates a new CreateEnvironmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateEnvironmentParams() *CreateEnvironmentParams {
	return &CreateEnvironmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateEnvironmentParamsWithTimeout creates a new CreateEnvironmentParams object
// with the ability to set a timeout on a request.
func NewCreateEnvironmentParamsWithTimeout(timeout time.Duration) *CreateEnvironmentParams {
	return &CreateEnvironmentParams{
		timeout: timeout,
	}
}

// NewCreateEnvironmentParamsWithContext creates a new CreateEnvironmentParams object
// with the ability to set a context for a request.
func NewCreateEnvironmentParamsWithContext(ctx context.Context) *CreateEnvironmentParams {
	return &CreateEnvironmentParams{
		Context: ctx,
	}
}

// NewCreateEnvironmentParamsWithHTTPClient creates a new CreateEnvironmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateEnvironmentParamsWithHTTPClient(client *http.Client) *CreateEnvironmentParams {
	return &CreateEnvironmentParams{
		HTTPClient: client,
	}
}

/*
CreateEnvironmentParams contains all the parameters to send to the API endpoint

	for the create environment operation.

	Typically these are written to a http.Request.
*/
type CreateEnvironmentParams struct {

	/* Body.

	   The canonical of the environment to create and its configuration.
	*/
	Body *models.NewEnvironment

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

// WithDefaults hydrates default values in the create environment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEnvironmentParams) WithDefaults() *CreateEnvironmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create environment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEnvironmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create environment params
func (o *CreateEnvironmentParams) WithTimeout(timeout time.Duration) *CreateEnvironmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create environment params
func (o *CreateEnvironmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create environment params
func (o *CreateEnvironmentParams) WithContext(ctx context.Context) *CreateEnvironmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create environment params
func (o *CreateEnvironmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create environment params
func (o *CreateEnvironmentParams) WithHTTPClient(client *http.Client) *CreateEnvironmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create environment params
func (o *CreateEnvironmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create environment params
func (o *CreateEnvironmentParams) WithBody(body *models.NewEnvironment) *CreateEnvironmentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create environment params
func (o *CreateEnvironmentParams) SetBody(body *models.NewEnvironment) {
	o.Body = body
}

// WithOrganizationCanonical adds the organizationCanonical to the create environment params
func (o *CreateEnvironmentParams) WithOrganizationCanonical(organizationCanonical string) *CreateEnvironmentParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the create environment params
func (o *CreateEnvironmentParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithProjectCanonical adds the projectCanonical to the create environment params
func (o *CreateEnvironmentParams) WithProjectCanonical(projectCanonical string) *CreateEnvironmentParams {
	o.SetProjectCanonical(projectCanonical)
	return o
}

// SetProjectCanonical adds the projectCanonical to the create environment params
func (o *CreateEnvironmentParams) SetProjectCanonical(projectCanonical string) {
	o.ProjectCanonical = projectCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *CreateEnvironmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project_canonical
	if err := r.SetPathParam("project_canonical", o.ProjectCanonical); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}