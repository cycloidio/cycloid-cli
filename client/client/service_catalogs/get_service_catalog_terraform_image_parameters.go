// Code generated by go-swagger; DO NOT EDIT.

package service_catalogs

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

// NewGetServiceCatalogTerraformImageParams creates a new GetServiceCatalogTerraformImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetServiceCatalogTerraformImageParams() *GetServiceCatalogTerraformImageParams {
	return &GetServiceCatalogTerraformImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceCatalogTerraformImageParamsWithTimeout creates a new GetServiceCatalogTerraformImageParams object
// with the ability to set a timeout on a request.
func NewGetServiceCatalogTerraformImageParamsWithTimeout(timeout time.Duration) *GetServiceCatalogTerraformImageParams {
	return &GetServiceCatalogTerraformImageParams{
		timeout: timeout,
	}
}

// NewGetServiceCatalogTerraformImageParamsWithContext creates a new GetServiceCatalogTerraformImageParams object
// with the ability to set a context for a request.
func NewGetServiceCatalogTerraformImageParamsWithContext(ctx context.Context) *GetServiceCatalogTerraformImageParams {
	return &GetServiceCatalogTerraformImageParams{
		Context: ctx,
	}
}

// NewGetServiceCatalogTerraformImageParamsWithHTTPClient creates a new GetServiceCatalogTerraformImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetServiceCatalogTerraformImageParamsWithHTTPClient(client *http.Client) *GetServiceCatalogTerraformImageParams {
	return &GetServiceCatalogTerraformImageParams{
		HTTPClient: client,
	}
}

/*
GetServiceCatalogTerraformImageParams contains all the parameters to send to the API endpoint

	for the get service catalog terraform image operation.

	Typically these are written to a http.Request.
*/
type GetServiceCatalogTerraformImageParams struct {

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

// WithDefaults hydrates default values in the get service catalog terraform image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServiceCatalogTerraformImageParams) WithDefaults() *GetServiceCatalogTerraformImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get service catalog terraform image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServiceCatalogTerraformImageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get service catalog terraform image params
func (o *GetServiceCatalogTerraformImageParams) WithTimeout(timeout time.Duration) *GetServiceCatalogTerraformImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service catalog terraform image params
func (o *GetServiceCatalogTerraformImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service catalog terraform image params
func (o *GetServiceCatalogTerraformImageParams) WithContext(ctx context.Context) *GetServiceCatalogTerraformImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service catalog terraform image params
func (o *GetServiceCatalogTerraformImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service catalog terraform image params
func (o *GetServiceCatalogTerraformImageParams) WithHTTPClient(client *http.Client) *GetServiceCatalogTerraformImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service catalog terraform image params
func (o *GetServiceCatalogTerraformImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationCanonical adds the organizationCanonical to the get service catalog terraform image params
func (o *GetServiceCatalogTerraformImageParams) WithOrganizationCanonical(organizationCanonical string) *GetServiceCatalogTerraformImageParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get service catalog terraform image params
func (o *GetServiceCatalogTerraformImageParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithServiceCatalogRef adds the serviceCatalogRef to the get service catalog terraform image params
func (o *GetServiceCatalogTerraformImageParams) WithServiceCatalogRef(serviceCatalogRef string) *GetServiceCatalogTerraformImageParams {
	o.SetServiceCatalogRef(serviceCatalogRef)
	return o
}

// SetServiceCatalogRef adds the serviceCatalogRef to the get service catalog terraform image params
func (o *GetServiceCatalogTerraformImageParams) SetServiceCatalogRef(serviceCatalogRef string) {
	o.ServiceCatalogRef = serviceCatalogRef
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceCatalogTerraformImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
