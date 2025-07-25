// Code generated by go-swagger; DO NOT EDIT.

package component_pipelines

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

// NewUpdatePipelineParams creates a new UpdatePipelineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdatePipelineParams() *UpdatePipelineParams {
	return &UpdatePipelineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePipelineParamsWithTimeout creates a new UpdatePipelineParams object
// with the ability to set a timeout on a request.
func NewUpdatePipelineParamsWithTimeout(timeout time.Duration) *UpdatePipelineParams {
	return &UpdatePipelineParams{
		timeout: timeout,
	}
}

// NewUpdatePipelineParamsWithContext creates a new UpdatePipelineParams object
// with the ability to set a context for a request.
func NewUpdatePipelineParamsWithContext(ctx context.Context) *UpdatePipelineParams {
	return &UpdatePipelineParams{
		Context: ctx,
	}
}

// NewUpdatePipelineParamsWithHTTPClient creates a new UpdatePipelineParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdatePipelineParamsWithHTTPClient(client *http.Client) *UpdatePipelineParams {
	return &UpdatePipelineParams{
		HTTPClient: client,
	}
}

/*
UpdatePipelineParams contains all the parameters to send to the API endpoint

	for the update pipeline operation.

	Typically these are written to a http.Request.
*/
type UpdatePipelineParams struct {

	/* Body.

	   The pipeline configuration
	*/
	Body *models.UpdatePipeline

	/* ComponentCanonical.

	   A canonical of a component.
	*/
	ComponentCanonical string

	/* EnvironmentCanonical.

	   The environment canonical to use as part of a path
	*/
	EnvironmentCanonical string

	/* InpathPipelineName.

	   A pipeline name
	*/
	InpathPipelineName string

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

// WithDefaults hydrates default values in the update pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePipelineParams) WithDefaults() *UpdatePipelineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePipelineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update pipeline params
func (o *UpdatePipelineParams) WithTimeout(timeout time.Duration) *UpdatePipelineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update pipeline params
func (o *UpdatePipelineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update pipeline params
func (o *UpdatePipelineParams) WithContext(ctx context.Context) *UpdatePipelineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update pipeline params
func (o *UpdatePipelineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update pipeline params
func (o *UpdatePipelineParams) WithHTTPClient(client *http.Client) *UpdatePipelineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update pipeline params
func (o *UpdatePipelineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update pipeline params
func (o *UpdatePipelineParams) WithBody(body *models.UpdatePipeline) *UpdatePipelineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update pipeline params
func (o *UpdatePipelineParams) SetBody(body *models.UpdatePipeline) {
	o.Body = body
}

// WithComponentCanonical adds the componentCanonical to the update pipeline params
func (o *UpdatePipelineParams) WithComponentCanonical(componentCanonical string) *UpdatePipelineParams {
	o.SetComponentCanonical(componentCanonical)
	return o
}

// SetComponentCanonical adds the componentCanonical to the update pipeline params
func (o *UpdatePipelineParams) SetComponentCanonical(componentCanonical string) {
	o.ComponentCanonical = componentCanonical
}

// WithEnvironmentCanonical adds the environmentCanonical to the update pipeline params
func (o *UpdatePipelineParams) WithEnvironmentCanonical(environmentCanonical string) *UpdatePipelineParams {
	o.SetEnvironmentCanonical(environmentCanonical)
	return o
}

// SetEnvironmentCanonical adds the environmentCanonical to the update pipeline params
func (o *UpdatePipelineParams) SetEnvironmentCanonical(environmentCanonical string) {
	o.EnvironmentCanonical = environmentCanonical
}

// WithInpathPipelineName adds the inpathPipelineName to the update pipeline params
func (o *UpdatePipelineParams) WithInpathPipelineName(inpathPipelineName string) *UpdatePipelineParams {
	o.SetInpathPipelineName(inpathPipelineName)
	return o
}

// SetInpathPipelineName adds the inpathPipelineName to the update pipeline params
func (o *UpdatePipelineParams) SetInpathPipelineName(inpathPipelineName string) {
	o.InpathPipelineName = inpathPipelineName
}

// WithOrganizationCanonical adds the organizationCanonical to the update pipeline params
func (o *UpdatePipelineParams) WithOrganizationCanonical(organizationCanonical string) *UpdatePipelineParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the update pipeline params
func (o *UpdatePipelineParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithProjectCanonical adds the projectCanonical to the update pipeline params
func (o *UpdatePipelineParams) WithProjectCanonical(projectCanonical string) *UpdatePipelineParams {
	o.SetProjectCanonical(projectCanonical)
	return o
}

// SetProjectCanonical adds the projectCanonical to the update pipeline params
func (o *UpdatePipelineParams) SetProjectCanonical(projectCanonical string) {
	o.ProjectCanonical = projectCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePipelineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param component_canonical
	if err := r.SetPathParam("component_canonical", o.ComponentCanonical); err != nil {
		return err
	}

	// path param environment_canonical
	if err := r.SetPathParam("environment_canonical", o.EnvironmentCanonical); err != nil {
		return err
	}

	// path param inpath_pipeline_name
	if err := r.SetPathParam("inpath_pipeline_name", o.InpathPipelineName); err != nil {
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
