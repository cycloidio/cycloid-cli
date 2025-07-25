// Code generated by go-swagger; DO NOT EDIT.

package organization_cloud_cost_management_accounts

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

// NewCreateCloudCostManagementAccountChildParams creates a new CreateCloudCostManagementAccountChildParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCloudCostManagementAccountChildParams() *CreateCloudCostManagementAccountChildParams {
	return &CreateCloudCostManagementAccountChildParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCloudCostManagementAccountChildParamsWithTimeout creates a new CreateCloudCostManagementAccountChildParams object
// with the ability to set a timeout on a request.
func NewCreateCloudCostManagementAccountChildParamsWithTimeout(timeout time.Duration) *CreateCloudCostManagementAccountChildParams {
	return &CreateCloudCostManagementAccountChildParams{
		timeout: timeout,
	}
}

// NewCreateCloudCostManagementAccountChildParamsWithContext creates a new CreateCloudCostManagementAccountChildParams object
// with the ability to set a context for a request.
func NewCreateCloudCostManagementAccountChildParamsWithContext(ctx context.Context) *CreateCloudCostManagementAccountChildParams {
	return &CreateCloudCostManagementAccountChildParams{
		Context: ctx,
	}
}

// NewCreateCloudCostManagementAccountChildParamsWithHTTPClient creates a new CreateCloudCostManagementAccountChildParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCloudCostManagementAccountChildParamsWithHTTPClient(client *http.Client) *CreateCloudCostManagementAccountChildParams {
	return &CreateCloudCostManagementAccountChildParams{
		HTTPClient: client,
	}
}

/*
CreateCloudCostManagementAccountChildParams contains all the parameters to send to the API endpoint

	for the create cloud cost management account child operation.

	Typically these are written to a http.Request.
*/
type CreateCloudCostManagementAccountChildParams struct {

	// Body.
	Body *models.NewCloudCostManagementAccountChild

	/* OrganizationCanonical.

	   A canonical of an organization.
	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create cloud cost management account child params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCloudCostManagementAccountChildParams) WithDefaults() *CreateCloudCostManagementAccountChildParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create cloud cost management account child params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCloudCostManagementAccountChildParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create cloud cost management account child params
func (o *CreateCloudCostManagementAccountChildParams) WithTimeout(timeout time.Duration) *CreateCloudCostManagementAccountChildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cloud cost management account child params
func (o *CreateCloudCostManagementAccountChildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cloud cost management account child params
func (o *CreateCloudCostManagementAccountChildParams) WithContext(ctx context.Context) *CreateCloudCostManagementAccountChildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cloud cost management account child params
func (o *CreateCloudCostManagementAccountChildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cloud cost management account child params
func (o *CreateCloudCostManagementAccountChildParams) WithHTTPClient(client *http.Client) *CreateCloudCostManagementAccountChildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cloud cost management account child params
func (o *CreateCloudCostManagementAccountChildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create cloud cost management account child params
func (o *CreateCloudCostManagementAccountChildParams) WithBody(body *models.NewCloudCostManagementAccountChild) *CreateCloudCostManagementAccountChildParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create cloud cost management account child params
func (o *CreateCloudCostManagementAccountChildParams) SetBody(body *models.NewCloudCostManagementAccountChild) {
	o.Body = body
}

// WithOrganizationCanonical adds the organizationCanonical to the create cloud cost management account child params
func (o *CreateCloudCostManagementAccountChildParams) WithOrganizationCanonical(organizationCanonical string) *CreateCloudCostManagementAccountChildParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the create cloud cost management account child params
func (o *CreateCloudCostManagementAccountChildParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCloudCostManagementAccountChildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
