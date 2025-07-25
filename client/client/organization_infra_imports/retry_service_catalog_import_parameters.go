// Code generated by go-swagger; DO NOT EDIT.

package organization_infra_imports

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

// NewRetryServiceCatalogImportParams creates a new RetryServiceCatalogImportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRetryServiceCatalogImportParams() *RetryServiceCatalogImportParams {
	return &RetryServiceCatalogImportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRetryServiceCatalogImportParamsWithTimeout creates a new RetryServiceCatalogImportParams object
// with the ability to set a timeout on a request.
func NewRetryServiceCatalogImportParamsWithTimeout(timeout time.Duration) *RetryServiceCatalogImportParams {
	return &RetryServiceCatalogImportParams{
		timeout: timeout,
	}
}

// NewRetryServiceCatalogImportParamsWithContext creates a new RetryServiceCatalogImportParams object
// with the ability to set a context for a request.
func NewRetryServiceCatalogImportParamsWithContext(ctx context.Context) *RetryServiceCatalogImportParams {
	return &RetryServiceCatalogImportParams{
		Context: ctx,
	}
}

// NewRetryServiceCatalogImportParamsWithHTTPClient creates a new RetryServiceCatalogImportParams object
// with the ability to set a custom HTTPClient for a request.
func NewRetryServiceCatalogImportParamsWithHTTPClient(client *http.Client) *RetryServiceCatalogImportParams {
	return &RetryServiceCatalogImportParams{
		HTTPClient: client,
	}
}

/*
RetryServiceCatalogImportParams contains all the parameters to send to the API endpoint

	for the retry service catalog import operation.

	Typically these are written to a http.Request.
*/
type RetryServiceCatalogImportParams struct {

	/* OrganizationCanonical.

	   A canonical of an organization.
	*/
	OrganizationCanonical string

	/* ServiceCatalogRef.

	   A Service Catalog name
	*/
	ServiceCatalogRef string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the retry service catalog import params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetryServiceCatalogImportParams) WithDefaults() *RetryServiceCatalogImportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the retry service catalog import params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetryServiceCatalogImportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the retry service catalog import params
func (o *RetryServiceCatalogImportParams) WithTimeout(timeout time.Duration) *RetryServiceCatalogImportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retry service catalog import params
func (o *RetryServiceCatalogImportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retry service catalog import params
func (o *RetryServiceCatalogImportParams) WithContext(ctx context.Context) *RetryServiceCatalogImportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retry service catalog import params
func (o *RetryServiceCatalogImportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retry service catalog import params
func (o *RetryServiceCatalogImportParams) WithHTTPClient(client *http.Client) *RetryServiceCatalogImportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retry service catalog import params
func (o *RetryServiceCatalogImportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationCanonical adds the organizationCanonical to the retry service catalog import params
func (o *RetryServiceCatalogImportParams) WithOrganizationCanonical(organizationCanonical string) *RetryServiceCatalogImportParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the retry service catalog import params
func (o *RetryServiceCatalogImportParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithServiceCatalogRef adds the serviceCatalogRef to the retry service catalog import params
func (o *RetryServiceCatalogImportParams) WithServiceCatalogRef(serviceCatalogRef string) *RetryServiceCatalogImportParams {
	o.SetServiceCatalogRef(serviceCatalogRef)
	return o
}

// SetServiceCatalogRef adds the serviceCatalogRef to the retry service catalog import params
func (o *RetryServiceCatalogImportParams) SetServiceCatalogRef(serviceCatalogRef string) {
	o.ServiceCatalogRef = serviceCatalogRef
}

// WriteToRequest writes these params to a swagger request
func (o *RetryServiceCatalogImportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	// path param service_catalog_ref
	if err := r.SetPathParam("service_catalog_ref", o.ServiceCatalogRef); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
