// Code generated by go-swagger; DO NOT EDIT.

package organization_infrastructure_policies

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
)

// NewDeleteInfraPolicyParams creates a new DeleteInfraPolicyParams object
// with the default values initialized.
func NewDeleteInfraPolicyParams() *DeleteInfraPolicyParams {
	var ()
	return &DeleteInfraPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInfraPolicyParamsWithTimeout creates a new DeleteInfraPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInfraPolicyParamsWithTimeout(timeout time.Duration) *DeleteInfraPolicyParams {
	var ()
	return &DeleteInfraPolicyParams{

		timeout: timeout,
	}
}

// NewDeleteInfraPolicyParamsWithContext creates a new DeleteInfraPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInfraPolicyParamsWithContext(ctx context.Context) *DeleteInfraPolicyParams {
	var ()
	return &DeleteInfraPolicyParams{

		Context: ctx,
	}
}

// NewDeleteInfraPolicyParamsWithHTTPClient creates a new DeleteInfraPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInfraPolicyParamsWithHTTPClient(client *http.Client) *DeleteInfraPolicyParams {
	var ()
	return &DeleteInfraPolicyParams{
		HTTPClient: client,
	}
}

/*DeleteInfraPolicyParams contains all the parameters to send to the API endpoint
for the delete infra policy operation typically these are written to a http.Request
*/
type DeleteInfraPolicyParams struct {

	/*InfraPolicyCanonical
	  The canonical of an InfraPolicy.

	*/
	InfraPolicyCanonical string
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete infra policy params
func (o *DeleteInfraPolicyParams) WithTimeout(timeout time.Duration) *DeleteInfraPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete infra policy params
func (o *DeleteInfraPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete infra policy params
func (o *DeleteInfraPolicyParams) WithContext(ctx context.Context) *DeleteInfraPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete infra policy params
func (o *DeleteInfraPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete infra policy params
func (o *DeleteInfraPolicyParams) WithHTTPClient(client *http.Client) *DeleteInfraPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete infra policy params
func (o *DeleteInfraPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfraPolicyCanonical adds the infraPolicyCanonical to the delete infra policy params
func (o *DeleteInfraPolicyParams) WithInfraPolicyCanonical(infraPolicyCanonical string) *DeleteInfraPolicyParams {
	o.SetInfraPolicyCanonical(infraPolicyCanonical)
	return o
}

// SetInfraPolicyCanonical adds the infraPolicyCanonical to the delete infra policy params
func (o *DeleteInfraPolicyParams) SetInfraPolicyCanonical(infraPolicyCanonical string) {
	o.InfraPolicyCanonical = infraPolicyCanonical
}

// WithOrganizationCanonical adds the organizationCanonical to the delete infra policy params
func (o *DeleteInfraPolicyParams) WithOrganizationCanonical(organizationCanonical string) *DeleteInfraPolicyParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the delete infra policy params
func (o *DeleteInfraPolicyParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInfraPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param infra_policy_canonical
	if err := r.SetPathParam("infra_policy_canonical", o.InfraPolicyCanonical); err != nil {
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