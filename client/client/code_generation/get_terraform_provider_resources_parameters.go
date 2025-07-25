// Code generated by go-swagger; DO NOT EDIT.

package code_generation

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

// NewGetTerraformProviderResourcesParams creates a new GetTerraformProviderResourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTerraformProviderResourcesParams() *GetTerraformProviderResourcesParams {
	return &GetTerraformProviderResourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTerraformProviderResourcesParamsWithTimeout creates a new GetTerraformProviderResourcesParams object
// with the ability to set a timeout on a request.
func NewGetTerraformProviderResourcesParamsWithTimeout(timeout time.Duration) *GetTerraformProviderResourcesParams {
	return &GetTerraformProviderResourcesParams{
		timeout: timeout,
	}
}

// NewGetTerraformProviderResourcesParamsWithContext creates a new GetTerraformProviderResourcesParams object
// with the ability to set a context for a request.
func NewGetTerraformProviderResourcesParamsWithContext(ctx context.Context) *GetTerraformProviderResourcesParams {
	return &GetTerraformProviderResourcesParams{
		Context: ctx,
	}
}

// NewGetTerraformProviderResourcesParamsWithHTTPClient creates a new GetTerraformProviderResourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTerraformProviderResourcesParamsWithHTTPClient(client *http.Client) *GetTerraformProviderResourcesParams {
	return &GetTerraformProviderResourcesParams{
		HTTPClient: client,
	}
}

/*
GetTerraformProviderResourcesParams contains all the parameters to send to the API endpoint

	for the get terraform provider resources operation.

	Typically these are written to a http.Request.
*/
type GetTerraformProviderResourcesParams struct {

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

	/* ProviderCanonical.

	   A canonical of a Provider
	*/
	ProviderCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get terraform provider resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTerraformProviderResourcesParams) WithDefaults() *GetTerraformProviderResourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get terraform provider resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTerraformProviderResourcesParams) SetDefaults() {
	var (
		pageIndexDefault = uint32(1)

		pageSizeDefault = uint32(1000)
	)

	val := GetTerraformProviderResourcesParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get terraform provider resources params
func (o *GetTerraformProviderResourcesParams) WithTimeout(timeout time.Duration) *GetTerraformProviderResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get terraform provider resources params
func (o *GetTerraformProviderResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get terraform provider resources params
func (o *GetTerraformProviderResourcesParams) WithContext(ctx context.Context) *GetTerraformProviderResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get terraform provider resources params
func (o *GetTerraformProviderResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get terraform provider resources params
func (o *GetTerraformProviderResourcesParams) WithHTTPClient(client *http.Client) *GetTerraformProviderResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get terraform provider resources params
func (o *GetTerraformProviderResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationCanonical adds the organizationCanonical to the get terraform provider resources params
func (o *GetTerraformProviderResourcesParams) WithOrganizationCanonical(organizationCanonical string) *GetTerraformProviderResourcesParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get terraform provider resources params
func (o *GetTerraformProviderResourcesParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithPageIndex adds the pageIndex to the get terraform provider resources params
func (o *GetTerraformProviderResourcesParams) WithPageIndex(pageIndex *uint32) *GetTerraformProviderResourcesParams {
	o.SetPageIndex(pageIndex)
	return o
}

// SetPageIndex adds the pageIndex to the get terraform provider resources params
func (o *GetTerraformProviderResourcesParams) SetPageIndex(pageIndex *uint32) {
	o.PageIndex = pageIndex
}

// WithPageSize adds the pageSize to the get terraform provider resources params
func (o *GetTerraformProviderResourcesParams) WithPageSize(pageSize *uint32) *GetTerraformProviderResourcesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get terraform provider resources params
func (o *GetTerraformProviderResourcesParams) SetPageSize(pageSize *uint32) {
	o.PageSize = pageSize
}

// WithProviderCanonical adds the providerCanonical to the get terraform provider resources params
func (o *GetTerraformProviderResourcesParams) WithProviderCanonical(providerCanonical string) *GetTerraformProviderResourcesParams {
	o.SetProviderCanonical(providerCanonical)
	return o
}

// SetProviderCanonical adds the providerCanonical to the get terraform provider resources params
func (o *GetTerraformProviderResourcesParams) SetProviderCanonical(providerCanonical string) {
	o.ProviderCanonical = providerCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *GetTerraformProviderResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param provider_canonical
	if err := r.SetPathParam("provider_canonical", o.ProviderCanonical); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
