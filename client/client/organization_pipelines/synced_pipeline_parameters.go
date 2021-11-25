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

	strfmt "github.com/go-openapi/strfmt"
)

// NewSyncedPipelineParams creates a new SyncedPipelineParams object
// with the default values initialized.
func NewSyncedPipelineParams() *SyncedPipelineParams {
	var ()
	return &SyncedPipelineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSyncedPipelineParamsWithTimeout creates a new SyncedPipelineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSyncedPipelineParamsWithTimeout(timeout time.Duration) *SyncedPipelineParams {
	var ()
	return &SyncedPipelineParams{

		timeout: timeout,
	}
}

// NewSyncedPipelineParamsWithContext creates a new SyncedPipelineParams object
// with the default values initialized, and the ability to set a context for a request
func NewSyncedPipelineParamsWithContext(ctx context.Context) *SyncedPipelineParams {
	var ()
	return &SyncedPipelineParams{

		Context: ctx,
	}
}

// NewSyncedPipelineParamsWithHTTPClient creates a new SyncedPipelineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSyncedPipelineParamsWithHTTPClient(client *http.Client) *SyncedPipelineParams {
	var ()
	return &SyncedPipelineParams{
		HTTPClient: client,
	}
}

/*SyncedPipelineParams contains all the parameters to send to the API endpoint
for the synced pipeline operation typically these are written to a http.Request
*/
type SyncedPipelineParams struct {

	/*InpathPipelineName
	  A pipeline name

	*/
	InpathPipelineName string
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the synced pipeline params
func (o *SyncedPipelineParams) WithTimeout(timeout time.Duration) *SyncedPipelineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the synced pipeline params
func (o *SyncedPipelineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the synced pipeline params
func (o *SyncedPipelineParams) WithContext(ctx context.Context) *SyncedPipelineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the synced pipeline params
func (o *SyncedPipelineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the synced pipeline params
func (o *SyncedPipelineParams) WithHTTPClient(client *http.Client) *SyncedPipelineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the synced pipeline params
func (o *SyncedPipelineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInpathPipelineName adds the inpathPipelineName to the synced pipeline params
func (o *SyncedPipelineParams) WithInpathPipelineName(inpathPipelineName string) *SyncedPipelineParams {
	o.SetInpathPipelineName(inpathPipelineName)
	return o
}

// SetInpathPipelineName adds the inpathPipelineName to the synced pipeline params
func (o *SyncedPipelineParams) SetInpathPipelineName(inpathPipelineName string) {
	o.InpathPipelineName = inpathPipelineName
}

// WithOrganizationCanonical adds the organizationCanonical to the synced pipeline params
func (o *SyncedPipelineParams) WithOrganizationCanonical(organizationCanonical string) *SyncedPipelineParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the synced pipeline params
func (o *SyncedPipelineParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *SyncedPipelineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
