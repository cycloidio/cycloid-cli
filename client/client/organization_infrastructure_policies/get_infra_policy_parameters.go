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

// NewGetInfraPolicyParams creates a new GetInfraPolicyParams object
// with the default values initialized.
func NewGetInfraPolicyParams() *GetInfraPolicyParams {
	var ()
	return &GetInfraPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInfraPolicyParamsWithTimeout creates a new GetInfraPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInfraPolicyParamsWithTimeout(timeout time.Duration) *GetInfraPolicyParams {
	var ()
	return &GetInfraPolicyParams{

		timeout: timeout,
	}
}

// NewGetInfraPolicyParamsWithContext creates a new GetInfraPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInfraPolicyParamsWithContext(ctx context.Context) *GetInfraPolicyParams {
	var ()
	return &GetInfraPolicyParams{

		Context: ctx,
	}
}

// NewGetInfraPolicyParamsWithHTTPClient creates a new GetInfraPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInfraPolicyParamsWithHTTPClient(client *http.Client) *GetInfraPolicyParams {
	var ()
	return &GetInfraPolicyParams{
		HTTPClient: client,
	}
}

/*GetInfraPolicyParams contains all the parameters to send to the API endpoint
for the get infra policy operation typically these are written to a http.Request
*/
type GetInfraPolicyParams struct {

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

// WithTimeout adds the timeout to the get infra policy params
func (o *GetInfraPolicyParams) WithTimeout(timeout time.Duration) *GetInfraPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get infra policy params
func (o *GetInfraPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get infra policy params
func (o *GetInfraPolicyParams) WithContext(ctx context.Context) *GetInfraPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get infra policy params
func (o *GetInfraPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get infra policy params
func (o *GetInfraPolicyParams) WithHTTPClient(client *http.Client) *GetInfraPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get infra policy params
func (o *GetInfraPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfraPolicyCanonical adds the infraPolicyCanonical to the get infra policy params
func (o *GetInfraPolicyParams) WithInfraPolicyCanonical(infraPolicyCanonical string) *GetInfraPolicyParams {
	o.SetInfraPolicyCanonical(infraPolicyCanonical)
	return o
}

// SetInfraPolicyCanonical adds the infraPolicyCanonical to the get infra policy params
func (o *GetInfraPolicyParams) SetInfraPolicyCanonical(infraPolicyCanonical string) {
	o.InfraPolicyCanonical = infraPolicyCanonical
}

// WithOrganizationCanonical adds the organizationCanonical to the get infra policy params
func (o *GetInfraPolicyParams) WithOrganizationCanonical(organizationCanonical string) *GetInfraPolicyParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get infra policy params
func (o *GetInfraPolicyParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *GetInfraPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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