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
)

// NewGetWatchRuleParams creates a new GetWatchRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWatchRuleParams() *GetWatchRuleParams {
	return &GetWatchRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWatchRuleParamsWithTimeout creates a new GetWatchRuleParams object
// with the ability to set a timeout on a request.
func NewGetWatchRuleParamsWithTimeout(timeout time.Duration) *GetWatchRuleParams {
	return &GetWatchRuleParams{
		timeout: timeout,
	}
}

// NewGetWatchRuleParamsWithContext creates a new GetWatchRuleParams object
// with the ability to set a context for a request.
func NewGetWatchRuleParamsWithContext(ctx context.Context) *GetWatchRuleParams {
	return &GetWatchRuleParams{
		Context: ctx,
	}
}

// NewGetWatchRuleParamsWithHTTPClient creates a new GetWatchRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWatchRuleParamsWithHTTPClient(client *http.Client) *GetWatchRuleParams {
	return &GetWatchRuleParams{
		HTTPClient: client,
	}
}

/*
GetWatchRuleParams contains all the parameters to send to the API endpoint

	for the get watch rule operation.

	Typically these are written to a http.Request.
*/
type GetWatchRuleParams struct {

	/* WatchRuleCanonical.

	   A canonical of a watch rule.
	*/
	WatchRuleCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get watch rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWatchRuleParams) WithDefaults() *GetWatchRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get watch rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWatchRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get watch rule params
func (o *GetWatchRuleParams) WithTimeout(timeout time.Duration) *GetWatchRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get watch rule params
func (o *GetWatchRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get watch rule params
func (o *GetWatchRuleParams) WithContext(ctx context.Context) *GetWatchRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get watch rule params
func (o *GetWatchRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get watch rule params
func (o *GetWatchRuleParams) WithHTTPClient(client *http.Client) *GetWatchRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get watch rule params
func (o *GetWatchRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWatchRuleCanonical adds the watchRuleCanonical to the get watch rule params
func (o *GetWatchRuleParams) WithWatchRuleCanonical(watchRuleCanonical string) *GetWatchRuleParams {
	o.SetWatchRuleCanonical(watchRuleCanonical)
	return o
}

// SetWatchRuleCanonical adds the watchRuleCanonical to the get watch rule params
func (o *GetWatchRuleParams) SetWatchRuleCanonical(watchRuleCanonical string) {
	o.WatchRuleCanonical = watchRuleCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *GetWatchRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param watch_rule_canonical
	if err := r.SetPathParam("watch_rule_canonical", o.WatchRuleCanonical); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
