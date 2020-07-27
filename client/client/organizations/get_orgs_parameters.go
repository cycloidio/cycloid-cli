// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewGetOrgsParams creates a new GetOrgsParams object
// with the default values initialized.
func NewGetOrgsParams() *GetOrgsParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(10)
	)
	return &GetOrgsParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrgsParamsWithTimeout creates a new GetOrgsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrgsParamsWithTimeout(timeout time.Duration) *GetOrgsParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(10)
	)
	return &GetOrgsParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetOrgsParamsWithContext creates a new GetOrgsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrgsParamsWithContext(ctx context.Context) *GetOrgsParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(10)
	)
	return &GetOrgsParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetOrgsParamsWithHTTPClient creates a new GetOrgsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrgsParamsWithHTTPClient(client *http.Client) *GetOrgsParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(10)
	)
	return &GetOrgsParams{
		PageIndex:  &pageIndexDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetOrgsParams contains all the parameters to send to the API endpoint
for the get orgs operation typically these are written to a http.Request
*/
type GetOrgsParams struct {

	/*OrderBy
	  Allows to order the list of items. Example usage: field_name:asc


	*/
	OrderBy *string
	/*OrganizationCreatedAt
	  Search by organization's creation date

	*/
	OrganizationCreatedAt *strfmt.DateTime
	/*OrganizationName
	  Search by the organization's name

	*/
	OrganizationName *string
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

// WithTimeout adds the timeout to the get orgs params
func (o *GetOrgsParams) WithTimeout(timeout time.Duration) *GetOrgsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get orgs params
func (o *GetOrgsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get orgs params
func (o *GetOrgsParams) WithContext(ctx context.Context) *GetOrgsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get orgs params
func (o *GetOrgsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get orgs params
func (o *GetOrgsParams) WithHTTPClient(client *http.Client) *GetOrgsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get orgs params
func (o *GetOrgsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderBy adds the orderBy to the get orgs params
func (o *GetOrgsParams) WithOrderBy(orderBy *string) *GetOrgsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get orgs params
func (o *GetOrgsParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithOrganizationCreatedAt adds the organizationCreatedAt to the get orgs params
func (o *GetOrgsParams) WithOrganizationCreatedAt(organizationCreatedAt *strfmt.DateTime) *GetOrgsParams {
	o.SetOrganizationCreatedAt(organizationCreatedAt)
	return o
}

// SetOrganizationCreatedAt adds the organizationCreatedAt to the get orgs params
func (o *GetOrgsParams) SetOrganizationCreatedAt(organizationCreatedAt *strfmt.DateTime) {
	o.OrganizationCreatedAt = organizationCreatedAt
}

// WithOrganizationName adds the organizationName to the get orgs params
func (o *GetOrgsParams) WithOrganizationName(organizationName *string) *GetOrgsParams {
	o.SetOrganizationName(organizationName)
	return o
}

// SetOrganizationName adds the organizationName to the get orgs params
func (o *GetOrgsParams) SetOrganizationName(organizationName *string) {
	o.OrganizationName = organizationName
}

// WithPageIndex adds the pageIndex to the get orgs params
func (o *GetOrgsParams) WithPageIndex(pageIndex *uint32) *GetOrgsParams {
	o.SetPageIndex(pageIndex)
	return o
}

// SetPageIndex adds the pageIndex to the get orgs params
func (o *GetOrgsParams) SetPageIndex(pageIndex *uint32) {
	o.PageIndex = pageIndex
}

// WithPageSize adds the pageSize to the get orgs params
func (o *GetOrgsParams) WithPageSize(pageSize *uint32) *GetOrgsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get orgs params
func (o *GetOrgsParams) SetPageSize(pageSize *uint32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrgsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.OrderBy != nil {

		// query param order_by
		var qrOrderBy string
		if o.OrderBy != nil {
			qrOrderBy = *o.OrderBy
		}
		qOrderBy := qrOrderBy
		if qOrderBy != "" {
			if err := r.SetQueryParam("order_by", qOrderBy); err != nil {
				return err
			}
		}

	}

	if o.OrganizationCreatedAt != nil {

		// query param organization_created_at
		var qrOrganizationCreatedAt strfmt.DateTime
		if o.OrganizationCreatedAt != nil {
			qrOrganizationCreatedAt = *o.OrganizationCreatedAt
		}
		qOrganizationCreatedAt := qrOrganizationCreatedAt.String()
		if qOrganizationCreatedAt != "" {
			if err := r.SetQueryParam("organization_created_at", qOrganizationCreatedAt); err != nil {
				return err
			}
		}

	}

	if o.OrganizationName != nil {

		// query param organization_name
		var qrOrganizationName string
		if o.OrganizationName != nil {
			qrOrganizationName = *o.OrganizationName
		}
		qOrganizationName := qrOrganizationName
		if qOrganizationName != "" {
			if err := r.SetQueryParam("organization_name", qOrganizationName); err != nil {
				return err
			}
		}

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
