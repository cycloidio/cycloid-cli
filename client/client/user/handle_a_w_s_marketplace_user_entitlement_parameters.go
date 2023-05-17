// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewHandleAWSMarketplaceUserEntitlementParams creates a new HandleAWSMarketplaceUserEntitlementParams object
// with the default values initialized.
func NewHandleAWSMarketplaceUserEntitlementParams() *HandleAWSMarketplaceUserEntitlementParams {

	return &HandleAWSMarketplaceUserEntitlementParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHandleAWSMarketplaceUserEntitlementParamsWithTimeout creates a new HandleAWSMarketplaceUserEntitlementParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHandleAWSMarketplaceUserEntitlementParamsWithTimeout(timeout time.Duration) *HandleAWSMarketplaceUserEntitlementParams {

	return &HandleAWSMarketplaceUserEntitlementParams{

		timeout: timeout,
	}
}

// NewHandleAWSMarketplaceUserEntitlementParamsWithContext creates a new HandleAWSMarketplaceUserEntitlementParams object
// with the default values initialized, and the ability to set a context for a request
func NewHandleAWSMarketplaceUserEntitlementParamsWithContext(ctx context.Context) *HandleAWSMarketplaceUserEntitlementParams {

	return &HandleAWSMarketplaceUserEntitlementParams{

		Context: ctx,
	}
}

// NewHandleAWSMarketplaceUserEntitlementParamsWithHTTPClient creates a new HandleAWSMarketplaceUserEntitlementParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHandleAWSMarketplaceUserEntitlementParamsWithHTTPClient(client *http.Client) *HandleAWSMarketplaceUserEntitlementParams {

	return &HandleAWSMarketplaceUserEntitlementParams{
		HTTPClient: client,
	}
}

/*HandleAWSMarketplaceUserEntitlementParams contains all the parameters to send to the API endpoint
for the handle a w s marketplace user entitlement operation typically these are written to a http.Request
*/
type HandleAWSMarketplaceUserEntitlementParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the handle a w s marketplace user entitlement params
func (o *HandleAWSMarketplaceUserEntitlementParams) WithTimeout(timeout time.Duration) *HandleAWSMarketplaceUserEntitlementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the handle a w s marketplace user entitlement params
func (o *HandleAWSMarketplaceUserEntitlementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the handle a w s marketplace user entitlement params
func (o *HandleAWSMarketplaceUserEntitlementParams) WithContext(ctx context.Context) *HandleAWSMarketplaceUserEntitlementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the handle a w s marketplace user entitlement params
func (o *HandleAWSMarketplaceUserEntitlementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the handle a w s marketplace user entitlement params
func (o *HandleAWSMarketplaceUserEntitlementParams) WithHTTPClient(client *http.Client) *HandleAWSMarketplaceUserEntitlementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the handle a w s marketplace user entitlement params
func (o *HandleAWSMarketplaceUserEntitlementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HandleAWSMarketplaceUserEntitlementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
