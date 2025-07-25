// Code generated by go-swagger; DO NOT EDIT.

package organization_resource_pools

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

// NewUpdateResourcePoolParams creates a new UpdateResourcePoolParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateResourcePoolParams() *UpdateResourcePoolParams {
	return &UpdateResourcePoolParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateResourcePoolParamsWithTimeout creates a new UpdateResourcePoolParams object
// with the ability to set a timeout on a request.
func NewUpdateResourcePoolParamsWithTimeout(timeout time.Duration) *UpdateResourcePoolParams {
	return &UpdateResourcePoolParams{
		timeout: timeout,
	}
}

// NewUpdateResourcePoolParamsWithContext creates a new UpdateResourcePoolParams object
// with the ability to set a context for a request.
func NewUpdateResourcePoolParamsWithContext(ctx context.Context) *UpdateResourcePoolParams {
	return &UpdateResourcePoolParams{
		Context: ctx,
	}
}

// NewUpdateResourcePoolParamsWithHTTPClient creates a new UpdateResourcePoolParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateResourcePoolParamsWithHTTPClient(client *http.Client) *UpdateResourcePoolParams {
	return &UpdateResourcePoolParams{
		HTTPClient: client,
	}
}

/*
UpdateResourcePoolParams contains all the parameters to send to the API endpoint

	for the update resource pool operation.

	Typically these are written to a http.Request.
*/
type UpdateResourcePoolParams struct {

	/* Body.

	   The information of the organization's resource_pool to update.
	*/
	Body *models.NewResourcePool

	/* OrganizationCanonical.

	   A canonical of an organization.
	*/
	OrganizationCanonical string

	/* ResourcePoolCanonical.

	   Organization Resource Pool canonical
	*/
	ResourcePoolCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update resource pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateResourcePoolParams) WithDefaults() *UpdateResourcePoolParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update resource pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateResourcePoolParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update resource pool params
func (o *UpdateResourcePoolParams) WithTimeout(timeout time.Duration) *UpdateResourcePoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update resource pool params
func (o *UpdateResourcePoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update resource pool params
func (o *UpdateResourcePoolParams) WithContext(ctx context.Context) *UpdateResourcePoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update resource pool params
func (o *UpdateResourcePoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update resource pool params
func (o *UpdateResourcePoolParams) WithHTTPClient(client *http.Client) *UpdateResourcePoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update resource pool params
func (o *UpdateResourcePoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update resource pool params
func (o *UpdateResourcePoolParams) WithBody(body *models.NewResourcePool) *UpdateResourcePoolParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update resource pool params
func (o *UpdateResourcePoolParams) SetBody(body *models.NewResourcePool) {
	o.Body = body
}

// WithOrganizationCanonical adds the organizationCanonical to the update resource pool params
func (o *UpdateResourcePoolParams) WithOrganizationCanonical(organizationCanonical string) *UpdateResourcePoolParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the update resource pool params
func (o *UpdateResourcePoolParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithResourcePoolCanonical adds the resourcePoolCanonical to the update resource pool params
func (o *UpdateResourcePoolParams) WithResourcePoolCanonical(resourcePoolCanonical string) *UpdateResourcePoolParams {
	o.SetResourcePoolCanonical(resourcePoolCanonical)
	return o
}

// SetResourcePoolCanonical adds the resourcePoolCanonical to the update resource pool params
func (o *UpdateResourcePoolParams) SetResourcePoolCanonical(resourcePoolCanonical string) {
	o.ResourcePoolCanonical = resourcePoolCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateResourcePoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param resource_pool_canonical
	if err := r.SetPathParam("resource_pool_canonical", o.ResourcePoolCanonical); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
