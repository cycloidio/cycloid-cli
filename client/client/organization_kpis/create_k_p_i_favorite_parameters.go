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
	"github.com/go-openapi/strfmt"
)

// NewCreateKPIFavoriteParams creates a new CreateKPIFavoriteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateKPIFavoriteParams() *CreateKPIFavoriteParams {
	return &CreateKPIFavoriteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateKPIFavoriteParamsWithTimeout creates a new CreateKPIFavoriteParams object
// with the ability to set a timeout on a request.
func NewCreateKPIFavoriteParamsWithTimeout(timeout time.Duration) *CreateKPIFavoriteParams {
	return &CreateKPIFavoriteParams{
		timeout: timeout,
	}
}

// NewCreateKPIFavoriteParamsWithContext creates a new CreateKPIFavoriteParams object
// with the ability to set a context for a request.
func NewCreateKPIFavoriteParamsWithContext(ctx context.Context) *CreateKPIFavoriteParams {
	return &CreateKPIFavoriteParams{
		Context: ctx,
	}
}

// NewCreateKPIFavoriteParamsWithHTTPClient creates a new CreateKPIFavoriteParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateKPIFavoriteParamsWithHTTPClient(client *http.Client) *CreateKPIFavoriteParams {
	return &CreateKPIFavoriteParams{
		HTTPClient: client,
	}
}

/*
CreateKPIFavoriteParams contains all the parameters to send to the API endpoint

	for the create k p i favorite operation.

	Typically these are written to a http.Request.
*/
type CreateKPIFavoriteParams struct {

	/* KpiCanonical.

	   A canonical of a kpi.
	*/
	KpiCanonical string

	/* OrganizationCanonical.

	   A canonical of an organization.
	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create k p i favorite params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateKPIFavoriteParams) WithDefaults() *CreateKPIFavoriteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create k p i favorite params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateKPIFavoriteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create k p i favorite params
func (o *CreateKPIFavoriteParams) WithTimeout(timeout time.Duration) *CreateKPIFavoriteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create k p i favorite params
func (o *CreateKPIFavoriteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create k p i favorite params
func (o *CreateKPIFavoriteParams) WithContext(ctx context.Context) *CreateKPIFavoriteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create k p i favorite params
func (o *CreateKPIFavoriteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create k p i favorite params
func (o *CreateKPIFavoriteParams) WithHTTPClient(client *http.Client) *CreateKPIFavoriteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create k p i favorite params
func (o *CreateKPIFavoriteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKpiCanonical adds the kpiCanonical to the create k p i favorite params
func (o *CreateKPIFavoriteParams) WithKpiCanonical(kpiCanonical string) *CreateKPIFavoriteParams {
	o.SetKpiCanonical(kpiCanonical)
	return o
}

// SetKpiCanonical adds the kpiCanonical to the create k p i favorite params
func (o *CreateKPIFavoriteParams) SetKpiCanonical(kpiCanonical string) {
	o.KpiCanonical = kpiCanonical
}

// WithOrganizationCanonical adds the organizationCanonical to the create k p i favorite params
func (o *CreateKPIFavoriteParams) WithOrganizationCanonical(organizationCanonical string) *CreateKPIFavoriteParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the create k p i favorite params
func (o *CreateKPIFavoriteParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *CreateKPIFavoriteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param kpi_canonical
	if err := r.SetPathParam("kpi_canonical", o.KpiCanonical); err != nil {
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
