// Code generated by go-swagger; DO NOT EDIT.

package organization_invitations

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

// NewGetInvitationsParams creates a new GetInvitationsParams object
// with the default values initialized.
func NewGetInvitationsParams() *GetInvitationsParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(10)
	)
	return &GetInvitationsParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInvitationsParamsWithTimeout creates a new GetInvitationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInvitationsParamsWithTimeout(timeout time.Duration) *GetInvitationsParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(10)
	)
	return &GetInvitationsParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetInvitationsParamsWithContext creates a new GetInvitationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInvitationsParamsWithContext(ctx context.Context) *GetInvitationsParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(10)
	)
	return &GetInvitationsParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetInvitationsParamsWithHTTPClient creates a new GetInvitationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInvitationsParamsWithHTTPClient(client *http.Client) *GetInvitationsParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(10)
	)
	return &GetInvitationsParams{
		PageIndex:  &pageIndexDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetInvitationsParams contains all the parameters to send to the API endpoint
for the get invitations operation typically these are written to a http.Request
*/
type GetInvitationsParams struct {

	/*InvitationCreatedAt
	  Search by Invitation's creation date

	*/
	InvitationCreatedAt *uint64
	/*InvitationState
	  Search by Invitation's state

	*/
	InvitationState *string
	/*OrderBy
	  Allows to order the list of items. Example usage: field_name:asc


	*/
	OrderBy *string
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

// WithTimeout adds the timeout to the get invitations params
func (o *GetInvitationsParams) WithTimeout(timeout time.Duration) *GetInvitationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invitations params
func (o *GetInvitationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invitations params
func (o *GetInvitationsParams) WithContext(ctx context.Context) *GetInvitationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invitations params
func (o *GetInvitationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invitations params
func (o *GetInvitationsParams) WithHTTPClient(client *http.Client) *GetInvitationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invitations params
func (o *GetInvitationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvitationCreatedAt adds the invitationCreatedAt to the get invitations params
func (o *GetInvitationsParams) WithInvitationCreatedAt(invitationCreatedAt *uint64) *GetInvitationsParams {
	o.SetInvitationCreatedAt(invitationCreatedAt)
	return o
}

// SetInvitationCreatedAt adds the invitationCreatedAt to the get invitations params
func (o *GetInvitationsParams) SetInvitationCreatedAt(invitationCreatedAt *uint64) {
	o.InvitationCreatedAt = invitationCreatedAt
}

// WithInvitationState adds the invitationState to the get invitations params
func (o *GetInvitationsParams) WithInvitationState(invitationState *string) *GetInvitationsParams {
	o.SetInvitationState(invitationState)
	return o
}

// SetInvitationState adds the invitationState to the get invitations params
func (o *GetInvitationsParams) SetInvitationState(invitationState *string) {
	o.InvitationState = invitationState
}

// WithOrderBy adds the orderBy to the get invitations params
func (o *GetInvitationsParams) WithOrderBy(orderBy *string) *GetInvitationsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get invitations params
func (o *GetInvitationsParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithOrganizationCanonical adds the organizationCanonical to the get invitations params
func (o *GetInvitationsParams) WithOrganizationCanonical(organizationCanonical string) *GetInvitationsParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get invitations params
func (o *GetInvitationsParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithPageIndex adds the pageIndex to the get invitations params
func (o *GetInvitationsParams) WithPageIndex(pageIndex *uint32) *GetInvitationsParams {
	o.SetPageIndex(pageIndex)
	return o
}

// SetPageIndex adds the pageIndex to the get invitations params
func (o *GetInvitationsParams) SetPageIndex(pageIndex *uint32) {
	o.PageIndex = pageIndex
}

// WithPageSize adds the pageSize to the get invitations params
func (o *GetInvitationsParams) WithPageSize(pageSize *uint32) *GetInvitationsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get invitations params
func (o *GetInvitationsParams) SetPageSize(pageSize *uint32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvitationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.InvitationCreatedAt != nil {

		// query param invitation_created_at
		var qrInvitationCreatedAt uint64
		if o.InvitationCreatedAt != nil {
			qrInvitationCreatedAt = *o.InvitationCreatedAt
		}
		qInvitationCreatedAt := swag.FormatUint64(qrInvitationCreatedAt)
		if qInvitationCreatedAt != "" {
			if err := r.SetQueryParam("invitation_created_at", qInvitationCreatedAt); err != nil {
				return err
			}
		}

	}

	if o.InvitationState != nil {

		// query param invitation_state
		var qrInvitationState string
		if o.InvitationState != nil {
			qrInvitationState = *o.InvitationState
		}
		qInvitationState := qrInvitationState
		if qInvitationState != "" {
			if err := r.SetQueryParam("invitation_state", qInvitationState); err != nil {
				return err
			}
		}

	}

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
