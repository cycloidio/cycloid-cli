// Code generated by go-swagger; DO NOT EDIT.

package organization_external_backends

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

// NewGetExternalBackendsParams creates a new GetExternalBackendsParams object
// with the default values initialized.
func NewGetExternalBackendsParams() *GetExternalBackendsParams {
	var ()
	return &GetExternalBackendsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetExternalBackendsParamsWithTimeout creates a new GetExternalBackendsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetExternalBackendsParamsWithTimeout(timeout time.Duration) *GetExternalBackendsParams {
	var ()
	return &GetExternalBackendsParams{

		timeout: timeout,
	}
}

// NewGetExternalBackendsParamsWithContext creates a new GetExternalBackendsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetExternalBackendsParamsWithContext(ctx context.Context) *GetExternalBackendsParams {
	var ()
	return &GetExternalBackendsParams{

		Context: ctx,
	}
}

// NewGetExternalBackendsParamsWithHTTPClient creates a new GetExternalBackendsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetExternalBackendsParamsWithHTTPClient(client *http.Client) *GetExternalBackendsParams {
	var ()
	return &GetExternalBackendsParams{
		HTTPClient: client,
	}
}

/*GetExternalBackendsParams contains all the parameters to send to the API endpoint
for the get external backends operation typically these are written to a http.Request
*/
type GetExternalBackendsParams struct {

	/*Environment
	  The environment canonical to use a query filter

	*/
	Environment *string
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string
	/*Pproject
	  A canonical of a project used for filtering.

	*/
	Pproject *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get external backends params
func (o *GetExternalBackendsParams) WithTimeout(timeout time.Duration) *GetExternalBackendsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get external backends params
func (o *GetExternalBackendsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get external backends params
func (o *GetExternalBackendsParams) WithContext(ctx context.Context) *GetExternalBackendsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get external backends params
func (o *GetExternalBackendsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get external backends params
func (o *GetExternalBackendsParams) WithHTTPClient(client *http.Client) *GetExternalBackendsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get external backends params
func (o *GetExternalBackendsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironment adds the environment to the get external backends params
func (o *GetExternalBackendsParams) WithEnvironment(environment *string) *GetExternalBackendsParams {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the get external backends params
func (o *GetExternalBackendsParams) SetEnvironment(environment *string) {
	o.Environment = environment
}

// WithOrganizationCanonical adds the organizationCanonical to the get external backends params
func (o *GetExternalBackendsParams) WithOrganizationCanonical(organizationCanonical string) *GetExternalBackendsParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get external backends params
func (o *GetExternalBackendsParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithPproject adds the pproject to the get external backends params
func (o *GetExternalBackendsParams) WithPproject(pproject *string) *GetExternalBackendsParams {
	o.SetPproject(pproject)
	return o
}

// SetPproject adds the pproject to the get external backends params
func (o *GetExternalBackendsParams) SetPproject(pproject *string) {
	o.Pproject = pproject
}

// WriteToRequest writes these params to a swagger request
func (o *GetExternalBackendsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Environment != nil {

		// query param environment
		var qrEnvironment string
		if o.Environment != nil {
			qrEnvironment = *o.Environment
		}
		qEnvironment := qrEnvironment
		if qEnvironment != "" {
			if err := r.SetQueryParam("environment", qEnvironment); err != nil {
				return err
			}
		}

	}

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	if o.Pproject != nil {

		// query param pproject
		var qrPproject string
		if o.Pproject != nil {
			qrPproject = *o.Pproject
		}
		qPproject := qrPproject
		if qPproject != "" {
			if err := r.SetQueryParam("pproject", qPproject); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
