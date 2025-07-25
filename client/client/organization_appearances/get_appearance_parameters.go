// Code generated by go-swagger; DO NOT EDIT.

package organization_appearances

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

// NewGetAppearanceParams creates a new GetAppearanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAppearanceParams() *GetAppearanceParams {
	return &GetAppearanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppearanceParamsWithTimeout creates a new GetAppearanceParams object
// with the ability to set a timeout on a request.
func NewGetAppearanceParamsWithTimeout(timeout time.Duration) *GetAppearanceParams {
	return &GetAppearanceParams{
		timeout: timeout,
	}
}

// NewGetAppearanceParamsWithContext creates a new GetAppearanceParams object
// with the ability to set a context for a request.
func NewGetAppearanceParamsWithContext(ctx context.Context) *GetAppearanceParams {
	return &GetAppearanceParams{
		Context: ctx,
	}
}

// NewGetAppearanceParamsWithHTTPClient creates a new GetAppearanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAppearanceParamsWithHTTPClient(client *http.Client) *GetAppearanceParams {
	return &GetAppearanceParams{
		HTTPClient: client,
	}
}

/*
GetAppearanceParams contains all the parameters to send to the API endpoint

	for the get appearance operation.

	Typically these are written to a http.Request.
*/
type GetAppearanceParams struct {

	/* AppearanceCanonical.

	   A canonical of an appearance.
	*/
	AppearanceCanonical string

	/* OrganizationCanonical.

	   A canonical of an organization.
	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get appearance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppearanceParams) WithDefaults() *GetAppearanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get appearance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppearanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get appearance params
func (o *GetAppearanceParams) WithTimeout(timeout time.Duration) *GetAppearanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get appearance params
func (o *GetAppearanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get appearance params
func (o *GetAppearanceParams) WithContext(ctx context.Context) *GetAppearanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get appearance params
func (o *GetAppearanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get appearance params
func (o *GetAppearanceParams) WithHTTPClient(client *http.Client) *GetAppearanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get appearance params
func (o *GetAppearanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppearanceCanonical adds the appearanceCanonical to the get appearance params
func (o *GetAppearanceParams) WithAppearanceCanonical(appearanceCanonical string) *GetAppearanceParams {
	o.SetAppearanceCanonical(appearanceCanonical)
	return o
}

// SetAppearanceCanonical adds the appearanceCanonical to the get appearance params
func (o *GetAppearanceParams) SetAppearanceCanonical(appearanceCanonical string) {
	o.AppearanceCanonical = appearanceCanonical
}

// WithOrganizationCanonical adds the organizationCanonical to the get appearance params
func (o *GetAppearanceParams) WithOrganizationCanonical(organizationCanonical string) *GetAppearanceParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get appearance params
func (o *GetAppearanceParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppearanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appearance_canonical
	if err := r.SetPathParam("appearance_canonical", o.AppearanceCanonical); err != nil {
		return err
	}

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
