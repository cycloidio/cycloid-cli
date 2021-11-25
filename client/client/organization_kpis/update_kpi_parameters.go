// Code generated by go-swagger; DO NOT EDIT.

package organization_kpis

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

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// NewUpdateKpiParams creates a new UpdateKpiParams object
// with the default values initialized.
func NewUpdateKpiParams() *UpdateKpiParams {
	var ()
	return &UpdateKpiParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateKpiParamsWithTimeout creates a new UpdateKpiParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateKpiParamsWithTimeout(timeout time.Duration) *UpdateKpiParams {
	var ()
	return &UpdateKpiParams{

		timeout: timeout,
	}
}

// NewUpdateKpiParamsWithContext creates a new UpdateKpiParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateKpiParamsWithContext(ctx context.Context) *UpdateKpiParams {
	var ()
	return &UpdateKpiParams{

		Context: ctx,
	}
}

// NewUpdateKpiParamsWithHTTPClient creates a new UpdateKpiParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateKpiParamsWithHTTPClient(client *http.Client) *UpdateKpiParams {
	var ()
	return &UpdateKpiParams{
		HTTPClient: client,
	}
}

/*UpdateKpiParams contains all the parameters to send to the API endpoint
for the update kpi operation typically these are written to a http.Request
*/
type UpdateKpiParams struct {

	/*Begin
	  The unix timestamp in seconds, which indicate the start of the time range.

	*/
	Begin *uint64
	/*Body
	  The information of the KPI new data

	*/
	Body *models.NewKPI
	/*End
	  The unix timestamp in seconds, which indicate the end of the time range.

	*/
	End *uint64
	/*FetchData
	  Flag to retrieve KPIs' data upon retrieveing KPIs themselves


	*/
	FetchData *bool
	/*KpiCanonical
	  A canonical of a kpi.

	*/
	KpiCanonical string
	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update kpi params
func (o *UpdateKpiParams) WithTimeout(timeout time.Duration) *UpdateKpiParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update kpi params
func (o *UpdateKpiParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update kpi params
func (o *UpdateKpiParams) WithContext(ctx context.Context) *UpdateKpiParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update kpi params
func (o *UpdateKpiParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update kpi params
func (o *UpdateKpiParams) WithHTTPClient(client *http.Client) *UpdateKpiParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update kpi params
func (o *UpdateKpiParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBegin adds the begin to the update kpi params
func (o *UpdateKpiParams) WithBegin(begin *uint64) *UpdateKpiParams {
	o.SetBegin(begin)
	return o
}

// SetBegin adds the begin to the update kpi params
func (o *UpdateKpiParams) SetBegin(begin *uint64) {
	o.Begin = begin
}

// WithBody adds the body to the update kpi params
func (o *UpdateKpiParams) WithBody(body *models.NewKPI) *UpdateKpiParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update kpi params
func (o *UpdateKpiParams) SetBody(body *models.NewKPI) {
	o.Body = body
}

// WithEnd adds the end to the update kpi params
func (o *UpdateKpiParams) WithEnd(end *uint64) *UpdateKpiParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the update kpi params
func (o *UpdateKpiParams) SetEnd(end *uint64) {
	o.End = end
}

// WithFetchData adds the fetchData to the update kpi params
func (o *UpdateKpiParams) WithFetchData(fetchData *bool) *UpdateKpiParams {
	o.SetFetchData(fetchData)
	return o
}

// SetFetchData adds the fetchData to the update kpi params
func (o *UpdateKpiParams) SetFetchData(fetchData *bool) {
	o.FetchData = fetchData
}

// WithKpiCanonical adds the kpiCanonical to the update kpi params
func (o *UpdateKpiParams) WithKpiCanonical(kpiCanonical string) *UpdateKpiParams {
	o.SetKpiCanonical(kpiCanonical)
	return o
}

// SetKpiCanonical adds the kpiCanonical to the update kpi params
func (o *UpdateKpiParams) SetKpiCanonical(kpiCanonical string) {
	o.KpiCanonical = kpiCanonical
}

// WithOrganizationCanonical adds the organizationCanonical to the update kpi params
func (o *UpdateKpiParams) WithOrganizationCanonical(organizationCanonical string) *UpdateKpiParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the update kpi params
func (o *UpdateKpiParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateKpiParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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

	if o.FetchData != nil {

		// query param fetch_data
		var qrFetchData bool
		if o.FetchData != nil {
			qrFetchData = *o.FetchData
		}
		qFetchData := swag.FormatBool(qrFetchData)
		if qFetchData != "" {
			if err := r.SetQueryParam("fetch_data", qFetchData); err != nil {
				return err
			}
		}

	}

	// path param kpi_canonical
	if err := r.SetPathParam("kpi_canonical", o.KpiCanonical); err != nil {
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
