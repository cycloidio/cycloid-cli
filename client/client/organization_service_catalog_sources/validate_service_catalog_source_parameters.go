// Code generated by go-swagger; DO NOT EDIT.

package organization_service_catalog_sources

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

// NewValidateServiceCatalogSourceParams creates a new ValidateServiceCatalogSourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateServiceCatalogSourceParams() *ValidateServiceCatalogSourceParams {
	return &ValidateServiceCatalogSourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateServiceCatalogSourceParamsWithTimeout creates a new ValidateServiceCatalogSourceParams object
// with the ability to set a timeout on a request.
func NewValidateServiceCatalogSourceParamsWithTimeout(timeout time.Duration) *ValidateServiceCatalogSourceParams {
	return &ValidateServiceCatalogSourceParams{
		timeout: timeout,
	}
}

// NewValidateServiceCatalogSourceParamsWithContext creates a new ValidateServiceCatalogSourceParams object
// with the ability to set a context for a request.
func NewValidateServiceCatalogSourceParamsWithContext(ctx context.Context) *ValidateServiceCatalogSourceParams {
	return &ValidateServiceCatalogSourceParams{
		Context: ctx,
	}
}

// NewValidateServiceCatalogSourceParamsWithHTTPClient creates a new ValidateServiceCatalogSourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateServiceCatalogSourceParamsWithHTTPClient(client *http.Client) *ValidateServiceCatalogSourceParams {
	return &ValidateServiceCatalogSourceParams{
		HTTPClient: client,
	}
}

/*
ValidateServiceCatalogSourceParams contains all the parameters to send to the API endpoint

	for the validate service catalog source operation.

	Typically these are written to a http.Request.
*/
type ValidateServiceCatalogSourceParams struct {

	/* OrganizationCanonical.

	   A canonical of an organization.
	*/
	OrganizationCanonical string

	/* ServiceCatalogSourceCanonical.

	   Organization Service Catalog Sources canonical
	*/
	ServiceCatalogSourceCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate service catalog source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateServiceCatalogSourceParams) WithDefaults() *ValidateServiceCatalogSourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate service catalog source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateServiceCatalogSourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate service catalog source params
func (o *ValidateServiceCatalogSourceParams) WithTimeout(timeout time.Duration) *ValidateServiceCatalogSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate service catalog source params
func (o *ValidateServiceCatalogSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate service catalog source params
func (o *ValidateServiceCatalogSourceParams) WithContext(ctx context.Context) *ValidateServiceCatalogSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate service catalog source params
func (o *ValidateServiceCatalogSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate service catalog source params
func (o *ValidateServiceCatalogSourceParams) WithHTTPClient(client *http.Client) *ValidateServiceCatalogSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate service catalog source params
func (o *ValidateServiceCatalogSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationCanonical adds the organizationCanonical to the validate service catalog source params
func (o *ValidateServiceCatalogSourceParams) WithOrganizationCanonical(organizationCanonical string) *ValidateServiceCatalogSourceParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the validate service catalog source params
func (o *ValidateServiceCatalogSourceParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithServiceCatalogSourceCanonical adds the serviceCatalogSourceCanonical to the validate service catalog source params
func (o *ValidateServiceCatalogSourceParams) WithServiceCatalogSourceCanonical(serviceCatalogSourceCanonical string) *ValidateServiceCatalogSourceParams {
	o.SetServiceCatalogSourceCanonical(serviceCatalogSourceCanonical)
	return o
}

// SetServiceCatalogSourceCanonical adds the serviceCatalogSourceCanonical to the validate service catalog source params
func (o *ValidateServiceCatalogSourceParams) SetServiceCatalogSourceCanonical(serviceCatalogSourceCanonical string) {
	o.ServiceCatalogSourceCanonical = serviceCatalogSourceCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateServiceCatalogSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	// path param service_catalog_source_canonical
	if err := r.SetPathParam("service_catalog_source_canonical", o.ServiceCatalogSourceCanonical); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
