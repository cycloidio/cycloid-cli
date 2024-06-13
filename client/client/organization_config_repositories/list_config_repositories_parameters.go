// Code generated by go-swagger; DO NOT EDIT.

package organization_config_repositories

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

// NewListConfigRepositoriesParams creates a new ListConfigRepositoriesParams object
// with the default values initialized.
func NewListConfigRepositoriesParams() *ListConfigRepositoriesParams {
	var (
		defaultVarDefault = bool(false)
		pageIndexDefault  = uint32(1)
		pageSizeDefault   = uint32(1000)
	)
	return &ListConfigRepositoriesParams{
		Default:   &defaultVarDefault,
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListConfigRepositoriesParamsWithTimeout creates a new ListConfigRepositoriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListConfigRepositoriesParamsWithTimeout(timeout time.Duration) *ListConfigRepositoriesParams {
	var (
		defaultVarDefault = bool(false)
		pageIndexDefault  = uint32(1)
		pageSizeDefault   = uint32(1000)
	)
	return &ListConfigRepositoriesParams{
		Default:   &defaultVarDefault,
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		timeout: timeout,
	}
}

// NewListConfigRepositoriesParamsWithContext creates a new ListConfigRepositoriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListConfigRepositoriesParamsWithContext(ctx context.Context) *ListConfigRepositoriesParams {
	var (
		defaultDefault   = bool(false)
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &ListConfigRepositoriesParams{
		Default:   &defaultDefault,
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		Context: ctx,
	}
}

// NewListConfigRepositoriesParamsWithHTTPClient creates a new ListConfigRepositoriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListConfigRepositoriesParamsWithHTTPClient(client *http.Client) *ListConfigRepositoriesParams {
	var (
		defaultDefault   = bool(false)
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &ListConfigRepositoriesParams{
		Default:    &defaultDefault,
		PageIndex:  &pageIndexDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*ListConfigRepositoriesParams contains all the parameters to send to the API endpoint
for the list config repositories operation typically these are written to a http.Request
*/
type ListConfigRepositoriesParams struct {

	/*Default
	  Value describing whether to return default

	*/
	Default *bool
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

// WithTimeout adds the timeout to the list config repositories params
func (o *ListConfigRepositoriesParams) WithTimeout(timeout time.Duration) *ListConfigRepositoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list config repositories params
func (o *ListConfigRepositoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list config repositories params
func (o *ListConfigRepositoriesParams) WithContext(ctx context.Context) *ListConfigRepositoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list config repositories params
func (o *ListConfigRepositoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list config repositories params
func (o *ListConfigRepositoriesParams) WithHTTPClient(client *http.Client) *ListConfigRepositoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list config repositories params
func (o *ListConfigRepositoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDefault adds the defaultVar to the list config repositories params
func (o *ListConfigRepositoriesParams) WithDefault(defaultVar *bool) *ListConfigRepositoriesParams {
	o.SetDefault(defaultVar)
	return o
}

// SetDefault adds the default to the list config repositories params
func (o *ListConfigRepositoriesParams) SetDefault(defaultVar *bool) {
	o.Default = defaultVar
}

// WithOrganizationCanonical adds the organizationCanonical to the list config repositories params
func (o *ListConfigRepositoriesParams) WithOrganizationCanonical(organizationCanonical string) *ListConfigRepositoriesParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the list config repositories params
func (o *ListConfigRepositoriesParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithPageIndex adds the pageIndex to the list config repositories params
func (o *ListConfigRepositoriesParams) WithPageIndex(pageIndex *uint32) *ListConfigRepositoriesParams {
	o.SetPageIndex(pageIndex)
	return o
}

// SetPageIndex adds the pageIndex to the list config repositories params
func (o *ListConfigRepositoriesParams) SetPageIndex(pageIndex *uint32) {
	o.PageIndex = pageIndex
}

// WithPageSize adds the pageSize to the list config repositories params
func (o *ListConfigRepositoriesParams) WithPageSize(pageSize *uint32) *ListConfigRepositoriesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list config repositories params
func (o *ListConfigRepositoriesParams) SetPageSize(pageSize *uint32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *ListConfigRepositoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Default != nil {

		// query param default
		var qrDefault bool
		if o.Default != nil {
			qrDefault = *o.Default
		}
		qDefault := swag.FormatBool(qrDefault)
		if qDefault != "" {
			if err := r.SetQueryParam("default", qDefault); err != nil {
				return err
			}
		}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
