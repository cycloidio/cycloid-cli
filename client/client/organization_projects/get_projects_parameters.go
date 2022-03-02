// Code generated by go-swagger; DO NOT EDIT.

package organization_projects

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

// NewGetProjectsParams creates a new GetProjectsParams object
// with the default values initialized.
func NewGetProjectsParams() *GetProjectsParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &GetProjectsParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectsParamsWithTimeout creates a new GetProjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProjectsParamsWithTimeout(timeout time.Duration) *GetProjectsParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &GetProjectsParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetProjectsParamsWithContext creates a new GetProjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProjectsParamsWithContext(ctx context.Context) *GetProjectsParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &GetProjectsParams{
		PageIndex: &pageIndexDefault,
		PageSize:  &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetProjectsParamsWithHTTPClient creates a new GetProjectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProjectsParamsWithHTTPClient(client *http.Client) *GetProjectsParams {
	var (
		pageIndexDefault = uint32(1)
		pageSizeDefault  = uint32(1000)
	)
	return &GetProjectsParams{
		PageIndex:  &pageIndexDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetProjectsParams contains all the parameters to send to the API endpoint
for the get projects operation typically these are written to a http.Request
*/
type GetProjectsParams struct {

	/*Favorite
	  Flag to retrieve favorite data from the members favorite list.


	*/
	Favorite *bool
	/*OrderBy
	  Allows to order the list of items. Example usage: field_name:asc


	*/
	OrderBy *string
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
	/*ProjectConfigRepositoryCanonical
	  Search by project's config repository's ID

	*/
	ProjectConfigRepositoryCanonical *uint32
	/*ProjectCreatedAt
	  Search by project's creation date

	*/
	ProjectCreatedAt *uint64
	/*ProjectDescription
	  Search by project's description

	*/
	ProjectDescription *string
	/*ProjectName
	  Search by project's name

	*/
	ProjectName *string
	/*ServiceCatalogSourceCanonical
	  Organization Service Catalog Sources canonical

	*/
	ServiceCatalogSourceCanonical *string
	/*UserID
	  Search by entity's owner

	*/
	UserID *uint32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get projects params
func (o *GetProjectsParams) WithTimeout(timeout time.Duration) *GetProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get projects params
func (o *GetProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get projects params
func (o *GetProjectsParams) WithContext(ctx context.Context) *GetProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get projects params
func (o *GetProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get projects params
func (o *GetProjectsParams) WithHTTPClient(client *http.Client) *GetProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get projects params
func (o *GetProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFavorite adds the favorite to the get projects params
func (o *GetProjectsParams) WithFavorite(favorite *bool) *GetProjectsParams {
	o.SetFavorite(favorite)
	return o
}

// SetFavorite adds the favorite to the get projects params
func (o *GetProjectsParams) SetFavorite(favorite *bool) {
	o.Favorite = favorite
}

// WithOrderBy adds the orderBy to the get projects params
func (o *GetProjectsParams) WithOrderBy(orderBy *string) *GetProjectsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get projects params
func (o *GetProjectsParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithOrganizationCanonical adds the organizationCanonical to the get projects params
func (o *GetProjectsParams) WithOrganizationCanonical(organizationCanonical string) *GetProjectsParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get projects params
func (o *GetProjectsParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithPageIndex adds the pageIndex to the get projects params
func (o *GetProjectsParams) WithPageIndex(pageIndex *uint32) *GetProjectsParams {
	o.SetPageIndex(pageIndex)
	return o
}

// SetPageIndex adds the pageIndex to the get projects params
func (o *GetProjectsParams) SetPageIndex(pageIndex *uint32) {
	o.PageIndex = pageIndex
}

// WithPageSize adds the pageSize to the get projects params
func (o *GetProjectsParams) WithPageSize(pageSize *uint32) *GetProjectsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get projects params
func (o *GetProjectsParams) SetPageSize(pageSize *uint32) {
	o.PageSize = pageSize
}

// WithProjectConfigRepositoryCanonical adds the projectConfigRepositoryCanonical to the get projects params
func (o *GetProjectsParams) WithProjectConfigRepositoryCanonical(projectConfigRepositoryCanonical *uint32) *GetProjectsParams {
	o.SetProjectConfigRepositoryCanonical(projectConfigRepositoryCanonical)
	return o
}

// SetProjectConfigRepositoryCanonical adds the projectConfigRepositoryCanonical to the get projects params
func (o *GetProjectsParams) SetProjectConfigRepositoryCanonical(projectConfigRepositoryCanonical *uint32) {
	o.ProjectConfigRepositoryCanonical = projectConfigRepositoryCanonical
}

// WithProjectCreatedAt adds the projectCreatedAt to the get projects params
func (o *GetProjectsParams) WithProjectCreatedAt(projectCreatedAt *uint64) *GetProjectsParams {
	o.SetProjectCreatedAt(projectCreatedAt)
	return o
}

// SetProjectCreatedAt adds the projectCreatedAt to the get projects params
func (o *GetProjectsParams) SetProjectCreatedAt(projectCreatedAt *uint64) {
	o.ProjectCreatedAt = projectCreatedAt
}

// WithProjectDescription adds the projectDescription to the get projects params
func (o *GetProjectsParams) WithProjectDescription(projectDescription *string) *GetProjectsParams {
	o.SetProjectDescription(projectDescription)
	return o
}

// SetProjectDescription adds the projectDescription to the get projects params
func (o *GetProjectsParams) SetProjectDescription(projectDescription *string) {
	o.ProjectDescription = projectDescription
}

// WithProjectName adds the projectName to the get projects params
func (o *GetProjectsParams) WithProjectName(projectName *string) *GetProjectsParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the get projects params
func (o *GetProjectsParams) SetProjectName(projectName *string) {
	o.ProjectName = projectName
}

// WithServiceCatalogSourceCanonical adds the serviceCatalogSourceCanonical to the get projects params
func (o *GetProjectsParams) WithServiceCatalogSourceCanonical(serviceCatalogSourceCanonical *string) *GetProjectsParams {
	o.SetServiceCatalogSourceCanonical(serviceCatalogSourceCanonical)
	return o
}

// SetServiceCatalogSourceCanonical adds the serviceCatalogSourceCanonical to the get projects params
func (o *GetProjectsParams) SetServiceCatalogSourceCanonical(serviceCatalogSourceCanonical *string) {
	o.ServiceCatalogSourceCanonical = serviceCatalogSourceCanonical
}

// WithUserID adds the userID to the get projects params
func (o *GetProjectsParams) WithUserID(userID *uint32) *GetProjectsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get projects params
func (o *GetProjectsParams) SetUserID(userID *uint32) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.OrderBy != nil {

		// query param order_by
		var qrOrderBy string
		if o.OrderBy != nil {
			qrOrderBy = *o.OrderBy
		}
		qOrderBy := qrOrderBy
		if qOrderBy != "" {
			if err := r.SetQueryParam("order_by", qOrderBy); err != nil {
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

	if o.ProjectConfigRepositoryCanonical != nil {

		// query param project_config_repository_canonical
		var qrProjectConfigRepositoryCanonical uint32
		if o.ProjectConfigRepositoryCanonical != nil {
			qrProjectConfigRepositoryCanonical = *o.ProjectConfigRepositoryCanonical
		}
		qProjectConfigRepositoryCanonical := swag.FormatUint32(qrProjectConfigRepositoryCanonical)
		if qProjectConfigRepositoryCanonical != "" {
			if err := r.SetQueryParam("project_config_repository_canonical", qProjectConfigRepositoryCanonical); err != nil {
				return err
			}
		}

	}

	if o.ProjectCreatedAt != nil {

		// query param project_created_at
		var qrProjectCreatedAt uint64
		if o.ProjectCreatedAt != nil {
			qrProjectCreatedAt = *o.ProjectCreatedAt
		}
		qProjectCreatedAt := swag.FormatUint64(qrProjectCreatedAt)
		if qProjectCreatedAt != "" {
			if err := r.SetQueryParam("project_created_at", qProjectCreatedAt); err != nil {
				return err
			}
		}

	}

	if o.ProjectDescription != nil {

		// query param project_description
		var qrProjectDescription string
		if o.ProjectDescription != nil {
			qrProjectDescription = *o.ProjectDescription
		}
		qProjectDescription := qrProjectDescription
		if qProjectDescription != "" {
			if err := r.SetQueryParam("project_description", qProjectDescription); err != nil {
				return err
			}
		}

	}

	if o.ProjectName != nil {

		// query param project_name
		var qrProjectName string
		if o.ProjectName != nil {
			qrProjectName = *o.ProjectName
		}
		qProjectName := qrProjectName
		if qProjectName != "" {
			if err := r.SetQueryParam("project_name", qProjectName); err != nil {
				return err
			}
		}

	}

	if o.ServiceCatalogSourceCanonical != nil {

		// query param service_catalog_source_canonical
		var qrServiceCatalogSourceCanonical string
		if o.ServiceCatalogSourceCanonical != nil {
			qrServiceCatalogSourceCanonical = *o.ServiceCatalogSourceCanonical
		}
		qServiceCatalogSourceCanonical := qrServiceCatalogSourceCanonical
		if qServiceCatalogSourceCanonical != "" {
			if err := r.SetQueryParam("service_catalog_source_canonical", qServiceCatalogSourceCanonical); err != nil {
				return err
			}
		}

	}

	if o.UserID != nil {

		// query param user_id
		var qrUserID uint32
		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := swag.FormatUint32(qrUserID)
		if qUserID != "" {
			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
