// Code generated by go-swagger; DO NOT EDIT.

package organization_api_keys

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

// NewGetAPIKeysParams creates a new GetAPIKeysParams object
// with the default values initialized.
func NewGetAPIKeysParams() *GetAPIKeysParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &GetAPIKeysParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIKeysParamsWithTimeout creates a new GetAPIKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIKeysParamsWithTimeout(timeout time.Duration) *GetAPIKeysParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &GetAPIKeysParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetAPIKeysParamsWithContext creates a new GetAPIKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIKeysParamsWithContext(ctx context.Context) *GetAPIKeysParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &GetAPIKeysParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetAPIKeysParamsWithHTTPClient creates a new GetAPIKeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIKeysParamsWithHTTPClient(client *http.Client) *GetAPIKeysParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &GetAPIKeysParams{
		PageIndex:  &pageIndexDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetAPIKeysParams contains all the parameters to send to the API endpoint
for the get API keys operation typically these are written to a http.Request
*/
type GetAPIKeysParams struct {

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
	/*UserID
	  Search by entity's owner

	*/
	UserID *uint32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API keys params
func (o *GetAPIKeysParams) WithTimeout(timeout time.Duration) *GetAPIKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API keys params
func (o *GetAPIKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API keys params
func (o *GetAPIKeysParams) WithContext(ctx context.Context) *GetAPIKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API keys params
func (o *GetAPIKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API keys params
func (o *GetAPIKeysParams) WithHTTPClient(client *http.Client) *GetAPIKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API keys params
func (o *GetAPIKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationCanonical adds the organizationCanonical to the get API keys params
func (o *GetAPIKeysParams) WithOrganizationCanonical(organizationCanonical string) *GetAPIKeysParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get API keys params
func (o *GetAPIKeysParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithPageIndex adds the pageIndex to the get API keys params
func (o *GetAPIKeysParams) WithPageIndex(pageIndex *uint32) *GetAPIKeysParams {
	o.SetPageIndex(pageIndex)
	return o
}

// SetPageIndex adds the pageIndex to the get API keys params
func (o *GetAPIKeysParams) SetPageIndex(pageIndex *uint32) {
	o.PageIndex = pageIndex
}

// WithPageSize adds the pageSize to the get API keys params
func (o *GetAPIKeysParams) WithPageSize(pageSize *uint32) *GetAPIKeysParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get API keys params
func (o *GetAPIKeysParams) SetPageSize(pageSize *uint32) {
	o.PageSize = pageSize
}

// WithUserID adds the userID to the get API keys params
func (o *GetAPIKeysParams) WithUserID(userID *uint32) *GetAPIKeysParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get API keys params
func (o *GetAPIKeysParams) SetUserID(userID *uint32) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.UserID != nil {

		// query param user_id
		var qrUserID uint32
		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := swag.FormatUint32(qrUserID)
		if qUserID != "" {
			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
