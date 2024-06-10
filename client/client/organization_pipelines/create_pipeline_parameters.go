// Code generated by go-swagger; DO NOT EDIT.

package organization_pipelines

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

// NewCreatePipelineParams creates a new CreatePipelineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePipelineParams() *CreatePipelineParams {
	return &CreatePipelineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePipelineParamsWithTimeout creates a new CreatePipelineParams object
// with the ability to set a timeout on a request.
func NewCreatePipelineParamsWithTimeout(timeout time.Duration) *CreatePipelineParams {
	return &CreatePipelineParams{
		timeout: timeout,
	}
}

// NewCreatePipelineParamsWithContext creates a new CreatePipelineParams object
// with the ability to set a context for a request.
func NewCreatePipelineParamsWithContext(ctx context.Context) *CreatePipelineParams {
	return &CreatePipelineParams{
		Context: ctx,
	}
}

// NewCreatePipelineParamsWithHTTPClient creates a new CreatePipelineParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePipelineParamsWithHTTPClient(client *http.Client) *CreatePipelineParams {
	return &CreatePipelineParams{
		HTTPClient: client,
	}
}

/*
CreatePipelineParams contains all the parameters to send to the API endpoint

	for the create pipeline operation.

	Typically these are written to a http.Request.
*/
type CreatePipelineParams struct {

	/* Body.

	   The configuration of the pipeline to create.
	*/
	Body *models.NewPipeline

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

// WithDefaults hydrates default values in the create pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePipelineParams) WithDefaults() *CreatePipelineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePipelineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create pipeline params
func (o *CreatePipelineParams) WithTimeout(timeout time.Duration) *CreatePipelineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create pipeline params
func (o *CreatePipelineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create pipeline params
func (o *CreatePipelineParams) WithContext(ctx context.Context) *CreatePipelineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create pipeline params
func (o *CreatePipelineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create pipeline params
func (o *CreatePipelineParams) WithHTTPClient(client *http.Client) *CreatePipelineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create pipeline params
func (o *CreatePipelineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create pipeline params
func (o *CreatePipelineParams) WithBody(body *models.NewPipeline) *CreatePipelineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create pipeline params
func (o *CreatePipelineParams) SetBody(body *models.NewPipeline) {
	o.Body = body
}

// WithOrganizationCanonical adds the organizationCanonical to the create pipeline params
func (o *CreatePipelineParams) WithOrganizationCanonical(organizationCanonical string) *CreatePipelineParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the create pipeline params
func (o *CreatePipelineParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithProjectCanonical adds the projectCanonical to the create pipeline params
func (o *CreatePipelineParams) WithProjectCanonical(projectCanonical string) *CreatePipelineParams {
	o.SetProjectCanonical(projectCanonical)
	return o
}

// SetProjectCanonical adds the projectCanonical to the create pipeline params
func (o *CreatePipelineParams) SetProjectCanonical(projectCanonical string) {
	o.ProjectCanonical = projectCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePipelineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
