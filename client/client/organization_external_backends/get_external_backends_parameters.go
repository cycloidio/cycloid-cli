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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetExternalBackendsParams creates a new GetExternalBackendsParams object
// with the default values initialized.
func NewGetExternalBackendsParams() *GetExternalBackendsParams {
	var ()
	return &GetExternalBackendsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetExternalBackendsParamsWithTimeout creates a new GetExternalBackendsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetExternalBackendsParamsWithTimeout(timeout time.Duration) *GetExternalBackendsParams {
	var ()
	return &GetExternalBackendsParams{

		timeout: timeout,
	}
}

// NewGetExternalBackendsParamsWithContext creates a new GetExternalBackendsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetExternalBackendsParamsWithContext(ctx context.Context) *GetExternalBackendsParams {
	var ()
	return &GetExternalBackendsParams{

		Context: ctx,
	}
}

// NewGetExternalBackendsParamsWithHTTPClient creates a new GetExternalBackendsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetExternalBackendsParamsWithHTTPClient(client *http.Client) *GetExternalBackendsParams {
	var ()
	return &GetExternalBackendsParams{
		HTTPClient: client,
	}
}

/*GetExternalBackendsParams contains all the parameters to send to the API endpoint
for the get external backends operation typically these are written to a http.Request
*/
type GetExternalBackendsParams struct {

	/*EnvironmentCanonical
	  A list of environments' canonical to filter from

	*/
	EnvironmentCanonical *string
	/*ExternalBackendDefault
	  Filter for default Terraform External Backend

	*/
	ExternalBackendDefault *bool
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string
	/*ProjectCanonical
	  A list of projects' canonical to filter from

	*/
	ProjectCanonical *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get external backends params
func (o *GetExternalBackendsParams) WithTimeout(timeout time.Duration) *GetExternalBackendsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get external backends params
func (o *GetExternalBackendsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get external backends params
func (o *GetExternalBackendsParams) WithContext(ctx context.Context) *GetExternalBackendsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get external backends params
func (o *GetExternalBackendsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get external backends params
func (o *GetExternalBackendsParams) WithHTTPClient(client *http.Client) *GetExternalBackendsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get external backends params
func (o *GetExternalBackendsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentCanonical adds the environmentCanonical to the get external backends params
func (o *GetExternalBackendsParams) WithEnvironmentCanonical(environmentCanonical *string) *GetExternalBackendsParams {
	o.SetEnvironmentCanonical(environmentCanonical)
	return o
}

// SetEnvironmentCanonical adds the environmentCanonical to the get external backends params
func (o *GetExternalBackendsParams) SetEnvironmentCanonical(environmentCanonical *string) {
	o.EnvironmentCanonical = environmentCanonical
}

// WithExternalBackendDefault adds the externalBackendDefault to the get external backends params
func (o *GetExternalBackendsParams) WithExternalBackendDefault(externalBackendDefault *bool) *GetExternalBackendsParams {
	o.SetExternalBackendDefault(externalBackendDefault)
	return o
}

// SetExternalBackendDefault adds the externalBackendDefault to the get external backends params
func (o *GetExternalBackendsParams) SetExternalBackendDefault(externalBackendDefault *bool) {
	o.ExternalBackendDefault = externalBackendDefault
}

// WithOrganizationCanonical adds the organizationCanonical to the get external backends params
func (o *GetExternalBackendsParams) WithOrganizationCanonical(organizationCanonical string) *GetExternalBackendsParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get external backends params
func (o *GetExternalBackendsParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithProjectCanonical adds the projectCanonical to the get external backends params
func (o *GetExternalBackendsParams) WithProjectCanonical(projectCanonical *string) *GetExternalBackendsParams {
	o.SetProjectCanonical(projectCanonical)
	return o
}

// SetProjectCanonical adds the projectCanonical to the get external backends params
func (o *GetExternalBackendsParams) SetProjectCanonical(projectCanonical *string) {
	o.ProjectCanonical = projectCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *GetExternalBackendsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EnvironmentCanonical != nil {

		// query param environment_canonical
		var qrEnvironmentCanonical string
		if o.EnvironmentCanonical != nil {
			qrEnvironmentCanonical = *o.EnvironmentCanonical
		}
		qEnvironmentCanonical := qrEnvironmentCanonical
		if qEnvironmentCanonical != "" {
			if err := r.SetQueryParam("environment_canonical", qEnvironmentCanonical); err != nil {
				return err
			}
		}

	}

	if o.ExternalBackendDefault != nil {

		// query param external_backend_default
		var qrExternalBackendDefault bool
		if o.ExternalBackendDefault != nil {
			qrExternalBackendDefault = *o.ExternalBackendDefault
		}
		qExternalBackendDefault := swag.FormatBool(qrExternalBackendDefault)
		if qExternalBackendDefault != "" {
			if err := r.SetQueryParam("external_backend_default", qExternalBackendDefault); err != nil {
				return err
			}
		}

	}

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	if o.ProjectCanonical != nil {

		// query param project_canonical
		var qrProjectCanonical string
		if o.ProjectCanonical != nil {
			qrProjectCanonical = *o.ProjectCanonical
		}
		qProjectCanonical := qrProjectCanonical
		if qProjectCanonical != "" {
			if err := r.SetQueryParam("project_canonical", qProjectCanonical); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
