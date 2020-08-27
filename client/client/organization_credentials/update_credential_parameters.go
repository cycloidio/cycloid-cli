// Code generated by go-swagger; DO NOT EDIT.

package organization_credentials

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

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// NewUpdateCredentialParams creates a new UpdateCredentialParams object
// with the default values initialized.
func NewUpdateCredentialParams() *UpdateCredentialParams {
	var ()
	return &UpdateCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCredentialParamsWithTimeout creates a new UpdateCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCredentialParamsWithTimeout(timeout time.Duration) *UpdateCredentialParams {
	var ()
	return &UpdateCredentialParams{

		timeout: timeout,
	}
}

// NewUpdateCredentialParamsWithContext creates a new UpdateCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCredentialParamsWithContext(ctx context.Context) *UpdateCredentialParams {
	var ()
	return &UpdateCredentialParams{

		Context: ctx,
	}
}

// NewUpdateCredentialParamsWithHTTPClient creates a new UpdateCredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCredentialParamsWithHTTPClient(client *http.Client) *UpdateCredentialParams {
	var ()
	return &UpdateCredentialParams{
		HTTPClient: client,
	}
}

/*UpdateCredentialParams contains all the parameters to send to the API endpoint
for the update credential operation typically these are written to a http.Request
*/
type UpdateCredentialParams struct {

	/*Body
	  The information of the organization to update.

	*/
	Body *models.UpdateCredential
	/*CredentialID
	  A Credential id

	*/
	CredentialID uint32
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update credential params
func (o *UpdateCredentialParams) WithTimeout(timeout time.Duration) *UpdateCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update credential params
func (o *UpdateCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update credential params
func (o *UpdateCredentialParams) WithContext(ctx context.Context) *UpdateCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update credential params
func (o *UpdateCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update credential params
func (o *UpdateCredentialParams) WithHTTPClient(client *http.Client) *UpdateCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update credential params
func (o *UpdateCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update credential params
func (o *UpdateCredentialParams) WithBody(body *models.UpdateCredential) *UpdateCredentialParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update credential params
func (o *UpdateCredentialParams) SetBody(body *models.UpdateCredential) {
	o.Body = body
}

// WithCredentialID adds the credentialID to the update credential params
func (o *UpdateCredentialParams) WithCredentialID(credentialID uint32) *UpdateCredentialParams {
	o.SetCredentialID(credentialID)
	return o
}

// SetCredentialID adds the credentialId to the update credential params
func (o *UpdateCredentialParams) SetCredentialID(credentialID uint32) {
	o.CredentialID = credentialID
}

// WithOrganizationCanonical adds the organizationCanonical to the update credential params
func (o *UpdateCredentialParams) WithOrganizationCanonical(organizationCanonical string) *UpdateCredentialParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the update credential params
func (o *UpdateCredentialParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param credential_id
	if err := r.SetPathParam("credential_id", swag.FormatUint32(o.CredentialID)); err != nil {
		return err
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