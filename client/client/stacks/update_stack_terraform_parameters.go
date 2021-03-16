// Code generated by go-swagger; DO NOT EDIT.

package stacks

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

// NewUpdateStackTerraformParams creates a new UpdateStackTerraformParams object
// with the default values initialized.
func NewUpdateStackTerraformParams() *UpdateStackTerraformParams {
	var ()
	return &UpdateStackTerraformParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateStackTerraformParamsWithTimeout creates a new UpdateStackTerraformParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateStackTerraformParamsWithTimeout(timeout time.Duration) *UpdateStackTerraformParams {
	var ()
	return &UpdateStackTerraformParams{

		timeout: timeout,
	}
}

// NewUpdateStackTerraformParamsWithContext creates a new UpdateStackTerraformParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateStackTerraformParamsWithContext(ctx context.Context) *UpdateStackTerraformParams {
	var ()
	return &UpdateStackTerraformParams{

		Context: ctx,
	}
}

// NewUpdateStackTerraformParamsWithHTTPClient creates a new UpdateStackTerraformParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateStackTerraformParamsWithHTTPClient(client *http.Client) *UpdateStackTerraformParams {
	var ()
	return &UpdateStackTerraformParams{
		HTTPClient: client,
	}
}

/*UpdateStackTerraformParams contains all the parameters to send to the API endpoint
for the update stack terraform operation typically these are written to a http.Request
*/
type UpdateStackTerraformParams struct {

	/*Body
	  The information of the Stack Terraform.

	*/
	Body *models.TerraformJSONConfig
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string
	/*StackRef
	  A Stack name

	*/
	StackRef string
	/*UseCaseCanonical
	  A use case canonical

	*/
	UseCaseCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update stack terraform params
func (o *UpdateStackTerraformParams) WithTimeout(timeout time.Duration) *UpdateStackTerraformParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update stack terraform params
func (o *UpdateStackTerraformParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update stack terraform params
func (o *UpdateStackTerraformParams) WithContext(ctx context.Context) *UpdateStackTerraformParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update stack terraform params
func (o *UpdateStackTerraformParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update stack terraform params
func (o *UpdateStackTerraformParams) WithHTTPClient(client *http.Client) *UpdateStackTerraformParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update stack terraform params
func (o *UpdateStackTerraformParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update stack terraform params
func (o *UpdateStackTerraformParams) WithBody(body *models.TerraformJSONConfig) *UpdateStackTerraformParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update stack terraform params
func (o *UpdateStackTerraformParams) SetBody(body *models.TerraformJSONConfig) {
	o.Body = body
}

// WithOrganizationCanonical adds the organizationCanonical to the update stack terraform params
func (o *UpdateStackTerraformParams) WithOrganizationCanonical(organizationCanonical string) *UpdateStackTerraformParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the update stack terraform params
func (o *UpdateStackTerraformParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithStackRef adds the stackRef to the update stack terraform params
func (o *UpdateStackTerraformParams) WithStackRef(stackRef string) *UpdateStackTerraformParams {
	o.SetStackRef(stackRef)
	return o
}

// SetStackRef adds the stackRef to the update stack terraform params
func (o *UpdateStackTerraformParams) SetStackRef(stackRef string) {
	o.StackRef = stackRef
}

// WithUseCaseCanonical adds the useCaseCanonical to the update stack terraform params
func (o *UpdateStackTerraformParams) WithUseCaseCanonical(useCaseCanonical string) *UpdateStackTerraformParams {
	o.SetUseCaseCanonical(useCaseCanonical)
	return o
}

// SetUseCaseCanonical adds the useCaseCanonical to the update stack terraform params
func (o *UpdateStackTerraformParams) SetUseCaseCanonical(useCaseCanonical string) {
	o.UseCaseCanonical = useCaseCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateStackTerraformParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param stack_ref
	if err := r.SetPathParam("stack_ref", o.StackRef); err != nil {
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
