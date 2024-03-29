// Code generated by go-swagger; DO NOT EDIT.

package organization_service_catalog_sources

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

// NewGetServiceCatalogSourcesParams creates a new GetServiceCatalogSourcesParams object
// with the default values initialized.
func NewGetServiceCatalogSourcesParams() *GetServiceCatalogSourcesParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &GetServiceCatalogSourcesParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceCatalogSourcesParamsWithTimeout creates a new GetServiceCatalogSourcesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServiceCatalogSourcesParamsWithTimeout(timeout time.Duration) *GetServiceCatalogSourcesParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &GetServiceCatalogSourcesParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetServiceCatalogSourcesParamsWithContext creates a new GetServiceCatalogSourcesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServiceCatalogSourcesParamsWithContext(ctx context.Context) *GetServiceCatalogSourcesParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &GetServiceCatalogSourcesParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetServiceCatalogSourcesParamsWithHTTPClient creates a new GetServiceCatalogSourcesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServiceCatalogSourcesParamsWithHTTPClient(client *http.Client) *GetServiceCatalogSourcesParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &GetServiceCatalogSourcesParams{
		PageIndex:  &pageIndexDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetServiceCatalogSourcesParams contains all the parameters to send to the API endpoint
for the get service catalog sources operation typically these are written to a http.Request
*/
type GetServiceCatalogSourcesParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get service catalog sources params
func (o *GetServiceCatalogSourcesParams) WithTimeout(timeout time.Duration) *GetServiceCatalogSourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service catalog sources params
func (o *GetServiceCatalogSourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service catalog sources params
func (o *GetServiceCatalogSourcesParams) WithContext(ctx context.Context) *GetServiceCatalogSourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service catalog sources params
func (o *GetServiceCatalogSourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service catalog sources params
func (o *GetServiceCatalogSourcesParams) WithHTTPClient(client *http.Client) *GetServiceCatalogSourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service catalog sources params
func (o *GetServiceCatalogSourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationCanonical adds the organizationCanonical to the get service catalog sources params
func (o *GetServiceCatalogSourcesParams) WithOrganizationCanonical(organizationCanonical string) *GetServiceCatalogSourcesParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get service catalog sources params
func (o *GetServiceCatalogSourcesParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithPageIndex adds the pageIndex to the get service catalog sources params
func (o *GetServiceCatalogSourcesParams) WithPageIndex(pageIndex *uint32) *GetServiceCatalogSourcesParams {
	o.SetPageIndex(pageIndex)
	return o
}

// SetPageIndex adds the pageIndex to the get service catalog sources params
func (o *GetServiceCatalogSourcesParams) SetPageIndex(pageIndex *uint32) {
	o.PageIndex = pageIndex
}

// WithPageSize adds the pageSize to the get service catalog sources params
func (o *GetServiceCatalogSourcesParams) WithPageSize(pageSize *uint32) *GetServiceCatalogSourcesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get service catalog sources params
func (o *GetServiceCatalogSourcesParams) SetPageSize(pageSize *uint32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceCatalogSourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
