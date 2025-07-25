// Code generated by go-swagger; DO NOT EDIT.

package organization_cloud_cost_management_filter_vaules

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
)

// NewGetCloudCostManagementFilterVaulesParams creates a new GetCloudCostManagementFilterVaulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCloudCostManagementFilterVaulesParams() *GetCloudCostManagementFilterVaulesParams {
	return &GetCloudCostManagementFilterVaulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCloudCostManagementFilterVaulesParamsWithTimeout creates a new GetCloudCostManagementFilterVaulesParams object
// with the ability to set a timeout on a request.
func NewGetCloudCostManagementFilterVaulesParamsWithTimeout(timeout time.Duration) *GetCloudCostManagementFilterVaulesParams {
	return &GetCloudCostManagementFilterVaulesParams{
		timeout: timeout,
	}
}

// NewGetCloudCostManagementFilterVaulesParamsWithContext creates a new GetCloudCostManagementFilterVaulesParams object
// with the ability to set a context for a request.
func NewGetCloudCostManagementFilterVaulesParamsWithContext(ctx context.Context) *GetCloudCostManagementFilterVaulesParams {
	return &GetCloudCostManagementFilterVaulesParams{
		Context: ctx,
	}
}

// NewGetCloudCostManagementFilterVaulesParamsWithHTTPClient creates a new GetCloudCostManagementFilterVaulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCloudCostManagementFilterVaulesParamsWithHTTPClient(client *http.Client) *GetCloudCostManagementFilterVaulesParams {
	return &GetCloudCostManagementFilterVaulesParams{
		HTTPClient: client,
	}
}

/*
GetCloudCostManagementFilterVaulesParams contains all the parameters to send to the API endpoint

	for the get cloud cost management filter vaules operation.

	Typically these are written to a http.Request.
*/
type GetCloudCostManagementFilterVaulesParams struct {

	/* Begin.

	   The unix timestamp in milliseconds, which indicate the start of the time range.

	   Format: uint64
	*/
	Begin *uint64

	/* Categories.

	   The names of the categories that you can use to filter and group your results.
	*/
	Categories []string

	/* Components.

	   The names of the components that you can use to filter your results
	*/
	Components []string

	/* Dashboard.

	   Specifies if the filter values are for the dashboard
	*/
	Dashboard *bool

	/* End.

	   The unix timestamp in milliseconds, which indicate the end of the time range.

	   Format: uint64
	*/
	End *uint64

	/* Environments.

	   The names of the environments that you can use to filter your results
	*/
	Environments []string

	/* GroupBy.

	   Represents a group when you specify a group by criteria, or in the response to a query with a specific grouping.
	*/
	GroupBy []string

	/* LinkedAccounts.

	   The ids of the linked accounts that you can use to filter your results
	*/
	LinkedAccounts []string

	/* MasterAccounts.

	   The ids of the master accounts that you can use to filter your results
	*/
	MasterAccounts []string

	/* OrganizationCanonical.

	   A canonical of an organization.
	*/
	OrganizationCanonical string

	/* Projects.

	   The names of the projects that you can use to filter your results
	*/
	Projects []string

	/* Regions.

	   The names of the regions that you can use to filter your results
	*/
	Regions []string

	/* ResourceTagging.

	   Filter results by only tagged or not tagged resources
	*/
	ResourceTagging *string

	/* Services.

	   The names of the services that you can use to filter and group your results.
	*/
	Services []string

	/* Tags.

	   The key and value of a tag concatenated by a ;.
	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cloud cost management filter vaules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCloudCostManagementFilterVaulesParams) WithDefaults() *GetCloudCostManagementFilterVaulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cloud cost management filter vaules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCloudCostManagementFilterVaulesParams) SetDefaults() {
	var (
		dashboardDefault = bool(false)
	)

	val := GetCloudCostManagementFilterVaulesParams{
		Dashboard: &dashboardDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) WithTimeout(timeout time.Duration) *GetCloudCostManagementFilterVaulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) WithContext(ctx context.Context) *GetCloudCostManagementFilterVaulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) WithHTTPClient(client *http.Client) *GetCloudCostManagementFilterVaulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBegin adds the begin to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) WithBegin(begin *uint64) *GetCloudCostManagementFilterVaulesParams {
	o.SetBegin(begin)
	return o
}

// SetBegin adds the begin to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) SetBegin(begin *uint64) {
	o.Begin = begin
}

// WithCategories adds the categories to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) WithCategories(categories []string) *GetCloudCostManagementFilterVaulesParams {
	o.SetCategories(categories)
	return o
}

// SetCategories adds the categories to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) SetCategories(categories []string) {
	o.Categories = categories
}

// WithComponents adds the components to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) WithComponents(components []string) *GetCloudCostManagementFilterVaulesParams {
	o.SetComponents(components)
	return o
}

// SetComponents adds the components to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) SetComponents(components []string) {
	o.Components = components
}

// WithDashboard adds the dashboard to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) WithDashboard(dashboard *bool) *GetCloudCostManagementFilterVaulesParams {
	o.SetDashboard(dashboard)
	return o
}

// SetDashboard adds the dashboard to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) SetDashboard(dashboard *bool) {
	o.Dashboard = dashboard
}

// WithEnd adds the end to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) WithEnd(end *uint64) *GetCloudCostManagementFilterVaulesParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) SetEnd(end *uint64) {
	o.End = end
}

// WithEnvironments adds the environments to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) WithEnvironments(environments []string) *GetCloudCostManagementFilterVaulesParams {
	o.SetEnvironments(environments)
	return o
}

// SetEnvironments adds the environments to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) SetEnvironments(environments []string) {
	o.Environments = environments
}

// WithGroupBy adds the groupBy to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) WithGroupBy(groupBy []string) *GetCloudCostManagementFilterVaulesParams {
	o.SetGroupBy(groupBy)
	return o
}

// SetGroupBy adds the groupBy to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) SetGroupBy(groupBy []string) {
	o.GroupBy = groupBy
}

// WithLinkedAccounts adds the linkedAccounts to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) WithLinkedAccounts(linkedAccounts []string) *GetCloudCostManagementFilterVaulesParams {
	o.SetLinkedAccounts(linkedAccounts)
	return o
}

// SetLinkedAccounts adds the linkedAccounts to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) SetLinkedAccounts(linkedAccounts []string) {
	o.LinkedAccounts = linkedAccounts
}

// WithMasterAccounts adds the masterAccounts to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) WithMasterAccounts(masterAccounts []string) *GetCloudCostManagementFilterVaulesParams {
	o.SetMasterAccounts(masterAccounts)
	return o
}

// SetMasterAccounts adds the masterAccounts to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) SetMasterAccounts(masterAccounts []string) {
	o.MasterAccounts = masterAccounts
}

// WithOrganizationCanonical adds the organizationCanonical to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) WithOrganizationCanonical(organizationCanonical string) *GetCloudCostManagementFilterVaulesParams {
	o.SetOrganizationCanonical(organizationCanonical)
	return o
}

// SetOrganizationCanonical adds the organizationCanonical to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) SetOrganizationCanonical(organizationCanonical string) {
	o.OrganizationCanonical = organizationCanonical
}

// WithProjects adds the projects to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) WithProjects(projects []string) *GetCloudCostManagementFilterVaulesParams {
	o.SetProjects(projects)
	return o
}

// SetProjects adds the projects to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) SetProjects(projects []string) {
	o.Projects = projects
}

// WithRegions adds the regions to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) WithRegions(regions []string) *GetCloudCostManagementFilterVaulesParams {
	o.SetRegions(regions)
	return o
}

// SetRegions adds the regions to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) SetRegions(regions []string) {
	o.Regions = regions
}

// WithResourceTagging adds the resourceTagging to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) WithResourceTagging(resourceTagging *string) *GetCloudCostManagementFilterVaulesParams {
	o.SetResourceTagging(resourceTagging)
	return o
}

// SetResourceTagging adds the resourceTagging to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) SetResourceTagging(resourceTagging *string) {
	o.ResourceTagging = resourceTagging
}

// WithServices adds the services to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) WithServices(services []string) *GetCloudCostManagementFilterVaulesParams {
	o.SetServices(services)
	return o
}

// SetServices adds the services to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) SetServices(services []string) {
	o.Services = services
}

// WithTags adds the tags to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) WithTags(tags []string) *GetCloudCostManagementFilterVaulesParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get cloud cost management filter vaules params
func (o *GetCloudCostManagementFilterVaulesParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *GetCloudCostManagementFilterVaulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Categories != nil {

		// binding items for categories
		joinedCategories := o.bindParamCategories(reg)

		// query array param categories
		if err := r.SetQueryParam("categories", joinedCategories...); err != nil {
			return err
		}
	}

	if o.Components != nil {

		// binding items for components
		joinedComponents := o.bindParamComponents(reg)

		// query array param components
		if err := r.SetQueryParam("components", joinedComponents...); err != nil {
			return err
		}
	}

	if o.Dashboard != nil {

		// query param dashboard
		var qrDashboard bool

		if o.Dashboard != nil {
			qrDashboard = *o.Dashboard
		}
		qDashboard := swag.FormatBool(qrDashboard)
		if qDashboard != "" {

			if err := r.SetQueryParam("dashboard", qDashboard); err != nil {
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

	if o.Environments != nil {

		// binding items for environments
		joinedEnvironments := o.bindParamEnvironments(reg)

		// query array param environments
		if err := r.SetQueryParam("environments", joinedEnvironments...); err != nil {
			return err
		}
	}

	if o.GroupBy != nil {

		// binding items for group_by
		joinedGroupBy := o.bindParamGroupBy(reg)

		// query array param group_by
		if err := r.SetQueryParam("group_by", joinedGroupBy...); err != nil {
			return err
		}
	}

	if o.LinkedAccounts != nil {

		// binding items for linked_accounts
		joinedLinkedAccounts := o.bindParamLinkedAccounts(reg)

		// query array param linked_accounts
		if err := r.SetQueryParam("linked_accounts", joinedLinkedAccounts...); err != nil {
			return err
		}
	}

	if o.MasterAccounts != nil {

		// binding items for master_accounts
		joinedMasterAccounts := o.bindParamMasterAccounts(reg)

		// query array param master_accounts
		if err := r.SetQueryParam("master_accounts", joinedMasterAccounts...); err != nil {
			return err
		}
	}

	// path param organization_canonical
	if err := r.SetPathParam("organization_canonical", o.OrganizationCanonical); err != nil {
		return err
	}

	if o.Projects != nil {

		// binding items for projects
		joinedProjects := o.bindParamProjects(reg)

		// query array param projects
		if err := r.SetQueryParam("projects", joinedProjects...); err != nil {
			return err
		}
	}

	if o.Regions != nil {

		// binding items for regions
		joinedRegions := o.bindParamRegions(reg)

		// query array param regions
		if err := r.SetQueryParam("regions", joinedRegions...); err != nil {
			return err
		}
	}

	if o.ResourceTagging != nil {

		// query param resource_tagging
		var qrResourceTagging string

		if o.ResourceTagging != nil {
			qrResourceTagging = *o.ResourceTagging
		}
		qResourceTagging := qrResourceTagging
		if qResourceTagging != "" {

			if err := r.SetQueryParam("resource_tagging", qResourceTagging); err != nil {
				return err
			}
		}
	}

	if o.Services != nil {

		// binding items for services
		joinedServices := o.bindParamServices(reg)

		// query array param services
		if err := r.SetQueryParam("services", joinedServices...); err != nil {
			return err
		}
	}

	if o.Tags != nil {

		// binding items for tags
		joinedTags := o.bindParamTags(reg)

		// query array param tags
		if err := r.SetQueryParam("tags", joinedTags...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetCloudCostManagementFilterVaules binds the parameter categories
func (o *GetCloudCostManagementFilterVaulesParams) bindParamCategories(formats strfmt.Registry) []string {
	categoriesIR := o.Categories

	var categoriesIC []string
	for _, categoriesIIR := range categoriesIR { // explode []string

		categoriesIIV := categoriesIIR // string as string
		categoriesIC = append(categoriesIC, categoriesIIV)
	}

	// items.CollectionFormat: "multi"
	categoriesIS := swag.JoinByFormat(categoriesIC, "multi")

	return categoriesIS
}

// bindParamGetCloudCostManagementFilterVaules binds the parameter components
func (o *GetCloudCostManagementFilterVaulesParams) bindParamComponents(formats strfmt.Registry) []string {
	componentsIR := o.Components

	var componentsIC []string
	for _, componentsIIR := range componentsIR { // explode []string

		componentsIIV := componentsIIR // string as string
		componentsIC = append(componentsIC, componentsIIV)
	}

	// items.CollectionFormat: "multi"
	componentsIS := swag.JoinByFormat(componentsIC, "multi")

	return componentsIS
}

// bindParamGetCloudCostManagementFilterVaules binds the parameter environments
func (o *GetCloudCostManagementFilterVaulesParams) bindParamEnvironments(formats strfmt.Registry) []string {
	environmentsIR := o.Environments

	var environmentsIC []string
	for _, environmentsIIR := range environmentsIR { // explode []string

		environmentsIIV := environmentsIIR // string as string
		environmentsIC = append(environmentsIC, environmentsIIV)
	}

	// items.CollectionFormat: "multi"
	environmentsIS := swag.JoinByFormat(environmentsIC, "multi")

	return environmentsIS
}

// bindParamGetCloudCostManagementFilterVaules binds the parameter group_by
func (o *GetCloudCostManagementFilterVaulesParams) bindParamGroupBy(formats strfmt.Registry) []string {
	groupByIR := o.GroupBy

	var groupByIC []string
	for _, groupByIIR := range groupByIR { // explode []string

		groupByIIV := groupByIIR // string as string
		groupByIC = append(groupByIC, groupByIIV)
	}

	// items.CollectionFormat: "multi"
	groupByIS := swag.JoinByFormat(groupByIC, "multi")

	return groupByIS
}

// bindParamGetCloudCostManagementFilterVaules binds the parameter linked_accounts
func (o *GetCloudCostManagementFilterVaulesParams) bindParamLinkedAccounts(formats strfmt.Registry) []string {
	linkedAccountsIR := o.LinkedAccounts

	var linkedAccountsIC []string
	for _, linkedAccountsIIR := range linkedAccountsIR { // explode []string

		linkedAccountsIIV := linkedAccountsIIR // string as string
		linkedAccountsIC = append(linkedAccountsIC, linkedAccountsIIV)
	}

	// items.CollectionFormat: "multi"
	linkedAccountsIS := swag.JoinByFormat(linkedAccountsIC, "multi")

	return linkedAccountsIS
}

// bindParamGetCloudCostManagementFilterVaules binds the parameter master_accounts
func (o *GetCloudCostManagementFilterVaulesParams) bindParamMasterAccounts(formats strfmt.Registry) []string {
	masterAccountsIR := o.MasterAccounts

	var masterAccountsIC []string
	for _, masterAccountsIIR := range masterAccountsIR { // explode []string

		masterAccountsIIV := masterAccountsIIR // string as string
		masterAccountsIC = append(masterAccountsIC, masterAccountsIIV)
	}

	// items.CollectionFormat: "multi"
	masterAccountsIS := swag.JoinByFormat(masterAccountsIC, "multi")

	return masterAccountsIS
}

// bindParamGetCloudCostManagementFilterVaules binds the parameter projects
func (o *GetCloudCostManagementFilterVaulesParams) bindParamProjects(formats strfmt.Registry) []string {
	projectsIR := o.Projects

	var projectsIC []string
	for _, projectsIIR := range projectsIR { // explode []string

		projectsIIV := projectsIIR // string as string
		projectsIC = append(projectsIC, projectsIIV)
	}

	// items.CollectionFormat: "multi"
	projectsIS := swag.JoinByFormat(projectsIC, "multi")

	return projectsIS
}

// bindParamGetCloudCostManagementFilterVaules binds the parameter regions
func (o *GetCloudCostManagementFilterVaulesParams) bindParamRegions(formats strfmt.Registry) []string {
	regionsIR := o.Regions

	var regionsIC []string
	for _, regionsIIR := range regionsIR { // explode []string

		regionsIIV := regionsIIR // string as string
		regionsIC = append(regionsIC, regionsIIV)
	}

	// items.CollectionFormat: "multi"
	regionsIS := swag.JoinByFormat(regionsIC, "multi")

	return regionsIS
}

// bindParamGetCloudCostManagementFilterVaules binds the parameter services
func (o *GetCloudCostManagementFilterVaulesParams) bindParamServices(formats strfmt.Registry) []string {
	servicesIR := o.Services

	var servicesIC []string
	for _, servicesIIR := range servicesIR { // explode []string

		servicesIIV := servicesIIR // string as string
		servicesIC = append(servicesIC, servicesIIV)
	}

	// items.CollectionFormat: "multi"
	servicesIS := swag.JoinByFormat(servicesIC, "multi")

	return servicesIS
}

// bindParamGetCloudCostManagementFilterVaules binds the parameter tags
func (o *GetCloudCostManagementFilterVaulesParams) bindParamTags(formats strfmt.Registry) []string {
	tagsIR := o.Tags

	var tagsIC []string
	for _, tagsIIR := range tagsIR { // explode []string

		tagsIIV := tagsIIR // string as string
		tagsIC = append(tagsIC, tagsIIV)
	}

	// items.CollectionFormat: "multi"
	tagsIS := swag.JoinByFormat(tagsIC, "multi")

	return tagsIS
}
