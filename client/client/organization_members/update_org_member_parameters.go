// Code generated by go-swagger; DO NOT EDIT.

package organization_members

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

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// NewUpdateOrgMemberParams creates a new UpdateOrgMemberParams object
// with the default values initialized.
func NewUpdateOrgMemberParams() *UpdateOrgMemberParams {
	var ()
	return &UpdateOrgMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrgMemberParamsWithTimeout creates a new UpdateOrgMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateOrgMemberParamsWithTimeout(timeout time.Duration) *UpdateOrgMemberParams {
	var ()
	return &UpdateOrgMemberParams{

		timeout: timeout,
	}
}

// NewUpdateOrgMemberParamsWithContext creates a new UpdateOrgMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateOrgMemberParamsWithContext(ctx context.Context) *UpdateOrgMemberParams {
	var ()
	return &UpdateOrgMemberParams{

		Context: ctx,
	}
}

// NewUpdateOrgMemberParamsWithHTTPClient creates a new UpdateOrgMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateOrgMemberParamsWithHTTPClient(client *http.Client) *UpdateOrgMemberParams {
	var ()
	return &UpdateOrgMemberParams{
		HTTPClient: client,
	}
}

/*UpdateOrgMemberParams contains all the parameters to send to the API endpoint
for the update org member operation typically these are written to a http.Request
*/
type UpdateOrgMemberParams struct {

	/*Body
	  The member information to be updated.

	*/
	Body *models.MemberAssignation
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string
	/*Username
	  A username

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update org member params
func (o *UpdateOrgMemberParams) WithTimeout(timeout time.Duration) *UpdateOrgMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update org member params
func (o *UpdateOrgMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update org member params
func (o *UpdateOrgMemberParams) WithContext(ctx context.Context) *UpdateOrgMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update org member params
func (o *UpdateOrgMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update org member params
func (o *UpdateOrgMemberParams) WithHTTPClient(client *http.Client) *UpdateOrgMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update org member params
func (o *UpdateOrgMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update org member params
func (o *UpdateOrgMemberParams) WithBody(body *models.MemberAssignation) *UpdateOrgMemberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update org member params
func (o *UpdateOrgMemberParams) SetBody(body *models.MemberAssignation) {
	o.Body = body
}

// WithOrganizationCanonical adds the organizationCanonical to the update org member params
func (o *UpdateOrgMemberParams) WithOrganizationCanonical(organizationCanonical string) *UpdateOrgMemberParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the update org member params
func (o *UpdateOrgMemberParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithUsername adds the username to the update org member params
func (o *UpdateOrgMemberParams) WithUsername(username string) *UpdateOrgMemberParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the update org member params
func (o *UpdateOrgMemberParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrgMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
