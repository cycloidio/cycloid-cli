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

	"github.com/cycloidio/cycloid-cli/client/models"
)

// NewCreateServiceCatalogFromTemplateParams creates a new CreateServiceCatalogFromTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateServiceCatalogFromTemplateParams() *CreateServiceCatalogFromTemplateParams {
	return &CreateServiceCatalogFromTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateServiceCatalogFromTemplateParamsWithTimeout creates a new CreateServiceCatalogFromTemplateParams object
// with the ability to set a timeout on a request.
func NewCreateServiceCatalogFromTemplateParamsWithTimeout(timeout time.Duration) *CreateServiceCatalogFromTemplateParams {
	return &CreateServiceCatalogFromTemplateParams{
		timeout: timeout,
	}
}

// NewCreateServiceCatalogFromTemplateParamsWithContext creates a new CreateServiceCatalogFromTemplateParams object
// with the ability to set a context for a request.
func NewCreateServiceCatalogFromTemplateParamsWithContext(ctx context.Context) *CreateServiceCatalogFromTemplateParams {
	return &CreateServiceCatalogFromTemplateParams{
		Context: ctx,
	}
}

// NewCreateServiceCatalogFromTemplateParamsWithHTTPClient creates a new CreateServiceCatalogFromTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateServiceCatalogFromTemplateParamsWithHTTPClient(client *http.Client) *CreateServiceCatalogFromTemplateParams {
	return &CreateServiceCatalogFromTemplateParams{
		HTTPClient: client,
	}
}

/*
CreateServiceCatalogFromTemplateParams contains all the parameters to send to the API endpoint

	for the create service catalog from template operation.

	Typically these are written to a http.Request.
*/
type CreateServiceCatalogFromTemplateParams struct {

	/* Body.

	   The information of the ServiceCatalog.
	*/
	Body *models.NewServiceCatalogFromTemplate

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

// WithDefaults hydrates default values in the create service catalog from template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateServiceCatalogFromTemplateParams) WithDefaults() *CreateServiceCatalogFromTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create service catalog from template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateServiceCatalogFromTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create service catalog from template params
func (o *CreateServiceCatalogFromTemplateParams) WithTimeout(timeout time.Duration) *CreateServiceCatalogFromTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create service catalog from template params
func (o *CreateServiceCatalogFromTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create service catalog from template params
func (o *CreateServiceCatalogFromTemplateParams) WithContext(ctx context.Context) *CreateServiceCatalogFromTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create service catalog from template params
func (o *CreateServiceCatalogFromTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create service catalog from template params
func (o *CreateServiceCatalogFromTemplateParams) WithHTTPClient(client *http.Client) *CreateServiceCatalogFromTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create service catalog from template params
func (o *CreateServiceCatalogFromTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create service catalog from template params
func (o *CreateServiceCatalogFromTemplateParams) WithBody(body *models.NewServiceCatalogFromTemplate) *CreateServiceCatalogFromTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create service catalog from template params
func (o *CreateServiceCatalogFromTemplateParams) SetBody(body *models.NewServiceCatalogFromTemplate) {
	o.Body = body
}

// WithOrganizationCanonical adds the organizationCanonical to the create service catalog from template params
func (o *CreateServiceCatalogFromTemplateParams) WithOrganizationCanonical(organizationCanonical string) *CreateServiceCatalogFromTemplateParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the create service catalog from template params
func (o *CreateServiceCatalogFromTemplateParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithServiceCatalogRef adds the serviceCatalogRef to the create service catalog from template params
func (o *CreateServiceCatalogFromTemplateParams) WithServiceCatalogRef(serviceCatalogRef string) *CreateServiceCatalogFromTemplateParams {
	o.SetServiceCatalogRef(serviceCatalogRef)
	return o
}

// SetServiceCatalogRef adds the serviceCatalogRef to the create service catalog from template params
func (o *CreateServiceCatalogFromTemplateParams) SetServiceCatalogRef(serviceCatalogRef string) {
	o.ServiceCatalogRef = serviceCatalogRef
}

// WriteToRequest writes these params to a swagger request
func (o *CreateServiceCatalogFromTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
