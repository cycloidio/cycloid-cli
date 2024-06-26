// Code generated by go-swagger; DO NOT EDIT.

package cost_estimation

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

// NewCostEstimateTfPlanParams creates a new CostEstimateTfPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCostEstimateTfPlanParams() *CostEstimateTfPlanParams {
	return &CostEstimateTfPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCostEstimateTfPlanParamsWithTimeout creates a new CostEstimateTfPlanParams object
// with the ability to set a timeout on a request.
func NewCostEstimateTfPlanParamsWithTimeout(timeout time.Duration) *CostEstimateTfPlanParams {
	return &CostEstimateTfPlanParams{
		timeout: timeout,
	}
}

// NewCostEstimateTfPlanParamsWithContext creates a new CostEstimateTfPlanParams object
// with the ability to set a context for a request.
func NewCostEstimateTfPlanParamsWithContext(ctx context.Context) *CostEstimateTfPlanParams {
	return &CostEstimateTfPlanParams{
		Context: ctx,
	}
}

// NewCostEstimateTfPlanParamsWithHTTPClient creates a new CostEstimateTfPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewCostEstimateTfPlanParamsWithHTTPClient(client *http.Client) *CostEstimateTfPlanParams {
	return &CostEstimateTfPlanParams{
		HTTPClient: client,
	}
}

/*
CostEstimateTfPlanParams contains all the parameters to send to the API endpoint

	for the cost estimate tf plan operation.

	Typically these are written to a http.Request.
*/
type CostEstimateTfPlanParams struct {

	// Body.
	Body *models.TerraformPlanInput

	/* OrganizationCanonical.

	   A canonical of an organization.
	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cost estimate tf plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CostEstimateTfPlanParams) WithDefaults() *CostEstimateTfPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cost estimate tf plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CostEstimateTfPlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cost estimate tf plan params
func (o *CostEstimateTfPlanParams) WithTimeout(timeout time.Duration) *CostEstimateTfPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cost estimate tf plan params
func (o *CostEstimateTfPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cost estimate tf plan params
func (o *CostEstimateTfPlanParams) WithContext(ctx context.Context) *CostEstimateTfPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cost estimate tf plan params
func (o *CostEstimateTfPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cost estimate tf plan params
func (o *CostEstimateTfPlanParams) WithHTTPClient(client *http.Client) *CostEstimateTfPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cost estimate tf plan params
func (o *CostEstimateTfPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cost estimate tf plan params
func (o *CostEstimateTfPlanParams) WithBody(body *models.TerraformPlanInput) *CostEstimateTfPlanParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cost estimate tf plan params
func (o *CostEstimateTfPlanParams) SetBody(body *models.TerraformPlanInput) {
	o.Body = body
}

// WithOrganizationCanonical adds the organizationCanonical to the cost estimate tf plan params
func (o *CostEstimateTfPlanParams) WithOrganizationCanonical(organizationCanonical string) *CostEstimateTfPlanParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the cost estimate tf plan params
func (o *CostEstimateTfPlanParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *CostEstimateTfPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
