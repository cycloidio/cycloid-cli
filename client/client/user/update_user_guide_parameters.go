// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewUpdateUserGuideParams creates a new UpdateUserGuideParams object
// with the default values initialized.
func NewUpdateUserGuideParams() *UpdateUserGuideParams {
	var ()
	return &UpdateUserGuideParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserGuideParamsWithTimeout creates a new UpdateUserGuideParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateUserGuideParamsWithTimeout(timeout time.Duration) *UpdateUserGuideParams {
	var ()
	return &UpdateUserGuideParams{

		timeout: timeout,
	}
}

// NewUpdateUserGuideParamsWithContext creates a new UpdateUserGuideParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateUserGuideParamsWithContext(ctx context.Context) *UpdateUserGuideParams {
	var ()
	return &UpdateUserGuideParams{

		Context: ctx,
	}
}

// NewUpdateUserGuideParamsWithHTTPClient creates a new UpdateUserGuideParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateUserGuideParamsWithHTTPClient(client *http.Client) *UpdateUserGuideParams {
	var ()
	return &UpdateUserGuideParams{
		HTTPClient: client,
	}
}

/*UpdateUserGuideParams contains all the parameters to send to the API endpoint
for the update user guide operation typically these are written to a http.Request
*/
type UpdateUserGuideParams struct {

	/*Body
	  The guide's progress JSON schema

	*/
	Body models.UserGuide

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update user guide params
func (o *UpdateUserGuideParams) WithTimeout(timeout time.Duration) *UpdateUserGuideParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user guide params
func (o *UpdateUserGuideParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user guide params
func (o *UpdateUserGuideParams) WithContext(ctx context.Context) *UpdateUserGuideParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user guide params
func (o *UpdateUserGuideParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user guide params
func (o *UpdateUserGuideParams) WithHTTPClient(client *http.Client) *UpdateUserGuideParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user guide params
func (o *UpdateUserGuideParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update user guide params
func (o *UpdateUserGuideParams) WithBody(body models.UserGuide) *UpdateUserGuideParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update user guide params
func (o *UpdateUserGuideParams) SetBody(body models.UserGuide) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserGuideParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
