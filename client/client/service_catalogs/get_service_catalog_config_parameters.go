// Code generated by go-swagger; DO NOT EDIT.

package service_catalogs

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
)

// NewGetServiceCatalogConfigParams creates a new GetServiceCatalogConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetServiceCatalogConfigParams() *GetServiceCatalogConfigParams {
	return &GetServiceCatalogConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceCatalogConfigParamsWithTimeout creates a new GetServiceCatalogConfigParams object
// with the ability to set a timeout on a request.
func NewGetServiceCatalogConfigParamsWithTimeout(timeout time.Duration) *GetServiceCatalogConfigParams {
	return &GetServiceCatalogConfigParams{
		timeout: timeout,
	}
}

// NewGetServiceCatalogConfigParamsWithContext creates a new GetServiceCatalogConfigParams object
// with the ability to set a context for a request.
func NewGetServiceCatalogConfigParamsWithContext(ctx context.Context) *GetServiceCatalogConfigParams {
	return &GetServiceCatalogConfigParams{
		Context: ctx,
	}
}

// NewGetServiceCatalogConfigParamsWithHTTPClient creates a new GetServiceCatalogConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetServiceCatalogConfigParamsWithHTTPClient(client *http.Client) *GetServiceCatalogConfigParams {
	return &GetServiceCatalogConfigParams{
		HTTPClient: client,
	}
}

/*
GetServiceCatalogConfigParams contains all the parameters to send to the API endpoint

	for the get service catalog config operation.

	Typically these are written to a http.Request.
*/
type GetServiceCatalogConfigParams struct {

	/* ComponentCanonical.

	   A canonical of a component.
	*/
	ComponentCanonical *string

	/* ComponentName.

	   A name of a component.
	*/
	ComponentName *string

	/* EnvironmentCanonical.

	   A list of environments' canonical to filter from
	*/
	EnvironmentCanonical *string

	/* OrganizationCanonical.

	   A canonical of an organization.
	*/
	OrganizationCanonical string

	/* ProjectCanonical.

	   A list of projects' canonical to filter from
	*/
	ProjectCanonical *string

	/* ServiceCatalogRef.

	   A Service Catalog name
	*/
	ServiceCatalogRef string

	/* UseCase.

	   A use case of a stack to be selectd from the stack config
	*/
	UseCase *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get service catalog config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServiceCatalogConfigParams) WithDefaults() *GetServiceCatalogConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get service catalog config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServiceCatalogConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get service catalog config params
func (o *GetServiceCatalogConfigParams) WithTimeout(timeout time.Duration) *GetServiceCatalogConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service catalog config params
func (o *GetServiceCatalogConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service catalog config params
func (o *GetServiceCatalogConfigParams) WithContext(ctx context.Context) *GetServiceCatalogConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service catalog config params
func (o *GetServiceCatalogConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service catalog config params
func (o *GetServiceCatalogConfigParams) WithHTTPClient(client *http.Client) *GetServiceCatalogConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service catalog config params
func (o *GetServiceCatalogConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComponentCanonical adds the componentCanonical to the get service catalog config params
func (o *GetServiceCatalogConfigParams) WithComponentCanonical(componentCanonical *string) *GetServiceCatalogConfigParams {
	o.SetComponentCanonical(componentCanonical)
	return o
}

// SetComponentCanonical adds the componentCanonical to the get service catalog config params
func (o *GetServiceCatalogConfigParams) SetComponentCanonical(componentCanonical *string) {
	o.ComponentCanonical = componentCanonical
}

// WithComponentName adds the componentName to the get service catalog config params
func (o *GetServiceCatalogConfigParams) WithComponentName(componentName *string) *GetServiceCatalogConfigParams {
	o.SetComponentName(componentName)
	return o
}

// SetComponentName adds the componentName to the get service catalog config params
func (o *GetServiceCatalogConfigParams) SetComponentName(componentName *string) {
	o.ComponentName = componentName
}

// WithEnvironmentCanonical adds the environmentCanonical to the get service catalog config params
func (o *GetServiceCatalogConfigParams) WithEnvironmentCanonical(environmentCanonical *string) *GetServiceCatalogConfigParams {
	o.SetEnvironmentCanonical(environmentCanonical)
	return o
}

// SetEnvironmentCanonical adds the environmentCanonical to the get service catalog config params
func (o *GetServiceCatalogConfigParams) SetEnvironmentCanonical(environmentCanonical *string) {
	o.EnvironmentCanonical = environmentCanonical
}

// WithOrganizationCanonical adds the organizationCanonical to the get service catalog config params
func (o *GetServiceCatalogConfigParams) WithOrganizationCanonical(organizationCanonical string) *GetServiceCatalogConfigParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get service catalog config params
func (o *GetServiceCatalogConfigParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithProjectCanonical adds the projectCanonical to the get service catalog config params
func (o *GetServiceCatalogConfigParams) WithProjectCanonical(projectCanonical *string) *GetServiceCatalogConfigParams {
	o.SetProjectCanonical(projectCanonical)
	return o
}

// SetProjectCanonical adds the projectCanonical to the get service catalog config params
func (o *GetServiceCatalogConfigParams) SetProjectCanonical(projectCanonical *string) {
	o.ProjectCanonical = projectCanonical
}

// WithServiceCatalogRef adds the serviceCatalogRef to the get service catalog config params
func (o *GetServiceCatalogConfigParams) WithServiceCatalogRef(serviceCatalogRef string) *GetServiceCatalogConfigParams {
	o.SetServiceCatalogRef(serviceCatalogRef)
	return o
}

// SetServiceCatalogRef adds the serviceCatalogRef to the get service catalog config params
func (o *GetServiceCatalogConfigParams) SetServiceCatalogRef(serviceCatalogRef string) {
	o.ServiceCatalogRef = serviceCatalogRef
}

// WithUseCase adds the useCase to the get service catalog config params
func (o *GetServiceCatalogConfigParams) WithUseCase(useCase *string) *GetServiceCatalogConfigParams {
	o.SetUseCase(useCase)
	return o
}

// SetUseCase adds the useCase to the get service catalog config params
func (o *GetServiceCatalogConfigParams) SetUseCase(useCase *string) {
	o.UseCase = useCase
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceCatalogConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.ComponentName != nil {

		// query param component_name
		var qrComponentName string

		if o.ComponentName != nil {
			qrComponentName = *o.ComponentName
		}
		qComponentName := qrComponentName
		if qComponentName != "" {

			if err := r.SetQueryParam("component_name", qComponentName); err != nil {
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

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	if o.ProjectCanonical != nil {

		// query param project_canonical
		var qrProjectCanonical string

		if o.ProjectCanonical != nil {
			qrProjectCanonical = *o.ProjectCanonical
		}
		qProjectCanonical := qrProjectCanonical
		if qProjectCanonical != "" {

			if err := r.SetQueryParam("project_canonical", qProjectCanonical); err != nil {
				return err
			}
		}
	}

	// path param service_catalog_ref
	if err := r.SetPathParam("service_catalog_ref", o.ServiceCatalogRef); err != nil {
		return err
	}

	if o.UseCase != nil {

		// query param use_case
		var qrUseCase string

		if o.UseCase != nil {
			qrUseCase = *o.UseCase
		}
		qUseCase := qrUseCase
		if qUseCase != "" {

			if err := r.SetQueryParam("use_case", qUseCase); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
