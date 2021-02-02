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

	strfmt "github.com/go-openapi/strfmt"
)

// NewRefreshServiceCatalogSourceParams creates a new RefreshServiceCatalogSourceParams object
// with the default values initialized.
func NewRefreshServiceCatalogSourceParams() *RefreshServiceCatalogSourceParams {
	var ()
	return &RefreshServiceCatalogSourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRefreshServiceCatalogSourceParamsWithTimeout creates a new RefreshServiceCatalogSourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRefreshServiceCatalogSourceParamsWithTimeout(timeout time.Duration) *RefreshServiceCatalogSourceParams {
	var ()
	return &RefreshServiceCatalogSourceParams{

		timeout: timeout,
	}
}

// NewRefreshServiceCatalogSourceParamsWithContext creates a new RefreshServiceCatalogSourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewRefreshServiceCatalogSourceParamsWithContext(ctx context.Context) *RefreshServiceCatalogSourceParams {
	var ()
	return &RefreshServiceCatalogSourceParams{

		Context: ctx,
	}
}

// NewRefreshServiceCatalogSourceParamsWithHTTPClient creates a new RefreshServiceCatalogSourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRefreshServiceCatalogSourceParamsWithHTTPClient(client *http.Client) *RefreshServiceCatalogSourceParams {
	var ()
	return &RefreshServiceCatalogSourceParams{
		HTTPClient: client,
	}
}

/*RefreshServiceCatalogSourceParams contains all the parameters to send to the API endpoint
for the refresh service catalog source operation typically these are written to a http.Request
*/
type RefreshServiceCatalogSourceParams struct {

	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string
	/*ServiceCatalogSourceCanonical
	  Organization Service Catalog Sources canonical

	*/
	ServiceCatalogSourceCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the refresh service catalog source params
func (o *RefreshServiceCatalogSourceParams) WithTimeout(timeout time.Duration) *RefreshServiceCatalogSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the refresh service catalog source params
func (o *RefreshServiceCatalogSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the refresh service catalog source params
func (o *RefreshServiceCatalogSourceParams) WithContext(ctx context.Context) *RefreshServiceCatalogSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the refresh service catalog source params
func (o *RefreshServiceCatalogSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the refresh service catalog source params
func (o *RefreshServiceCatalogSourceParams) WithHTTPClient(client *http.Client) *RefreshServiceCatalogSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the refresh service catalog source params
func (o *RefreshServiceCatalogSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationCanonical adds the organizationCanonical to the refresh service catalog source params
func (o *RefreshServiceCatalogSourceParams) WithOrganizationCanonical(organizationCanonical string) *RefreshServiceCatalogSourceParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the refresh service catalog source params
func (o *RefreshServiceCatalogSourceParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithServiceCatalogSourceCanonical adds the serviceCatalogSourceCanonical to the refresh service catalog source params
func (o *RefreshServiceCatalogSourceParams) WithServiceCatalogSourceCanonical(serviceCatalogSourceCanonical string) *RefreshServiceCatalogSourceParams {
	o.SetServiceCatalogSourceCanonical(serviceCatalogSourceCanonical)
	return o
}

// SetServiceCatalogSourceCanonical adds the serviceCatalogSourceCanonical to the refresh service catalog source params
func (o *RefreshServiceCatalogSourceParams) SetServiceCatalogSourceCanonical(serviceCatalogSourceCanonical string) {
	o.ServiceCatalogSourceCanonical = serviceCatalogSourceCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *RefreshServiceCatalogSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
