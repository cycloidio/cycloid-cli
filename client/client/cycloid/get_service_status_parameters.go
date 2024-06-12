// Code generated by go-swagger; DO NOT EDIT.

package cycloid

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
)

// NewGetServiceStatusParams creates a new GetServiceStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetServiceStatusParams() *GetServiceStatusParams {
	return &GetServiceStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceStatusParamsWithTimeout creates a new GetServiceStatusParams object
// with the ability to set a timeout on a request.
func NewGetServiceStatusParamsWithTimeout(timeout time.Duration) *GetServiceStatusParams {
	return &GetServiceStatusParams{
		timeout: timeout,
	}
}

// NewGetServiceStatusParamsWithContext creates a new GetServiceStatusParams object
// with the ability to set a context for a request.
func NewGetServiceStatusParamsWithContext(ctx context.Context) *GetServiceStatusParams {
	return &GetServiceStatusParams{
		Context: ctx,
	}
}

// NewGetServiceStatusParamsWithHTTPClient creates a new GetServiceStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetServiceStatusParamsWithHTTPClient(client *http.Client) *GetServiceStatusParams {
	return &GetServiceStatusParams{
		HTTPClient: client,
	}
}

/*
GetServiceStatusParams contains all the parameters to send to the API endpoint

	for the get service status operation.

	Typically these are written to a http.Request.
*/
type GetServiceStatusParams struct {

	/* ServiceStatusCanonical.

	   The canonical of the service you want to get the status from
	*/
	ServiceStatusCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get service status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServiceStatusParams) WithDefaults() *GetServiceStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get service status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServiceStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get service status params
func (o *GetServiceStatusParams) WithTimeout(timeout time.Duration) *GetServiceStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service status params
func (o *GetServiceStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service status params
func (o *GetServiceStatusParams) WithContext(ctx context.Context) *GetServiceStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service status params
func (o *GetServiceStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service status params
func (o *GetServiceStatusParams) WithHTTPClient(client *http.Client) *GetServiceStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service status params
func (o *GetServiceStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceStatusCanonical adds the serviceStatusCanonical to the get service status params
func (o *GetServiceStatusParams) WithServiceStatusCanonical(serviceStatusCanonical string) *GetServiceStatusParams {
	o.SetServiceStatusCanonical(serviceStatusCanonical)
	return o
}

// SetServiceStatusCanonical adds the serviceStatusCanonical to the get service status params
func (o *GetServiceStatusParams) SetServiceStatusCanonical(serviceStatusCanonical string) {
	o.ServiceStatusCanonical = serviceStatusCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param service_status_canonical
	if err := r.SetPathParam("service_status_canonical", o.ServiceStatusCanonical); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
