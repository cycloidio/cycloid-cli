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
)

// NewRenamePipelineParams creates a new RenamePipelineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRenamePipelineParams() *RenamePipelineParams {
	return &RenamePipelineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRenamePipelineParamsWithTimeout creates a new RenamePipelineParams object
// with the ability to set a timeout on a request.
func NewRenamePipelineParamsWithTimeout(timeout time.Duration) *RenamePipelineParams {
	return &RenamePipelineParams{
		timeout: timeout,
	}
}

// NewRenamePipelineParamsWithContext creates a new RenamePipelineParams object
// with the ability to set a context for a request.
func NewRenamePipelineParamsWithContext(ctx context.Context) *RenamePipelineParams {
	return &RenamePipelineParams{
		Context: ctx,
	}
}

// NewRenamePipelineParamsWithHTTPClient creates a new RenamePipelineParams object
// with the ability to set a custom HTTPClient for a request.
func NewRenamePipelineParamsWithHTTPClient(client *http.Client) *RenamePipelineParams {
	return &RenamePipelineParams{
		HTTPClient: client,
	}
}

/*
RenamePipelineParams contains all the parameters to send to the API endpoint

	for the rename pipeline operation.

	Typically these are written to a http.Request.
*/
type RenamePipelineParams struct {

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

	/* PipelineName.

	   A pipeline name
	*/
	PipelineName string

	/* ProjectCanonical.

	   A canonical of a project.
	*/
	ProjectCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rename pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RenamePipelineParams) WithDefaults() *RenamePipelineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rename pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RenamePipelineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rename pipeline params
func (o *RenamePipelineParams) WithTimeout(timeout time.Duration) *RenamePipelineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rename pipeline params
func (o *RenamePipelineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rename pipeline params
func (o *RenamePipelineParams) WithContext(ctx context.Context) *RenamePipelineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rename pipeline params
func (o *RenamePipelineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rename pipeline params
func (o *RenamePipelineParams) WithHTTPClient(client *http.Client) *RenamePipelineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rename pipeline params
func (o *RenamePipelineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComponentCanonical adds the componentCanonical to the rename pipeline params
func (o *RenamePipelineParams) WithComponentCanonical(componentCanonical string) *RenamePipelineParams {
	o.SetComponentCanonical(componentCanonical)
	return o
}

// SetComponentCanonical adds the componentCanonical to the rename pipeline params
func (o *RenamePipelineParams) SetComponentCanonical(componentCanonical string) {
	o.ComponentCanonical = componentCanonical
}

// WithEnvironmentCanonical adds the environmentCanonical to the rename pipeline params
func (o *RenamePipelineParams) WithEnvironmentCanonical(environmentCanonical string) *RenamePipelineParams {
	o.SetEnvironmentCanonical(environmentCanonical)
	return o
}

// SetEnvironmentCanonical adds the environmentCanonical to the rename pipeline params
func (o *RenamePipelineParams) SetEnvironmentCanonical(environmentCanonical string) {
	o.EnvironmentCanonical = environmentCanonical
}

// WithInpathPipelineName adds the inpathPipelineName to the rename pipeline params
func (o *RenamePipelineParams) WithInpathPipelineName(inpathPipelineName string) *RenamePipelineParams {
	o.SetInpathPipelineName(inpathPipelineName)
	return o
}

// SetInpathPipelineName adds the inpathPipelineName to the rename pipeline params
func (o *RenamePipelineParams) SetInpathPipelineName(inpathPipelineName string) {
	o.InpathPipelineName = inpathPipelineName
}

// WithOrganizationCanonical adds the organizationCanonical to the rename pipeline params
func (o *RenamePipelineParams) WithOrganizationCanonical(organizationCanonical string) *RenamePipelineParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the rename pipeline params
func (o *RenamePipelineParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithPipelineName adds the pipelineName to the rename pipeline params
func (o *RenamePipelineParams) WithPipelineName(pipelineName string) *RenamePipelineParams {
	o.SetPipelineName(pipelineName)
	return o
}

// SetPipelineName adds the pipelineName to the rename pipeline params
func (o *RenamePipelineParams) SetPipelineName(pipelineName string) {
	o.PipelineName = pipelineName
}

// WithProjectCanonical adds the projectCanonical to the rename pipeline params
func (o *RenamePipelineParams) WithProjectCanonical(projectCanonical string) *RenamePipelineParams {
	o.SetProjectCanonical(projectCanonical)
	return o
}

// SetProjectCanonical adds the projectCanonical to the rename pipeline params
func (o *RenamePipelineParams) SetProjectCanonical(projectCanonical string) {
	o.ProjectCanonical = projectCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *RenamePipelineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param inpath_pipeline_name
	if err := r.SetPathParam("inpath_pipeline_name", o.InpathPipelineName); err != nil {
		return err
	}

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	// query param pipeline_name
	qrPipelineName := o.PipelineName
	qPipelineName := qrPipelineName
	if qPipelineName != "" {

		if err := r.SetQueryParam("pipeline_name", qPipelineName); err != nil {
			return err
		}
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
