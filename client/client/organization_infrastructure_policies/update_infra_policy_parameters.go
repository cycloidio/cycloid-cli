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

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// NewUpdateInfraPolicyParams creates a new UpdateInfraPolicyParams object
// with the default values initialized.
func NewUpdateInfraPolicyParams() *UpdateInfraPolicyParams {
	var ()
	return &UpdateInfraPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateInfraPolicyParamsWithTimeout creates a new UpdateInfraPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateInfraPolicyParamsWithTimeout(timeout time.Duration) *UpdateInfraPolicyParams {
	var ()
	return &UpdateInfraPolicyParams{

		timeout: timeout,
	}
}

// NewUpdateInfraPolicyParamsWithContext creates a new UpdateInfraPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateInfraPolicyParamsWithContext(ctx context.Context) *UpdateInfraPolicyParams {
	var ()
	return &UpdateInfraPolicyParams{

		Context: ctx,
	}
}

// NewUpdateInfraPolicyParamsWithHTTPClient creates a new UpdateInfraPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateInfraPolicyParamsWithHTTPClient(client *http.Client) *UpdateInfraPolicyParams {
	var ()
	return &UpdateInfraPolicyParams{
		HTTPClient: client,
	}
}

/*UpdateInfraPolicyParams contains all the parameters to send to the API endpoint
for the update infra policy operation typically these are written to a http.Request
*/
type UpdateInfraPolicyParams struct {

	/*Body
	  The information of the organization to update.

	*/
	Body *models.UpdateInfraPolicy
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

// WithTimeout adds the timeout to the update infra policy params
func (o *UpdateInfraPolicyParams) WithTimeout(timeout time.Duration) *UpdateInfraPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update infra policy params
func (o *UpdateInfraPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update infra policy params
func (o *UpdateInfraPolicyParams) WithContext(ctx context.Context) *UpdateInfraPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update infra policy params
func (o *UpdateInfraPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update infra policy params
func (o *UpdateInfraPolicyParams) WithHTTPClient(client *http.Client) *UpdateInfraPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update infra policy params
func (o *UpdateInfraPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update infra policy params
func (o *UpdateInfraPolicyParams) WithBody(body *models.UpdateInfraPolicy) *UpdateInfraPolicyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update infra policy params
func (o *UpdateInfraPolicyParams) SetBody(body *models.UpdateInfraPolicy) {
	o.Body = body
}

// WithInfraPolicyCanonical adds the infraPolicyCanonical to the update infra policy params
func (o *UpdateInfraPolicyParams) WithInfraPolicyCanonical(infraPolicyCanonical string) *UpdateInfraPolicyParams {
	o.SetInfraPolicyCanonical(infraPolicyCanonical)
	return o
}

// SetInfraPolicyCanonical adds the infraPolicyCanonical to the update infra policy params
func (o *UpdateInfraPolicyParams) SetInfraPolicyCanonical(infraPolicyCanonical string) {
	o.InfraPolicyCanonical = infraPolicyCanonical
}

// WithOrganizationCanonical adds the organizationCanonical to the update infra policy params
func (o *UpdateInfraPolicyParams) WithOrganizationCanonical(organizationCanonical string) *UpdateInfraPolicyParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the update infra policy params
func (o *UpdateInfraPolicyParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateInfraPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
