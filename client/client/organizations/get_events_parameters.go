// Code generated by go-swagger; DO NOT EDIT.

package organizations

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
)

// NewGetEventsParams creates a new GetEventsParams object
// with the default values initialized.
func NewGetEventsParams() *GetEventsParams {
	var ()
	return &GetEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEventsParamsWithTimeout creates a new GetEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEventsParamsWithTimeout(timeout time.Duration) *GetEventsParams {
	var ()
	return &GetEventsParams{

		timeout: timeout,
	}
}

// NewGetEventsParamsWithContext creates a new GetEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEventsParamsWithContext(ctx context.Context) *GetEventsParams {
	var ()
	return &GetEventsParams{

		Context: ctx,
	}
}

// NewGetEventsParamsWithHTTPClient creates a new GetEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEventsParamsWithHTTPClient(client *http.Client) *GetEventsParams {
	var ()
	return &GetEventsParams{
		HTTPClient: client,
	}
}

/*GetEventsParams contains all the parameters to send to the API endpoint
for the get events operation typically these are written to a http.Request
*/
type GetEventsParams struct {

	/*Begin
	  The unix timestamp in milliseconds, which indicate the start of the time range.

	*/
	Begin *uint64
	/*End
	  The unix timestamp in milliseconds, which indicate the end of the time range.

	*/
	End *uint64
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string
	/*Severity
	  Specify the severities of the events to be requested. The returned events must have one of the specified severities.

	*/
	Severity []string
	/*Type
	  Specify the types of the events to be requested. The returned events must have one of the specified types.

	*/
	Type []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get events params
func (o *GetEventsParams) WithTimeout(timeout time.Duration) *GetEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get events params
func (o *GetEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get events params
func (o *GetEventsParams) WithContext(ctx context.Context) *GetEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get events params
func (o *GetEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get events params
func (o *GetEventsParams) WithHTTPClient(client *http.Client) *GetEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get events params
func (o *GetEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBegin adds the begin to the get events params
func (o *GetEventsParams) WithBegin(begin *uint64) *GetEventsParams {
	o.SetBegin(begin)
	return o
}

// SetBegin adds the begin to the get events params
func (o *GetEventsParams) SetBegin(begin *uint64) {
	o.Begin = begin
}

// WithEnd adds the end to the get events params
func (o *GetEventsParams) WithEnd(end *uint64) *GetEventsParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get events params
func (o *GetEventsParams) SetEnd(end *uint64) {
	o.End = end
}

// WithOrganizationCanonical adds the organizationCanonical to the get events params
func (o *GetEventsParams) WithOrganizationCanonical(organizationCanonical string) *GetEventsParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get events params
func (o *GetEventsParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithSeverity adds the severity to the get events params
func (o *GetEventsParams) WithSeverity(severity []string) *GetEventsParams {
	o.SetSeverity(severity)
	return o
}

// SetSeverity adds the severity to the get events params
func (o *GetEventsParams) SetSeverity(severity []string) {
	o.Severity = severity
}

// WithType adds the typeVar to the get events params
func (o *GetEventsParams) WithType(typeVar []string) *GetEventsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get events params
func (o *GetEventsParams) SetType(typeVar []string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Begin != nil {

		// query param begin
		var qrBegin uint64
		if o.Begin != nil {
			qrBegin = *o.Begin
		}
		qBegin := swag.FormatUint64(qrBegin)
		if qBegin != "" {
			if err := r.SetQueryParam("begin", qBegin); err != nil {
				return err
			}
		}

	}

	if o.End != nil {

		// query param end
		var qrEnd uint64
		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := swag.FormatUint64(qrEnd)
		if qEnd != "" {
			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
		}

	}

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	valuesSeverity := o.Severity

	joinedSeverity := swag.JoinByFormat(valuesSeverity, "multi")
	// query array param severity
	if err := r.SetQueryParam("severity", joinedSeverity...); err != nil {
		return err
	}

	valuesType := o.Type

	joinedType := swag.JoinByFormat(valuesType, "multi")
	// query array param type
	if err := r.SetQueryParam("type", joinedType...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
