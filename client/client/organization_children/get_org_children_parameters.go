// Code generated by go-swagger; DO NOT EDIT.

package organization_children

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

// NewGetOrgChildrenParams creates a new GetOrgChildrenParams object
// with the default values initialized.
func NewGetOrgChildrenParams() *GetOrgChildrenParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &GetOrgChildrenParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrgChildrenParamsWithTimeout creates a new GetOrgChildrenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrgChildrenParamsWithTimeout(timeout time.Duration) *GetOrgChildrenParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &GetOrgChildrenParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetOrgChildrenParamsWithContext creates a new GetOrgChildrenParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrgChildrenParamsWithContext(ctx context.Context) *GetOrgChildrenParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &GetOrgChildrenParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetOrgChildrenParamsWithHTTPClient creates a new GetOrgChildrenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrgChildrenParamsWithHTTPClient(client *http.Client) *GetOrgChildrenParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &GetOrgChildrenParams{
		PageIndex:  &pageIndexDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetOrgChildrenParams contains all the parameters to send to the API endpoint
for the get org children operation typically these are written to a http.Request
*/
type GetOrgChildrenParams struct {

	/*OrderBy
	  Allows to order the list of items. Example usage: field_name:asc


	*/
	OrderBy *string
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string
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

// WithTimeout adds the timeout to the get org children params
func (o *GetOrgChildrenParams) WithTimeout(timeout time.Duration) *GetOrgChildrenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get org children params
func (o *GetOrgChildrenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get org children params
func (o *GetOrgChildrenParams) WithContext(ctx context.Context) *GetOrgChildrenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get org children params
func (o *GetOrgChildrenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get org children params
func (o *GetOrgChildrenParams) WithHTTPClient(client *http.Client) *GetOrgChildrenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get org children params
func (o *GetOrgChildrenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderBy adds the orderBy to the get org children params
func (o *GetOrgChildrenParams) WithOrderBy(orderBy *string) *GetOrgChildrenParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get org children params
func (o *GetOrgChildrenParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithOrganizationCanonical adds the organizationCanonical to the get org children params
func (o *GetOrgChildrenParams) WithOrganizationCanonical(organizationCanonical string) *GetOrgChildrenParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get org children params
func (o *GetOrgChildrenParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithOrganizationCreatedAt adds the organizationCreatedAt to the get org children params
func (o *GetOrgChildrenParams) WithOrganizationCreatedAt(organizationCreatedAt *strfmt.DateTime) *GetOrgChildrenParams {
	o.SetOrganizationCreatedAt(organizationCreatedAt)
	return o
}

// SetOrganizationCreatedAt adds the organizationCreatedAt to the get org children params
func (o *GetOrgChildrenParams) SetOrganizationCreatedAt(organizationCreatedAt *strfmt.DateTime) {
	o.OrganizationCreatedAt = organizationCreatedAt
}

// WithOrganizationName adds the organizationName to the get org children params
func (o *GetOrgChildrenParams) WithOrganizationName(organizationName *string) *GetOrgChildrenParams {
	o.SetOrganizationName(organizationName)
	return o
}

// SetOrganizationName adds the organizationName to the get org children params
func (o *GetOrgChildrenParams) SetOrganizationName(organizationName *string) {
	o.OrganizationName = organizationName
}

// WithPageIndex adds the pageIndex to the get org children params
func (o *GetOrgChildrenParams) WithPageIndex(pageIndex *uint32) *GetOrgChildrenParams {
	o.SetPageIndex(pageIndex)
	return o
}

// SetPageIndex adds the pageIndex to the get org children params
func (o *GetOrgChildrenParams) SetPageIndex(pageIndex *uint32) {
	o.PageIndex = pageIndex
}

// WithPageSize adds the pageSize to the get org children params
func (o *GetOrgChildrenParams) WithPageSize(pageSize *uint32) *GetOrgChildrenParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get org children params
func (o *GetOrgChildrenParams) SetPageSize(pageSize *uint32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrgChildrenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
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
