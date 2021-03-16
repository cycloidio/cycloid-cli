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

// NewCreateStackParams creates a new CreateStackParams object
// with the default values initialized.
func NewCreateStackParams() *CreateStackParams {
	var ()
	return &CreateStackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateStackParamsWithTimeout creates a new CreateStackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateStackParamsWithTimeout(timeout time.Duration) *CreateStackParams {
	var ()
	return &CreateStackParams{

		timeout: timeout,
	}
}

// NewCreateStackParamsWithContext creates a new CreateStackParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateStackParamsWithContext(ctx context.Context) *CreateStackParams {
	var ()
	return &CreateStackParams{

		Context: ctx,
	}
}

// NewCreateStackParamsWithHTTPClient creates a new CreateStackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateStackParamsWithHTTPClient(client *http.Client) *CreateStackParams {
	var ()
	return &CreateStackParams{
		HTTPClient: client,
	}
}

/*CreateStackParams contains all the parameters to send to the API endpoint
for the create stack operation typically these are written to a http.Request
*/
type CreateStackParams struct {

	/*Body
	  The information of the Stack.

	*/
	Body *models.NewStack
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create stack params
func (o *CreateStackParams) WithTimeout(timeout time.Duration) *CreateStackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create stack params
func (o *CreateStackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create stack params
func (o *CreateStackParams) WithContext(ctx context.Context) *CreateStackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create stack params
func (o *CreateStackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create stack params
func (o *CreateStackParams) WithHTTPClient(client *http.Client) *CreateStackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create stack params
func (o *CreateStackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create stack params
func (o *CreateStackParams) WithBody(body *models.NewStack) *CreateStackParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create stack params
func (o *CreateStackParams) SetBody(body *models.NewStack) {
	o.Body = body
}

// WithOrganizationCanonical adds the organizationCanonical to the create stack params
func (o *CreateStackParams) WithOrganizationCanonical(organizationCanonical string) *CreateStackParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the create stack params
func (o *CreateStackParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *CreateStackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
