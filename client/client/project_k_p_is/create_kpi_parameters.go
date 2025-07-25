// Code generated by go-swagger; DO NOT EDIT.

package project_k_p_is

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cycloidio/cycloid-cli/client/models"
)

// NewCreateKpiParams creates a new CreateKpiParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateKpiParams() *CreateKpiParams {
	return &CreateKpiParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateKpiParamsWithTimeout creates a new CreateKpiParams object
// with the ability to set a timeout on a request.
func NewCreateKpiParamsWithTimeout(timeout time.Duration) *CreateKpiParams {
	return &CreateKpiParams{
		timeout: timeout,
	}
}

// NewCreateKpiParamsWithContext creates a new CreateKpiParams object
// with the ability to set a context for a request.
func NewCreateKpiParamsWithContext(ctx context.Context) *CreateKpiParams {
	return &CreateKpiParams{
		Context: ctx,
	}
}

// NewCreateKpiParamsWithHTTPClient creates a new CreateKpiParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateKpiParamsWithHTTPClient(client *http.Client) *CreateKpiParams {
	return &CreateKpiParams{
		HTTPClient: client,
	}
}

/*
CreateKpiParams contains all the parameters to send to the API endpoint

	for the create kpi operation.

	Typically these are written to a http.Request.
*/
type CreateKpiParams struct {

	/* Begin.

	   The unix timestamp in milliseconds, which indicate the start of the time range.

	   Format: uint64
	*/
	Begin *uint64

	/* Body.

	   The information of the KPI
	*/
	Body *models.NewKPI

	/* ComponentCanonical.

	   A canonical of a component.
	*/
	ComponentCanonical *string

	/* End.

	   The unix timestamp in milliseconds, which indicate the end of the time range.

	   Format: uint64
	*/
	End *uint64

	/* EnvironmentCanonical.

	   A list of environments' canonical to filter from
	*/
	EnvironmentCanonical *string

	/* Favorite.

	   Flag to retrieve favorite data from the members favorite list.

	*/
	Favorite *bool

	/* FetchData.

	   Flag to retrieve KPIs' data upon retrieving KPIs themselves

	*/
	FetchData *bool

	/* OrganizationCanonical.

	   A canonical of an organization.
	*/
	OrganizationCanonical string

	/* PageIndex.

	   The page number to request. The first page is 1.

	   Format: uint32
	   Default: 1
	*/
	PageIndex *uint32

	/* PageSize.

	   The number of items at most which the response can have.

	   Format: uint32
	   Default: 1000
	*/
	PageSize *uint32

	/* ProjectCanonical.

	   A canonical of a project.
	*/
	ProjectCanonical string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create kpi params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateKpiParams) WithDefaults() *CreateKpiParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create kpi params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateKpiParams) SetDefaults() {
	var (
		pageIndexDefault = uint32(1)

		pageSizeDefault = uint32(1000)
	)

	val := CreateKpiParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create kpi params
func (o *CreateKpiParams) WithTimeout(timeout time.Duration) *CreateKpiParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create kpi params
func (o *CreateKpiParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create kpi params
func (o *CreateKpiParams) WithContext(ctx context.Context) *CreateKpiParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create kpi params
func (o *CreateKpiParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create kpi params
func (o *CreateKpiParams) WithHTTPClient(client *http.Client) *CreateKpiParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create kpi params
func (o *CreateKpiParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBegin adds the begin to the create kpi params
func (o *CreateKpiParams) WithBegin(begin *uint64) *CreateKpiParams {
	o.SetBegin(begin)
	return o
}

// SetBegin adds the begin to the create kpi params
func (o *CreateKpiParams) SetBegin(begin *uint64) {
	o.Begin = begin
}

// WithBody adds the body to the create kpi params
func (o *CreateKpiParams) WithBody(body *models.NewKPI) *CreateKpiParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create kpi params
func (o *CreateKpiParams) SetBody(body *models.NewKPI) {
	o.Body = body
}

// WithComponentCanonical adds the componentCanonical to the create kpi params
func (o *CreateKpiParams) WithComponentCanonical(componentCanonical *string) *CreateKpiParams {
	o.SetComponentCanonical(componentCanonical)
	return o
}

// SetComponentCanonical adds the componentCanonical to the create kpi params
func (o *CreateKpiParams) SetComponentCanonical(componentCanonical *string) {
	o.ComponentCanonical = componentCanonical
}

// WithEnd adds the end to the create kpi params
func (o *CreateKpiParams) WithEnd(end *uint64) *CreateKpiParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the create kpi params
func (o *CreateKpiParams) SetEnd(end *uint64) {
	o.End = end
}

// WithEnvironmentCanonical adds the environmentCanonical to the create kpi params
func (o *CreateKpiParams) WithEnvironmentCanonical(environmentCanonical *string) *CreateKpiParams {
	o.SetEnvironmentCanonical(environmentCanonical)
	return o
}

// SetEnvironmentCanonical adds the environmentCanonical to the create kpi params
func (o *CreateKpiParams) SetEnvironmentCanonical(environmentCanonical *string) {
	o.EnvironmentCanonical = environmentCanonical
}

// WithFavorite adds the favorite to the create kpi params
func (o *CreateKpiParams) WithFavorite(favorite *bool) *CreateKpiParams {
	o.SetFavorite(favorite)
	return o
}

// SetFavorite adds the favorite to the create kpi params
func (o *CreateKpiParams) SetFavorite(favorite *bool) {
	o.Favorite = favorite
}

// WithFetchData adds the fetchData to the create kpi params
func (o *CreateKpiParams) WithFetchData(fetchData *bool) *CreateKpiParams {
	o.SetFetchData(fetchData)
	return o
}

// SetFetchData adds the fetchData to the create kpi params
func (o *CreateKpiParams) SetFetchData(fetchData *bool) {
	o.FetchData = fetchData
}

// WithOrganizationCanonical adds the organizationCanonical to the create kpi params
func (o *CreateKpiParams) WithOrganizationCanonical(organizationCanonical string) *CreateKpiParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the create kpi params
func (o *CreateKpiParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithPageIndex adds the pageIndex to the create kpi params
func (o *CreateKpiParams) WithPageIndex(pageIndex *uint32) *CreateKpiParams {
	o.SetPageIndex(pageIndex)
	return o
}

// SetPageIndex adds the pageIndex to the create kpi params
func (o *CreateKpiParams) SetPageIndex(pageIndex *uint32) {
	o.PageIndex = pageIndex
}

// WithPageSize adds the pageSize to the create kpi params
func (o *CreateKpiParams) WithPageSize(pageSize *uint32) *CreateKpiParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the create kpi params
func (o *CreateKpiParams) SetPageSize(pageSize *uint32) {
	o.PageSize = pageSize
}

// WithProjectCanonical adds the projectCanonical to the create kpi params
func (o *CreateKpiParams) WithProjectCanonical(projectCanonical string) *CreateKpiParams {
	o.SetProjectCanonical(projectCanonical)
	return o
}

// SetProjectCanonical adds the projectCanonical to the create kpi params
func (o *CreateKpiParams) SetProjectCanonical(projectCanonical string) {
	o.ProjectCanonical = projectCanonical
}

// WriteToRequest writes these params to a swagger request
func (o *CreateKpiParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ComponentCanonical != nil {

		// query param component_canonical
		var qrComponentCanonical string

		if o.ComponentCanonical != nil {
			qrComponentCanonical = *o.ComponentCanonical
		}
		qComponentCanonical := qrComponentCanonical
		if qComponentCanonical != "" {

			if err := r.SetQueryParam("component_canonical", qComponentCanonical); err != nil {
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

	if o.EnvironmentCanonical != nil {

		// query param environment_canonical
		var qrEnvironmentCanonical string

		if o.EnvironmentCanonical != nil {
			qrEnvironmentCanonical = *o.EnvironmentCanonical
		}
		qEnvironmentCanonical := qrEnvironmentCanonical
		if qEnvironmentCanonical != "" {

			if err := r.SetQueryParam("environment_canonical", qEnvironmentCanonical); err != nil {
				return err
			}
		}
	}

	if o.Favorite != nil {

		// query param favorite
		var qrFavorite bool

		if o.Favorite != nil {
			qrFavorite = *o.Favorite
		}
		qFavorite := swag.FormatBool(qrFavorite)
		if qFavorite != "" {

			if err := r.SetQueryParam("favorite", qFavorite); err != nil {
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

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	if o.PageIndex != nil {

		// query param page_index
		var qrPageIndex uint32

		if o.PageIndex != nil {
			qrPageIndex = *o.PageIndex
		}
		qPageIndex := swag.FormatUint32(qrPageIndex)
		if qPageIndex != "" {

			if err := r.SetQueryParam("page_index", qPageIndex); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize uint32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatUint32(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
	}

	// path param project_canonical
	if err := r.SetPathParam("project_canonical", o.ProjectCanonical); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
