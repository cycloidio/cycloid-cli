// Code generated by go-swagger; DO NOT EDIT.

package organization_pipelines_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetJobsParams creates a new GetJobsParams object
// with the default values initialized.
func NewGetJobsParams() *GetJobsParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &GetJobsParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetJobsParamsWithTimeout creates a new GetJobsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetJobsParamsWithTimeout(timeout time.Duration) *GetJobsParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &GetJobsParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetJobsParamsWithContext creates a new GetJobsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetJobsParamsWithContext(ctx context.Context) *GetJobsParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &GetJobsParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetJobsParamsWithHTTPClient creates a new GetJobsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetJobsParamsWithHTTPClient(client *http.Client) *GetJobsParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &GetJobsParams{
		PageIndex:  &pageIndexDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetJobsParams contains all the parameters to send to the API endpoint
for the get jobs operation typically these are written to a http.Request
*/
type GetJobsParams struct {

	/*InpathPipelineName
	  A pipeline name

	*/
	InpathPipelineName string
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string
	/*PageIndex
	  The page number to request. The first page is 1.

	*/
	PageIndex *uint32
	/*PageSize
	  The number of items at most which the response can have.

	*/
	PageSize *uint32
	/*ProjectCanonical
	  A canonical of a project.

	*/
	ProjectCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get jobs params
func (o *GetJobsParams) WithTimeout(timeout time.Duration) *GetJobsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get jobs params
func (o *GetJobsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get jobs params
func (o *GetJobsParams) WithContext(ctx context.Context) *GetJobsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get jobs params
func (o *GetJobsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get jobs params
func (o *GetJobsParams) WithHTTPClient(client *http.Client) *GetJobsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get jobs params
func (o *GetJobsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInpathPipelineName adds the inpathPipelineName to the get jobs params
func (o *GetJobsParams) WithInpathPipelineName(inpathPipelineName string) *GetJobsParams {
	o.SetInpathPipelineName(inpathPipelineName)
	return o
}

// SetInpathPipelineName adds the inpathPipelineName to the get jobs params
func (o *GetJobsParams) SetInpathPipelineName(inpathPipelineName string) {
	o.InpathPipelineName = inpathPipelineName
}

// WithOrganizationCanonical adds the organizationCanonical to the get jobs params
func (o *GetJobsParams) WithOrganizationCanonical(organizationCanonical string) *GetJobsParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get jobs params
func (o *GetJobsParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithPageIndex adds the pageIndex to the get jobs params
func (o *GetJobsParams) WithPageIndex(pageIndex *uint32) *GetJobsParams {
	o.SetPageIndex(pageIndex)
	return o
}

// SetPageIndex adds the pageIndex to the get jobs params
func (o *GetJobsParams) SetPageIndex(pageIndex *uint32) {
	o.PageIndex = pageIndex
}

// WithPageSize adds the pageSize to the get jobs params
func (o *GetJobsParams) WithPageSize(pageSize *uint32) *GetJobsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get jobs params
func (o *GetJobsParams) SetPageSize(pageSize *uint32) {
	o.PageSize = pageSize
}

// WithProjectCanonical adds the projectCanonical to the get jobs params
func (o *GetJobsParams) WithProjectCanonical(projectCanonical string) *GetJobsParams {
	o.SetProjectCanonical(projectCanonical)
	return o
}

// SetProjectCanonical adds the projectCanonical to the get jobs params
func (o *GetJobsParams) SetProjectCanonical(projectCanonical string) {
	o.ProjectCanonical = projectCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *GetJobsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.PageIndex != nil {

		// query param page_index
		var qrPageIndex uint32
		if o.PageIndex != nil {
			qrPageIndex = *o.PageIndex
		}
		qPageIndex := swag.FormatUint32(qrPageIndex)
		if qPageIndex != "" {
			if err := r.SetQueryParam("page_index", qPageIndex); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize uint32
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatUint32(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
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
