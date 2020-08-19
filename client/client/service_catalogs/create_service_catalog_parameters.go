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

// NewCreateServiceCatalogParams creates a new CreateServiceCatalogParams object
// with the default values initialized.
func NewCreateServiceCatalogParams() *CreateServiceCatalogParams {
	var ()
	return &CreateServiceCatalogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateServiceCatalogParamsWithTimeout creates a new CreateServiceCatalogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateServiceCatalogParamsWithTimeout(timeout time.Duration) *CreateServiceCatalogParams {
	var ()
	return &CreateServiceCatalogParams{

		timeout: timeout,
	}
}

// NewCreateServiceCatalogParamsWithContext creates a new CreateServiceCatalogParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateServiceCatalogParamsWithContext(ctx context.Context) *CreateServiceCatalogParams {
	var ()
	return &CreateServiceCatalogParams{

		Context: ctx,
	}
}

// NewCreateServiceCatalogParamsWithHTTPClient creates a new CreateServiceCatalogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateServiceCatalogParamsWithHTTPClient(client *http.Client) *CreateServiceCatalogParams {
	var ()
	return &CreateServiceCatalogParams{
		HTTPClient: client,
	}
}

/*CreateServiceCatalogParams contains all the parameters to send to the API endpoint
for the create service catalog operation typically these are written to a http.Request
*/
type CreateServiceCatalogParams struct {

	/*Body
	  The information of the ServiceCatalog.

	*/
	Body *models.NewServiceCatalog
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create service catalog params
func (o *CreateServiceCatalogParams) WithTimeout(timeout time.Duration) *CreateServiceCatalogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create service catalog params
func (o *CreateServiceCatalogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create service catalog params
func (o *CreateServiceCatalogParams) WithContext(ctx context.Context) *CreateServiceCatalogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create service catalog params
func (o *CreateServiceCatalogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create service catalog params
func (o *CreateServiceCatalogParams) WithHTTPClient(client *http.Client) *CreateServiceCatalogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create service catalog params
func (o *CreateServiceCatalogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create service catalog params
func (o *CreateServiceCatalogParams) WithBody(body *models.NewServiceCatalog) *CreateServiceCatalogParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create service catalog params
func (o *CreateServiceCatalogParams) SetBody(body *models.NewServiceCatalog) {
	o.Body = body
}

// WithOrganizationCanonical adds the organizationCanonical to the create service catalog params
func (o *CreateServiceCatalogParams) WithOrganizationCanonical(organizationCanonical string) *CreateServiceCatalogParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the create service catalog params
func (o *CreateServiceCatalogParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *CreateServiceCatalogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
