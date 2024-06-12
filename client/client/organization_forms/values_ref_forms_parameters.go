// Code generated by go-swagger; DO NOT EDIT.

package organization_forms

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

// NewValuesRefFormsParams creates a new ValuesRefFormsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValuesRefFormsParams() *ValuesRefFormsParams {
	return &ValuesRefFormsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValuesRefFormsParamsWithTimeout creates a new ValuesRefFormsParams object
// with the ability to set a timeout on a request.
func NewValuesRefFormsParamsWithTimeout(timeout time.Duration) *ValuesRefFormsParams {
	return &ValuesRefFormsParams{
		timeout: timeout,
	}
}

// NewValuesRefFormsParamsWithContext creates a new ValuesRefFormsParams object
// with the ability to set a context for a request.
func NewValuesRefFormsParamsWithContext(ctx context.Context) *ValuesRefFormsParams {
	return &ValuesRefFormsParams{
		Context: ctx,
	}
}

// NewValuesRefFormsParamsWithHTTPClient creates a new ValuesRefFormsParams object
// with the ability to set a custom HTTPClient for a request.
func NewValuesRefFormsParamsWithHTTPClient(client *http.Client) *ValuesRefFormsParams {
	return &ValuesRefFormsParams{
		HTTPClient: client,
	}
}

/*
ValuesRefFormsParams contains all the parameters to send to the API endpoint

	for the values ref forms operation.

	Typically these are written to a http.Request.
*/
type ValuesRefFormsParams struct {

	/* Body.

	   The content of the forms file to be validated.
	*/
	Body *models.FormsValuesRef

	/* OrganizationCanonical.

	   A canonical of an organization.
	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the values ref forms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValuesRefFormsParams) WithDefaults() *ValuesRefFormsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the values ref forms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValuesRefFormsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the values ref forms params
func (o *ValuesRefFormsParams) WithTimeout(timeout time.Duration) *ValuesRefFormsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the values ref forms params
func (o *ValuesRefFormsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the values ref forms params
func (o *ValuesRefFormsParams) WithContext(ctx context.Context) *ValuesRefFormsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the values ref forms params
func (o *ValuesRefFormsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the values ref forms params
func (o *ValuesRefFormsParams) WithHTTPClient(client *http.Client) *ValuesRefFormsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the values ref forms params
func (o *ValuesRefFormsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the values ref forms params
func (o *ValuesRefFormsParams) WithBody(body *models.FormsValuesRef) *ValuesRefFormsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the values ref forms params
func (o *ValuesRefFormsParams) SetBody(body *models.FormsValuesRef) {
	o.Body = body
}

// WithOrganizationCanonical adds the organizationCanonical to the values ref forms params
func (o *ValuesRefFormsParams) WithOrganizationCanonical(organizationCanonical string) *ValuesRefFormsParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the values ref forms params
func (o *ValuesRefFormsParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *ValuesRefFormsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
