// Code generated by go-swagger; DO NOT EDIT.

package organization_roles

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
	"github.com/go-openapi/swag"
)

// NewGetRolesParams creates a new GetRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRolesParams() *GetRolesParams {
	return &GetRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRolesParamsWithTimeout creates a new GetRolesParams object
// with the ability to set a timeout on a request.
func NewGetRolesParamsWithTimeout(timeout time.Duration) *GetRolesParams {
	return &GetRolesParams{
		timeout: timeout,
	}
}

// NewGetRolesParamsWithContext creates a new GetRolesParams object
// with the ability to set a context for a request.
func NewGetRolesParamsWithContext(ctx context.Context) *GetRolesParams {
	return &GetRolesParams{
		Context: ctx,
	}
}

// NewGetRolesParamsWithHTTPClient creates a new GetRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRolesParamsWithHTTPClient(client *http.Client) *GetRolesParams {
	return &GetRolesParams{
		HTTPClient: client,
	}
}

/*
GetRolesParams contains all the parameters to send to the API endpoint

	for the get roles operation.

	Typically these are written to a http.Request.
*/
type GetRolesParams struct {

	/* OrganizationCanonical.

	   A canonical of an organization.
	*/
	OrganizationCanonical string

	/* PageIndex.

	   The page number to request. The first page is 1.

	   Format: uint32
	   Default: 1
	*/
	PageIndex *uint32

	/* PageSize.

	   The number of items at most which the response can have.

	   Format: uint32
	   Default: 1000
	*/
	PageSize *uint32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRolesParams) WithDefaults() *GetRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRolesParams) SetDefaults() {
	var (
		pageIndexDefault = uint32(1)

		pageSizeDefault = uint32(1000)
	)

	val := GetRolesParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get roles params
func (o *GetRolesParams) WithTimeout(timeout time.Duration) *GetRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get roles params
func (o *GetRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get roles params
func (o *GetRolesParams) WithContext(ctx context.Context) *GetRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get roles params
func (o *GetRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get roles params
func (o *GetRolesParams) WithHTTPClient(client *http.Client) *GetRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get roles params
func (o *GetRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationCanonical adds the organizationCanonical to the get roles params
func (o *GetRolesParams) WithOrganizationCanonical(organizationCanonical string) *GetRolesParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get roles params
func (o *GetRolesParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithPageIndex adds the pageIndex to the get roles params
func (o *GetRolesParams) WithPageIndex(pageIndex *uint32) *GetRolesParams {
	o.SetPageIndex(pageIndex)
	return o
}

// SetPageIndex adds the pageIndex to the get roles params
func (o *GetRolesParams) SetPageIndex(pageIndex *uint32) {
	o.PageIndex = pageIndex
}

// WithPageSize adds the pageSize to the get roles params
func (o *GetRolesParams) WithPageSize(pageSize *uint32) *GetRolesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get roles params
func (o *GetRolesParams) SetPageSize(pageSize *uint32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
