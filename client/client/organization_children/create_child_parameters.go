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
	"github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/client/models"
)

// NewCreateChildParams creates a new CreateChildParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateChildParams() *CreateChildParams {
	return &CreateChildParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateChildParamsWithTimeout creates a new CreateChildParams object
// with the ability to set a timeout on a request.
func NewCreateChildParamsWithTimeout(timeout time.Duration) *CreateChildParams {
	return &CreateChildParams{
		timeout: timeout,
	}
}

// NewCreateChildParamsWithContext creates a new CreateChildParams object
// with the ability to set a context for a request.
func NewCreateChildParamsWithContext(ctx context.Context) *CreateChildParams {
	return &CreateChildParams{
		Context: ctx,
	}
}

// NewCreateChildParamsWithHTTPClient creates a new CreateChildParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateChildParamsWithHTTPClient(client *http.Client) *CreateChildParams {
	return &CreateChildParams{
		HTTPClient: client,
	}
}

/*
CreateChildParams contains all the parameters to send to the API endpoint

	for the create child operation.

	Typically these are written to a http.Request.
*/
type CreateChildParams struct {

	/* Body.

	   The information of the organization to create.
	*/
	Body *models.NewOrganization

	/* OrganizationCanonical.

	   A canonical of an organization.
	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create child params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateChildParams) WithDefaults() *CreateChildParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create child params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateChildParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create child params
func (o *CreateChildParams) WithTimeout(timeout time.Duration) *CreateChildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create child params
func (o *CreateChildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create child params
func (o *CreateChildParams) WithContext(ctx context.Context) *CreateChildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create child params
func (o *CreateChildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create child params
func (o *CreateChildParams) WithHTTPClient(client *http.Client) *CreateChildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create child params
func (o *CreateChildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create child params
func (o *CreateChildParams) WithBody(body *models.NewOrganization) *CreateChildParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create child params
func (o *CreateChildParams) SetBody(body *models.NewOrganization) {
	o.Body = body
}

// WithOrganizationCanonical adds the organizationCanonical to the create child params
func (o *CreateChildParams) WithOrganizationCanonical(organizationCanonical string) *CreateChildParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the create child params
func (o *CreateChildParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *CreateChildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
