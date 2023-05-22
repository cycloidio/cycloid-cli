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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListServiceCatalogsParams creates a new ListServiceCatalogsParams object
// with the default values initialized.
func NewListServiceCatalogsParams() *ListServiceCatalogsParams {
	var (
		pageIndexDefault              = uint32(1)
		pageSizeDefault               = uint32(1000)
		serviceCatalogTemplateDefault = bool(false)
	)
	return &ListServiceCatalogsParams{
		PageIndex:              &pageIndexDefault,
		PageSize:               &pageSizeDefault,
		ServiceCatalogTemplate: &serviceCatalogTemplateDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListServiceCatalogsParamsWithTimeout creates a new ListServiceCatalogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListServiceCatalogsParamsWithTimeout(timeout time.Duration) *ListServiceCatalogsParams {
	var (
		pageIndexDefault              = uint32(1)
		pageSizeDefault               = uint32(1000)
		serviceCatalogTemplateDefault = bool(false)
	)
	return &ListServiceCatalogsParams{
		PageIndex:              &pageIndexDefault,
		PageSize:               &pageSizeDefault,
		ServiceCatalogTemplate: &serviceCatalogTemplateDefault,

		timeout: timeout,
	}
}

// NewListServiceCatalogsParamsWithContext creates a new ListServiceCatalogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListServiceCatalogsParamsWithContext(ctx context.Context) *ListServiceCatalogsParams {
	var (
		pageIndexDefault              = uint32(1)
		pageSizeDefault               = uint32(1000)
		serviceCatalogTemplateDefault = bool(false)
	)
	return &ListServiceCatalogsParams{
		PageIndex:              &pageIndexDefault,
		PageSize:               &pageSizeDefault,
		ServiceCatalogTemplate: &serviceCatalogTemplateDefault,

		Context: ctx,
	}
}

// NewListServiceCatalogsParamsWithHTTPClient creates a new ListServiceCatalogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListServiceCatalogsParamsWithHTTPClient(client *http.Client) *ListServiceCatalogsParams {
	var (
		pageIndexDefault              = uint32(1)
		pageSizeDefault               = uint32(1000)
		serviceCatalogTemplateDefault = bool(false)
	)
	return &ListServiceCatalogsParams{
		PageIndex:              &pageIndexDefault,
		PageSize:               &pageSizeDefault,
		ServiceCatalogTemplate: &serviceCatalogTemplateDefault,
		HTTPClient:             client,
	}
}

/*ListServiceCatalogsParams contains all the parameters to send to the API endpoint
for the list service catalogs operation typically these are written to a http.Request
*/
type ListServiceCatalogsParams struct {

	/*OrganizationCanonical
	  A canonical of an organization.

	*/
	OrganizationCanonical string
	/*PageIndex
	  The page number to request. The first page is 1.

	*/
	PageIndex *uint32
	/*PageSize
	  The number of items at most which the response can have.

	*/
	PageSize *uint32
	/*ServiceCatalogOwn
	  Filters the Service Catalogs to only show the ones owned by the User Organization


	*/
	ServiceCatalogOwn *bool
	/*ServiceCatalogStatus
	  The status of the catalog service used for filtering.

	*/
	ServiceCatalogStatus *string
	/*ServiceCatalogTemplate
	  Filters the Service Catalogs to only show the ones that are templates


	*/
	ServiceCatalogTemplate *bool
	/*ServiceCatalogTrusted
	  Filters the Service Catalogs to only show the ones that are from trusted source (Cycloid)


	*/
	ServiceCatalogTrusted *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list service catalogs params
func (o *ListServiceCatalogsParams) WithTimeout(timeout time.Duration) *ListServiceCatalogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list service catalogs params
func (o *ListServiceCatalogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list service catalogs params
func (o *ListServiceCatalogsParams) WithContext(ctx context.Context) *ListServiceCatalogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list service catalogs params
func (o *ListServiceCatalogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list service catalogs params
func (o *ListServiceCatalogsParams) WithHTTPClient(client *http.Client) *ListServiceCatalogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list service catalogs params
func (o *ListServiceCatalogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationCanonical adds the organizationCanonical to the list service catalogs params
func (o *ListServiceCatalogsParams) WithOrganizationCanonical(organizationCanonical string) *ListServiceCatalogsParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the list service catalogs params
func (o *ListServiceCatalogsParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithPageIndex adds the pageIndex to the list service catalogs params
func (o *ListServiceCatalogsParams) WithPageIndex(pageIndex *uint32) *ListServiceCatalogsParams {
	o.SetPageIndex(pageIndex)
	return o
}

// SetPageIndex adds the pageIndex to the list service catalogs params
func (o *ListServiceCatalogsParams) SetPageIndex(pageIndex *uint32) {
	o.PageIndex = pageIndex
}

// WithPageSize adds the pageSize to the list service catalogs params
func (o *ListServiceCatalogsParams) WithPageSize(pageSize *uint32) *ListServiceCatalogsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list service catalogs params
func (o *ListServiceCatalogsParams) SetPageSize(pageSize *uint32) {
	o.PageSize = pageSize
}

// WithServiceCatalogOwn adds the serviceCatalogOwn to the list service catalogs params
func (o *ListServiceCatalogsParams) WithServiceCatalogOwn(serviceCatalogOwn *bool) *ListServiceCatalogsParams {
	o.SetServiceCatalogOwn(serviceCatalogOwn)
	return o
}

// SetServiceCatalogOwn adds the serviceCatalogOwn to the list service catalogs params
func (o *ListServiceCatalogsParams) SetServiceCatalogOwn(serviceCatalogOwn *bool) {
	o.ServiceCatalogOwn = serviceCatalogOwn
}

// WithServiceCatalogStatus adds the serviceCatalogStatus to the list service catalogs params
func (o *ListServiceCatalogsParams) WithServiceCatalogStatus(serviceCatalogStatus *string) *ListServiceCatalogsParams {
	o.SetServiceCatalogStatus(serviceCatalogStatus)
	return o
}

// SetServiceCatalogStatus adds the serviceCatalogStatus to the list service catalogs params
func (o *ListServiceCatalogsParams) SetServiceCatalogStatus(serviceCatalogStatus *string) {
	o.ServiceCatalogStatus = serviceCatalogStatus
}

// WithServiceCatalogTemplate adds the serviceCatalogTemplate to the list service catalogs params
func (o *ListServiceCatalogsParams) WithServiceCatalogTemplate(serviceCatalogTemplate *bool) *ListServiceCatalogsParams {
	o.SetServiceCatalogTemplate(serviceCatalogTemplate)
	return o
}

// SetServiceCatalogTemplate adds the serviceCatalogTemplate to the list service catalogs params
func (o *ListServiceCatalogsParams) SetServiceCatalogTemplate(serviceCatalogTemplate *bool) {
	o.ServiceCatalogTemplate = serviceCatalogTemplate
}

// WithServiceCatalogTrusted adds the serviceCatalogTrusted to the list service catalogs params
func (o *ListServiceCatalogsParams) WithServiceCatalogTrusted(serviceCatalogTrusted *bool) *ListServiceCatalogsParams {
	o.SetServiceCatalogTrusted(serviceCatalogTrusted)
	return o
}

// SetServiceCatalogTrusted adds the serviceCatalogTrusted to the list service catalogs params
func (o *ListServiceCatalogsParams) SetServiceCatalogTrusted(serviceCatalogTrusted *bool) {
	o.ServiceCatalogTrusted = serviceCatalogTrusted
}

// WriteToRequest writes these params to a swagger request
func (o *ListServiceCatalogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.ServiceCatalogOwn != nil {

		// query param service_catalog_own
		var qrServiceCatalogOwn bool
		if o.ServiceCatalogOwn != nil {
			qrServiceCatalogOwn = *o.ServiceCatalogOwn
		}
		qServiceCatalogOwn := swag.FormatBool(qrServiceCatalogOwn)
		if qServiceCatalogOwn != "" {
			if err := r.SetQueryParam("service_catalog_own", qServiceCatalogOwn); err != nil {
				return err
			}
		}

	}

	if o.ServiceCatalogStatus != nil {

		// query param service_catalog_status
		var qrServiceCatalogStatus string
		if o.ServiceCatalogStatus != nil {
			qrServiceCatalogStatus = *o.ServiceCatalogStatus
		}
		qServiceCatalogStatus := qrServiceCatalogStatus
		if qServiceCatalogStatus != "" {
			if err := r.SetQueryParam("service_catalog_status", qServiceCatalogStatus); err != nil {
				return err
			}
		}

	}

	if o.ServiceCatalogTemplate != nil {

		// query param service_catalog_template
		var qrServiceCatalogTemplate bool
		if o.ServiceCatalogTemplate != nil {
			qrServiceCatalogTemplate = *o.ServiceCatalogTemplate
		}
		qServiceCatalogTemplate := swag.FormatBool(qrServiceCatalogTemplate)
		if qServiceCatalogTemplate != "" {
			if err := r.SetQueryParam("service_catalog_template", qServiceCatalogTemplate); err != nil {
				return err
			}
		}

	}

	if o.ServiceCatalogTrusted != nil {

		// query param service_catalog_trusted
		var qrServiceCatalogTrusted bool
		if o.ServiceCatalogTrusted != nil {
			qrServiceCatalogTrusted = *o.ServiceCatalogTrusted
		}
		qServiceCatalogTrusted := swag.FormatBool(qrServiceCatalogTrusted)
		if qServiceCatalogTrusted != "" {
			if err := r.SetQueryParam("service_catalog_trusted", qServiceCatalogTrusted); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
