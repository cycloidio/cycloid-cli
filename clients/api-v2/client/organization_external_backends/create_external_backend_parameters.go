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

	models "github.com/cycloidio/youdeploy-cli/clients/api-v2/models"
)

// NewCreateExternalBackendParams creates a new CreateExternalBackendParams object
// with the default values initialized.
func NewCreateExternalBackendParams() *CreateExternalBackendParams {
	var ()
	return &CreateExternalBackendParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateExternalBackendParamsWithTimeout creates a new CreateExternalBackendParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateExternalBackendParamsWithTimeout(timeout time.Duration) *CreateExternalBackendParams {
	var ()
	return &CreateExternalBackendParams{

		timeout: timeout,
	}
}

// NewCreateExternalBackendParamsWithContext creates a new CreateExternalBackendParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateExternalBackendParamsWithContext(ctx context.Context) *CreateExternalBackendParams {
	var ()
	return &CreateExternalBackendParams{

		Context: ctx,
	}
}

// NewCreateExternalBackendParamsWithHTTPClient creates a new CreateExternalBackendParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateExternalBackendParamsWithHTTPClient(client *http.Client) *CreateExternalBackendParams {
	var ()
	return &CreateExternalBackendParams{
		HTTPClient: client,
	}
}

/*CreateExternalBackendParams contains all the parameters to send to the API endpoint
for the create external backend operation typically these are written to a http.Request
*/
type CreateExternalBackendParams struct {

	/*Body
	  The information of the external backend

	*/
	Body *models.NewExternalBackend
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

// WithTimeout adds the timeout to the create external backend params
func (o *CreateExternalBackendParams) WithTimeout(timeout time.Duration) *CreateExternalBackendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create external backend params
func (o *CreateExternalBackendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create external backend params
func (o *CreateExternalBackendParams) WithContext(ctx context.Context) *CreateExternalBackendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create external backend params
func (o *CreateExternalBackendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create external backend params
func (o *CreateExternalBackendParams) WithHTTPClient(client *http.Client) *CreateExternalBackendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create external backend params
func (o *CreateExternalBackendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create external backend params
func (o *CreateExternalBackendParams) WithBody(body *models.NewExternalBackend) *CreateExternalBackendParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create external backend params
func (o *CreateExternalBackendParams) SetBody(body *models.NewExternalBackend) {
	o.Body = body
}

// WithEnvironment adds the environment to the create external backend params
func (o *CreateExternalBackendParams) WithEnvironment(environment *string) *CreateExternalBackendParams {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the create external backend params
func (o *CreateExternalBackendParams) SetEnvironment(environment *string) {
	o.Environment = environment
}

// WithOrganizationCanonical adds the organizationCanonical to the create external backend params
func (o *CreateExternalBackendParams) WithOrganizationCanonical(organizationCanonical string) *CreateExternalBackendParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the create external backend params
func (o *CreateExternalBackendParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithPproject adds the pproject to the create external backend params
func (o *CreateExternalBackendParams) WithPproject(pproject *string) *CreateExternalBackendParams {
	o.SetPproject(pproject)
	return o
}

// SetPproject adds the pproject to the create external backend params
func (o *CreateExternalBackendParams) SetPproject(pproject *string) {
	o.Pproject = pproject
}

// WriteToRequest writes these params to a swagger request
func (o *CreateExternalBackendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
