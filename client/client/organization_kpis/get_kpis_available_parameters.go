// Code generated by go-swagger; DO NOT EDIT.

package organization_kpis

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

// NewGetKpisAvailableParams creates a new GetKpisAvailableParams object
// with the default values initialized.
func NewGetKpisAvailableParams() *GetKpisAvailableParams {
	var ()
	return &GetKpisAvailableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKpisAvailableParamsWithTimeout creates a new GetKpisAvailableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKpisAvailableParamsWithTimeout(timeout time.Duration) *GetKpisAvailableParams {
	var ()
	return &GetKpisAvailableParams{

		timeout: timeout,
	}
}

// NewGetKpisAvailableParamsWithContext creates a new GetKpisAvailableParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKpisAvailableParamsWithContext(ctx context.Context) *GetKpisAvailableParams {
	var ()
	return &GetKpisAvailableParams{

		Context: ctx,
	}
}

// NewGetKpisAvailableParamsWithHTTPClient creates a new GetKpisAvailableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKpisAvailableParamsWithHTTPClient(client *http.Client) *GetKpisAvailableParams {
	var ()
	return &GetKpisAvailableParams{
		HTTPClient: client,
	}
}

/*GetKpisAvailableParams contains all the parameters to send to the API endpoint
for the get kpis available operation typically these are written to a http.Request
*/
type GetKpisAvailableParams struct {

	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get kpis available params
func (o *GetKpisAvailableParams) WithTimeout(timeout time.Duration) *GetKpisAvailableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kpis available params
func (o *GetKpisAvailableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kpis available params
func (o *GetKpisAvailableParams) WithContext(ctx context.Context) *GetKpisAvailableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kpis available params
func (o *GetKpisAvailableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kpis available params
func (o *GetKpisAvailableParams) WithHTTPClient(client *http.Client) *GetKpisAvailableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kpis available params
func (o *GetKpisAvailableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationCanonical adds the organizationCanonical to the get kpis available params
func (o *GetKpisAvailableParams) WithOrganizationCanonical(organizationCanonical string) *GetKpisAvailableParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get kpis available params
func (o *GetKpisAvailableParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *GetKpisAvailableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
