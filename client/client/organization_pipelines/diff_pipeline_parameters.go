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

// NewDiffPipelineParams creates a new DiffPipelineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDiffPipelineParams() *DiffPipelineParams {
	return &DiffPipelineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDiffPipelineParamsWithTimeout creates a new DiffPipelineParams object
// with the ability to set a timeout on a request.
func NewDiffPipelineParamsWithTimeout(timeout time.Duration) *DiffPipelineParams {
	return &DiffPipelineParams{
		timeout: timeout,
	}
}

// NewDiffPipelineParamsWithContext creates a new DiffPipelineParams object
// with the ability to set a context for a request.
func NewDiffPipelineParamsWithContext(ctx context.Context) *DiffPipelineParams {
	return &DiffPipelineParams{
		Context: ctx,
	}
}

// NewDiffPipelineParamsWithHTTPClient creates a new DiffPipelineParams object
// with the ability to set a custom HTTPClient for a request.
func NewDiffPipelineParamsWithHTTPClient(client *http.Client) *DiffPipelineParams {
	return &DiffPipelineParams{
		HTTPClient: client,
	}
}

/*
DiffPipelineParams contains all the parameters to send to the API endpoint

	for the diff pipeline operation.

	Typically these are written to a http.Request.
*/
type DiffPipelineParams struct {

	/* Body.

	   The pipeline configuration
	*/
	Body *models.UpdatePipeline

	/* InpathPipelineName.

	   A pipeline name
	*/
	InpathPipelineName string

	/* OrganizationCanonical.

	   A canonical of an organization.
	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the diff pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiffPipelineParams) WithDefaults() *DiffPipelineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the diff pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiffPipelineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the diff pipeline params
func (o *DiffPipelineParams) WithTimeout(timeout time.Duration) *DiffPipelineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the diff pipeline params
func (o *DiffPipelineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the diff pipeline params
func (o *DiffPipelineParams) WithContext(ctx context.Context) *DiffPipelineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the diff pipeline params
func (o *DiffPipelineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the diff pipeline params
func (o *DiffPipelineParams) WithHTTPClient(client *http.Client) *DiffPipelineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the diff pipeline params
func (o *DiffPipelineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the diff pipeline params
func (o *DiffPipelineParams) WithBody(body *models.UpdatePipeline) *DiffPipelineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the diff pipeline params
func (o *DiffPipelineParams) SetBody(body *models.UpdatePipeline) {
	o.Body = body
}

// WithInpathPipelineName adds the inpathPipelineName to the diff pipeline params
func (o *DiffPipelineParams) WithInpathPipelineName(inpathPipelineName string) *DiffPipelineParams {
	o.SetInpathPipelineName(inpathPipelineName)
	return o
}

// SetInpathPipelineName adds the inpathPipelineName to the diff pipeline params
func (o *DiffPipelineParams) SetInpathPipelineName(inpathPipelineName string) {
	o.InpathPipelineName = inpathPipelineName
}

// WithOrganizationCanonical adds the organizationCanonical to the diff pipeline params
func (o *DiffPipelineParams) WithOrganizationCanonical(organizationCanonical string) *DiffPipelineParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the diff pipeline params
func (o *DiffPipelineParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *DiffPipelineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param inpath_pipeline_name
	if err := r.SetPathParam("inpath_pipeline_name", o.InpathPipelineName); err != nil {
		return err
	}

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
