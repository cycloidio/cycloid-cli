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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// NewUpdateServiceCatalogTerraformParams creates a new UpdateServiceCatalogTerraformParams object
// with the default values initialized.
func NewUpdateServiceCatalogTerraformParams() *UpdateServiceCatalogTerraformParams {
	var ()
	return &UpdateServiceCatalogTerraformParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateServiceCatalogTerraformParamsWithTimeout creates a new UpdateServiceCatalogTerraformParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateServiceCatalogTerraformParamsWithTimeout(timeout time.Duration) *UpdateServiceCatalogTerraformParams {
	var ()
	return &UpdateServiceCatalogTerraformParams{

		timeout: timeout,
	}
}

// NewUpdateServiceCatalogTerraformParamsWithContext creates a new UpdateServiceCatalogTerraformParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateServiceCatalogTerraformParamsWithContext(ctx context.Context) *UpdateServiceCatalogTerraformParams {
	var ()
	return &UpdateServiceCatalogTerraformParams{

		Context: ctx,
	}
}

// NewUpdateServiceCatalogTerraformParamsWithHTTPClient creates a new UpdateServiceCatalogTerraformParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateServiceCatalogTerraformParamsWithHTTPClient(client *http.Client) *UpdateServiceCatalogTerraformParams {
	var ()
	return &UpdateServiceCatalogTerraformParams{
		HTTPClient: client,
	}
}

/*UpdateServiceCatalogTerraformParams contains all the parameters to send to the API endpoint
for the update service catalog terraform operation typically these are written to a http.Request
*/
type UpdateServiceCatalogTerraformParams struct {

	/*Body
	  The information of the ServiceCatalog Terraform.

	*/
	Body *models.TerraformJSONConfig
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string
	/*ServiceCatalogRef
	  A Service Catalog name

	*/
	ServiceCatalogRef string
	/*UseCaseCanonical
	  A use case canonical

	*/
	UseCaseCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update service catalog terraform params
func (o *UpdateServiceCatalogTerraformParams) WithTimeout(timeout time.Duration) *UpdateServiceCatalogTerraformParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update service catalog terraform params
func (o *UpdateServiceCatalogTerraformParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update service catalog terraform params
func (o *UpdateServiceCatalogTerraformParams) WithContext(ctx context.Context) *UpdateServiceCatalogTerraformParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update service catalog terraform params
func (o *UpdateServiceCatalogTerraformParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update service catalog terraform params
func (o *UpdateServiceCatalogTerraformParams) WithHTTPClient(client *http.Client) *UpdateServiceCatalogTerraformParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update service catalog terraform params
func (o *UpdateServiceCatalogTerraformParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update service catalog terraform params
func (o *UpdateServiceCatalogTerraformParams) WithBody(body *models.TerraformJSONConfig) *UpdateServiceCatalogTerraformParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update service catalog terraform params
func (o *UpdateServiceCatalogTerraformParams) SetBody(body *models.TerraformJSONConfig) {
	o.Body = body
}

// WithOrganizationCanonical adds the organizationCanonical to the update service catalog terraform params
func (o *UpdateServiceCatalogTerraformParams) WithOrganizationCanonical(organizationCanonical string) *UpdateServiceCatalogTerraformParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the update service catalog terraform params
func (o *UpdateServiceCatalogTerraformParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithServiceCatalogRef adds the serviceCatalogRef to the update service catalog terraform params
func (o *UpdateServiceCatalogTerraformParams) WithServiceCatalogRef(serviceCatalogRef string) *UpdateServiceCatalogTerraformParams {
	o.SetServiceCatalogRef(serviceCatalogRef)
	return o
}

// SetServiceCatalogRef adds the serviceCatalogRef to the update service catalog terraform params
func (o *UpdateServiceCatalogTerraformParams) SetServiceCatalogRef(serviceCatalogRef string) {
	o.ServiceCatalogRef = serviceCatalogRef
}

// WithUseCaseCanonical adds the useCaseCanonical to the update service catalog terraform params
func (o *UpdateServiceCatalogTerraformParams) WithUseCaseCanonical(useCaseCanonical string) *UpdateServiceCatalogTerraformParams {
	o.SetUseCaseCanonical(useCaseCanonical)
	return o
}

// SetUseCaseCanonical adds the useCaseCanonical to the update service catalog terraform params
func (o *UpdateServiceCatalogTerraformParams) SetUseCaseCanonical(useCaseCanonical string) {
	o.UseCaseCanonical = useCaseCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateServiceCatalogTerraformParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param use_case_canonical
	if err := r.SetPathParam("use_case_canonical", o.UseCaseCanonical); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
