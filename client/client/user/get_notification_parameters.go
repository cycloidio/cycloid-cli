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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetNotificationParams creates a new GetNotificationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNotificationParams() *GetNotificationParams {
	return &GetNotificationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNotificationParamsWithTimeout creates a new GetNotificationParams object
// with the ability to set a timeout on a request.
func NewGetNotificationParamsWithTimeout(timeout time.Duration) *GetNotificationParams {
	return &GetNotificationParams{
		timeout: timeout,
	}
}

// NewGetNotificationParamsWithContext creates a new GetNotificationParams object
// with the ability to set a context for a request.
func NewGetNotificationParamsWithContext(ctx context.Context) *GetNotificationParams {
	return &GetNotificationParams{
		Context: ctx,
	}
}

// NewGetNotificationParamsWithHTTPClient creates a new GetNotificationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNotificationParamsWithHTTPClient(client *http.Client) *GetNotificationParams {
	return &GetNotificationParams{
		HTTPClient: client,
	}
}

/*
GetNotificationParams contains all the parameters to send to the API endpoint

	for the get notification operation.

	Typically these are written to a http.Request.
*/
type GetNotificationParams struct {

	/* NotificationID.

	   Notification ID

	   Format: uint32
	*/
	NotificationID uint32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get notification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNotificationParams) WithDefaults() *GetNotificationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get notification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNotificationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get notification params
func (o *GetNotificationParams) WithTimeout(timeout time.Duration) *GetNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get notification params
func (o *GetNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get notification params
func (o *GetNotificationParams) WithContext(ctx context.Context) *GetNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get notification params
func (o *GetNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get notification params
func (o *GetNotificationParams) WithHTTPClient(client *http.Client) *GetNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get notification params
func (o *GetNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNotificationID adds the notificationID to the get notification params
func (o *GetNotificationParams) WithNotificationID(notificationID uint32) *GetNotificationParams {
	o.SetNotificationID(notificationID)
	return o
}

// SetNotificationID adds the notificationId to the get notification params
func (o *GetNotificationParams) SetNotificationID(notificationID uint32) {
	o.NotificationID = notificationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param notification_id
	if err := r.SetPathParam("notification_id", swag.FormatUint32(o.NotificationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
