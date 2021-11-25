// Code generated by go-swagger; DO NOT EDIT.

package organization_children

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

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// NewCreateOrgChildParams creates a new CreateOrgChildParams object
// with the default values initialized.
func NewCreateOrgChildParams() *CreateOrgChildParams {
	var ()
	return &CreateOrgChildParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrgChildParamsWithTimeout creates a new CreateOrgChildParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateOrgChildParamsWithTimeout(timeout time.Duration) *CreateOrgChildParams {
	var ()
	return &CreateOrgChildParams{

		timeout: timeout,
	}
}

// NewCreateOrgChildParamsWithContext creates a new CreateOrgChildParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateOrgChildParamsWithContext(ctx context.Context) *CreateOrgChildParams {
	var ()
	return &CreateOrgChildParams{

		Context: ctx,
	}
}

// NewCreateOrgChildParamsWithHTTPClient creates a new CreateOrgChildParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateOrgChildParamsWithHTTPClient(client *http.Client) *CreateOrgChildParams {
	var ()
	return &CreateOrgChildParams{
		HTTPClient: client,
	}
}

/*CreateOrgChildParams contains all the parameters to send to the API endpoint
for the create org child operation typically these are written to a http.Request
*/
type CreateOrgChildParams struct {

	/*Body
	  The information of the organization to create.

	*/
	Body *models.NewOrganization
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create org child params
func (o *CreateOrgChildParams) WithTimeout(timeout time.Duration) *CreateOrgChildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create org child params
func (o *CreateOrgChildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create org child params
func (o *CreateOrgChildParams) WithContext(ctx context.Context) *CreateOrgChildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create org child params
func (o *CreateOrgChildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create org child params
func (o *CreateOrgChildParams) WithHTTPClient(client *http.Client) *CreateOrgChildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create org child params
func (o *CreateOrgChildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create org child params
func (o *CreateOrgChildParams) WithBody(body *models.NewOrganization) *CreateOrgChildParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create org child params
func (o *CreateOrgChildParams) SetBody(body *models.NewOrganization) {
	o.Body = body
}

// WithOrganizationCanonical adds the organizationCanonical to the create org child params
func (o *CreateOrgChildParams) WithOrganizationCanonical(organizationCanonical string) *CreateOrgChildParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the create org child params
func (o *CreateOrgChildParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrgChildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
